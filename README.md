# Creavio

> Privacy-focused, community-driven platform for **chat, voice, video, and streaming** — built with Go microservices and Next.js.

Sign up for early access at [creavio.app](https://creavio.app).

Author: [Rahul Bharati](https://rahulbharati.dev) | Twitter: [@BharatiRahul](https://twitter.com/BharatiRahul) | GitHub: [@rahul-bharati](https://github.com/rahul-bharati)

---
## Table of Contents
- [Vision](#vision)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Architecture](#architecture)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [1) Clone the repository](#1-clone-the-repository)
  - [2) Check your local Kubernetes is ready](#2-check-your-local-kubernetes-is-ready)
  - [3) Create the dev env files (Kustomize will turn these into a ConfigMap/Secret)](#3-create-the-dev-env-files-kustomize-will-turn-these-into-a-configmapsecret)
  - [4) Start the dev environment with Tilt](#4-start-the-dev-environment-with-tilt)
  - [5) Access the services](#5-access-the-services)
  - [6) Stop the dev environment](#6-stop-the-dev-environment)
  - [Useful Links](#useful-links)

---

## Vision

Most current platforms prioritize gathering data over protecting users.  
Private chats can be read, mined, and even used against you.

Creavio began as a simple creator’s economy tool (like Patreon or Ko-fi),  
but as I refined my ideas, it became more complex — and I eventually  
paused development because I couldn’t settle on the right path.

### Now my vision is clear:

**Creavio will be a community-first, privacy-focused platform for communication and collaboration.**  
Its mission is to empower communities and organizations to fully own their data  
while enjoying modern features — with less intrusion.

## Features
- **Privacy-first**: End-to-end encryption for chats, voice, video, and streams.
- **Community-driven**: Open-source and user feedback-driven development.
- **Multi-modal communication**: Seamless integration of chat, voice, video, and streaming
- **Creator tools**: Monetization options for creators, including subscriptions and donations.
- **Customizable**: Themes, plugins, and integrations to tailor the experience.
- **Cross-platform**: Accessible on web, mobile, and desktop.
- **Scalable**: Built with Go microservices for performance and reliability.
- **Secure**: Regular security audits and updates to protect user data.
- **User-friendly**: Intuitive UI/UX designed for all skill levels.
- **Open-source**: Transparent development process with community contributions.

## Tech Stack
- **Backend**: Go microservices for scalability and performance.
- **Frontend**: Next.js for a modern, responsive user interface.
- **Orchestration**: Docker and Kubernetes for containerization and deployment.
- **Database**: PostgreSQL for reliable data storage.
- **Real-time Communication**: WebRTC and WebSockets for low-latency interactions.
- **Storage**: Object storage (e.g., Cloudflare R2, DigitalOcean Spaces) for media
- **Authentication**: OAuth 2.0 and JWT for secure user authentication using Keycloak or third-party providers like Auth0 or Firebase auth.
- **CI/CD**: GitHub Actions for automated testing and deployment.
- **Monitoring**: Prometheus, Opentelemetry and Grafana for system monitoring and alerting.
- **Version Control**: Git for source code management.

> The tech stack decisions are not final and may evolve as the project progresses and any finalization will be based on ADRs and RFCs.

## Architecture

- **Core Components**: Microservices for modularity, API Gateway for routing, Service Mesh for secure service-to-service calls.
- **Networking & Communication**: Message Broker (NATS/RabbitMQ), WebRTC for real-time, CDN for fast content delivery.
- **Deployment & Ops**: Load Balancer for HA, CI/CD pipelines, monitoring & logging for observability.


See [ARCHITECTURE.md](./docs/architecture/README.md) for more details.

## Getting Started

Run the full stack locally using Docker Desktop’s Kubernetes and Tilt.

### Prerequisites

- **Docker Desktop** (with Kubernetes enabled)  
  - Install: [Docker Desktop](https://www.docker.com/products/docker-desktop)  
  - Enable Kubernetes: *Docker Desktop → Settings → Kubernetes → Enable Kubernetes → Apply & Restart*  
    Docs: https://docs.docker.com/desktop/kubernetes/


- **kubectl** (Kubernetes CLI)  
  -Install: [Kubernetes CLI](https://kubernetes.io/docs/tasks/tools/)


- **Tilt** (fast local dev for Kubernetes)  
  Install: [Tilt](https://docs.tilt.dev/install.html)


- **Git**  
  Install: https://git-scm.com/downloads

> **Windows:** Use **WSL2** (Ubuntu recommended) with Docker Desktop’s WSL integration: [WSL](https://learn.microsoft.com/en-us/windows/wsl/install) + [Docker WSL2](https://learn.microsoft.com/windows/wsl/)

---

### 1) Clone the repository

```bash
    git clone https://github.com/rahul-bharati/creavio-app.git
    cd creavio-app
```

### 2) Check your local Kubernetes is ready

```bash
    kubectl version --client
    kubectl get nodes
```
You should see your local cluster node in Ready state (Docker Desktop’s built-in cluster).

### 3) Create the dev env files (Kustomize will turn these into a ConfigMap/Secret)

Create `infra/k8s/dev/config.env` (non-secret config):

```env
    USER_SERVICE_URL=http://creavio-user-service-dev:8001
    NOTIFICATION_SERVICE_URL=http://creavio-notification-service-dev:8002
    LOG_LEVEL=debug
```

Create `infra/k8s/dev/secrets.env` (secret config - **do not commit to git and mappings will be added in secrets.env.example**):

```env
    JWT_SECRET=change-me
    DB_PASSWORD=local-dev-password
```
**Note:** These files are watched by Tilt; changes will roll out automatically.

### 4) Start the dev environment with Tilt

```bash
    tilt up
```
Tilt will
- Build the Docker images,
- Apply the Kubernetes manifests (via Kustomize) into the creavio-dev namespace,
- Port-forward services so you can hit them from your laptop.

Open the Tilt UI (it prints a URL) and wait for all resources to turn green.

### 5) Access the services

- Frontend: [http://localhost:3000](http://localhost:3000)
- API Gateway: [http://localhost:8000](http://localhost:8000)
- User Service: [http://localhost:8001](http://localhost:8001)
- Notification Service: [http://localhost:8002](http://localhost:8002)

(Port-forwards are defined in `Tiltfile`)

### 6) Stop the dev environment

```bash
    tilt down
```
This removes the Kubernetes resources that Tilt created.

### Useful Links
- [Tilt Docs](https://docs.tilt.dev/)
- [Kubernetes Docs](https://kubernetes.io/docs/home/)
- [Kustomize Docs](https://kustomize.io/)
- [Docker Desktop Docs](https://docs.docker.com/desktop/)

---