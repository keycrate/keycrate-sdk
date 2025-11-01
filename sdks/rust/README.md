# Keycrate Rust SDK

License authentication SDK for Rust projects.

## Installation

Add this to your `Cargo.toml`:

```toml
[dependencies]
keycrate = "1.0.0"
```

## Usage

### Authenticate with License Key

```rust
use keycrate::{LicenseAuthClient, AuthenticateOptions};

#[tokio::main]
async fn main() {
    let client = LicenseAuthClient::new(
        "https://api.keycrate.dev",
        "your-app-id"
    );

    let opts = AuthenticateOptions {
        license: Some("your-license-key".to_string()),
        ..Default::default()
    };

    match client.authenticate(opts).await {
        Ok(result) => {
            if result.success {
                println!("License verified!");
            } else {
                println!("Error: {}", result.message);
            }
        }
        Err(e) => eprintln!("Request failed: {}", e),
    }
}
```

### Authenticate with Username/Password

```rust
let opts = AuthenticateOptions {
    username: Some("user123".to_string()),
    password: Some("password123".to_string()),
    ..Default::default()
};

let result = client.authenticate(opts).await?;
```

### Authenticate with HWID (Hardware ID)

```rust
let opts = AuthenticateOptions {
    license: Some("your-license-key".to_string()),
    hwid: Some("device-id-12345".to_string()),
    ..Default::default()
};

let result = client.authenticate(opts).await?;
```

### Register Credentials

```rust
use keycrate::RegisterOptions;

let opts = RegisterOptions {
    license: "your-license-key".to_string(),
    username: "newuser@example.com".to_string(),
    password: "securepassword".to_string(),
};

let result = client.register(opts).await?;

if result.success {
    println!("Registration successful!");
} else {
    println!("Error: {}", result.message);
}
```

## API Reference

### `LicenseAuthClient::new(host, app_id)`

Creates a new Keycrate client.

**Parameters:**

-   `host` (impl Into<String>): Base URL of the Keycrate API
-   `app_id` (impl Into<String>): Your application ID

**Returns:** `LicenseAuthClient` instance

### `client.authenticate(opts) -> Result<AuthResponse, Box<dyn Error>>`

Authenticate using either a license key or username/password.

**Parameters:**

-   `opts.license` (Option<String>): License key
-   `opts.username` (Option<String>): Username
-   `opts.password` (Option<String>): Password
-   `opts.hwid` (Option<String>): Hardware ID (optional)

**Returns:** `Result<AuthResponse, Box<dyn Error>>`

### `client.register(opts) -> Result<AuthResponse, Box<dyn Error>>`

Register credentials for a license.

**Parameters:**

-   `opts.license` (String, required): License key
-   `opts.username` (String, required): Username
-   `opts.password` (String, required): Password

**Returns:** `Result<AuthResponse, Box<dyn Error>>`

## Response Structure

All methods return an `AuthResponse`:

```rust
pub struct AuthResponse {
    pub success: bool,
    pub message: String,
    pub data: Option<serde_json::Value>,
}
```

## License

MIT
