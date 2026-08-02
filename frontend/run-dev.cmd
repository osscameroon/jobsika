@echo off
cd /d "%~dp0"
set BASE_URL=http://localhost:7000
node node_modules\nuxt\bin\nuxt.js --hostname 127.0.0.1 --port 3000
