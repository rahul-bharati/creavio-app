# ADR 002: Application Architecture

**Date:** 2025-09-20  
**Status:** Accepted _(by self, as solo maintainer at this stage)_

---

## Context

As the project grows, it is essential to establish a clear and maintainable architecture for the application. This architecture will guide development, ensure consistency, and facilitate future enhancements.

## Problem Statement

We must choose an application architecture that supports growth, scalability, and maintainability from day one. The chosen architecture should support scalability, maintainability, and ease of understanding for future contributors.

## Options
1. **Monolithic Architecture**: A single unified codebase where all components are tightly integrated.
   - *Pros*: Simplicity in deployment, easier to manage for small teams.
   - *Cons*: Can become unwieldy as the application grows, harder to scale individual components.
   - *Decision*: Rejected due to scalability concerns and also the need for modularity. Plus I want to personally learn more about microservices.
2. **Microservices Architecture**: The application is divided into small, independent services that communicate over a network.
   - *Pros*: Scalability, flexibility in technology choices, easier to manage complex applications.
   - *Cons*: Increased complexity in deployment and management, requires robust inter-service communication.
   - *Decision*: Accepted. This architecture aligns with the project's goals for scalability and maintainability.
3. **Serverless Architecture**: Utilizing cloud services to run code in response to events without managing servers.
   - *Pros*: Reduced operational overhead, automatic scaling, cost-effective for variable workloads.
   - *Cons*: Vendor lock-in, potential cold start latency, limited control over the environment.
   - *Decision*: Rejected. While serverless has its advantages, the need for more control and flexibility in this project outweighs the benefits.

## Decision

After evaluating the options, we have decided to adopt a **Microservices Architecture** for the application. This decision is based on the need for scalability, maintainability, and the ability to manage complex functionalities effectively.

## Consequences
- **Positive**: Modular design allows independent scaling, deployment, and future replacement of services (e.g., switching to Rust for video streaming).
- **Negative**: Adds operational overhead (service discovery, monitoring, orchestration).
- **Mitigation**: Introduce Kubernetes, service gateway, and observability stack (Prometheus, OpenTelemetry) early in development to avoid drift.


> This architecture will be revisited periodically to ensure it continues to meet the project's needs as it evolves.

## Future Considerations

- Depending upon needs, we may explore hybrid architectures or integrate serverless components for specific use cases.

## References

- Builds on learnings from [ADR 001: Language and Framework Selection](./ADR-001-language-and-framework-selection.md)
- This ADR influences: [ADR 003: Database Strategy](./ADR-003-database-strategy.md)
- [Microservices Architecture](https://martinfowler.com/articles/microservices.html) by Martin Fowler
- [Building Microservices](https://www.oreilly.com/library/view/building-microservices/9781491950340/) by Sam Newman
- [The Twelve-Factor App](https://12factor.net/) - Best practices for building modern web applications

---