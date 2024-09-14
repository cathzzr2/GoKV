# GoKV 📊

GoKV is a high-performance key-value store implemented in Go, leveraging the [bitcask](https://medium.com/@arpitbhayani/bitcask-a-log-structured-fast-kv-store-c6c728a9536b) model. It is designed for simplicity, speed, and flexibility, with optimization strategies to enhance data retrieval and storage. GoKV supports Redis data structures and offers a user-friendly HTTP API for interaction.

<img width="300" alt="Screenshot 2024-09-07 at 22 34 12" src="https://github.com/user-attachments/assets/009ba424-68f4-4ba8-8fee-20c972a150ad">

## Features
- **Basic Key-Value Storage:** Supports standard operations like read, write, and delete.
- **In-Memory Indexing:** Accelerates read operations by minimizing disk access.
- **File Management & Optimization:** Includes efficient I/O handling and merge operations.
- **Redis Compatibility:** Supports Redis data structures and protocols.
- **Backup & Testing:** Ensures reliable data storage and retrieval with backup mechanisms and thorough testing.
- **HTTP API:** A simple, RESTful API for interacting with the key-value store.

## File Structure
`/benchmark`: Contains benchmarking scripts for performance evaluation.

`/data`: Stores database files.

`/examples`: Includes example implementations of GoKV.

`/http`: Defines HTTP routes and CRUD handlers.

`/index`: Manages key indexing for faster lookups.

`/redis`: Implements Redis protocol compatibility.

`/utils`: Utility functions to support core features.

## Installation:
```bash
git clone https://github.com/cathzzr2/GoKV.git
cd GoKV
go get
go run ./cmd/server
```
