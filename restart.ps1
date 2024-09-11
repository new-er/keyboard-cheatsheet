fyne build

# Check if main.exe is running
$process = Get-Process -Name "main" -ErrorAction SilentlyContinue

# If the process is running, kill it
if ($process) {
    Stop-Process -Name "main" -Force
    Write-Output "main.exe has been stopped."
}

# Start main.exe
Start-Process -FilePath "main.exe"
Write-Output "main.exe has been restarted."

