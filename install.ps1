$link = "https://github.com/unknowndevelopment-ld/Equilotl/releases/latest/download/EquilotlCli.exe"

$outfile = "$env:TEMP\Hellocord-CLI.exe"

Write-Output "Downloading installer to $outfile"

Invoke-WebRequest -Uri "$link" -OutFile "$outfile"

Write-Output ""

Start-Process -Wait -NoNewWindow -FilePath "$outfile"

# Cleanup
Remove-Item -Force "$outfile"
