# Keycrate Go SDK

License authentication SDK for Go projects.

## Installation

```bash
go get github.com/keycrate/keycrate-sdk/sdks/go
```

## Usage

### Authenticate with License Key

```go
package main

import (
	"fmt"
	"log"
	"github.com/keycrate/keycrate-sdk/sdks/go"
)

func main() {
	client := keycrate.New("https://api.keycrate.dev", "your-app-id")

	result, err := client.Authenticate(keycrate.AuthenticateOptions{
		License: "your-license-key",
	})

	if err != nil {
		log.Fatal(err)
	}

	if result.Success {
		fmt.Println("License verified!")
	} else {
		fmt.Println("Error:", result.Message)
	}
}
```

### Authenticate with Username/Password

```go
result, err := client.Authenticate(keycrate.AuthenticateOptions{
	Username: "user@example.com",
	Password: "password123",
})
```

### Authenticate with HWID (Hardware ID)

```go
result, err := client.Authenticate(keycrate.AuthenticateOptions{
	License: "your-license-key", // Licese or Username and Password
	HWID:    "device-id-12345",
})
```

### Register Credentials

```go
result, err := client.Register(keycrate.RegisterOptions{
	License:  "your-license-key",
	Username: "newuser@example.com",
	Password: "securepassword",
})

if result.Success {
	fmt.Println("Registration successful!")
} else {
	fmt.Println("Error:", result.Message)
}
```

## API Reference

### `New(host, appID string) *Client`

Creates and returns a new Keycrate client.

**Parameters:**

-   `host` (string): Base URL of the Keycrate API
-   `appID` (string): Your application ID

**Returns:** `*Client` instance

### `client.Authenticate(opts AuthenticateOptions) (*AuthResponse, error)`

Authenticate using either a license key or username/password.

**Parameters:**

-   `opts.License` (string): License key
-   `opts.Username` (string): Username
-   `opts.Password` (string): Password
-   `opts.HWID` (string): Hardware ID (optional)

**Returns:** `*AuthResponse` with success status and message

### `client.Register(opts RegisterOptions) (*AuthResponse, error)`

Register credentials for a license.

**Parameters:**

-   `opts.License` (string, required): License key
-   `opts.Username` (string, required): Username
-   `opts.Password` (string, required): Password

**Returns:** `*AuthResponse` with success status and message

## Response Structure

All methods return an `AuthResponse`:

```go
type AuthResponse struct {
	Success bool                   // true if operation succeeded
	Message string                 // Response message
	Data    map[string]interface{} // Optional response data
}
```

## License

MIT
