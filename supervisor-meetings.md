25/09/2025

Agenda:
- PROJECT GOALS -
    - Embed sustainability into Kubernetes operations by extending the platform with carbon-aware scheduling, autoscaling, governance, and observability
    - Provide actionable visibility into the environmental impact of workloads through real-time metrics, dashboards, and emissions reporting.
    - Simplify adoption by delivering the solution as an easily deployable, CLI-driven stack for cloud environments (AWS, Azure).
    - Validate effectiveness by demonstrating measurable emissions reductions and trade-offs compared to standard Kubernetes operation.
    
- OBJECTIVES - 
    - Design and implement custom APIs (CRDs) for carbon sources, green windows, workload sustainability policies, and emissions reporting.
    - Develop controllers in Go (using Kubebuilder) that fetch carbon intensity data (ElectricityMap/WattTime APIs), apply workload policies, and update status accordingly.
    - Integrate carbon signals into scheduling and autoscaling by implementing a scheduler plugin/extender
    - Enforce governance rules via OPA Gatekeeper, ensuring workloads comply with sustainability and resource efficiency standards.
    - Build observability dashboards with Prometheus and Grafana to visualize carbon intensity, workload emissions, and sustainability policy effects.
    - Package a deployment workflow using a Cobra-based CLI that provisions clusters (Kind locally for testing, Cluster APIs on AWS/Azure) and installs the sustainability stack.
    - Evaluate the system by comparing baseline (vanilla Kubernetes) vs. sustainability-enhanced operation, measuring energy use, emissions, and performance trade-offs.

Notes: 
- Currently no schedule or deadlines for FYP, will most likely be confirmed once supervisors are assigned (which will take place after all FYP proposals are submitted)
- Articles on current discussions of sustainable Kubernetes workloads.

Actionable Items:
- Investigate how tasks will be tracked (Agile method, through Jira, Trello or another service).
- Expand on Objectives: What can be done now, in parallel, what tasks are coming from the objectives, etc.
- Investigate scheduler (bin-packing problem)
- Review articles sent on Slack
