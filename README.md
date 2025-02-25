# defi-examples-go

This repository contains example scripts that leverage Blockdaemon's DEFI API using the Go SDK. These examples demonstrate how to interact with various DeFi functionalities across different blockchain networks.

## Requirements

- Go 1.21 or higher
- A Blockdaemon API Key

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/your-username/defi-examples-go
cd defi-examples-go
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up your environment:
```bash
cp .env.example .env
```

1. Add your [Blockdaemon API Key](https://www.blockdaemon.com/api) to the `.env` file:
```properties
BLOCKDAEMON_API_KEY=your_api_key_here
DEFI_API_BASE_PATH=https://svc.blockdaemon.com/defi/v1
```

## Available Examples

- `chains`: Lists supported blockchain networks

### Running Examples

Run any example using `go run`:

```bash
go run examples/chains/main.go    # Get supported blockchain networks
```

## Environment Variables

The following environment variables are required:

- `BLOCKDAEMON_API_KEY`: Your Blockdaemon API key
- `DEFI_API_BASE_PATH`: API base URL (defaults to https://svc.blockdaemon.com/defi/v1)

## Legal Warning ⚠️🚨

These example applications are provided as learning resources for the DEFI API. They are experimental and should not be used in production environments. For more information, please contact our team through our [support page](https://www.blockdaemon.com/support).

## Common Issues and Troubleshooting

### API Key Issues
- Ensure your API key is correctly set in the `.env` file
- Verify your API key has the necessary permissions

### Debug Logging
To enable more verbose logging, set the log level in your code:

```go
log.SetFlags(log.LstdFlags | log.Lshortfile)
```

## Need Help? 👋

Contact us through [email](support@blockdaemon.com) or our [support page](https://www.blockdaemon.com/support) for any assistance you may need.
