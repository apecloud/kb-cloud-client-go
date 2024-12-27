# kb-cloud-client-go

This repository contains a Go API client for the KubeBlocks Cloud API.

## Requirements

* Go 1.22+

## Installation

```bash
go get github.com/apecloud/kb-cloud-client-go
```

## Layout

The client library contains the following packages:

### The KubeBlocks Cloud API Client

The main API client is located in the `api/kbcloud` directory. Import it with:

```go
import "github.com/apecloud/kb-cloud-client-go/api/kbcloud"
```

### The Admin API Client 

The admin API client is located in the `api/kbcloud/admin` directory. Import it with:

```go
import "github.com/apecloud/kb-cloud-client-go/api/kbcloud/admin" 
```

## Authentication

The client supports authentication via API Key & Secret. The recommended way is to use environment variables with NewDefaultContext:

```go
// Use environment variables:
// KB_CLOUD_API_KEY_NAME - API key name
// KB_CLOUD_API_KEY_SECRET - API key secret
// KB_CLOUD_SITE - Optional site configuration
ctx := common.NewDefaultContext(context.Background())
```

You can also configure authentication directly through context:

### API Key & Secret via Context

```go
ctx := context.WithValue(
    context.Background(),
    common.ContextDigestAuth,
    common.DigestAuth{
        UserName: os.Getenv("KB_CLOUD_API_KEY_NAME"),
        Password: os.Getenv("KB_CLOUD_API_KEY_SECRET"),
    },
)
```

## Getting Started

For complete examples, check out the [examples](./examples) directory.

Here's an example using API key authentication:

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/apecloud/kb-cloud-client-go/api/common"
    "github.com/apecloud/kb-cloud-client-go/api/kbcloud"
)

func main() {
    // Set up authentication context
    ctx := context.WithValue(
        context.Background(),
        common.ContextDigestAuth,
        common.DigestAuth{
            UserName: os.Getenv("KB_CLOUD_API_KEY_NAME"),
            Password: os.Getenv("KB_CLOUD_API_KEY_SECRET"),
        },
    )
    
    // Optional: Set site configuration
    ctx = context.WithValue(
        ctx,
        common.ContextServerVariables,
        map[string]string{"site": os.Getenv("KB_CLOUD_SITE")},
    )
    
    orgName := "myorg"
	client := common.NewAPIClient(configuration)
	fmt.Println("Listing environments...")
	api := kbcloud.NewEnvironmentApi(client)
	envs, resp, err := api.ListEnvironment(ctx, orgName)
	if err != nil {
		log.Fatalf("Error listing environments: %v\nResponse: %v", err, resp)
	}
	fmt.Printf("Environments: %+v\n\n", envs)
}
```

### Create Database

```go
func createDatabase() {
    // Create context with default configuration
    ctx := common.NewDefaultContext(context.Background())
    
    configuration := kbcloud.NewConfiguration() 
    client := kbcloud.NewAPIClient(configuration)

    database := admin.Database{
        Name: "mydb",
        CharacterSet: "utf8mb4",
        Collate: "utf8mb4_general_ci",
    }

    resp, err := client.DatabaseApi.CreateDatabase(ctx, "myorg", "mycluster", database)
    if err != nil {
        log.Fatal(err)
    }
}
```

### Configuration Options

#### Environment Variables

The client supports the following environment variables:

```bash
# Authentication
KB_CLOUD_API_KEY_NAME=your-api-key-name
KB_CLOUD_API_KEY_SECRET=your-api-key-secret

# Site configuration
KB_CLOUD_SITE=https://api-test.apecloud.com

# HTTP proxy settings
HTTP_PROXY=http://proxy.example.com:8080
HTTPS_PROXY=http://proxy.example.com:8080
```

#### Server Configuration

Configure server URL and variables:

```go
configuration := kbcloud.NewConfiguration()

// Set server URL
configuration.Host = "api.example.com"

// Set server variables
configuration.ServerVariables = map[string]string{
    "site": os.Getenv("KB_CLOUD_SITE"),
}

// Or use context to set server variables
ctx := context.WithValue(
    context.Background(),
    common.ContextServerVariables,
    map[string]string{"site": os.Getenv("KB_CLOUD_SITE")},
)
```

#### Authentication

Configure authentication using context:

```go
// Bearer Token
ctx := context.WithValue(
    context.Background(),
    common.ContextAccessToken,
    "your-token",
)

// Basic Auth
ctx := context.WithValue(
    context.Background(),
    common.ContextBasicAuth,
    common.BasicAuth{
        UserName: "user",
        Password: "pass",
    },
)

// API Key
ctx := context.WithValue(
    context.Background(),
    common.ContextAPIKeys,
    map[string]common.APIKey{
        "apiKey": {
            Key: "your-api-key",
            Prefix: "Bearer",
        },
    },
)
```

#### Debug Mode

Enable debug logging:

```go
configuration := kbcloud.NewConfiguration()
configuration.Debug = true
```

#### Retry Configuration

Configure retry behavior:

```go
configuration := kbcloud.NewConfiguration()

// Configure retry settings
configuration.RetryConfiguration = common.RetryConfiguration{
    EnableRetry: true,           // Enable/disable retry
    MaxRetries: 3,               // Maximum number of retries
    RetryInterval: 2,            // Base retry interval in seconds
    MaxRetryInterval: 30,        // Maximum retry interval in seconds
    BackOffBase: 2,              // Exponential backoff base
    RequestTimeout: 30,          // Request timeout in seconds
}
```

#### HTTP Client Configuration

Configure HTTP client:

```go
configuration := kbcloud.NewConfiguration()

// Custom HTTP client
configuration.HTTPClient = &http.Client{
    Timeout: time.Second * 30,
    Transport: &http.Transport{
        MaxIdleConns: 10,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout: 30 * time.Second,
    },
}

// Skip TLS verification (not recommended for production)
ctx := context.WithValue(
    context.Background(),
    common.ContextInsecureSkipVerify,
    true,
)
```

#### User Agent

Configure custom user agent:

```go
configuration := kbcloud.NewConfiguration()
configuration.UserAgent = "MyApp/1.0.0"
```

#### Default Headers

Add default headers to all requests:

```go
configuration := kbcloud.NewConfiguration()
configuration.AddDefaultHeader("Custom-Header", "value")
```

## Development

### Generate API Client

To regenerate the API client code:

```bash
./generate.sh
```

## Contributing

We welcome contributions! Please feel free to submit a Pull Request.

## License

This library is distributed under the Apache 2.0 license found in the [LICENSE](./LICENSE) file.