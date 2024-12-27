package main

import (
	"fmt"
	"log"

	"github.com/apecloud/kb-cloud-client-go/api/common"
	"github.com/apecloud/kb-cloud-client-go/api/kbcloud"
	"github.com/apecloud/kb-cloud-client-go/examples/auth"
)

// This example demonstrates how to use the KubeBlocks Cloud API client.
//
// To run this example:
//
// 1. Set up your environment variables:
//
//    export KB_CLOUD_API_KEY_NAME="your-api-key-name"
//    export KB_CLOUD_API_KEY_SECRET="your-api-key-secret"
//    export KB_CLOUD_SITE="https://api-test.apecloud.com"
//
// 2. Run the example:
//
//    go run examples/main.go
//
// The example will list all environments in your organization.

// orgName is your organization name in KubeBlocks Cloud
// You can find it in the KubeBlocks Cloud console
const orgName = "my-org"

func main() {
	// Create a new authenticated context using environment variables:
	// - KB_CLOUD_API_KEY_NAME: Your API key name
	// - KB_CLOUD_API_KEY_SECRET: Your API key secret
	// - KB_CLOUD_SITE: API endpoint (e.g., https://api-test.apecloud.com)
	ctx := auth.NewAuthContext()

	// Create a new configuration with debug mode enabled
	// This will print API requests and responses to stdout
	configuration := common.NewConfiguration()
	configuration.Debug = true

	// Create a new API client using the configuration
	client := common.NewAPIClient(configuration)

	// Create a new Environment API instance
	// This demonstrates how to create and use specific API endpoints
	fmt.Println("Listing environments...")
	api := kbcloud.NewEnvironmentApi(client)

	// Call the ListEnvironment API
	// This will list all environments in your organization
	envs, resp, err := api.ListEnvironment(ctx, orgName)
	if err != nil {
		// Handle errors and print both the error and the full HTTP response
		log.Fatalf("Error listing environments: %v\nResponse: %v", err, resp)
	}

	// Print the environments
	fmt.Printf("Environments: %+v\n\n", envs)
}
