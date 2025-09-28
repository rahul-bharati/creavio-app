# ADR 006: Authentication and Identity

**Date:** 2025-09-28  
**Status:** Accepted _(by self, as solo maintainer at this stage)_

---

## Context

Creavio requires a robust, secure, and flexible authentication and identity management system. The solution must:

- Support user registration, login, password management and MFA.
- Allow integration with external identity providers (Google, Github, etc.).
- Be self-hostable to align with privacy-first and open-source principles.
- Scale with the platform while being lightweight in development and production.
- Provide a good developer experience with clear documentation and SDKs.
- Support role-based access control (RBAC) for future admin and moderation features.

Two major open-source identity solutions are considered: **Keycloak** and **FusionAuth**.
Additionally, cloud-hosted identity providers (e.g., Auth0, Firebase, Supabase Auth) and custom in-house implementations were reviewed.

---

## Problem Statement

We need to decide which identity provider to adopt as the foundation of Creavio’s authentication system.
The choice must balance **security, flexibility, developer experience, and resource usage**.

---

## Options Considered

### 1. Keycloak
- **Pros:**
  - Mature and widely adopted in the industry.
  - Supports a wide range of protocols (OAuth2, OpenID Connect, SAML).
  - Built-in support for social logins and MFA.
  - Strong community and Red Hat Backing.
  - Extensible with custom providers and themes.
- **Cons:**
  - Heavier resource usage, which may be overkill for initial stages.
  - More complex to set up and manage.
  - Heavy JVM-based footprint, resource intensive.
  - Database coupling (PostgreSQL, MySQL, etc.) adds operational overhead.

### 2. FusionAuth
- **Pros:**
  - Lightweight and easy to set up.
  - Modern UI and developer-friendly documentation.
  - Supports OAuth2, OpenID Connect, SAML, and social logins.
  - Built-in support for MFA and passwordless login.
  - Self-hostable with a simple Docker setup.
  - More flexible database options (PostgreSQL, MySQL, etc.) with less overhead
  - Flexible tenent model for multi-tenant applications.
- **Cons:**
  - Smaller community compared to Keycloak.
  - Less enterprise adoption, which may impact long-term support.
  - Some advanced features require a paid license.

### 3. Cloud-Hosted Providers (Auth0, Firebase Auth, Supabase Auth)
- **Pros:**
  - Quick to set up with minimal configuration.
  - Rich ecosystem integrations (social logins, enterprise SSO).
  - Scalable and managed by the provider.
  - Good developer experience with SDKs and documentation.
- **Cons:**
  - Not self-hostable, conflicting with privacy-first goals.
  - Potentially high costs as the user base grows.
  - Vendor lock-in and less control over data.
  - Limited customization compared to self-hosted solutions.

### 4. Custom Authentication Implementation
- **Pros:**
  - Full control over features and data.
  - Tailored specifically to Creavio’s needs.
  - No third-party dependencies.
- **Cons:**
  - Significant development and maintenance effort.
  - Higher risk of security vulnerabilities.
  - Lacks the maturity and features of established solutions.
  - Requires ongoing updates to keep up with security best practices.

---

## Decision

We will adopt **FusionAuth** as the primary identity provider for Creavio.

### Rationale:
- **Balance of Features and Resource Usage**: FusionAuth provides a good balance between features and resource consumption, making it suitable for both development and production environments.
- **Developer Experience**: FusionAuth’s modern UI and documentation will facilitate easier integration and onboarding for developers.
- **Self-Hosting**: Aligns with Creavio’s privacy-first and open-source principles.
- **Flexibility**: Supports a wide range of authentication protocols and social logins, with built-in MFA.
- **Scalability**: Can grow with the platform while remaining manageable.
- **Future-Proofing**: The flexible tenant model allows for potential multi-tenant applications in the future.

---

## Consequences


- **Positives**
    - Faster integration with services through API-first approach.
    - Lower infra overhead compared to Keycloak.
    - Clear migration path to enterprise features if needed.
- **Negatives**
    - Smaller ecosystem than Keycloak.
    - Potential future lock-in if FusionAuth enterprise features are required.
- **Mitigation**
    - Keep identity logic abstracted behind an internal service layer.
    - Ensure OpenID Connect / OAuth2 compliance so migration to Keycloak or Auth0 remains possible.  

---

## Future Considerations

- Explore integration with **service-to-service auth** (mTLS or OPA sidecars).
- Define a standard for user roles, permissions, and multi-tenant org support.
- Plan for **federated identity providers** (Google, GitHub, SSO) in later stages.
- Revisit if scaling requirements exceed FusionAuth’s operational sweet spot.
- Cloud identity providers may be reconsidered for enterprise clients who require SOC2-compliant hosted identity.

---

## References

- Builds on [ADR-001 Language and Framework Selection](./ADR-001-language-and-framework-selection.md) and [ADR-002 Application Architecture](./ADR-002-application-architecture.md).
- [FusionAuth Documentation](https://fusionauth.io/docs/)
- [Keycloak Documentation](https://www.keycloak.org/documentation)
- [Auth0 Documentation](https://auth0.com/docs)
- [Firebase Authentication](https://firebase.google.com/docs/auth)
- [Supabase Auth](https://supabase.com/docs/guides/auth)
- [OAuth2.1 RFC Draft](https://datatracker.ietf.org/doc/html/draft-ietf-oauth-v2-1-07)
- [OpenID Connect Specification](https://openid.net/specs/openid-connect-core-1_0.html)
- [NIST Digital Identity Guidelines](https://pages.nist.gov/800-63-3/)

---