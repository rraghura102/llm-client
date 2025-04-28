# LLM Command Line Client

A Proof of Concept (PoC) of a Golang based command line client to chat with the [`llm-server`](https://github.com/rraghura102/llm-server) interactively.

⚠️ NOTE
This code is not production-ready and is intended solely for proof-of-concept (PoC) and demonstration purposes. It lacks production-grade features such as authentication, request limits, error handling, hardening, and full model lifecycle management.

# Prequisites

1) [`llm-server`](https://github.com/rraghura102/llm-server) is running.

2) Clone Repository

```
git clone https://github.com/rraghura102/llm-server.git
```

# MAC ARM Build

1) Build and Run

```
cd llm-client

go mod init llm-client
go mod tidy
go build llm-client
```

# Run 

```
./llm-client
```

![llm-client screenshot](llm-client-screenshot.png)
