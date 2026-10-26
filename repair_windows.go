//go:build windows && !cli

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const discordInstallerURL = "https://discord.com/api/download?platform=win"

func repairDiscordInstall(di *DiscordInstall) error {
	// Stop Discord before invoking its installer so it can replace the damaged build.
	PreparePatch(di)

	client := &http.Client{Timeout: 10 * time.Minute}
	res, err := client.Get(discordInstallerURL)
	if err != nil {
		return fmt.Errorf("download official Discord installer: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("download official Discord installer: %s", res.Status)
	}
	const maxInstallerSize = 250 << 20
	if res.ContentLength > maxInstallerSize {
		return fmt.Errorf("official Discord installer is unexpectedly large (%d bytes)", res.ContentLength)
	}

	file, err := os.CreateTemp("", "Hellocord-Discord-Setup-*.exe")
	if err != nil {
		return fmt.Errorf("create temporary installer file: %w", err)
	}
	installerPath := file.Name()
	size, copyErr := io.Copy(file, io.LimitReader(res.Body, maxInstallerSize+1))
	closeErr := file.Close()
	if copyErr != nil {
		_ = os.Remove(installerPath)
		return fmt.Errorf("save official Discord installer: %w", copyErr)
	}
	if closeErr != nil {
		_ = os.Remove(installerPath)
		return fmt.Errorf("finish saving official Discord installer: %w", closeErr)
	}
	if size < 1<<20 || size > maxInstallerSize {
		_ = os.Remove(installerPath)
		return fmt.Errorf("downloaded installer has an invalid size (%d bytes)", size)
	}
	if res.ContentLength >= 0 && size != res.ContentLength {
		_ = os.Remove(installerPath)
		return fmt.Errorf("download was incomplete: got %d of %d bytes", size, res.ContentLength)
	}

	if err := exec.Command(installerPath).Start(); err != nil {
		_ = os.Remove(installerPath)
		return fmt.Errorf("start official Discord installer: %w", err)
	}
	return nil
}
