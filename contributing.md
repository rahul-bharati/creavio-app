# Contributing to Creavio

> **Public code, closed contributions (for now).**  
> This repository is published for transparency and learning. External pull requests are not being accepted at this stage while the product and architecture evolve toward a commercial release.

---

## 👀 How you can help today

- **Feedback & ideas:** Open an Issue with context and rationale. Use the templates (ADR / RFC / Enhancement / Bug).
- **Bugs:** If you can reproduce it, open an Issue with steps, expected vs actual, logs if possible.
- **Security:** Please do **not** open a public issue. Email **contact@rahulbharati.dev** with details. No SLA is guaranteed during pre-beta.

> Note: Issues may be triaged or closed without action while the roadmap is in flux.

---

## 🚧 What is not accepted right now

- **Pull Requests from external contributors** (code or docs).
- **New features via PR**.
- **Large refactors or architectural changes via PR**.

Creavio is a privacy-first commercial product. Contribution policies will be revisited after the closed beta.

---

## 🧭 Project governance (how changes happen)

- Architectural decisions are recorded as **ADRs** under `docs/architecture/ADR/`.
- Subsystem designs are proposed as **RFCs** under `docs/architecture/RFC/`.
- The **maintainer** (repo owner) is the sole approver and may mark ADRs Accepted/Finalized or Superseded.
- The engineering roadmap is tracked via **GitHub Projects** and **Milestones**.

If you have a proposal, open an **Issue** of type _Enhancement_ or _RFC suggestion_. 
The maintainer may turn it into a formal ADR/RFC when appropriate.

---

## 📝 Issue guidelines

Please include:
- **Context / Problem statement**
- **Why now / Impact**
- **Proposed approach (optional)**
- **Alternatives considered (optional)**
- **Acceptance criteria** (what “done” looks like)

Use labels: `bug`, `enhancement`, `rfc-suggestion`, `question`, `security`, etc.

---

## 🔒 Security policy (short)

- Report vulnerabilities to **contact@rahulbharati.dev** with steps to reproduce and affected versions/commits.
- Please allow reasonable time for investigation and remediation before any disclosure.
- Do not test on systems you don’t own or have permission to test.

---

## 📜 License & usage

- Source code is licensed under **MIT** (see `LICENSE`) unless stated otherwise in subfolders.
- While the code is public, **commercial rights are reserved by the project owner**; Creavio will be offered as a commercial product. Using the code does not grant trademark rights.

---

## 📦 Repo hygiene (for reference)

- Default branch: `main` (protected).
- All changes flow via PRs created and approved by the maintainer.
- ADRs are immutable once Accepted; changes come via a new ADR that **Supersedes** the old one.

---

## 📣 Contact

- General: contact@rahulbharati.dev (Add "General inquiry" in the subject)
- Security: contact@rahulbharati.dev (Add "Security issue" in the subject)
- Website: https://creavio.app

Thank you for your interest in Creavio! Watching or starring the repo is the best way to follow along until external contributions open up.
