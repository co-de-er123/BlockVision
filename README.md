# 🚀 **BlockVision – Ethereum Blockchain Explorer**

**BlockVision** is a real-time full-stack blockchain explorer that visualizes Ethereum transaction details, block data, and network health. Designed for both developers and general users, it bridges the transparency gap in blockchain interactions with speed and usability.

---

## ✨ **Features**

- 🔎 **Real-Time Block Explorer** – Track the latest blocks with metadata.
- 💸 **Transaction Lookup** – Fetch full details of any Ethereum transaction.
- 📊 **Network Health Dashboard** – Monitor block times, gas fees, and peer data.
- ⚡ **Redis Caching** – Boosts performance and response times.
- 🧑‍💻 **RESTful API** – Developer-friendly endpoints to integrate with your apps.
- 🌐 **Modern UI** – Built using React.js with a clean and intuitive interface.

---

## 🧰 **Tech Stack**

| **Layer**      | **Technologies Used**                    |
|----------------|-------------------------------------------|
| **Frontend**   | React.js, Axios, CSS                      |
| **Backend**    | Go (Golang), Gin Framework, Web3 API      |
| **Blockchain** | Ethereum, Infura or Geth                  |
| **Caching**    | Redis                                     |
| **DevOps**     | Docker, .env (dotenv)                     |

---

## 🛠️ **Backend Setup (Golang)**

### 🔧 **Prerequisites**

- Go 1.21+
- Redis server (local/cloud)
- Ethereum node access: [Infura](https://infura.io), [Alchemy](https://www.alchemy.com), or local Geth

### ▶️ **Run Backend**

```bash
cd backend
go mod tidy
go run main.go
