# ADR 003: Database Strategy

**Date:** 2025-09-20  
**Status:** Accepted _(by self, as solo maintainer at this stage)_

---

## Context

We need to decide on which database to use and how to manage it for our project. As the project is built using microservices architecture, we need a database solution that is scalable, reliable, and easy to manage.

## Problem Statement

We must select a database technology and structuring strategy that ensures scalability, reliability, and clear data ownership for microservices.

## Database Options Considered

1. **PostgreSQL**
    - Pros: Open-source, robust, supports complex queries, strong community support.
    - Cons: Requires more setup and maintenance compared to managed solutions.
2. **MySQL**
    - Pros: Widely used, good performance for read-heavy workloads, strong community support.
    - Cons: Less feature-rich than PostgreSQL, potential licensing issues with certain versions.
3. **MongoDB**
    - Pros: Flexible schema, good for unstructured data, easy to scale horizontally.
    - Cons: Less suitable for complex transactions, potential consistency issues.

## Database Structure Strategy

1. **Database per Service**: Each microservice will have its own dedicated database instance.
    - Pros: Clear separation of data, easier to manage service-specific data models.
    - Cons: Increased complexity in managing multiple databases, potential for data duplication.
2. **Shared Database**: All microservices share a single database instance.
    - Pros: Easier to manage a single database, reduced overhead.
    - Cons: Risk of tight coupling between services, harder to scale individual services.
3. **Shared Database with Separate Schemas**: A single database instance with separate schemas for each microservice.
    - Pros: Balance between separation and manageability, easier to enforce data boundaries.
    - Cons: Still some risk of coupling, more complex than a single database.
4. **Hybrid Approach**: Critical services have dedicated databases, while less critical services share a database.
    - Pros: Flexibility to optimize for different service needs, better resource utilization.
    - Cons: Increased complexity in managing different strategies.

## Decision

After evaluating the options, we have decided to use **PostgreSQL** as our database solution due to its robustness, feature set, and strong community support.
MongoDB was rejected because strong transactional consistency is critical for Creavio (e.g., message delivery, payment flows).
We will implement per-service databases to ensure clear separation of data and easier management of service-specific data models. This approach aligns well with our microservices architecture and allows for better scalability and maintainability.

## Consequences

## Consequences
- **Positive**: Clear data boundaries per service; enables independent scaling and schema evolution; aligns with microservices autonomy.
- **Negative**: Higher operational overhead with many Postgres instances; potential for data duplication and cross-service queries.
- **Mitigation**:
  - Use IaC + migration tooling per service.
  - Apply asynchronous patterns (events, CDC) instead of cross-service joins.
  - Regularly audit duplication vs denormalization trade-offs.

## Future Considerations

- As the project grows, we may need to revisit our database strategy to ensure it continues to meet our needs.
- We will monitor the performance and scalability of our chosen solution and be open to adopting new technologies or strategies as necessary.

## References

- Builds on [ADR-002 Application Architecture](./ADR-002-application-architecture.md)
- This ADR influences ADR-004 Service Communication Strategy
- [PostgreSQL Official Website](https://www.postgresql.org/)
- [Microservices Database Patterns](https://microservices.io/patterns/data/database-per-service.html)
- [Database Scalability Strategies](https://www.digitalocean.com/community/tutorials/database-scalability-strategies)
- [Choosing the Right Database for Microservices](https://www.oreilly.com/library/view/designing-microservices/9781492034018/ch04.html)