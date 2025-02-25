package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    defiapi "github.com/Blockdaemon/defi-api-go-sdk"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    apiKey := os.Getenv("BLOCKDAEMON_API_KEY")
    if apiKey == "" {
        log.Fatal("BLOCKDAEMON_API_KEY environment variable is required")
    }

    configuration := defiapi.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer " + apiKey)
    basePath := os.Getenv("DEFI_API_BASE_PATH")
    if basePath == "" {
        basePath = "https://svc.blockdaemon.com/defi/v1"
    }
    configuration.Servers = []defiapi.ServerConfiguration{
        {
            URL: basePath,
        },
    }

    apiClient := defiapi.NewAPIClient(configuration)

    resp, r, err := apiClient.ChainsAPI.GetChains(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error calling ChainsAPI.GetChains: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
        return
    }

    // Pretty print the response
    prettyJSON, err := json.MarshalIndent(resp, "", "    ")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error formatting JSON: %v\n", err)
        return
    }
    
    fmt.Println("Supported Chains:")
    fmt.Println(string(prettyJSON))
}
