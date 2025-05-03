# 🚀 BlockVision – Blockchain Explorer

BlockVision is a full-stack blockchain explorer that provides real-time visibility into Ethereum transactions, block data, and network health. Designed for both developers and non-technical users, BlockVision enables transparent access to on-chain data with high performance and a clean UI.

---

## 📌 Features

- 🔎 **View Live Ethereum Blocks** – Track and browse the latest blocks and their metadata.
- 💸 **Search Transactions** – Input transaction hashes to get detailed insights instantly.
- ⛓️ **Network Health Monitoring** – View Ethereum network stats like block time, gas prices, etc.
- ⚡ **Fast Caching** – Redis-based caching for rapid data retrieval.
- 🧠 **Developer-Friendly API** – REST API built in Golang using Gin for easy integration.
- 🌐 **Modern UI** – Built with React.js for a smooth and intuitive user experience.

---

## 🛠️ Tech Stack

| Layer           | Technologies Used                      |
|----------------|------------------------------------------|
| **Frontend**    | React.js, Axios, CSS                    |
| **Backend**     | Golang, Gin, go-ethereum, Redis         |
| **Blockchain**  | Ethereum, Web3 API (Infura or Geth)     |
| **Caching**     | Redis                                   |
| **DevOps**      | Docker, dotenv                          |

---

## 📷 Screenshots

_Add screenshots or a demo GIF here to showcase the UI and functionality._

---

## ⚙️ Installation

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/blockvision.git
cd blockvision

## 🚀 Backend Setup (Golang)

### ✅ Prerequisites
- Go 1.21+
- Redis server (local or cloud)
- Ethereum node access (e.g., [Infura](https://infura.io/), [Alchemy](https://www.alchemy.com/), or local Geth)

### ▶️ Run Backend

```bash
cd backend
go mod tidy
go run main.go
