# Wipe Tool

A secure data wiping application with a Go backend and Tauri frontend.

if you don't know how to install tauri v2 look at the [Tauri Docs](https://v2.tauri.app/start/create-project/#manual-setup-tauri-cli) or something

This project provides a **desktop application** for secure disk and partition wiping.  
It combines a **Go backend** (device operations) with a **Tauri + Node.js frontend** (cross-platform UI).

**Warning:** This tool performs **irreversible data destruction**.  
Use at your own risk. Always double-check the selected device before running.

## Features
- Cross-platform UI built with **Tauri**.
- Backend written in **Go** for direct disk operations.
- Planned support for both **HDD (overwrite methods)** and **SSD/NVMe (secure erase / blkdiscard)**.
- Extensible architecture for adding new wiping algorithms.


## Supported Platforms
- **Linux** (primary target)
- **macOS** (limited — some commands differ)
- **Windows** (planned, not stable yet)

> **Note:** Device handling is Linux-oriented (`/dev/sdX`, `/dev/nvmeXnY`). Other platforms may not work properly yet.


## Prerequisites
Install the following before building:

- [Go](https://go.dev/dl/) **1.20+**
- [Rust](https://www.rust-lang.org/tools/install) (stable)
- [Node.js](https://nodejs.org/en/download/) **18+**
- [Tauri v2 prerequisites](https://tauri.app/v2/guides/getting-started/prerequisites) (C++ build tools, WebView2 for Windows, etc.)
- Linux users: `udev`, `nvme-cli`, `hdparm`, `util-linux` recommended


## Installation

Clone the repository:

    git clone https://github.com/shivjeet1/sih2025.git

As Always:

    cd sih2025/ui 
    npm install

To Run:

    tauri dev

To build:

    tauri build
