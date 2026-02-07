# mini-harbor-cli

A lightweight Go-based command-line tool that explores modern CLI design patterns, including authenticated workflows, command grouping, and real-time audit log streaming.

The project focuses on building a clean, extensible CLI architecture and handling long-running commands that continuously surface system activity in a developer-friendly way.

---

## ✨ Features

- Structured CLI built using Cobra
- Persistent authentication via local configuration
- Repository-scoped audit log streaming
- Polling-based real-time event streaming
- Stateful deduplication to prevent repeated logs
- Clear command hierarchy and flag-based scoping

---

## 📁 Project Structure

```text
mini-harbor-cli/
├── cmd/
│   ├── root.go
│   ├── login.go
│   ├── project.go
│   └── audit.go
├── pkg/
│   ├── auth/
│   ├── config/
│   └── utils/
├── main.go
├── go.mod
└── README.md
```

---

## 🚀 Getting Started

### Prerequisites

- Go 1.20 or higher
- A GitHub Personal Access Token (classic)

---

## 🔧 Build

Build the binary locally:

```bash
go build -o mini
```

---

## 🔐 Authentication

Authenticate once using a GitHub token:

```bash
./mini login --token <your_github_token>
```

The token is stored locally and reused for subsequent authenticated commands.

---

## 📡 Audit Log Streaming

Stream repository activity:

```bash
./mini audit stream --repo owner/name
```

Example:

```bash
./mini audit stream --repo DuttaNeel07/mini-harbor-cli
```

The command runs continuously and prints new audit events as they occur.  
Stop the stream using `Ctrl + C`.

---

## 🧠 How Streaming Works

- GitHub’s repository events API is used as a generic audit-log source
- The first poll establishes a baseline without printing historical events
- Subsequent polls stream only newly observed events
- Event IDs are tracked to ensure exactly-once output
- Polling is used to keep behavior predictable in terminal environments

Example output:

```text
[PushEvent] DuttaNeel07/mini-harbor-cli
[CreateEvent] DuttaNeel07/mini-harbor-cli
```

---

## 📌 Limitations & Future Improvements

This project is intended as a learning and exploration tool.

Possible improvements include:
- API rate-limit handling and backoff
- Pagination support
- Graceful shutdown using signal handling
- Structured or colored output
- Support for alternative backends beyond GitHub

---

## 📄 License

MIT License

---

## 👤 Author

Neel Dutta  
Computer Science student interested in systems, tooling, and open-source development
