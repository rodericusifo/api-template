# 🧱 API Template

[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://github.com/yourusername/yourrepo/actions)
[![Docker](https://img.shields.io/badge/docker-ready-blue)](https://www.docker.com/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Made with ❤️](https://img.shields.io/badge/Made%20with-%E2%9D%A4-red)](#)

---

A **containerized starter framework** tailored for developing robust and high-performance backend APIs. Built for **scalability**, **maintainability**, and **team productivity**.

This template includes:

- 🧱 Modular structure for API services
- 🐳 Docker for consistent environments
- 📁 SQL + cache + broker integrations
- 🛠️ Developer-friendly Makefile workflow
- ✅ Production-ready setup with best practices

---

## 📦 Features

- ✅ Rapid setup using `make`
- 🧩 Scalable microservices architecture
- 🐘 PostgreSQL with initialization scripts
- 🚀 Redpanda & LavinMQ support for brokers
- 🔐 Environment isolation via `.env` files
- 🧠 Redis caching support

---

## 🚀 Quick Start

### 🧰 Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- `make` utility

#### Install `make`

- **Windows**
  - Install [Chocolatey](https://chocolatey.org/install)
  - Then: `choco install make`

- **macOS / Linux**
  - Install [Homebrew](https://brew.sh)
  - Then: `brew install make`

---

### 🛠️ Installation

1. **Initialize Database Scripts**

   Edit your initdb scripts in `path/to/initdb/db.sh` by adding required databases:

   ```bash
   CREATE DATABASE service_a_db;
   CREATE DATABASE service_b_db;
   ```

> **Note:** Replace `<database-sql>` with your actual database type, such as `postgres`.

2. **Copy `.env` Files**

   ```bash
   make setup-env ENV=<ENV>
   ```

> Replace `<ENV>` with your desired environment name (e.g. `dev`, `prod`, etc.)

   Fill empty values:

   ```dotenv
   # <ENV>.env file
   FIELD_A="HelloWorld"
   FIELD_B=1
   FIELD_C=true
   ```

> **Note:** If you use LavinMQ as one of your message brokers, set `BROKER_LAVINMQ_MAIN_DEFAULT_PASSWORD` using a **hashed password**.  
> You can generate a hashed password by following these steps:  
>
> 1. Run a temporary LavinMQ container:
>
>    ```bash
>    docker run -d --name lavinmq_temp -p 5672:5672 -p 15672:15672 cloudamqp/lavinmq
>    ```
>
> 2. Generate the **hashed password**:
>
>    ```bash
>    docker exec -it lavinmq_temp lavinmqctl hash_password yourdesiredpassword
>    ```
>
>    Use the result to set `BROKER_LAVINMQ_MAIN_DEFAULT_PASSWORD`.
>
> 3. Remove the temporary container:
>
>    ```bash
>    docker rm -f lavinmq_temp
>    ```

3. **Define Services**

   Edit `api.mk`:

   ```makefile
   APPLICATIONS := auth-service author-service category-service book-service api-gateway
   BROKERS := redpanda lavinmq
   DATABASES_CACHE := redis
   DATABASES_SQL := postgres
   ```

4. **Set Up Git Hooks (Required for Development)**

   ```bash
   make setup-hooks
   ```

   This installs development hooks that enforce project standards:

   - ✅ **Pre-commit**: Ensures only one service is modified per commit
   - ✅ **Commit-msg**: Validates conventional commit format (`type(scope): description`)
     - Supported types: `build`, `chore`, `ci`, `docs`, `feat`, `fix`, `perf`, `refactor`, `revert`, `style`, `test`
     - Example: `feat(auth-service): add login feature`
   - ✅ **Pre-push**: Comprehensive version validation
     - Enforces semantic versioning (`x.y.z` format)
     - Prevents major version updates (except `0.x.x` → `1.0.0`)
     - Limits minor/patch increments to 1
     - Ensures VERSION files are updated from base branch (`prod-v1`)

   All hooks automatically skip validation for merge and revert commits.

> **Important**: All developers must run this command after cloning the repository.

---

## 🔄 GitHub Workflows (CI/CD)

This template includes automated GitHub workflows for streamlined development and deployment:

### 🛡️ **Pull Request Validations** 
*Triggered on: Pull Requests to `prod-v1`*

- ✅ **File validation**: Ensures required `CHANGELOG.md` and `VERSION` files exist for modified services
- ✅ **Version validation**: Validates semantic versioning rules and format
- ✅ **Changelog validation**: Verifies proper changelog format and content
- ✅ **Service isolation**: Enforces one service per pull request policy
- ✅ **Branch compliance**: Ensures changes follow base branch requirements

### 📝 **Auto Create or Update Changelogs**
*Triggered on: Push to `prod-v1` (application changes)*

- 🔄 **Automatic changelog updates**: Updates CHANGELOG.md files when applications are modified
- 📅 **Smart formatting**: Maintains proper changelog structure and formatting
- 🚫 **Loop prevention**: Ignores changelog file changes to prevent infinite loops
- ⚡ **Concurrent safe**: Uses concurrency controls to prevent conflicts

### 🏷️ **Auto Create Tags**
*Triggered on: Workflow completion (Auto Create or Update Changelogs)*

- 🚀 **Automatic tagging**: Creates Git tags based on VERSION file changes
- 📦 **Service versioning**: Tags format: `service-name/vX.Y.Z` (e.g., `auth-service/v1.2.0`)
- 🔄 **Version detection**: Compares current vs base branch versions
- ✨ **Zero-config**: Works automatically after changelog updates

### 🐳 **Push to Docker Hub**
*Triggered on: Workflow completion (Auto Create Tags)*

- 📤 **Automatic deployment**: Builds and pushes service images to Docker Hub after tagging
- 🏷️ **Multi-tag support**: Creates both version-specific and `latest` tags
- 🔐 **Secure authentication**: Uses Docker Hub secrets for registry access
- 📦 **Batch processing**: Handles multiple services in parallel

### 📦 **Push to GitHub Container Registry**
*Triggered on: Workflow completion (Auto Create Tags)*

- 🚀 **GHCR deployment**: Builds and pushes service images to GitHub Container Registry
- 🔒 **Built-in auth**: Uses GitHub tokens for seamless authentication
- 🏷️ **Tag management**: Same tagging strategy as Docker Hub workflow
- 🌐 **Registry flexibility**: Supports both public and private container registries

### 🎯 **Release Coordinator**
*Triggered on: Workflow completion (Auto Create Tags)*

- ⏱️ **Workflow orchestration**: Waits for all Docker build workflows to complete
- 🚀 **Release trigger**: Dispatches release creation once all builds are finished
- 📊 **Status monitoring**: Tracks the completion of both Docker Hub and GHCR pushes
- 🔄 **Reliable sequencing**: Ensures proper order of deployment operations

### 📋 **Auto Create Releases**
*Triggered on: Repository dispatch (docker-builds-completed)*

- 🎉 **Automated releases**: Creates GitHub releases with comprehensive version tables
- 📊 **Version comparison**: Shows old vs new versions for all applications
- 🏷️ **Release tagging**: Creates timestamped release tags
- 📝 **Rich notes**: Includes deployment commands and Docker registry information

### 🎯 **Workflow Chain Example**

The workflows follow this automated sequence:

```bash
1. Developer pushes to prod-v1
   ↓
2. Auto Create or Update Changelogs runs
   ↓
3. Auto Create Tags runs (creates version tags)
   ↓
4. Both Docker Hub & GHCR Push workflows run in parallel
   ↓
5. Release Coordinator waits for both to complete
   ↓
6. Auto Create Releases runs (creates GitHub release)
```

**Pull Request validation:**
```bash
# Happens automatically on PR creation/updates
git checkout -b feature/auth-improvements
# ... make changes to auth-service ...
git push origin feature/auth-improvements
# → Creates PR → Triggers Pull Request Validations workflow
```

---

## ▶️ Usage

### Start the application:

```bash
make start ENV=<ENV>
```

### Stop the application:

```bash
make stop ENV=<ENV>
```

> Make sure you already running `make setup-env ENV=<ENV>` command.
> Replace `<ENV>` with your desired environment name (e.g. `dev`, `prod`, etc.).

---

## 📚 API Documentation

Test your APIs in Postman:

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/10344918-49466c3c-b1e9-4b1c-b6d7-b33940099111?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D10344918-49466c3c-b1e9-4b1c-b6d7-b33940099111%26entityType%3Dcollection%26workspaceId%3D667868fa-663b-45d5-a9ec-252ff52cb9c8)

---

## 🧪 Tech Stack

- **Language**: Go, Python, etc.
- **Database**: PostgreSQL
- **Cache**: Redis
- **Broker**: Redpanda, LavinMQ
- **Containerization**: Docker
- **Workflow**: Makefile

---

## 🤝 Contributing

Contributions are welcome! To contribute:

1. Fork the repo
2. Create your feature branch (`git checkout -b feature/foo`)
3. Commit your changes (`git commit -am 'Add foo'`)
4. Push to the branch (`git push origin feature/foo`)
5. Create a new Pull Request

---

## 📄 License & Notice

This code is licensed under the [MIT License](./LICENSE).

See the [NOTICE](./NOTICE) file for additional legal and attribution information.

> © 2025 Rodericus Ifo Krista. Unauthorized commercial use or misrepresentation of authorship is prohibited.

---

## 💬 Feedback

Got questions or suggestions?

- Open an [issue](https://github.com/rodericusifo/api-template/issues)
- Or email us at: `rodericus1999@gmail.com`

---

> Made with ❤️ by [Rodericus Ifo Krista](https://www.linkedin.com/in/rodericus-ifo-krista)
