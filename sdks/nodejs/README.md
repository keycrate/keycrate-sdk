# Keycrate JavaScript SDK

License authentication SDK for JavaScript/TypeScript projects.

## Installation

```bash
npm install keycrate
```

## Usage

### Authenticate with License Key

```javascript
import { configurate } from "keycrate-js";

const client = configurate("https://api.keycrate.dev", "your-app-id");

const result = await client.authenticate({
    license: "your-license-key",
});

if (result.success) {
    console.log("License verified!");
} else {
    console.log("Error:", result.message);
}
```

### Authenticate with Username/Password

```javascript
const result = await client.authenticate({
    username: "user@example.com",
    password: "password123",
});
```

### Authenticate with HWID (Hardware ID)

```javascript
const result = await client.authenticate({
    license: "your-license-key",
    hwid: "device-id-12345",
});
```

### Register Credentials

```javascript
const result = await client.register({
    license: "your-license-key",
    username: "newuser@example.com",
    password: "securepassword",
});
```

## API Reference

### `configurate(host, appId)`

Creates and returns a client instance.

**Parameters:**

-   `host` (string): Base URL of the Keycrate API
-   `appId` (string): Your application ID

**Returns:** `LicenseAuthClient` instance

### `client.authenticate(options)`

Authenticate using either a license key or username/password.

**Parameters:**

-   `options.license` (string, optional): License key
-   `options.username` (string, optional): Username
-   `options.password` (string, optional): Password
-   `options.hwid` (string, optional): Hardware ID

**Returns:** Promise<{success: boolean, message: string, data?: object}>

### `client.register(options)`

Register credentials for a license.

**Parameters:**

-   `options.license` (string, required): License key
-   `options.username` (string, required): Username
-   `options.password` (string, required): Password

**Returns:** Promise<{success: boolean, message: string}>

## License

MIT
