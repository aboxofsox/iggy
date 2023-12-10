$Platforms = 'windows/amd64', 'windows/386', 'linux/amd64', 'linux/386', 'darwin/amd64'

$i = 0
foreach ($Platform in $Platforms) {
    $GOOS = $Platform.Split('/')[0]
    $Arch = $Platform.Split('/')[1]

    $env:GOOS = $GOOS
    $env:GOARCH = $Arch

    $Output = "bin/$GOOS/$Arch/iggy"

    $PercentComplete = [math]::Round($i/$Platforms.Length * 100)

    Write-Progress -Activity 'Building iggy' -Status "Building for $GOOS/$Arch" -PercentComplete $PercentComplete

    if ($env:GOOS -eq 'windows') {
        $Output += '.exe'
    }


    go build -o $Output

    if ($LASTEXITCODE -ne 0) {
        Write-Host "Build failed"
        exit $LASTEXITCODE
    }
}