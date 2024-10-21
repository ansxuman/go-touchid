# go-touchid
A simple Go Library for Touch ID authentication on darwin.

![GO Touch ID ](https://github.com/user-attachments/assets/f801d32a-58f2-4b7b-8749-5a0b60701cc7)


## Features

- Support for general device authentication and specific biometric authentication
- Serial authentication to reduce frequent prompts

## Installation

To install go-touchid

```bash
go get github.com/ansxuman/go-touchid
```

## Usage

Here's a quick example of how to use go-touchid:

```go
package main

import (
	"fmt"
	"time"

	"github.com/ansxuman/go-touchid"
)

func main() {
	// Authenticate using any available method (Touch ID or passcode)
	success, err := touchid.Auth(touchid.DeviceTypeAny, "Confirm Action")
	if err != nil {
		fmt.Printf("Authentication error: %v\n", err)
		return
	
	if success {
		fmt.Println("Authentication successful!")
	} else {
		fmt.Println("Authentication failed.")
	}

	// Use serial authentication to avoid frequent prompts
	success, err = touchid.SerialAuth(touchid.DeviceTypeBiometrics, "Confirm Action", 30*time.Second)	
	if err != nil {
		fmt.Printf("Serial authentication error: %v\n", err)
		return
	}
	if success {
		fmt.Println("Serial authentication successful!")
	} else {
		fmt.Println("Serial authentication failed.")
	}
}
}
```

## API

- `Auth(deviceType DeviceType, reason string) (bool, error)`
  Authenticates the user using the specified device type.

- `SerialAuth(deviceType DeviceType, reason string, timeout time.Duration) (bool, error)`
  Authenticates the user and caches the result for the specified timeout duration.

## Device Types

- `DeviceTypeAny`: Uses any available authentication method (Touch ID or passcode)
- `DeviceTypeBiometrics`: Specifically requests biometric authentication (Touch ID)

## License

This project is licensed under the [Apache-2.0 license](LICENSE).
