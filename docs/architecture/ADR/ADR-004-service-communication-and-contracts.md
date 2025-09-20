# ADR 004: Service Communication and Contracts

**Date:** 2025-09-21  
**Status:** Accepted _(by self, as solo maintainer at this stage)_

---

## Context

Creavio is being built from scratch using a microservices architecture.  
A core requirement is **real-time communication** (chat, audio, video), alongside strong privacy guarantees.  
To achieve this, we must decide early on how services talk to each other and how contracts are defined.  
The wrong choice could create tight coupling, performance bottlenecks, or reliability issues later.

---

## Problem Statement

We need a consistent approach for:
- **Synchronous requests** (service-to-service, client-to-service)
- **Asynchronous events** (decoupled, real-time interactions)
- **Contracts** (clear, versioned, machine-readable interfaces)
- **Reliability** (tracing, retries, idempotency, monitoring)

The goal is to balance **developer productivity** with **scalability and observability**.

---

## Decision

### Communication Patterns

| Channel                  | Technology Choice   | Scope / Usage                              | Contract Definition           | Notes                                   |
|--------------------------|---------------------|--------------------------------------------|------------------------------|-----------------------------------------|
| **External Requests**    | REST over HTTPS     | All traffic entering via API Gateway       | OpenAPI (Swagger)            | Standardized, client-friendly           |
| **Internal Requests**    | gRPC over HTTP/2    | Low-latency service-to-service calls       | Protocol Buffers (.proto)    | Binary, fast, strongly typed            |
| **Asynchronous Events**  | NATS JetStream      | Pub/Sub, event streaming, decoupled workflows | CloudEvents spec + Protobuf   | High throughput, persistent, replayable |

### Supporting Decisions

1. **API Gateway**
    - Single entry point for clients.
    - Responsibilities: auth, rate limiting, routing, logging, metrics.
    - Only REST is exposed externally; gRPC & NATS are internal-only.

2. **Service Contracts**
    - REST APIs: defined in **OpenAPI 3.1**.
    - gRPC APIs: defined in **Protocol Buffers**.
    - Events: adopt **CloudEvents** with Protobuf schema.
    - Contracts live in a dedicated repo for visibility + versioning.

3. **Reliability**
    - Retries with exponential backoff + jitter.
    - Idempotency keys for write operations.
    - Outbox/CDC pattern for reliable event publishing.

4. **Observability**
    - Distributed tracing: **W3C Trace Context (traceparent)** propagated in all calls.
    - Correlation IDs (`x-request-id`) for external requests.
    - Metrics: Prometheus + Grafana.
    - Logs: Loki/ELK stack, centralized.

5. **Versioning**
    - REST: URL-based (e.g., `/v1/...`).
    - gRPC: package namespace versioning.
    - Events: schema registry with versioned Protobuf files.

---

## Consequences

- **Positives**
    - Clear boundaries between synchronous and async flows.
    - Scales horizontally as features (chat, voice, video) grow.
    - Strong contracts reduce integration issues.

- **Negatives**
    - More infra complexity (API Gateway, broker, schema registry).
    - Steeper learning curve (gRPC + NATS for newcomers).

- **Mitigations**
    - Start with core services only; expand gradually.
    - Provide boilerplate templates for new services.
    - Invest in automation (CI/CD, linting for contracts, monitoring).

---

## Future Considerations

- Explore **GraphQL** at gateway layer for frontend flexibility.
- Adopt a **service mesh** (Istio/Linkerd) once services >10 for traffic shaping, mTLS, and fine-grained observability.
- Evaluate **WebRTC** for real-time audio/video transport layer (chat/video services).

---

## References

- Builds on [ADR 001: Language and Framework Selection](./ADR-001-language-and-framework-selection.md)
- Builds on [ADR 002: Application Architecture](./ADR-002-application-architecture.md)
- Influenced by [ADR 003: Database Strategy](./ADR-003-database-strategy.md)
- [Chris Richardson – Microservices Patterns](https://microservices.io/patterns/index.html)
- [Martin Fowler – Microservices](https://martinfowler.com/articles/microservices.html)
- [Google Cloud – API Design Guide](https://cloud.google.com/apis/design)
- [gRPC Docs](https://grpc.io/docs/)
- [NATS JetStream Docs](https://docs.nats.io/jetstream)
- [OpenAPI Specification](https://swagger.io/specification/)
- [CloudEvents Spec](https://cloudevents.io/)
- [Uber Engineering – Moving to gRPC](https://eng.uber.com/grpc/)
- [W3C Trace Context](https://www.w3.org/TR/trace-context/)
- [API Gateway Pattern – microservices.io](https://microservices.io/patterns/apigateway.html)

---

