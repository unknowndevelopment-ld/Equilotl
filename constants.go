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
	HellocordBlue         = color.RGBA{39, 105, 210, 0xff}
	HellocordBlueHovered  = color.RGBA{29, 85, 184, 0xff}
	HellocordWhite        = color.RGBA{255, 255, 255, 0xff}
	HellocordWhiteHover   = color.RGBA{231, 241, 255, 0xff}
	HellocordSurface      = color.RGBA{243, 247, 252, 0xff}
	HellocordSurfaceHover = color.RGBA{225, 237, 251, 0xff}
	HellocordText         = color.RGBA{31, 45, 65, 0xff}
	HellocordAlert        = color.RGBA{224, 238, 255, 0xff}
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
