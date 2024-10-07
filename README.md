# Project Structure

```
root_project/
├── cmd/                     // 1. Contains the entry points for your executables.
│   ├── api/
│   │   └── main.go          // Entry point for the API server
│   └── worker/
│       └── main.go          // Entry point for the worker program
|
|
├── internal/                // 2. Holds the application code that's not intended for public use.
│   ├── models/
│   │   ├── order.go         // Order model
│   │   ├── deposito.go      // Deposito model
│   │   └── societa.go       // Societa model
│   ├── db/
│   │   ├── db.go            // Database connection setup
│   │   └── migrations/      // Database migrations
│   ├── api/
│   │   └── handlers.go      // API route handlers
│   ├── api_client/
│   │   └── client.go        // API client for worker
│   └── xmlparser/
│       └── parser.go        // XML parsing and validation
|
|
├── pkg/                     // 3. Optional directory for packages that could be used by external applications (e.g., utility functions).
│   └── utils/
│       └── helpers.go       // Utility functions (if any)
|
|
├── frontend/                     // New frontend directory
│   ├── public/
│   │   └── index.html            // HTML template
│   ├── src/
│   │   ├── components/           // React components
│   │   ├── App.js                // Main React component
│   │   └── index.js              // Entry point for React
│   ├── package.json              // NPM package configuration
│   └── webpack.config.js         // Webpack configuration (if using Webpack)
|
├── go.mod
└── go.sum

```