
# Project README

## Table of Contents

1. [Introduction](#introduction)
2. [Prerequisites](#prerequisites)
3. [Setting Up the Environment](#setting-up-the-environment)
4. [Building and Running Locally](#building-and-running-locally)
5. [Using Docker](#using-docker)
6. [Makefile Commands](#makefile-commands)

---

## Introduction

This project is a Go-based service that implements a server-client architecture using TCP communication. The server provides a "Proof of Work" (PoW) challenge to the client, which must solve the challenge and receive a motivational quote as a reward.

---

## Prerequisites

- Go 1.23.2+
- Docker and Docker Compose
- `make` (for running Makefile commands)

---

## Setting Up the Environment

### `.env` File

Create a `.env` file in the root directory. The `.env` file is used to define environment variables required by the server and client.

Example `.env` file:
```
# Server Configuration
TCP_SERVER_PORT=8080
POW_PREFIX=00
POW_DIFFICULTY=4

# Logging
LOG_LEVEL=info
LOG_DEST=stdout
LOG_ADD_SOURCE=true

# Metrics
METRICS_ADDRESS=:9090
METRICS_NAMESPACE=words_of_wisdom
METRICS_SUBSYSTEM=server
```

Ensure you adjust the values as needed for your environment.

---

## Building and Running Locally

### Install Dependencies

Run the following command to install project dependencies:
```sh
make deps
```

### Build

To build the project:
```sh
make build
```

This will create an executable file at `./bin/service`.

### Run the Service Locally

To start the service locally:
```sh
make run
```

---

## Using Docker

### Build and Start Containers

Build and start the Docker containers using Docker Compose (also rebuilds the images if you have them):
```sh
make dc-reb
```

This command will:
1. Stop any running containers.
2. Build new Docker images.
3. Start the containers in detached mode.

### Restart Containers

To restart existing containers without rebuilding the images (also will restart them if they already run):
```sh
make dc-reup
```

### Stopping Containers

To stop and remove all containers:
```sh
docker-compose down
```

---

## Makefile Commands

Here’s a summary of available `make` commands:

| Command              | Description                                                 |
|----------------------|-------------------------------------------------------------|
| `make deps`          | Installs project dependencies.                              |
| `make lint`          | Runs `golangci-lint` for code analysis.                     |
| `make test`          | Runs all unit tests.                                        |
| `make int-test`      | Runs integration tests.                                     |
| `make build`         | Builds the project binary.                                  |
| `make run`           | Runs the service locally.                                   |
| `make tools`         | Generates required tools for the project.                   |
| `make codegen`       | Generates code (e.g., mocks, OpenAPI client).               |
| `make dc-reup`       | Restarts Docker Compose containers.                         |
| `make dc-reb`        | Rebuilds and restarts Docker Compose containers.            |

---

## Notes

1. Ensure you have the `.env` file in the project root before running the service locally or in Docker.
2. For Docker Compose, the `.env` file will be mounted into the container and used during runtime.

## Why This Specific PoW Algorithm?
This PoW algorithm was chosen for its simplicity and practicality:

Hash-Based: It relies on SHA-256, a widely-used and secure hashing function. This makes the algorithm easy to understand, implement, and verify.

Dynamic Difficulty: The difficulty level is represented by the number of leading zeros required in the hash, making it straightforward to adjust based on server needs or workload.

Efficient Validation: The server only checks the hash against the required difficulty, which is computationally inexpensive compared to generating a hash.

Compatibility: It works well with low-resource systems and is adaptable for scenarios where clients may have varied computational capabilities.

Educational Value: It's a simple yet effective demonstration of PoW concepts used in more complex systems like Bitcoin, without adding unnecessary complexity.
