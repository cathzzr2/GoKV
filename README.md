# GoKV 📊

GoKV is a high-performance key-value store implemented in Go leveraging [bitcask](https://medium.com/@arpitbhayani/bitcask-a-log-structured-fast-kv-store-c6c728a9536b) model. It is designed for simplicity, speed, and flexibility, featuring a range of optimization strategies to accelerate data retrieval and storage. GoKV also supports Redis data structures and provides a user-friendly HTTP API for easy interaction.

<img width="300" alt="Screenshot 2024-09-07 at 22 34 12" src="https://github.com/user-attachments/assets/009ba424-68f4-4ba8-8fee-20c972a150ad">

## Features:
- **Basic Key-Value Storage Operations:** GoKV supports the basic functionalities of key-value stores, including data read/write, merge operations, and WriteBatch for efficient bulk operations.

- **In-Memory Indexing:** By storing keys in memory, GoKV accelerates read operations, enabling quick data retrieval with minimal disk access.

- **File Management and Optimizations:** Implements common optimization strategies, including in-memory indexing, file I/O management, and the optimization of merge operations, enhancing the overall efficiency of data access.

- **Redis Compatibility:** GoKV extends its functionality by supporting Redis data structures and protocols, enabling seamless data migration and interaction with existing Redis-based systems.

- **Backup, Testing, and HTTP API Integration:** Provides reliable data backup mechanisms, thorough testing, and a simple HTTP API for CRUD operations, making GoKV versatile and developer-friendly.

- **User-Friendly HTTP API:** GoKV exposes an intuitive API that supports Create, Read, Update, and Delete operations through standard HTTP protocols, allowing developers to interact with the database effortlessly.

## Installation:
```bash
git clone https://github.com/cathzzr2/GoKV.git
cd GoKV
