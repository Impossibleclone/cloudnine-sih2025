package wipe

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"
    // "runtime"
    "time"
	"syscall"

    "sih2025/pkg/log"
)

type WipeCertificate struct {
    Device     string    `json:"device"`
    Passes     int       `json:"passes"`
    StartTime  time.Time `json:"start_time"`
    EndTime    time.Time `json:"end_time"`
    Duration   string    `json:"duration"`
    Platform   string    `json:"platform"`
    Method     string    `json:"method"`
    Signature  string    `json:"signature"`
    PublicKey  string    `json:"public_key"`
    Standards  []string  `json:"standards"`
}

func GenerateCertificate(device string, passes int, duration time.Duration, platform string) *WipeCertificate {
    cert := &WipeCertificate{
        Device:    device,
        Passes:    passes,
        StartTime: time.Now().Add(-duration),
        EndTime:   time.Now(),
        Duration:  duration.String(),
        Platform:  platform,
        Method:    "Multi-pass overwrite (NIST 800-88 compliant)",
        Standards: []string{"NIST SP 800-88"},
    }

    privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        log.Warn("Failed to generate keys for demo: %v", err)
        return cert
    }

    data, _ := json.Marshal(cert)
    hash := sha256.Sum256(data)
    sig, err := ecdsa.SignASN1(rand.Reader, privKey, hash[:])
    if err != nil {
        log.Warn("Signing failed: %v", err)
    } else {
        cert.Signature = hex.EncodeToString(sig)
        pubKeyBytes := elliptic.Marshal(elliptic.P256(), privKey.PublicKey.X, privKey.PublicKey.Y)
        cert.PublicKey = hex.EncodeToString(pubKeyBytes)
    }

    return cert
}

func choosemethod(device string) error {
	log.Info("Wiping device %s", device)
	err := eraseNvme(device)
	if err != nil {
		return err
	}
	err = eraseSata(device)
	if err != nil {
		return err
	}
	fmt.Println("Wipe complete")
	syscall.Sync()
	return nil
}

