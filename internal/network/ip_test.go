package network

import (
	"net"
	"testing"
)

// TestGetLocalIP ensures the function runs without panic and returns a valid IP structure if successful.
func TestGetLocalIP(t *testing.T) {
	ip, err := GetLocalIP()

	// Logic:
	// It is possible for this to return an error if the machine has no network interface.
	// However, if it returns an IP, it MUST be a valid IP string.
	if err == nil {
		parsedIP := net.ParseIP(ip)
		if parsedIP == nil {
			t.Errorf("GetLocalIP returned an invalid IP string: %s", ip)
		}
		if parsedIP.IsLoopback() {
			t.Errorf("GetLocalIP returned a loopback address (127.0.0.1), expected a LAN IP")
		}
	} else {
		// If error, just log it. We don't fail the test because CI environments
		// sometimes don't have standard network interfaces.
		t.Logf("GetLocalIP returned error (acceptable in some CI envs): %v", err)
	}
}
