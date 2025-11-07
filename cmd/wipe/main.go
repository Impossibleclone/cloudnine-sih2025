package main

import (
	"cloudnine-sih2025/internals/cert"
	"cloudnine-sih2025/internals/scanner"
	"cloudnine-sih2025/internals/wipe"
	"cloudnine-sih2025/pkg/log"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/jung-kurt/gofpdf"
)

var (
	device  string
	passes  int
	output  string
	listAll bool // This will be set from the -l flag
	// certDir string
)

func main() {

	// --- 1. Define command-line flags ---
	flag.BoolVar(&listAll, "l", false, "List all discoverable devices and exit")
	flag.StringVar(&output, "output", "wipe_certificate", "Base name for output certificate files (e.g., 'wipe_certificate')")

	// Parse all defined flags (e.g., -l, -output)
	flag.Parse()

	// --- 2. Handle the -l flag ---
	// If -l was used, list devices and exit immediately.
	if listAll {
		devices, err := scanner.DiscoverDevices()
		if err != nil {
			log.Fatal("Failed to discover devices: %v", err)
		}
		fmt.Println("Available devices:")
		for _, dev := range devices {
			// Print a more detailed, aligned view
			fmt.Printf("  Name: %-10s Model: %-20s Size: %s\n", dev.Name, dev.Model, dev.Size)
		}
		os.Exit(0)
	}

	// --- 3. Get the device (the positional argument) ---
	// flag.Args() returns any arguments that *weren't* flags.
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Error: No device specified.")
		fmt.Println("Usage: ./wipe-tool [options] <device_path>")
		fmt.Println("Example: ./wipe-tool /dev/sda")
		fmt.Println("\nOptions:")
		flag.PrintDefaults() // Prints all defined flags
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("Error: Too many devices specified. Please specify only one.")
		os.Exit(1)
	}

	// We have exactly one device argument. Set the global 'device' variable.
	device = args[0]

	// --- 4. Run the wipe process (rest of your original code) ---
	log.Init()
	log.Info("Starting secure wipe process for device: %s", device)
	log.Info("Platform: %s, Passes: %d", runtime.GOOS, passes)

	startTime := time.Now()

	fmt.Printf("Starting wipe process for %s...\n", device)
	err := wipe.Wipe(device)
	if err != nil {
		log.Error("Wipe failed: %v", err)
		os.Exit(1)
	}
	duration := time.Since(startTime)
	log.Info("Wipe completed successfully in %v", duration)

	// Generate and save the certificate
	certData := cert.GenerateCertificate(device, passes, duration, runtime.GOOS)
	if err := saveCertificate(certData, output); err != nil {
		log.Error("Failed to save certificate: %v", err)
		os.Exit(1)
	}

	log.Info("Certificate saved to: %s.pdf and %s.json", output, output)
	log.Info("Wipe process finished.")
}

func saveCertificate(cert *cert.Certificate, output string) error {
	jsonData, err := json.MarshalIndent(cert, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(output+".json", jsonData, 0644); err != nil {
		return err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Secure Wipe Certificate")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Device: %s", cert.Device))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Duration: %s", cert.Duration))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Platform: %s", cert.Platform))
	return pdf.OutputFileAndClose(output + ".pdf")
}
