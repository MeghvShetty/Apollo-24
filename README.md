## Apollo-24

Apollo-24 is a pathfinder project designed to automate administrative workflows between Jira and ServiceNow (SNOW) using APIs and webhooks. As a middleware platform, Apollo-24 aims to accelerate customer triage response times and eliminate redundant data entry, helping large organizations maintain workflow efficiency and data consistency.

---

## Project Overview

Apollo-24 provides a web interface for customers to submit requests. Once a request is submitted, the system generates a triage report (with a range from SIRC to SIR3) and creates a Jira ticket using a standardized template. The Jira ticket is linked within the tool, allowing customers to track their requests in real time.

### Key Features

- **Jira Automation**
  - Automatically builds Jira tickets with required fields:
    - SI number & link
    - BIA record information
    - Project description
    - Project owner name/email
    - Platform name
    - Lab name
    - GW1 date as due date
    - Solution architect name/email
    - Ticket name format: `SI - Lab_Project-Name` (e.g., `MI60007-James_Bond-Spy_Car`)
  - Adds comments to SC&C SI tasks with the SIR rating.
  - Listens for Jira events (via webhooks) to trigger dependent actions, such as creating related tickets in other Jira boards and updating ticket comments for traceability.

- **ServiceNow Integration**
  - Designed to synchronize workflows with ServiceNow, reducing manual re-entry and keeping data consistent across platforms.

- **Event-Driven Workflow**
  - Example: When a ticket in Jira Board A moves from **In Progress** to **Peer Review**, Apollo-24 creates a ticket in Jira Board B’s backlog and comments on the original ticket with the new ticket number for easy mapping.

---

## Technology Stack & Packages

Based on the repository structure and code artifacts, Apollo-24 primarily uses Go (Golang) for its backend logic:

| Component          | Technology / Package    | Purpose                                      |
|--------------------|------------------------|----------------------------------------------|
| Main Backend       | Go (Golang)            | Core application logic and API server        |
| API Routing        | Custom Go router       | Handles incoming API/webhook requests        |
| Templates          | Go templates           | For generating Jira ticket content           |
| Webhook Handling   | Go                     | Listens to Jira webhook events               |
| Data Structures    | JSON                   | For request/response and webhook payloads    |

**Key Files and Directories:**
- `apollo-24.go`: Main entry point for the application.
- `api/`, `logic/`, `router/`: Core Go source code for API endpoints, business logic, and routing.
- `templates/`: Stores templates for Jira ticket creation.
- `webhook.json`: Example Jira webhook payload for development/testing.

No additional package managers (like npm or pip) are present, confirming the backend is implemented in Go, with dependencies managed via `go.mod` and `go.sum`.

---

## Getting Started

### Prerequisites

- Go 1.18+ installed
- Jira and ServiceNow instances with API access
- **Jira API keys and ServiceNow API keys must be created and securely added to your environment or configuration files before running Apollo-24.**  
- Webhook configuration in Jira to point to Apollo-24’s webhook endpoint

### Setup

1. **Clone the repository:**
   ```sh
   git clone https://github.com/MeghvShetty/Apollo-24.git
   cd Apollo-24
   ```
2. **Install dependencies:**
   ```sh
   go mod tidy
   ```
3. **Configure environment:**
   - Set up environment variables or config files for Jira/SNOW API credentials and webhook URLs.
   - **Ensure your Jira and ServiceNow API keys are generated and stored securely.**
4. **Run the server:**
   ```sh
   go run apollo-24.go
   ```
5. **Configure Jira webhooks:**
   - In your Jira instance, set up a webhook to POST relevant events (e.g., issue updates, status changes) to your Apollo-24 server.

---

## Example Workflow

- **Customer submits request via website**
- **Apollo-24:**
  - Generates triage report (SIRC-SIR3)
  - Creates Jira ticket with all required fields and standardized naming
  - Links the Jira ticket for customer tracking
- **Jira Event (e.g., status change to Peer Review):**
  - Apollo-24 creates a dependent ticket in another Jira board
  - Comments on the original ticket with the new ticket reference

---

## Contribution

- Open to issues and pull requests.
- Please follow Go best practices and ensure all new features are covered by tests where possible.

---

## License

See [LICENSE](LICENSE) for details.

---

Apollo-24 is a proof-of-concept project to demonstrate the power of event-driven automation between enterprise platforms, reducing administrative overhead and improving workflow transparency.

