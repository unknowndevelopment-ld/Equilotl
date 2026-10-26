/*
 * SPDX-License-Identifier: GPL-3.0
 * Vencord Installer, a cross platform gui/cli app for installing Vencord
 * Copyright (c) 2023 Vendicated and Vencord contributors
 */

package main

import (
	"equilotl/buildinfo"
	"image/color"
)

const ReleaseUrl = "https://api.github.com/repos/unknowndevelopment-ld/Equicord/releases/latest"
const ReleaseUrlFallback = "https://api.github.com/repos/unknowndevelopment-ld/Equicord/releases?per_page=1"
const InstallerReleaseUrl = "https://api.github.com/repos/unknowndevelopment-ld/Equilotl/releases/latest"
const InstallerReleaseUrlFallback = "https://api.github.com/repos/unknowndevelopment-ld/Equilotl/releases?per_page=1"

var UserAgent = "Hellocord/" + buildinfo.InstallerGitHash + " (https://github.com/unknowndevelopment-ld/Equilotl)"

const SupportUrl = "https://github.com/unknowndevelopment-ld/Equicord/issues"

var (
	HellocordRed        = color.RGBA{190, 24, 42, 0xff}
	HellocordRedHovered = color.RGBA{225, 35, 54, 0xff}
	HellocordDark       = color.RGBA{25, 25, 29, 0xff}
	HellocordDarkHover  = color.RGBA{42, 42, 48, 0xff}
	HellocordAlert      = color.RGBA{105, 18, 29, 0xff}
)

var LinuxDiscordNames = []string{
	"Discord",
	"DiscordPTB",
	"DiscordCanary",
	"DiscordDevelopment",
	"discord",
	"discordptb",
	"discordcanary",
	"discorddevelopment",
	"discord-ptb",
	"discord-canary",
	"discord-development",
	// Flatpak
	"com.discordapp.Discord",
	"com.discordapp.DiscordPTB",
	"com.discordapp.DiscordCanary",
	"com.discordapp.DiscordDevelopment",
}
