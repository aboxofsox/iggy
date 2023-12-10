$URL = 'https://api.github.com/repos/aboxofsox/iggy/releases/latest'

$RootDir = Join-Path -Path $env:USERPROFILE -ChildPath 'iggy'

if (!(Test-Path -Path $RootDir)) {
    New-Item -Path $RootDir -ItemType Directory
}

$LatestRelease = Invoke-RestMethod -Uri $URL

$Assets = $LatestRelease.$Assets

foreach ($Asset in $Assets) {
    $Name = $Asset.name
    $URL = $Asset.browser_download_url

    $Output = Join-Path -Path $RootDir -ChildPath $Name

    Write-Host "Downloading $Name"

    Invoke-WebRequest -Uri $URL -OutFile $Output
}