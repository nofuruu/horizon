<script lang="ts">
	import * as Card from "$lib/components/ui/card/index.js";
	import * as Chart from "$lib/components/ui/chart/index.js";
	import { Badge } from "$lib/components/ui/badge/index.js";
	import { Skeleton } from "$lib/components/ui/skeleton/index.js";
	import AnimatedCounter from "$lib/components/animated-counter.svelte";
	import { AreaChart } from "layerchart";
	import { scaleUtc } from "d3-scale";
	import { mode } from "mode-watcher";
	import CpuIcon from "@lucide/svelte/icons/cpu";
	import MemoryStickIcon from "@lucide/svelte/icons/memory-stick";
	import HardDriveIcon from "@lucide/svelte/icons/hard-drive";
	import ServerIcon from "@lucide/svelte/icons/server";
	import ActivityIcon from "@lucide/svelte/icons/activity";
	import { onMount, onDestroy } from 'svelte';

	// --- State ---
	let metrics = $state<any>(null);
	let cpuHistory = $state<any[]>([]);
	let memoryHistory = $state<any[]>([]);
	let loading = $state(true);
	let error = $state<string | null>(null);
	let intervalId: number | null = null;

	// API endpoint
	const API_BASE = 'http://localhost:8080/api';

	// --- Chart configs ---
	const cpuChartConfig = {
		usage: { label: "CPU Usage", color: "var(--chart-1)" },
	} satisfies Chart.ChartConfig;

	const memoryChartConfig = {
		used: { label: "Used Memory", color: "var(--chart-3)" },
	} satisfies Chart.ChartConfig;

	// --- Fetch metrics ---
	async function fetchMetrics() {
		try {
			const response = await fetch(`${API_BASE}/metrics`);
			if (!response.ok) throw new Error('Failed to fetch metrics');
			
			const data = await response.json();
			metrics = data;
			
			// Add to history for charts (keep last 20 data points)
			const now = new Date();
			cpuHistory = [...cpuHistory, { 
				time: now, 
				usage: data.cpu.usage_percent 
			}].slice(-20);
			
			memoryHistory = [...memoryHistory, { 
				time: now, 
				used: data.memory.used_percent 
			}].slice(-20);
			
			loading = false;
			error = null;
		} catch (err) {
			error = err instanceof Error ? err.message : 'Unknown error';
			loading = false;
		}
	}

	// --- System info ---
	let systemInfo = $state<any>(null);
	async function fetchSystemInfo() {
		try {
			const response = await fetch(`${API_BASE}/system`);
			if (!response.ok) throw new Error('Failed to fetch system info');
			systemInfo = await response.json();
		} catch (err) {
			console.error('Failed to fetch system info:', err);
		}
	}

	// --- Lifecycle ---
	onMount(() => {
		fetchMetrics();
		fetchSystemInfo();
		// Update every 2 seconds
		intervalId = setInterval(fetchMetrics, 2000);
	});

	onDestroy(() => {
		if (intervalId) clearInterval(intervalId);
	});

	// --- Derived stats ---
	const stats = $derived(metrics ? [
		{
			title: "CPU Usage",
			value: metrics.cpu.usage_percent.toFixed(1),
			unit: "%",
			icon: CpuIcon,
			cores: metrics.cpu.cores,
			status: metrics.cpu.usage_percent > 80 ? 'critical' : metrics.cpu.usage_percent > 60 ? 'warning' : 'normal'
		},
		{
			title: "Memory Usage",
			value: metrics.memory.used_percent.toFixed(1),
			unit: "%",
			icon: MemoryStickIcon,
			total: (metrics.memory.total / (1024**3)).toFixed(1),
			used: (metrics.memory.used / (1024**3)).toFixed(1),
			status: metrics.memory.used_percent > 85 ? 'critical' : metrics.memory.used_percent > 70 ? 'warning' : 'normal'
		},
		{
			title: "Disk Usage",
			value: metrics.disk.used_percent.toFixed(1),
			unit: "%",
			icon: HardDriveIcon,
			total: (metrics.disk.total / (1024**3)).toFixed(0),
			used: (metrics.disk.used / (1024**3)).toFixed(0),
			status: metrics.disk.used_percent > 90 ? 'critical' : metrics.disk.used_percent > 75 ? 'warning' : 'normal'
		},
	] : []);

	// --- Helpers ---
	function formatUptime(seconds: number): string {
		const days = Math.floor(seconds / 86400);
		const hours = Math.floor((seconds % 86400) / 3600);
		const minutes = Math.floor((seconds % 3600) / 60);
		return `${days}d ${hours}h ${minutes}m`;
	}

	function getStatusBadge(status: string) {
		switch (status) {
			case 'critical':
				return 'destructive';
			case 'warning':
				return 'outline';
			default:
				return 'default';
		}
	}

	function getStatusColor(status: string) {
		switch (status) {
			case 'critical':
				return 'text-red-600 dark:text-red-400';
			case 'warning':
				return 'text-yellow-600 dark:text-yellow-400';
			default:
				return 'text-green-600 dark:text-green-400';
		}
	}
</script>

<svelte:head>
	<title>LinuxPulse - System Monitoring</title>
</svelte:head>

<div class="space-y-6">
	<!-- Header -->
	<div class="flex items-end justify-between">
		<div>
			<h1 class="text-3xl font-bold tracking-tight">System Monitor</h1>
			<p class="text-muted-foreground">Real-time Linux system metrics and performance</p>
		</div>
		<div class="flex items-center gap-2">
			<ActivityIcon class="size-4 text-green-500 animate-pulse" />
			<span class="text-sm text-muted-foreground">Live</span>
		</div>
	</div>

	{#if error}
		<Card.Root class="border-red-500">
			<Card.Content class="pt-6">
				<p class="text-red-600 dark:text-red-400">Error: {error}</p>
				<p class="text-sm text-muted-foreground mt-2">Make sure the Go backend is running on localhost:8080</p>
			</Card.Content>
		</Card.Root>
	{/if}

	<!-- Row 1: System Info -->
	{#if systemInfo}
		<Card.Root>
			<Card.Header>
				<Card.Title class="flex items-center gap-2">
					<ServerIcon class="size-5" />
					System Information
				</Card.Title>
			</Card.Header>
			<Card.Content class="grid gap-4 md:grid-cols-4">
				<div>
					<p class="text-sm text-muted-foreground">Hostname</p>
					<p class="text-lg font-semibold">{systemInfo.hostname}</p>
				</div>
				<div>
					<p class="text-sm text-muted-foreground">OS</p>
					<p class="text-lg font-semibold">{systemInfo.platform} {systemInfo.platform_version}</p>
				</div>
				<div>
					<p class="text-sm text-muted-foreground">Kernel</p>
					<p class="text-lg font-semibold">{systemInfo.kernel_version}</p>
				</div>
				<div>
					<p class="text-sm text-muted-foreground">Uptime</p>
					<p class="text-lg font-semibold">{formatUptime(systemInfo.uptime)}</p>
				</div>
			</Card.Content>
		</Card.Root>
	{/if}

	<!-- Row 2: KPI Cards -->
	{#if loading && !metrics}
		<div class="grid gap-4 md:grid-cols-3">
			{#each [1, 2, 3] as _}
				<Card.Root>
					<Card.Header class="pb-2">
						<Skeleton class="h-4 w-24" />
					</Card.Header>
					<Card.Content>
						<Skeleton class="h-10 w-20" />
						<Skeleton class="h-4 w-32 mt-2" />
					</Card.Content>
				</Card.Root>
			{/each}
		</div>
	{:else if metrics}
		<div class="grid gap-4 md:grid-cols-3">
			{#each stats as stat}
				<Card.Root class="relative overflow-hidden">
					<Card.Header class="flex flex-row items-center justify-between space-y-0 pb-2">
						<Card.Title class="text-sm font-medium">{stat.title}</Card.Title>
						<stat.icon class={`size-5 ${getStatusColor(stat.status)}`} />
					</Card.Header>
					<Card.Content>
						<div class="flex items-baseline gap-1">
							<span class="text-3xl font-bold">
								<AnimatedCounter value={parseFloat(stat.value)} decimals={1} />
							</span>
							<span class="text-xl text-muted-foreground">{stat.unit}</span>
						</div>
						
						{#if stat.cores}
							<p class="text-xs text-muted-foreground mt-2">{stat.cores} cores available</p>
						{:else if stat.total}
							<p class="text-xs text-muted-foreground mt-2">
								{stat.used} GB / {stat.total} GB used
							</p>
						{/if}

						<div class="mt-3">
							<Badge variant={getStatusBadge(stat.status)} class="text-[10px]">
								{stat.status.toUpperCase()}
							</Badge>
						</div>
					</Card.Content>
					
					<!-- Background indicator bar -->
					<div class="absolute bottom-0 left-0 right-0 h-1 bg-muted">
						<div 
							class="h-full transition-all duration-500 {
								stat.status === 'critical' ? 'bg-red-500' : 
								stat.status === 'warning' ? 'bg-yellow-500' : 
								'bg-green-500'
							}"
							style="width: {stat.value}%"
						></div>
					</div>
				</Card.Root>
			{/each}
		</div>
	{/if}

	<!-- Row 3: Charts -->
	<div class="grid gap-4 md:grid-cols-2">
		<!-- CPU Chart -->
		<Card.Root>
			<Card.Header>
				<Card.Title>CPU Usage History</Card.Title>
				<Card.Description>Real-time CPU utilization (last 40 seconds)</Card.Description>
			</Card.Header>
			<Card.Content>
				{#key mode.current}
					{#if cpuHistory.length > 0}
						<Chart.Container config={cpuChartConfig} class="h-[300px] w-full">
							<AreaChart
								data={cpuHistory}
								x="time"
								y="usage"
								xScale={scaleUtc()}
								props={{
									area: { 
										fill: 'var(--chart-1)',
										fillOpacity: 0.2
									},
									line: { 
										stroke: 'var(--chart-1)',
										strokeWidth: 2
									},
									xAxis: {
										format: (d: Date) => {
											return d.toLocaleTimeString('en-US', { 
												hour: '2-digit', 
												minute: '2-digit',
												second: '2-digit'
											});
										}
									},
									yAxis: {
										format: (d: number) => `${d.toFixed(0)}%`
									}
								}}
							>
								{#snippet tooltip()}
									<Chart.Tooltip />
								{/snippet}
							</AreaChart>
						</Chart.Container>
					{:else}
						<div class="flex h-[300px] items-center justify-center">
							<p class="text-muted-foreground">Collecting data...</p>
						</div>
					{/if}
				{/key}
			</Card.Content>
		</Card.Root>

		<!-- Memory Chart -->
		<Card.Root>
			<Card.Header>
				<Card.Title>Memory Usage History</Card.Title>
				<Card.Description>Real-time memory utilization (last 40 seconds)</Card.Description>
			</Card.Header>
			<Card.Content>
				{#key mode.current}
					{#if memoryHistory.length > 0}
						<Chart.Container config={memoryChartConfig} class="h-[300px] w-full">
							<AreaChart
								data={memoryHistory}
								x="time"
								y="used"
								xScale={scaleUtc()}
								props={{
									area: { 
										fill: 'var(--chart-3)',
										fillOpacity: 0.2
									},
									line: { 
										stroke: 'var(--chart-3)',
										strokeWidth: 2
									},
									xAxis: {
										format: (d: Date) => {
											return d.toLocaleTimeString('en-US', { 
												hour: '2-digit', 
												minute: '2-digit',
												second: '2-digit'
											});
										}
									},
									yAxis: {
										format: (d: number) => `${d.toFixed(0)}%`
									}
								}}
							>
								{#snippet tooltip()}
									<Chart.Tooltip />
								{/snippet}
							</AreaChart>
						</Chart.Container>
					{:else}
						<div class="flex h-[300px] items-center justify-center">
							<p class="text-muted-foreground">Collecting data...</p>
						</div>
					{/if}
				{/key}
			</Card.Content>
		</Card.Root>
	</div>
</div>