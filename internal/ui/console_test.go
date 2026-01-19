package ui

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintBanner(t *testing.T) {
	// 1. Setup "Fake Terminal" (Buffer)
	// Instead of printing to screen, code will print to this variable
	var buffer bytes.Buffer

	// 2. Define test data
	testURL := "http://192.168.1.10:8080"
	testMode := "Download Test"
	testPath := "/tmp/test-files"

	// 3. Run the function with our buffer
	PrintBanner(&buffer, testURL, testMode, testPath)

	// 4. Get the output as string
	output := buffer.String()

	// 5. Assertions (Check if output contains expected info)

	// Check Title
	if !strings.Contains(output, "QUICK SHARE") {
		t.Errorf("Output missing title. Got:\n%s", output)
	}

	// Check URL
	if !strings.Contains(output, testURL) {
		t.Errorf("Output missing URL %s. Got:\n%s", testURL, output)
	}

	// Check Mode
	if !strings.Contains(output, testMode) {
		t.Errorf("Output missing Mode %s. Got:\n%s", testMode, output)
	}

	// Check Path
	if !strings.Contains(output, testPath) {
		t.Errorf("Output missing Path %s. Got:\n%s", testPath, output)
	}

	// Check if QR Code logic was triggered (QR usually prints block characters)
	// We just check if the output is significantly long, implying QR data exists
	if len(output) < 100 {
		t.Errorf("Output seems too short, QR code might be missing.")
	}
}
