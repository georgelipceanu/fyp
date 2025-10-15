# 15/10/2025

Agenda:
- Previous Action Items -
	- Energy formula expansion:
 		- Derived from linear power model in article provided on Slack
   		- Could try other dynamic power functions outside of prototype if needed
		- Update to scheduler doc and actual doc
 	- Prototype:
  		- Working CRDs with Kubebuilder
    	- Scheduler WIP (hope to have a barebones working version by the end of the week)
     	- Attempting locally right now, will migrate to EKS once a baseline is established
  	- Document:
  		- Base format laid out (sections may be added in the future) 
  		- Introduction and Scheduler sections 
- Electricity Maps Response -
	- Told the sign up with academic email (which hasn't worked), waiting for response back

Notes:

Action Items:


# 09/10/2025

Agenda:
- Previous Action Items -
	- Scheduler Write-up - On Overleaf, very rough and won't be used in project but I feel that I understand the scheduler and how I want to go about it.
 	- Semester Plan - Starting from this week onwards:
		- Week 1-2: CRDs, Controller logic and Scheduler
  		- Week 3-4: Observability - Grafana, Prometheus and Kepler
    	- Week 5: CLI tooling
     	- Week 6+: Multi-Cluster Deployments
    - Template:
		- Introduction - Background + Motivation
    	- Review of Related Work
     	- System Architecture and Design
        - Scheduler Design and Implementation
        - Implementation of Supporting Components
        - Multi-Cluster Deployments
        - Testing and Evaluation
  	- Document to read:
  		- Google's large-scale operational framework that reduces the carbon footprint of data center computing by shifting workloads to times when local electricity grids are cleaner.
  		- 1–2% average reduction in carbon-related power use, and up to 8% reduction in some clusters with high renewable variability.
  	 	- Value in this field

Notes: 
- CPU measured in cores, how would it compare, expand on energy
- OpenShift compatibility/addons

Actionable Items:
- Expand on energy formula, explore alternatives.
- Begin work on actual document
- Begin prototyping

# 01/10/2025

Agenda:
- Previous articles for review -
	- Medium: 
		- Only Discusses how a project similar to mine could be done, not an actual finished project (which I originally thought it was).
		- Provides meaningful insight on how this can be done on a baseline level.
		- Does this project differ enough from the article for an FYP?
	- Bin-packing problem:  
		- Pods = items (CPU, memory, GPUs, etc.), Nodes = bins, Scheduler tries to spread load for resilience and balance in vanilla Kube.
		- Scheduler will also need to consider energy used, aiming to minimise nodes used while also maintaining availability.
	- Karmada: 
		- Providing multi-cluster deployments would also open up the option of workloads being deployed in certain regions depending on carbon data, allowing for greater depth in scheduler.
		- Seems to function very similar to HyperShift (which I worked on in my internship).

- APIs that could be used -
	- ElectricityMaps https://www.electricitymaps.com/:
		- Ideal since it offers all regions needed for gathering relevant carbon information.
		- Used to be free API calls to all regions with limited use (hourly updates), now it is reduced to one region per account.
		- Unlocking all regions costs 99.00 per month.
		- Discussions to revert or are ongoing https://github.com/thegreenwebfoundation/grid-aware-websites/issues/21.
		- Also contacted personally on current project idea and any potential leeway on this fee.
		- Doesn't offer sub-region level info.
	- CarbonIntensityUK https://carbonintensity.org.uk/:
		- Offers sub-region level info for UK only (eu-west-2)
		- Free.
	- Watttime https://watttime.org/: 
		- Only provides information for one region (CAISO_NORTH) with free tier.
		- Claims to provide information globally with Analyst tier, which the price of is not listed.
		- Contacted regarding this tier and why it is needed?
	- Using multiple APIs for different regions?
		- Allows system to be free (or close to free) but provides extra input from user deploying this.

- Objectives into tasks -
	- What is being done now?:
		- Currently working on Cloud Computing assignment that will investigate Prometheus and Grafana.
		- Also working on Mobile App development assignment that will include using previously mentioned APIs to investigate how workloads (or "cloud computing tasks" as they're called in this assignment) will be scheduled.
		- Investigating APIs and how these would be corresponded in CRDs.
		- Investigating other tools that could potentially be used to substitute the need for  
	- What needs to be done?: 
		- Finalise APIs being used (which is difficult since there I am waiting for responses from ElectricityMaps and Watttime).

Notes: 
- What information does AWS provide regarding energy/carbon data?
- Focus on single EKS cluster for prototype, expand to multi-region deployments if possible to get APIs and if possible in scope.
- How does the scheduler work?
- How will jobs be given a value for it to be based on? 

Actionable Items:
- Write up how scheduler works in Kubernetes, how it works and how it will be changed.
- Research additional inputs for scheduler.
- Read paper provided on Slack.
- Get template and plan out this semester.

# 25/09/2025

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
