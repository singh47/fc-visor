
# fc-visor

**fc-visor** is a CLI tool for monitoring and inspecting [Firecracker](https://github.com/firecracker-microvm/firecracker) microVMs. It offers metrics, info, and a top-like dashboard (beta version) via the Firecracker API socket.

---

## Features

- List running Firecracker VMs

- Fetch VM configuration and metrics (`/machine-config` and `/metrics`)

- Live `top`-style TUI

- Pluggable and lightweight

---
## 🛠️ Prerequisites

- Go 1.18+ installed

- Firecracker VMs running with exposed API sockets

- Unix-based OS (Linux/macOS)

---

## Installation

### Clone the repo

```bash
git  clone  https://github.com/yourname/fc-visor.git

cd  fc-visor

go  mod  tidy

go  build  -o  fc-visor
```

  

### Usage
##### List running Firecracker sockets
```bash
./fc-visor list
```
##### Inspect a VM's config
```bash
./fc-visor inspect --socket /path/to/api.socket
```
##### Fetch metrics from a VM
```bash
./fc-visor metrics --socket /path/to/api.socket
```
##### Run live top-like monitor - (unstable - a beta feature)
```bash
./fc-visor top --socket /path/to/api.socket
```
