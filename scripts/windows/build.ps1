$Platforms = 'windows/amd64', 'windows/386', 'linux/amd64', 'linux/386', 'darwin/amd64'

if (!(Test-Path -Path 'bin')) {
    New-Item -Path 'bin' -ItemType Directory
}

$GitTag = git describe --tags --abbrev=0
$Version = Get-Content -Path 'version.txt'

if (Compare-Versions -Version1 $GitTag -Version2 $Version -eq 0) {
    Write-Host 'No new version'
    exit 0
}

switch (Compare-Versions -Version1 $GitTag -Version2 $Version) {
    1 {
        Write-Host "New version $GitTag"
    }
    -1 {
        Write-Host "Version $Version is newer than $GitTag"
        exit 0
    }
}

$i = 0
foreach ($Platform in $Platforms) {
    $GOOS = $Platform.Split('/')[0]
    $Arch = $Platform.Split('/')[1]

    $env:GOOS = $GOOS
    $env:GOARCH = $Arch

    $Output = "bin/$GOOS/$Arch/iggy-$GOOS-$ARCH-$GitTag"

    $PercentComplete = [math]::Round($i/$Platforms.Length * 100)

    Write-Progress -Activity 'Building iggy' -Status "Building for $GOOS/$Arch" -PercentComplete $PercentComplete

    if ($env:GOOS -eq 'windows') {
        $Output += '.exe'
    }

    go build -ldflags "-X main.Version=$GitTag" -o $Output

    if ($LASTEXITCODE -ne 0) {
        Write-Host "Build failed"
        exit $LASTEXITCODE
    }

    $i++ 
}

Remove-Item -Path 'bin/checksums.json' -ErrorAction SilentlyContinue

$Checksums = @()

Get-ChildItem -Path 'bin' -Recurse -File | ForEach-Object {
    $Checksums += [PSCustomObject]@{
        Path = $_.Name
        Checksum = Get-FileHash -Path $_.FullName -Algorithm MD5 | Select-Object -Expand Hash
    }
}

$Checksums | ConvertTo-Json | Out-File 'bin/checksums.json'