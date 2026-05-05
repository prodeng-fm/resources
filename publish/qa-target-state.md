# Quality Assurance — Target State Document

This target state document describes how quality is built into the software lifecycle, across every class of software in the org. It covers the design principles that govern it, the controls at each phase in the lifecycle that assures it and the target architecture that realises it.

## 1 Quality across the lifecycle

Quality emerges from embedding controls into the systems that produce, run, and govern software — not from people remembering to follow process. The **control catalogue** defines what must be enforced; the **platform** is where enforcement happens.

The catalogue is pushed into platforms via policy-as-code or platform-native configuration — governance by design, not process compliance on top. The platform varies by class — a CI/CD release gate, a Palantir configuration constraint, a vendor SLA monitor for SaaS — but the principle of pushing controls into it doesn't change.

Today's QA standard sits mainly in Validate, with controls at the release gate. The scope of this target state is to extend the catalogue across the lifecycle, so controls run from intent to operation, per class of software. The four phases:

- **Define** — controls on intent: testable requirements, sign-off, traceability before code begins
- **Build** — controls on the artifact: shared validation (lint, tests, static analysis), signed immutable artefact
- **Validate** — controls on quality: functional, NFR, end-to-end validation, progressive promotion against production-like reality
- **Operate** — controls on operation: governed promotion, runtime guardrails, observability defaults, RCA feedback

*[Diagram 1 — Phases and the architecture beneath them. Top row: four phase boxes left-to-right — **Define** → **Build** → **Validate** → **Operate**. Forward arrows between adjacent boxes (code progression: intent → changed code → validated artefact → running system). One return arrow from Operate back to Define (operational feedback into intent). Beneath the phases: a **Platform** band — where controls become real, varying by class of software. Beneath the platform: a **Control catalogue** band — what must be true, universal across classes. Upward arrows from the catalogue, through the platform, into each phase box, show controls landing into enforcement points at every phase.]*

These phases are universal — every piece of software that reaches production passes through them. The controls that fill each phase vary by class; the table below shows what runs where.

| Software class | Define | Build | Validate | Operate |
|---|---|---|---|---|
| **Custom-build** | Requirement quality (testable acceptance criteria); requirement traceability | Static analysis and tests; signed immutable artefact | Progressive validation gates (functional, NFR, security); end-to-end test gate | Governed promotion (pipeline-enforced release readiness); observability defaults |
| **Low-code platforms** (Palantir, Power Platform) | Platform design validation (schemas, flows, connectors); requirement traceability | App and flow validation; versioned ALM artefacts | Functional/integration/security gates; DLP/access policy enforcement | Platform monitoring; configuration drift detection |
| **ML / analytical workloads** | Locked success criteria and fairness scope; requirement traceability | Code and data lineage; signed model artefacts | Drift/fairness/bias checks; shadow deployment gate | Governed model-registry promotion; drift threshold alerts |
| **COTS** (e.g. Confluence, Guidewire, Outlook on-prem) | Baseline config locked; integration contracts approved | Config-as-code with version control; integration code follows custom-build controls | Upgrade impact gate; integration boundary tests | Patch governance; integration health monitoring |
| **SaaS** (e.g. Stripe, Salesforce) | Tenant config baseline locked; data flow contracts approved | Tenant configuration under change control; integration code follows custom-build controls | Integration tests; configuration drift checks | Vendor change tracking; SLA monitoring |

## 2 Target state — SDLC custom-build

The rest of this document goes deep on **SDLC custom-build**. The design principles come first — the requirements the architecture has to meet — followed by the target architecture that realises them: components, integration, and controls across the lifecycle.

### Design Principles

Nine principles ensure quality across the lifecycle:

**1. The pipeline is the only route to production and the single authority on releasability.** All production changes — application, infrastructure, security; human-authored or AI-authored — pass through automated gates. Releasability is *computed* from policies embedded in code rather than decided ad-hoc; the pipeline produces the deterministic evidence on which any required human sign-off is based. The pipeline does not replace business or regulatory sign-off where it remains necessary — it makes that sign-off evidence-based, repeatable, and deterministic.

**2. Fast feedback through shared control implementation.** Code now changes faster, in more places, and not just by humans. Validation logic — compile, lint, static analysis, unit and integration tests — is defined once. Agents, engineers, and the pipeline execute the same controls. This is fundamental for agents. A coding agent or AI-ops loop iterates only as fast as the feedback it gets where it's working; without fast signal at the point of change, the agent becomes inefficient.

**3. The pipeline is deterministic.** The pipeline re-executes all defined controls and adds what cannot run on a workstation: end-to-end tests, non-functional acceptance, security scanning, environment-level compliance. Evidence is a byproduct of delivery, not assembled by the team after the fact.

**4. Validation is staged against progressively production-like reality.** A release candidate moves through environments that progressively approximate production: production-like systems, production-like data, production-like end-to-end flows. Each stage closes a category of risk earlier stages cannot. Anonymised production data is the basis for pre-production validation; synthetic data alone misses real-world shapes. UAT remains a structural gate where business accountability requires human judgment.

**5. Pipeline velocity is a security control.** Time-to-exploitation is shrinking, and many vulnerabilities (a Spring upgrade, a transitive CVE) cannot be patched in isolation; they require full functional and non-functional re-validation, so a lane that bypasses normal QA is the wrong instrument. The ambition is a pipeline fast enough that break-glass procedures are rare, audited exceptions rather than routine practice. Security velocity follows from pipeline velocity, not from a separate lane.

**6. Every artifact is traceable to the intent that produced it.** From any production artifact, the chain is verifiable: the controls that validated it, the code it contains, the test that validates it, and the requirements those specifications formalise. In an agent-driven world, this chain is the governance mechanism.

**7. Production outcomes feed back into the quality system itself.** Every defect that reaches production, every UAT regression, every gate slip is evidence of an incomplete upstream control — a missing specification, a weak test, a miscalibrated threshold. The architecture channels this evidence back to the control it exposes, so the quality system gets sharper from real use.

**8. Deterministic testing is automated; human effort moves to where judgment is required.** If a test is repeatable and its expected outcome is known, it belongs in the pipeline. Manual execution of deterministic tests is low-value work. The human QA role shifts to exploratory testing — investigating how the product behaves in scenarios that cannot be pre-scripted, using deep product knowledge. Exploration requires judgment and context that automation cannot provide.

**9. Universal coverage, proportionate depth.** Every application is in scope. Tier determines the depth of controls that apply, not whether they apply. No application opts out of proportionate quality expectations.

### Target Architecture

The architecture below realises the principles above for SDLC custom-build, organised across the four lifecycle phases from §1.

*[Diagram 2 — Phases with internals. Same four boxes as Diagram 1, opened to show contents.
- **Define box:** the work management system (e.g. ADO Boards) holding executable specifications, requirement-to-change traceability links, and a sign-off gate before downstream work begins.
- **Build box:** three internal sub-boxes — **workstation**, **agent sandbox**, **operational agents** — each containing a **harness** running shared controls at the point of authorship.
- **Validate box:** the **pipeline** as a sequence of stages — CI (compile, lint, unit and integration tests, static analysis) → functional → non-functional → security → artefact signing — emitting an immutable signed release candidate. The candidate is then promoted through progressively production-like **test environments** (functional, performance, pre-production), provisioned by **test data management** (anonymised production extracts, synthetic generation), with smoke validation at every deployment. **Exploratory testing** runs at tier-applicable stages.
- **Operate box:** infrastructure containers (compute, storage, database, networking) with **policy enforcement points** overlaid (drift detection, identity, encryption, network segmentation) and observability collectors emitting telemetry.

Underneath all four boxes, two cross-cutting bands:
- **Control catalogue** (band immediately below): central library of control definitions, with arrows running upward into each phase to show controls landing into enforcement points.
- **Enterprise knowledge** (band at the bottom): measurement chain collecting evidence — specifications and traceability from Define, build/coverage metrics from Build, validation results from Validate, runtime telemetry and incidents from Operate. A return-path arrow runs from Enterprise knowledge back up into Define (intent calibration) and into the Control catalogue (gate threshold refinement).]*

The architecture realises each phase through specific components, with two cross-cutting layers spanning all four. The platform team builds and maintains each; application teams consume them through designed interfaces.

#### Core Components

| Phase | Components |
|---|---|
| **Define.** Where intent is captured and approved before code begins. | • **Work management** — sign-off gate; every change links to an approved work item; traceability anchor for the rest of the architecture.<br>• **Executable specifications** — machine-validatable expression of intent, versioned alongside code so the same artefact drives implementation and validation. |
| **Build.** Where code is changed by humans or autonomous agents. | • **Harness** — runs the same controls as the pipeline at the point of authorship (workstation, agent sandbox, operational agents); parity guaranteed by shared definitions in the repository. |
| **Validate.** Where changed code is checked against intent before promotion. | • **Pipeline** — the single authority on releasability; re-executes harness controls and extends them in stages (CI → functional → NFR → security → artefact signing); produces immutable signed release candidates; smoke validation at every deployment.<br>• **Test environments** — production-like environments provisioned via the same IaC discipline as production; parity asserted before each stage; cross-system integration enabled where business processes span systems.<br>• **Test data management** — anonymised production extracts, synthetic generation, and versioning as a platform service; privacy and regulatory compliance enforced in the provisioning logic.<br>• **Exploratory testing** — human investigation of pre-script-resistant scenarios (edge cases, workflow combinations); required for tier-applicable applications. |
| **Operate.** Where validated code is deployed and operated. | • **Infrastructure platform** — compute, storage, database, networking via governed templates; manifests turn declarations into compliant infrastructure; runtime guardrails (drift detection, identity, encryption, network segmentation) enforced continuously via policy-as-code.<br>• **Observability** — structured logging, standard metrics, deployment visibility, monitoring against SLOs; applied by default through platform templates. |
| **Cross-cutting.** Two layers span all four phases. | • **Control catalogue** — central library separating *what's enforced* from *how it's enforced*; definitions authored per domain (QA, CI/CD, observability, infrastructure) platform-agnostically; implementations vary per platform (policy-as-code, platform-native configuration, or team-demonstrated evidence).<br>• **Enterprise knowledge** — the artefact chain across phases (specs, code, validation evidence, telemetry); feeds forward (specs become validation targets) and back (production evidence shapes specs and gate thresholds). |

#### How the Components Connect

Five design threads run through the components. Each thread crosses multiple components; the components cohere because they collectively realise all five threads.

**Thread 1 — Requirements → Executable Specifications → Evidence.** Intent is captured as **executable specifications** in **work management** with sign-off, drafted and validated against templates in the **Harness**, re-checked against implementation by the **Pipeline** as part of promotion, governed by definitions in the **Control catalogue**, and evidenced through traceability chains in **Enterprise knowledge**.

**Thread 2 — Application Code → Pipeline → Signed artifact.** Code can be authored in multiple places — workstations, agent sandboxes, operational agents — with the **Harness** running shared controls at the point of authorship. The same controls re-execute and extend in the **Pipeline** (specs, dependency integrity, code-level security), producing an immutable signed release candidate; gate definitions are sourced from the **Control catalogue**, and build evidence flows into **Enterprise knowledge**.

**Thread 3 — Infrastructure → Policy → Operate.** Workload manifests are authored alongside application code in the **Harness**, validated against schema and policy by the **Pipeline** (policy-as-code), governed by infrastructure controls maintained in the **Control catalogue**, provisioned as compliant infrastructure by the **Infrastructure platform** (production runtime and test environments alike), with runtime guardrails enforced continuously and reported via **Enterprise knowledge**.

**Thread 4 — Release candidate → Progressive validation → Production confidence.** A signed release candidate from the **Pipeline** is promoted through successive **test environments** provisioned with progressively production-like data by **test data management**. Functional, non-functional, and end-to-end business process validation run at each stage, with **exploratory testing** and UAT providing human-judgment sign-off for tier-applicable workflows. Validation results flow into **Enterprise knowledge** as evidence of pre-production readiness; gate thresholds are sourced from the **Control catalogue**.

**Thread 5 — Production evidence → Sharper specifications and controls.** Production telemetry and incidents flow from **Observability** into **Enterprise knowledge**, where they aggregate as evidence of where upstream control was incomplete. The evidence sharpens gate thresholds in the **Control catalogue** and intent in **executable specifications** held in **work management**; the **Harness** then enforces the sharper definitions at the point of authorship on the next change.

#### Controls across the lifecycle

Two actors run through every phase: the **platform team** builds the control plane (golden paths, shared control definitions, pipeline gates, policy-as-code); **application teams** work within it (specify intent, write code, declare infrastructure needs, respond to gate feedback). The pipeline is the interface where both meet. The controls listed below are **candidates** — the full catalogue is delivered as a separate program task.

##### Define

Intent must be verifiable before work begins. Specifications captured in **work management** as **executable specifications** link to approved work items; templates and validation tooling come from the platform team, and sign-off is a gate that blocks downstream work via the traceability chain. Workload manifests declare app needs (compute, storage, networking, dependencies) against governed infrastructure templates. Prior incident patterns and quality trends shape where specification effort is invested.

*Directional control areas at the gate → Build:*

| Control area | What it covers |
|---|---|
| Specification quality | Executable specs are well-formed, testable, and cover functional, non-functional, and monitoring requirements |
| Accountable sign-off | Where human sign-off is required, the accountable role is defined per application by its governance and captured in the requirement artefact. Sign-off is GitOps-style — approval given against pipeline-produced evidence, not as a separate review activity. Absent sign-off blocks downstream work via the traceability chain |
| Requirement traceability | Every change traces to an approved requirement; no orphan code enters the pipeline |
| Infrastructure declaration | Workload manifests (compute, storage, networking, dependencies) match schema and policy constraints |
| Monitoring intent declared | Observability requirements stated at design — monitoring is planned, not retrofitted |

##### Build

The same controls run as code is produced and in the pipeline. The **harness** runs compile, lint, static analysis, and tests locally for fast feedback; the **pipeline** re-runs them and produces an immutable signed artefact when they pass. Coverage is measured as percentage of specs with passing implementations, not lines of code; results flow as structured data into **enterprise knowledge**.

*Directional control areas at the gate → Validate:*

| Control area | What it covers |
|---|---|
| Code quality (static analysis) | Compile, lint, static analysis, style — defined once in repo, run locally and in pipeline with identical code |
| Test coverage against specifications | Unit and integration tests link to specs; coverage measured as percentage of specs with passing implementations, not lines of code |
| Dependency integrity | Vulnerability scanning, license compliance, supply-chain provenance on all third-party components |
| Infrastructure-as-code validation | Policy-as-code controls on platform-team and app-team infrastructure code; same CI/CD discipline for infra as for application code |
| Artifact integrity | Immutable signed release candidate produced; provenance metadata recorded and validated |

##### Validate

The artifact is immutable across promotions; only configuration and test data context change. Validation runs through progressive stages:

1. **Functional validation** — dev-accessible environment with real or stubbed downstream connections; validates behaviour against executable specifications.
2. **Non-functional validation (performance, resilience)** — production-scale environment with representative data volumes; validates load, resilience, and performance against declared NFRs.
3. **Pre-production validation** — environment with anonymised production data, consistent across applications for end-to-end flows; UAT signs off tier-applicable workflows where business or regulatory accountability requires human judgment.

The **pipeline** runs these stages, gating promotion at each. **Test environments** and **test data management** provision them; **exploratory testing** investigates pre-script-resistant scenarios. Quality metrics — coverage, automation ratio, end-to-end pass rate, defect leakage — flow into **enterprise knowledge** as trend indicators, separate from gate thresholds.

*Directional control areas at the gate → Operate:*

| Control area | What it covers |
|---|---|
| Functional validation (QA framework) | Executable specifications run deterministically against the release candidate; results link back to the specs |
| Non-functional validation (QA framework) | Performance, security, resilience, and accessibility acceptance tests against declared NFRs |
| End-to-end business process validation (QA framework) | Cross-system workflows validated in pre-production for tier-applicable applications; UAT sign-off on scenarios where business accountability requires it |
| Test environment fidelity | Each validation stage runs in a production-like environment provisioned via the same IaC discipline as production; configuration parity is asserted before validation runs |
| Test data integrity | Anonymised production data provisioned to pre-production with referential integrity preserved; privacy and regulatory compliance (PII, data residency, retention) enforced in masking and provisioning logic |
| Environment compliance | Integration testing, environment parity checks, and policy-as-code validation through progressive promotion |
| Smoke validation | Light, fast post-deployment validation that the deployment landed and core paths are reachable; same shape across environments; deployment without smoke green stops further validation |
| Regression discipline | Mandatory automated regression coverage executed before every release |
| Exploratory testing | Human, judgment-driven investigation of scenarios that cannot be pre-scripted — edge cases, workflow combinations, real-world usage patterns |

##### Operate

Validated artifacts are deployed through governed promotion. The **infrastructure platform** enforces runtime guardrails (drift detection, identity, encryption, network segmentation) continuously; **observability** is applied by default through platform templates. Material defects trigger full RCA via the traceability chain — *was a spec, control, or gate missing?* — while lower-severity issues aggregate for pattern detection. DORA metrics and RCA findings flow into **enterprise knowledge** as the return path closing the loop to specification investment and gate calibration.

*Directional control areas in operation:*

| Control area | What it covers |
|---|---|
| Governed promotion | Only pipelines promote artifacts to production; all upstream gates green; promotion recorded with provenance |
| Runtime guardrails | Drift detection, identity enforcement, encryption validation, network segmentation — enforced continuously via policy-as-code |
| Observability by default | Structured logging, standard metrics, deployment visibility — platform templates; unobservable services cannot be considered production-ready |
| Incident detection and response | Monitoring and alerting against SLOs; break-glass procedures are governed, time-limited, audited, and rare |
| RCA feedback loop | Material defects trigger full RCA — *what was missing — a spec, a control, or a gate?* Output updates the relevant upstream control; closure verified via the traceability chain. Lower-severity issues feed signal aggregation; pattern-based escalation triggers RCA when a category warrants it |
| Delivery health telemetry | DORA metrics (deployment frequency, lead time for changes, change failure rate, MTTR); change failure rate correlates deployment and incident data via the data-and-metrics workstream |
