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
