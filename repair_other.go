//go:build !windows && !cli

package main

import "errors"

func repairDiscordInstall(_ *DiscordInstall) error {
	return errors.New("automatic Discord repair is currently supported only on Windows")
}
