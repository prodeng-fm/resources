# Quality Assurance — Target State Document

This document details the QA target state for 2028 — what quality looks like when the platform enforces it, how we get there, and how teams experience the change. It is the detail artifact behind the QA PID. Audience: CTO, engineering directors, workstream leads.

---

## 1 Foundations

### Governing Principle

Quality is produced by the system itself, not by people following process. The pipeline — not a sign-off review board — determines whether software can ship. Control evidence is generated as a byproduct of delivery, not assembled after the fact. This applies equally whether a human or an AI agent produced the change. The actor changes; what must be true does not.

### Lifecycle Across Software Classes

The strategic intent paper frames how software in the org is *planned, built, verified, deployed, operated, and evidenced* — a single end-to-end lifecycle from intent through runtime observability. For the QA workstream this lifecycle has four phases:

- **Define & Design** — intent captured in a form that can be validated; critical functionality identified and signed off
- **Code & Build** — the artifact is produced (source code, configuration, model, integration); controls run during authorship and at submission
- **Validate & Test** — the artifact is proven against intent; functional, non-functional, and integration validation land here
- **Release & Operate** — the validated artifact is promoted, deployed, monitored; operational evidence feeds back to upstream controls

These phases are universal — every piece of software that reaches production passes through them in some form. What fills each phase varies materially by software class.

| Software class | Define & Design | Code & Build | Validate & Test | Release & Operate |
|---|---|---|---|---|
| **Custom-build** | Executable specifications, sign-off, work-item-to-requirement traceability | Code with shared controls (lint, static analysis, unit / integration tests) defined once in repo; immutable signed artifact | Pipeline gates: functional, NFR, security, dependency, environment compliance | Governed promotion; observability defaults; DORA; RCA feedback loop |
| **Low-code platforms** (Palantir, Power Platform) | Platform-native specifications and sign-off; flow / data-model intent captured | Application composed in platform; tests in platform's own framework; ALM artifacts versioned | Platform-native validation gates; promotion through dev / test / prod under platform ALM | Platform monitoring; drift detection; version visibility |
| **ML / analytical workloads** | Decision specifications — what decisions the model makes, success criteria, drift definitions | Training code, feature pipelines, validation datasets under same shared-control discipline as custom-build; model artifacts versioned in registry | Performance against decision specs; drift testing; fairness / bias checks; shadow or A/B validation | Model-registry promotion; production drift monitoring; retraining triggers |
| **COTS** (e.g. Confluence, Guidewire, Outlook on-prem) | Document critical business workflows the COTS supports; configuration intent; integration boundaries | Configuration captured as code where possible; integration code follows custom-build discipline | UAT for critical workflows; version / upgrade impact analysis; integration boundary tests | Vendor patch cadence; integration health monitoring; upgrade validation |
| **SaaS** (e.g. Stripe, Salesforce) | Document critical business workflows the SaaS supports; tenant / feature configuration intent; data flows | Tenant configuration changes under change control; integration code follows custom-build discipline | Integration tests; vendor SLA validation; configuration drift checks | Vendor change tracking; integration health monitoring; SLA monitoring |

Two cross-cutting elements run through every phase regardless of software class:

- **The control catalogue** (governance workstream) is the central register of controls. Each domain workstream authors its controls platform-agnostically; platform / class workstreams pick them up and implement through the enforcement mode appropriate to each class — policy-as-code, platform-native configuration, or team-demonstrated evidence. The currently approved QA standard's controls primarily describe release-time evidence; the target state extends the catalogue with lifecycle controls (Define & Design, Code & Build) per class.
- **Enterprise knowledge** (operational data, incident records, telemetry) feeds back into upstream controls regardless of class. Incidents for an Outlook integration inform the same kind of control sharpening as production incidents for a custom-build service — what was missing, what control should have caught it.

In 2026, criticality is assigned at APM level by governance, and large APMs are tested as critical end-to-end. Where an APM is hybrid — custom code, COTS components, SaaS integrations under one classification — the dominant class is assigned and guidelines describe how to apply the QA standard to sub-components. Per-class guidelines for the existing QA standard are a 2026 deliverable; full per-class target states and catalogue extensions follow from 2026 H2 onwards.

§2 below describes the SDLC target state in depth — its design principles, operating model, and architecture. Per-class target states for COTS and SaaS follow in their own documents, sequenced from 2027 onward.

---

## 2 SDLC Target State

This section describes the target state for **SDLC custom-build software** in concrete terms — the design principles that govern it, the operating model walked phase-by-phase, and the target architecture. Low-code platforms and ML / analytical workloads inherit the same shape with class-specific variants noted where they differ. Per-class architectures for COTS and SaaS — which differ enough that components like the harness and pipeline do not apply in the SDLC sense — are sequenced as 2027 deliverables in their own target states (see §1 *Lifecycle Across Software Classes* for their phase content).

### Design Principles

Eight architectural principles make the SDLC target state work:

**1. The pipeline is the only route to production and the single authority on releasability.** All production changes — application, infrastructure, security; human-authored or AI-authored — pass through automated gates. Releasability is *computed* from policies embedded in code rather than decided ad-hoc; the pipeline produces the deterministic evidence on which any required human sign-off is based, GitOps-style. The pipeline does not replace business or regulatory sign-off where it remains necessary — it makes that sign-off evidence-based rather than opinion-based. Emergency scenarios use a governed break-glass procedure — time-limited, audited, and automatically reverted.

**2. Fast feedback through shared control implementation.** Validation logic — compile, lint, static analysis, unit and integration tests — is defined once in the repository. Agents, engineers, and the pipeline execute the same controls. Local feedback is fast and trustworthy because it runs the same code as the pipeline.

**3. The pipeline is deterministic.** The pipeline re-executes all repo-defined controls and adds what cannot run on a workstation: end-to-end tests, non-functional acceptance, security scanning, environment-level compliance. Local feedback is a speed optimisation, not a gate.

**4. Production outcomes feed back into the quality system itself.** Every defect that reaches production, every UAT regression, every gate slip is evidence of an incomplete upstream control — a missing specification, a weak test, a miscalibrated threshold. The architecture channels this evidence back to the control it exposes, so the quality system gets sharper from real use. Quality is continuous, not stamped at release.

**5. Every artifact is traceable to the intent that produced it.** From any production artifact, the chain is verifiable: the controls that validated it, the code it contains, the executable specifications it satisfies, and the requirements those specifications formalise. In an agent-driven world, this chain is the governance mechanism. Evidence is a byproduct of delivery, not assembled by the team after the fact.

**6. Deterministic testing is automated; human effort moves to where judgment is required.** If a test is repeatable and its expected outcome is known, it belongs in the pipeline. Manual execution of deterministic tests is low-value work. The human QA role shifts to exploratory testing — investigating how the product behaves in scenarios that cannot be pre-scripted, using deep product knowledge. Exploration requires judgment and context that automation cannot provide.

**7. Universal coverage, proportionate depth.** Every application is in scope. Tier determines the depth of controls that apply, not whether they apply. No application opts out of proportionate quality expectations.

**8. Pipeline velocity makes emergency lanes the exception.** The pipeline is fast and automated enough that all releases — security patches, feature changes, urgent fixes — flow through the same gates. Hotfix lanes exist when the normal pipeline is too slow; the ambition is a pipeline fast enough that break-glass procedures are rare, audited exceptions rather than routine practice. Security velocity follows from pipeline velocity, not from a separate lane.

**A note on application architecture and scale.** These principles describe the target state for SDLC custom-build software in its modernised form — decomposed enough that local controls can mirror pipeline controls, fast enough that break-glass is exceptional. Modern application architecture (independently deployable services, modular test suites, infrastructure-as-code) reaches this directly. Legacy monolithic applications — where full-system local replication is infeasible and integration suites run for hours — need a transition path: progressive decomposition, staged or sampled local feedback, or compensating controls in the pipeline until the target is reachable. The principles do not change; the path to them does, and that path lives in the application team's modernisation roadmap supported by tech standards and the platform team.

### Target Operating Model

Two actors run through every phase: the **platform team** builds the control plane (golden paths, shared control definitions, pipeline gates, policy-as-code); **application teams** work within it (specify intent, write code, declare infrastructure needs, respond to gate feedback). The pipeline is the interface where both meet.

#### Define & Design

Intent must be verifiable before work begins.

- **Requirements.** Application teams specify intent as **executable specifications** — artefacts structured enough to drive validation tooling, rich enough for human review, and traceable through delivery. Form follows the system being specified: behavioural acceptance criteria (commonly Given / When / Then), event models or sequence diagrams for workflow and integration systems, structured examples with screenshots and reference flows for richer context, or specialised methodologies like spec-kit — whatever unambiguously expresses the intent and connects to validation. The platform team provides templates and validation tooling for approved formats; the team owns the content. Sign-off is a gate — if absent, downstream work is blocked. Today requirements are often informal and gaps discovered late; in the target state untestable requirements are caught at definition.

- **Code.** Every change starts with a work item linked to an approved requirement. No orphan code.

- **Infrastructure.** Application teams declare what their workload needs — compute, storage, networking, service dependencies — via a manifest. The platform team maintains manifest schema and governed templates that turn declarations into compliant infrastructure.

- **Enterprise knowledge.** Prior incident patterns and quality trends shape where specification effort is invested.

*Directional control areas at the gate → Code & Build:*

| Control area | What it covers |
|---|---|
| Specification quality | Executable specs are well-formed, testable, and cover functional, non-functional, and monitoring requirements |
| Accountable sign-off | Where human sign-off is required, the accountable role is defined per application by its governance and captured in the requirement artefact. Sign-off is GitOps-style — approval given against pipeline-produced evidence, not as a separate review activity. Absent sign-off blocks downstream work via the traceability chain |
| Requirement traceability | Every change traces to an approved requirement; no orphan code enters the pipeline |
| Infrastructure declaration | Workload manifests (compute, storage, networking, dependencies) match schema and policy constraints |
| Monitoring intent declared | Observability requirements stated at design — monitoring is planned, not retrofitted |

#### Code & Build

Shared controls prove their value. The same validation logic runs locally for fast feedback and again in the pipeline for enforcement.

- **Requirements.** Engineers and agents implement against signed-off specifications. Coverage is measured as percentage of specs with passing implementations, not percentage of code. Agents are constrained by specs and controls.

- **Code.** Controls — compile, lint, static analysis, unit and integration tests — are defined once in the repository. The platform team provides golden path pipeline templates and shared control definitions; application teams inherit and extend. Engineers and agents run controls locally; on commit, the pipeline re-executes. If they pass, an immutable signed artifact is produced.

- **Infrastructure.** Workload manifests are validated against schema and policy. Platform-team infrastructure code follows the same CI/CD discipline.

- **Enterprise knowledge.** Measurement is automatic at every control point. Coverage, scan results, compliance flow as structured data.

*Directional control areas at the gate → Validate & Test:*

| Control area | What it covers |
|---|---|
| Code quality (static analysis) | Compile, lint, static analysis, style — defined once in repo, run locally and in pipeline with identical code |
| Test coverage against specifications | Unit and integration tests link to specs; coverage measured as percentage of specs with passing implementations, not lines of code |
| Dependency integrity | Vulnerability scanning, license compliance, supply-chain provenance on all third-party components |
| Infrastructure-as-code validation | Policy-as-code controls on platform-team and app-team infrastructure code; same CI/CD discipline for infra as for application code |
| Artifact integrity | Immutable signed release candidate produced; provenance metadata recorded and validated |

#### Validate & Test

The release candidate is promoted through progressively demanding stages. The artifact does not change between stages — each stage adds confidence. The QA framework's three control areas land primarily here.

- **Requirements.** Specifications execute deterministically. Results are structured data. Every result links to the specification it validates. Regression is mandatory and automated.

- **Code.** Functional acceptance, NFR acceptance (performance, security scanning, load), dependency vulnerability analysis, environment compliance. All green = releasable. Any red = not.

- **Infrastructure.** The validated plan is promoted through integration testing, environment parity checks, and progressive environment promotion.

- **Enterprise knowledge.** Quality metrics — requirements coverage, test automation ratio, etc. — are computed as trend indicators, not pass/fail gates. Gate thresholds are separate enforcement mechanisms.

- **Human layer — exploratory testing.** Automation handles regression; humans handle exploration. Exploratory testing investigates how the product behaves in scenarios that cannot be pre-scripted — edge cases, workflow combinations, real-world usage patterns. It requires product knowledge, judgment, and creativity — and improves as tester understanding of the domain deepens.

*Directional control areas at the gate → Release & Operate:*

| Control area | What it covers |
|---|---|
| Functional validation (QA framework) | Executable specifications run deterministically against the release candidate; results link back to the specs |
| Non-functional validation (QA framework) | Performance, security, resilience, and accessibility acceptance tests against declared NFRs |
| End-to-end business process validation (QA framework) | Cross-system workflows validated in production-like environments for tier-applicable applications; business owner sign-off on scenarios |
| Environment compliance | Integration testing, environment parity checks, and policy-as-code validation through progressive promotion |
| Regression discipline | Mandatory automated regression coverage executed before every release |
| Exploratory testing | Human, judgment-driven investigation of scenarios that cannot be pre-scripted — edge cases, workflow combinations, real-world usage patterns |

#### Release & Operate

Validated artifacts are deployed through governed promotion. Operational data flows back into specification investment and gate calibration.

- **Requirements.** The traceability chain is verifiable: requirement → spec → implementation → test result → artifact → deployment → runtime. Material defects (incidents, release-blocking regressions, recurring patterns) feed back via full RCA — each one asks *what did we miss: a spec, a control, or a gate?* Lower-severity issues are captured automatically as feedback signals and aggregated for pattern detection; full RCA triggers when a category accumulates beyond a materiality threshold.

- **Code.** Observability is applied by default through platform templates — structured logging, standard metrics, deployment version visibility. Application teams inherit these defaults.

- **Infrastructure.** The platform enforces runtime guardrails — drift detection, identity, encryption, network segmentation — automatically and continuously.

- **Enterprise knowledge.** Delivery health metrics (DORA) measure system-level performance. Incident RCA feeds back into regression suites, specification gaps, and design priorities. This is the return path: operational insight feeding back into the front of the lifecycle.

*Directional control areas in operation:*

| Control area | What it covers |
|---|---|
| Governed promotion | Only pipelines promote artifacts to production; all upstream gates green; promotion recorded with provenance |
| Runtime guardrails | Drift detection, identity enforcement, encryption validation, network segmentation — enforced continuously via policy-as-code |
| Observability by default | Structured logging, standard metrics, deployment visibility — platform templates; unobservable services cannot be considered production-ready |
| Incident detection and response | Monitoring and alerting against SLOs; break-glass procedures are governed, time-limited, audited, and rare |
| RCA feedback loop | Material defects trigger full RCA — *what was missing — a spec, a control, or a gate?* Output updates the relevant upstream control; closure verified via the traceability chain. Lower-severity issues feed signal aggregation; pattern-based escalation triggers RCA when a category warrants it |
| Delivery health telemetry | DORA metrics (deployment frequency, lead time for changes, change failure rate, MTTR); change failure rate correlates deployment and incident data via the data-and-metrics workstream |

### Target Architecture

The components below realise the operating model walked above — their integration, the scope they span, and how controls mature toward continuous assurance. Platform-agnostic in definition; platform-specific in enforcement.

#### Scope of this Architecture

This architecture addresses **custom-build software** end-to-end — code authored in managed repositories, built and validated through CI/CD, promoted to production through governed pipelines.

- **Custom-build on ADO** (or equivalent CI/CD for managed code repositories) — primary scope
- **Low-code platforms** (Palantir, Power Platform) — inherit the SDLC shape with platform-native variants of the components below; per-class extension follows in 2026 H2
- **ML / analytical workloads** — inherit the SDLC shape with class-specific decision-spec, training-pipeline, and model-registry variants; per-class extension follows in 2026 H2
- **COTS and SaaS** — out of scope for *this architecture*; their target states are sequenced as 2027 deliverables (see §1 *Lifecycle Across Software Classes* for their phase content)
- **End-user computing (EUC) / End-User Applications (EUA)** — out of scope entirely; user-developed spreadsheets, personal macros, and ad-hoc tools not managed as business applications

For low-code and ML/analytical platforms in scope, the architecture must be **considered per class** — component responsibilities and enforcement mechanisms (policy-as-code, platform-native configuration, or team-demonstrated evidence) are defined by the platform team and the QA workstream together. This TS defines the architectural contract for custom-build; platform-specific architecture documents carry the implementation for each class.

"Platform" in this architecture is a concept — governed promotion with validated controls — not a specific tool. Custom-build on ADO implements it one way; Palantir implements an equivalent natively; Power Platform through its ALM; model registries through validation pipelines.

#### Core Components

Four integrated components make up the target architecture. The platform team builds and maintains each; application teams consume them through designed interfaces.

**Harness.** The developer and agent feedback loop that runs shared controls from the repository — the same code the pipeline runs. Gives engineers and agents fast, trustworthy feedback regardless of where the harness runs (workstation, cloud IDE, remote development VM, agent sandbox). Specifications, lint rules, unit tests, static analysis, and policy checks execute in the harness before reaching CI/CD. The defining property is parity with the pipeline, not location. The harness makes agent-driven development safe by giving agents the same controls as humans. *Owner:* platform team (shared controls); tooling teams (harness packaging).

**Platform.** The single authority on releasability. Re-executes the shared controls the harness ran locally, then adds what cannot run on a workstation: end-to-end tests, NFR acceptance, security scanning, environment-level compliance. Produces immutable signed release candidates with provenance; promotes through progressively demanding environments; enforces gates at each transition. The "platform" manifests differently per class — ADO CI/CD pipelines for custom-build, Palantir native promotion for that platform's apps, Power Platform ALM for low-code, model registry validation for ML — same concept each time. *Owner:* CI/CD workstream (custom-build architecture); platform teams (per-class implementations).

**Policy-as-code and the control catalogue.** The governance workstream maintains the central control catalogue — structure, maturity model, and the platform applicability matrix. Each domain workstream (QA, CI/CD, observability, infrastructure) authors its control definitions platform-agnostically and contributes them upward to the catalogue. Platform workstreams pick up controls and implement them through the enforcement mode appropriate to each platform:

- **Policy-as-code** where the platform supports mechanical enforcement (CI/CD pipelines with policy gates, infrastructure with OPA/Sentinel-equivalent, model registries with validation policies).
- **Platform-native configuration** where controls are built into the platform's own product (Palantir's native approval gates, Power Platform's ALM, SaaS vendor SLA monitoring).
- **Team-demonstrated evidence** where neither applies (COTS boundaries, SaaS configuration governance, or immature platform enforcement); teams evidence compliance against the same control definitions.

Teams on mature platforms inherit quality by default, and the cost of governance is low because evidence is a byproduct of delivery. Teams on immature or exception-type platforms evidence compliance manually — legitimate but transient as platforms mature. *Owner:* governance workstream (catalogue + matrix); domain workstreams (control definitions); platform teams (per-platform implementation).

**Enterprise knowledge.** The measurement chain spanning platform telemetry, QA tool output, incident data, and operational metrics — across software classes. DORA metrics, defect detection ratios, and estate-wide quality visibility flow from this layer, with automated return-path feedback from production into specification investment and gate calibration. Operational evidence applies regardless of class: incidents for an Outlook integration inform control sharpening as much as production incidents for a custom-build service. Ownership of the unified data layer is cross-workstream, currently unresolved at programme level (QA PID §8 dep #6). *Owner:* data and metrics workstream (interface); domain workstreams (data contracts).

#### How the Components Connect

Four design threads run through the four components. Each thread crosses multiple components; the components cohere because they collectively realise all four threads.

**Thread 1 — Requirements → Executable Specifications → Evidence.** Authored in the **Harness** (where specifications are drafted and validated against templates). Validated through the **Platform** (where specs are checked against implementation as part of promotion). Governed by the **Control catalogue** (where spec format, coverage, and sign-off controls are defined platform-agnostically). Evidenced via the **Enterprise knowledge** (where traceability chains and coverage metrics flow). In an agent-driven world, specifications are the governance mechanism: they constrain what agents build; the pipeline validates what they produce.

**Thread 2 — Application Code → Pipeline → Production.** Realised by the **Harness** (where code is written with fast, trustworthy feedback from shared controls) and the **Platform** (where the same controls re-execute for enforcement, producing immutable signed artifacts). Gate definitions and thresholds live in the **Control catalogue**. Delivery metrics and defect detection flow through the **Enterprise knowledge**. Key design decision: shared control implementation — validation logic defined once in the repository, executed identically locally and in the pipeline.

**Thread 3 — Infrastructure → Policy → Runtime.** Declared via the **Harness** (workload manifests authored alongside application code) and validated by the **Platform** (policy-as-code checks against schema and policy). Infrastructure controls authored and maintained in the **Control catalogue**. Runtime guardrails enforced continuously, reported via the **Enterprise knowledge**. The manifest-driven model means application teams declare what they need without writing infrastructure; the platform translates declarations into compliant, governed infrastructure.

**Thread 4 — Enterprise knowledge (return path).** Primarily lives in the **Enterprise knowledge** component. Feeds back into the **Harness** (incident patterns and operational evidence shape specification investment) and into the **Control catalogue** (gate thresholds refined from operational evidence). The unified data layer enables automated return-path feedback from production into design; ownership is currently unresolved at programme level (QA PID §8 dep #6).

#### Control Maturity Journey

Controls live in the **Control catalogue** component (see *Core Components*). Each control progresses through four stages of maturity. Progression is not uniform across platforms — the same control may be *gated* on a mature platform (e.g., ADO with policy-as-code on pipeline) and *defined* on an immature one (e.g., a newly-onboarded Palantir workspace). The architecture supports the full progression; the roadmap defines when each control arrives at each stage for each platform.

| | **Defined** | **Measured** | **Gated** | **Governed** |
|---|---|---|---|---|
| **What it means** | Control specified in the catalogue with clear criteria. Teams self-assess, supported by assessment tooling (QA tool, checker agents). Enforcement is procedural, not mechanical. | Control observable in platform telemetry. Compliance visible at individual application and estate level. Non-compliance reported (dashboards, scorecards) but not blocked. | Control enforced mechanically. Non-compliance blocks the relevant action (merge, promotion, sign-off). Break-glass exists only as governed, time-limited, audited exception (per principle 8). | Control continuously assured. Enforcement monitored; effectiveness measured (is the control still preventing the failures it was designed for?); definition evolves from operational evidence (per principle 4). Ownership named; audits occur. Mature steady-state. |
| **Example — *"code changes require peer review"*** | Catalogue specifies *"every change to critical custom-build must be peer-reviewed before merge"*; teams self-assess via QA tool assessment. | ADO reports number and percentage of merges with completed review; scorecard visible to team and at estate level. | ADO branch policy blocks merge without approved review; break-glass governed. | Review effectiveness tracked (do reviewed changes produce fewer production defects?); policy refined (e.g., two reviewers for tier-1 apps); break-glass use audited and trended down. |

The Measured → Gated transition is typically rolled out progressively: visibility (report violations) → soft enforcement (block new, grandfather existing) → hard enforcement (block everything). This staged approach lets platforms move controls to mechanical enforcement without breaking existing teams.
