# ADR 001: Language and Framework Selection for Creavio

**Date:** 2025-09-20  
**Status:** Accepted _(by self, as solo maintainer at this stage)_

---

## Context
Creavio is a privacy-focused platform for secure chats, audio, and video streaming. The tech stack must ensure security, scalability, and developer productivity.

## Problem Statement
The stack should support security, performance, and easy maintenance, and be accessible to developers.

## Backend Options
1. **Node.js/Express**
   - Pros: Large ecosystem, real-time support, single language for full stack.
   - Cons: Single-threaded, less performant for CPU-heavy tasks, scalability challenges.
2. **Python/Django**
   - Pros: Rapid development, secure, beginner-friendly.
   - Cons: Slower performance, less suited for real-time apps, GIL limits multithreading.
3. **Go**
   - Pros: High performance, efficient concurrency, easy deployment, strong standard library.
   - Cons: Smaller ecosystem, steeper learning curve, fewer language features.

## Frontend Options
1. **React**
   - Pros: Large community, reusable components, good performance.
   - Cons: Steep learning curve, requires extra libraries, frequent updates.
2. **Flutter**
   - Pros: Single codebase for web/mobile, fast development, strong performance.
   - Cons: Smaller ecosystem, less mature for web, limited third-party libraries.
3. **Next.js**
   - Pros: Built on React, server-side rendering, good performance/SEO, simplifies routing.
   - Cons: Adds complexity, learning curve for SSR, advanced use cases need extra config.

## Decision
- **Backend:** Go for performance, concurrency, and deployment.
- **Frontend:** Next.js for SSR, SEO, and React ecosystem.

## Consequences
- **Positive:** High concurrency, efficient deployment, SEO-friendly frontend.
- **Negative:** Smaller Go ecosystem, steeper learning curve.
- **Mitigation:** Use microservices to adopt other languages if needed.

## Alternatives
- **Backend:** Node.js avoided due to npm package concerns; Python better for scripting/data science.
- **Frontend:** React is base for Next.js; Flutter requires learning a new ecosystem.

## Future Considerations
Mobile tech (Flutter/React Native/Capacitor) will be decided after web MVP stabilizes.

## References
- This ADR influences: [ADR 002: Application Architecture](./ADR-002-application-architecture.md)
- See: [Next.js](https://nextjs.org/) | [Go](https://go.dev/)

---
