package touchid

import (
	"fmt"
	"testing"
	"time"
)

func TestAuth(t *testing.T) {
	tests := []struct {
		name       string
		deviceType DeviceType
		reason     string
		wantErr    bool
	}{
		{"DeviceTypeAny", DeviceTypeAny, "Confirm Action", false},
		{"DeviceTypeBiometrics", DeviceTypeBiometrics, "Confirm Action", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			success, err := Auth(tt.deviceType, tt.reason)
			if (err != nil) != tt.wantErr {
				t.Errorf("Auth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !success {
				t.Errorf("Auth() success = %v, want true", success)
			}
			fmt.Printf("%s: Authentication %s\n", tt.name, map[bool]string{true: "successful", false: "failed"}[success])
		})
	}
}

func TestSerialAuth(t *testing.T) {
	deviceType := DeviceTypeBiometrics
	reason := "Confirm Action"
	timeout := 10 * time.Second

	// First authentication
	success, err := SerialAuth(deviceType, reason, timeout)
	if err != nil {
		t.Fatalf("SerialAuth() error = %v", err)
	}
	if !success {
		t.Fatalf("SerialAuth() success = %v, want true", success)
	}
	fmt.Println("First authentication successful")

	// Second authentication (should not prompt)
	start := time.Now()
	success, err = SerialAuth(deviceType, reason, timeout)
	if err != nil {
		t.Fatalf("SerialAuth() error = %v", err)
	}
	if !success {
		t.Fatalf("SerialAuth() success = %v, want true", success)
	}
	if time.Since(start) > time.Second {
		t.Errorf("SerialAuth() took too long, expected instant response")
	}
	fmt.Println("Second authentication successful (no prompt)")

	// Wait for timeout to expire
	time.Sleep(timeout)

	// Third authentication (should prompt again)
	success, err = SerialAuth(deviceType, reason, timeout)
	if err != nil {
		t.Fatalf("SerialAuth() error = %v", err)
	}
	if !success {
		t.Fatalf("SerialAuth() success = %v, want true", success)
	}
	fmt.Println("Third authentication successful (after timeout)")
}
