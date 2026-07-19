Set-Location -LiteralPath $PSScriptRoot
& "C:\Program Files\nodejs\npx.cmd" nuxt --hostname 127.0.0.1 --port 3000 *>&1 | Tee-Object -FilePath (Join-Path $PSScriptRoot "nuxt-live.log")
