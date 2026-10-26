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
const InstallerReleaseUrl = "https://api.github.com/repos/unknowndevelopment-ld/Equilotl/releases?per_page=1"
const InstallerReleaseUrlFallback = "https://api.github.com/repos/unknowndevelopment-ld/Equilotl/releases?per_page=1"

var UserAgent = "Hellocord/" + buildinfo.InstallerGitHash + " (https://github.com/unknowndevelopment-ld/Equilotl)"

const SupportUrl = "https://github.com/unknowndevelopment-ld/Equicord/issues"

var (
	HellocordBlue         = color.RGBA{39, 105, 210, 0xff}
	HellocordBlueHovered  = color.RGBA{29, 85, 184, 0xff}
	HellocordWhite        = color.RGBA{20, 20, 24, 0xff}
	HellocordWhiteHover   = color.RGBA{36, 39, 46, 0xff}
	HellocordSurface      = color.RGBA{29, 30, 34, 0xff}
	HellocordSurfaceHover = color.RGBA{45, 48, 56, 0xff}
	HellocordText         = color.RGBA{240, 242, 245, 0xff}
	HellocordAlert        = color.RGBA{35, 43, 58, 0xff}
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
