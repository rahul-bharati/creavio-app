# ADR 005: Orchestration and Development Environment

**Date:** 2025-09-21  
**Status:** Accepted _(by self, as solo maintainer at this stage)_

---

## Context
Creavio follows a microservices-first approach. Running multiple services locally and in production requires orchestration, networking, and consistent developer tooling. Without a defined strategy, development can quickly become fragmented and deployments unreliable.

## Problem Statement
We need a clear orchestration and dev environment strategy that:
- Supports **local development** (fast feedback, easy onboarding).
- Scales to **production** (secure, stable, and cloud-ready).
- Provides consistency in **infrastructure**, **networking**, and **observability**.
- Defines a **standard for environment variables and secrets management**.

---

## Options Considered

### 1. Docker Compose Only
- **Pros:** Simple, easy for local dev, widely known.
- **Cons:** Not representative of production Kubernetes setup, limited orchestration features.

### 2. Kubernetes (minikube, kind, or k3d for local)
- **Pros:** Production-like environment, service discovery, scaling, built-in orchestration.
- **Cons:** Steeper learning curve, more resource-heavy, slower feedback loops if misconfigured.

### 3. Kubernetes with Tilt for Dev
- **Pros:** Same Kubernetes base, plus **live reloads, log streaming, and a dashboard**.
- **Cons:** Adds another tool to learn and maintain, some initial setup overhead.

---

## Decision
We will use **Kubernetes (k3d or minikube) for orchestration** and **Tilt.dev for local development**.
- All services will run inside Kubernetes namespaces.
- Tilt will provide **fast iteration**, **log streaming**, and **service health visibility**.
- Each service will include a **Dockerfile** and Kubernetes manifests (Helm or Kustomize can be added later).
- **Environment Variables & Secrets:**
    - Local: Managed via `.env` files loaded into Tilt (`.env.development`, `.env.test`).
    - Kubernetes Dev/Prod: Use **ConfigMaps** for non-sensitive config and **Secrets** for sensitive values.
    - Avoid committing `.env` files — only commit `.env.example` with placeholders.
    - Future: adopt **SealedSecrets** or **HashiCorp Vault** for secure production-grade secrets management.

---

## Consequences
- **Positive:**
    - Consistency between dev and prod.
    - Tilt improves developer experience (auto-reloads, logs, UI dashboard).
    - Environment variables standardized across services.
- **Negative:**
    - More tools to learn (K8s + Tilt + secret management).
    - Higher resource requirements locally.
- **Mitigations:**
    - Provide Makefile commands and Tilt configs to reduce friction.
    - Document environment variable conventions in `/docs/environment.md`.

---

## Future Considerations
- Add **service mesh (Istio/Linkerd)** once cross-service observability and retries are needed.
- Move to **managed Kubernetes (EKS/GKE/DigitalOcean K8s)** for production after MVP.
- Adopt **Helm or Kustomize** for repeatable deployments.
- Consider **GitOps (ArgoCD/Flux)** for environment automation.
- Introduce **Vault or Doppler** for enterprise-grade secrets handling.

---

## References
- Builds on:
    - [ADR 002: Application Architecture](./ADR-002-application-architecture.md)
    - [ADR 004: Service Communication and Contracts](./ADR-004-service-communication-and-contracts.md)
- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [Tilt.dev](https://tilt.dev/)
- [12-Factor App - Dev/Prod Parity](https://12factor.net/dev-prod-parity)
- [Kubernetes ConfigMaps](https://kubernetes.io/docs/concepts/configuration/configmap/)
- [Kubernetes Secrets](https://kubernetes.io/docs/concepts/configuration/secret/)
- [Bitnami SealedSecrets](https://github.com/bitnami-labs/sealed-secrets)
- [HashiCorp Vault](https://developer.hashicorp.com/vault/docs)

---
