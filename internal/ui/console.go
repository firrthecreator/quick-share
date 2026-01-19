package ui

import (
	"fmt"
	"io"

	"github.com/mdp/qrterminal/v3"
)

// PrintBanner displays the startup banner and information to the given writer.
// We use io.Writer to make it testable (we can capture output).
func PrintBanner(w io.Writer, url string, mode string, filePath string) {
	fmt.Fprintln(w, "QUICK SHARE - Instant File Sharing")
	fmt.Fprintf(w, "Mode:     %s\n", mode)
	if filePath != "" {
		fmt.Fprintf(w, "File/Dir: %s\n", filePath)
	}
	fmt.Fprintf(w, "URL:      %s\n", url)
	fmt.Fprintln(w, "Scan the QR code below with your mobile:")
	fmt.Fprintln(w, "")

	// Generate QR Code to the writer (w) instead of hardcoded stdout
	config := qrterminal.Config{
		Level:     qrterminal.M,
		Writer:    w,
		BlackChar: qrterminal.BLACK,
		WhiteChar: qrterminal.WHITE,
		QuietZone: 1,
	}
	qrterminal.GenerateWithConfig(url, config)

	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "Press Ctrl+C to stop the server.")
}
