package main

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"timestamp": time.Now().Unix(),
		})
	})

	api := r.Group("/api")
	{
		api.GET("/system", getSystemInfo)
		api.GET("/cpu", getCPUMetrics)
		api.GET("/memory", getMemoryMetrics)
		api.GET("/disk", getDiskMetrics)
		api.GET("/metrics", getAllMetrics)
	}

	port := ":8080"
	log.Printf("Server is running on port http://localhost%s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getSystemInfo(c *gin.Context) {
	info, err := host.Info()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"hostname":         info.Hostname,
		"os":               info.OS,
		"platform":         info.Platform,
		"platform_family":  info.PlatformFamily,
		"platform_version": info.PlatformVersion,
		"kernel_version":   info.KernelVersion,
		"architecture":     info.KernelArch,
		"uptime":           info.Uptime,
		"boot_time":        info.BootTime,
	})
}

func getCPUMetrics(c *gin.Context) {
	precentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	perCorePercent, _ := cpu.Percent(time.Second, true)

	c.JSON(http.StatusOK, gin.H{
		"usage_percent": precentages[0],
		"cores":         runtime.NumCPU(),
		"model":         cpuInfo[0].ModelName,
		"mhz":           cpuInfo[0].Mhz,
		"per_core":      perCorePercent,
	})
}

func getMemoryMetrics(c *gin.Context) {
	vnem, err := mem.VirtualMemory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":        vnem.Total,
		"available":    vnem.Available,
		"used":         vnem.Used,
		"used_percent": vnem.UsedPercent,
		"free":         vnem.Free,
	})
}

func getDiskMetrics(c *gin.Context) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var diskInfo []gin.H
	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue
		}

		diskInfo = append(diskInfo, gin.H{
			"device":       partition.Device,      
			"mountpoint":   partition.Mountpoint,  
			"fstype":       partition.Fstype,      
			"total":        usage.Total,
			"used":         usage.Used,
			"free":         usage.Free,
			"used_percent": usage.UsedPercent,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"disks": diskInfo,
	})
}

func getAllMetrics(c *gin.Context) {
	cpuPercent, _ := cpu.Percent(time.Second, false)
	vmem, _ := mem.VirtualMemory()
	diskUsage, _ := disk.Usage("/")

	c.JSON(http.StatusOK, gin.H{
		"cpu": gin.H{
			"usage_percent": cpuPercent[0],
			"cores":         runtime.NumCPU(),
		},
		"memory": gin.H{
			"total":        vmem.Total,
			"used":         vmem.Used,
			"used_percent": vmem.UsedPercent,
		},
		"disk": gin.H{
			"total":        diskUsage.Total,
			"used":         diskUsage.Used,
			"used_percent": diskUsage.UsedPercent,
		},
		"timestamp": time.Now().Unix(),
	})
}
