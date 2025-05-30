# Security Service Management Server

This is a backend server for managing security service queries, analytics, and admin authentication for a security services company. It is built with Go and provides a RESTful API for frontend integration.

## Features

- **Admin Authentication**: Secure login/logout with session cookies.
- **Query Management**: Add, update, and fetch customer queries for various security services.
- **Analytics**: Revenue analytics by service and month, including top services, pie/bar chart data, and growth trends.
- **Chatbot Integration**: Connects to Gemini API for AI-powered customer support.
- **CORS Support**: Allows frontend (e.g., React) to communicate securely.
- **JSON File Storage**: All data is stored in a simple database.json file.

## Folder Structure

```
handlers/               # All HTTP handler logic (API endpoints)
  ├── addNewQueries.go      # Handles adding new queries
  ├── analytics.go          # Analytics-related logic
  ├── auth.go               # Admin authentication handlers
  ├── chatbot.go            # Gemini chatbot integration
  ├── getAllQueries.go      # Fetching queries
  └── queryHandlers.go      # Miscellaneous query operations

models/                 # Data structures (Admin, Query, Analytics types)
  └── dataStructure.go      # Contains struct definitions

database.json           # Main data storage (queries, etc.)
.env                    # Environment variables (e.g., Gemini API key)
.gitignore              # Git ignored files list
go.mod                  # Go module definition
go.sum                  # Go dependencies checksum
main.go                 # Server entry point and route setup

```

## API Endpoints

| Endpoint                | Method | Description                                 |
|-------------------------|--------|---------------------------------------------|
| `/api/login`            | POST   | Admin login                                 |
| `/api/logout`           | POST   | Admin logout                                |
| `/api/check-login`      | GET    | Check login status                          |
| `/api/getAllQueries`    | GET    | Get all queries (sorted by ID descending)   |
| `/api/add-query`        | POST   | Add a new query                             |
| `/api/updateStatus`     | POST   | Update query status                         |
| `/api/analytics`        | GET    | Get analytics data for dashboard            |
| `/api/chat`             | POST   | Chatbot endpoint (Gemini API)               |

## Data Model

### Query Example

```json
{
  "id": 1,
  "name": "Alice",
  "email": "alice@example.com",
  "phone": "9991110001",
  "service": "Club Guards",
  "message": "Need guards for club event",
  "submitted_at": "2023-06-14T10:15:30Z",
  "numGuards": "4",
  "durationType": "hours",
  "durationValue": "10",
  "cameraRequired": false,
  "vehicleRequired": true,
  "firstAid": true,
  "walkieTalkie": false,
  "bulletProof": false,
  "fireSafety": true,
  "status": "Resolved",
  "cost": 4500
}
```

## Getting Started

### Prerequisites

- Go 1.18 or newer
- [Gemini API Key](https://ai.google.dev/gemini-api/docs/get-started) (for chatbot)

### Installation

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/security-service-server.git
   cd security-service-server
   ```

2. **Install dependencies:**
   ```sh
   go mod tidy
   ```

3. **Set up environment variables:**
   - Create a .env file in the root directory:
     ```
     GEMINI_API_KEY=your_gemini_api_key_here
     ```

4. **Run the server:**
   ```sh
   go run main.go
   ```
   The server will start on [http://localhost:8080](http://localhost:8080).

## Frontend Integration

- The server is CORS-enabled for `http://localhost:3000` (React default).
- Use `credentials: "include"` in your frontend fetch requests to handle authentication cookies.

## Customization

- **Admin Credentials:** Change the default admin email/password in [`models/dataStructure.go`](models/dataStructure.go).
- **Services & Pricing:** Update the chatbot system prompt in [`handlers/chatbot.go`](handlers/chatbot.go).
- **Data Storage:** All data is stored in `database.json`. For production, consider migrating to a real database.

**Feel free to fork and adapt for your own security service business!**
   The server will start on [http://localhost:8080](http://localhost:8080).

## Frontend Integration

- The server is CORS-enabled for `http://localhost:3000` (React default).
- Use `credentials: "include"` in your frontend fetch requests to handle authentication cookies.

## Customization

- **Admin Credentials:** Change the default admin email/password in [`models/dataStructure.go`](models/dataStructure.go).
- **Services & Pricing:** Update the chatbot system prompt in [`handlers/chatbot.go`](handlers/chatbot.go).
- **Data Storage:** All data is stored in `database.json`. For production, consider migrating to a real database.

**Feel free to fork and adapt for your own security service business!**
