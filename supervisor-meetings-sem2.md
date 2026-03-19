# 19/03/2026

Agenda:

- Haven't had much time to work this week because of other modules
  - Currently working on:
    - Finishing off operator (CarbonInfo and EmissionReport)
    - Finishing off tests
  - AS FAR AS OVERALL PROJECT, Still need to work on
    - CLI interface
    - Visuals for whole project in motion
- Showcase SUBMISSION FOR 25th MARCH 2PM
  - Only Submission box for poster
  - Poster will most likely contain same summary as the submission
- SUNLIFE Awards due 10th April:
  - 3 Categories - Innovation, Enterprise, Intellectual Curiosity
   
Notes:

- Time better spent on Presentations/tidying up/Visuals rather than porting to cloud providers
- Logged into comp expo, poster due 25th March
- Sunlife awards - Probably intellectual curiosity

Action Items:

- GET POSTER SENT BEFORE SUBMISSION Monday
- Filling out rest of showcase submission
- Finish off current work
- Revise work schedule for next Thursday review


# 12/02/2026

Agenda:

- Initialising the cluster with Kepler and Prometheus to take in and store node energy data
- Briefly stepping through CRDs  
- Running the operator (with the 3/5 controllers running, WorkloadPolicyController, KeplerMetricsSyncController and ReschedulingController) and running the scheduler
- Providing demo scenario where Node A is more efficient to schedule to than Node B, and applying a WorkloadPolicy to this for the scheduler to use when scheduling the demo workload
- Potentially rescheduling too (which I had working but broke when I had to go back to change some other workload policy logic), where Node B becomes more efficient to use over Node A and the rescheduling policy in the workload policy allows for the "reschedule" (pod eviction) to take place

Notes:

- Didn't work, need to find why
- Need greater visualisation to show the intricacies of what is being done
- Need to think of where to go next once fixed
- AI tools that can make visualisation easy? graphical or animated scenario of the total flow involved in scheduler

Action Items:

- Fix the workloadpolicycontroller issue
- Investigate tools to make visualisation more apparent and show to anyone how the scheduler/operator works itself
- Run a schedule for the week


# 05/02/2026

Agenda:

- Action Items from Last week:
  - Demo step through done for 12/02/2025
  - Issues with Jira, now switching to Trello
- CRDs updated and controllers mostly complete with minimal function, now rewriting scheduler scoring plugin to account for new changes

Notes:

- Demo at 4:30 12/02/2025
- Showcase submissions after reading week


Action Items:

- Pre-Write agenda for demo
- DEMO FOR WEEK 4 12/02/2025 half 4 (provisional)
- Richard to check recorded demo duration/details for project submission

# 29/01/2026

Agenda:

- Action Items from Last week:
  - Demo step through done for 12/02/2025
  - Issues with Jira:
    - Errors loading views
    - New UI

Notes:

- Step through demo for week 4 12/02/2025 half 4 (provisional)
  - Scheduler workflow
  - Controller logs/environment interaction
- Updating CRDs at the moment to be in line with current report
- Could switch to Trello if Jira is still acting up


Action Items:

- DEMO FOR WEEK 4 12/02/2025 half 4 (provisional)
- Share Jira or Trello


# 21/01/2026

Agenda:
- Semester 1 presentation and document

Notes:

- Presentation and things to add to document:
  - General Kubernetes background (Appendix/Glossary)
  - Add a description after system architecture diagram (similar to users stories but more lower level)
  - Testing strategies (scheduling, rescheduling)
  - Step through demo

Action Items:

- Create Jira board and start writing epics, stories, etc.
- Develop a testing scenario and prepare a step through to go through the actual work flow of how scheduler and operator/controllers work (set date before mid-term)
- High level breakdowns of work for the semester
