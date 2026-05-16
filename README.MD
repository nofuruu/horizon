# Horizon 
## Monitoring dashboard untuk sistem linux

> Sebuah program untuk menjalankan kode dari go untuk mengecek resource komputer

* Note : program ini belum diuji/ditest apabila dijalankan di windows
---

- Program dari go akan dicompile menjadi sebuah program/binary program yang akan dijalankan di mesin linux (dapat berjalan di background sebagai service). Karena program berjalan langsung pada sistem operasi linux maka program memiliki akses ke resource komputer(file program)

- untuk melihat penggunaan RAM maka program akan mengakses resource komputer via file /proc/meminfo pada sistem linux

- untuk melihat beban CPU via /proc/stat

- untuk melihat beban suhu dari desktop maka membaca file dari /sys/class/thermal

---
## Example

### Cara Kerja Sistem
![Example](assets/example.png)
---
### Dashboard 
![Example](assets/dashboard_example.png)
---
### Cpu Usage 
![Example](assets/cpu_example.png)
