package wipe

import (
	"cloudnine-sih2025/pkg/log"
	"fmt"
	"strings"
	"syscall"
)

func Wipe(device string) error {
	log.Info("Wiping device %s", device)
	var err error
	// Check the device name to determine its type
	if strings.Contains(device, "nvme") {
		// Device name contains "nvme" (e.g., "/dev/nvme0n1")
		log.Info("Detected NVMe device. Calling eraseNvme...", device)
		err = eraseNvme(device)

	} else if strings.Contains(device, "sd") {
		// Device name contains "sd" (e.g., "/dev/sda")
		log.Info("Detected SATA/SCSI device. Calling eraseSata...", device)
		err = eraseSata(device)

	} else {
		// Handle unknown or unsupported device types
		log.Error("Unknown or unsupported device type for wiping: %s", device)
		return fmt.Errorf("unsupported device type: %s", device)
	}
	if err != nil {
		log.Error("Error wiping device %s: %v", device, err)
		return err
	}

	log.Info("Wipe complete for device %s", device)
	syscall.Sync()
	return nil
}
