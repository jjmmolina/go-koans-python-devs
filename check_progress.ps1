# Script para verificar progreso en los Go Koans (PowerShell)

Write-Host "🔍 Verificando progreso en Go Koans..." -ForegroundColor Cyan
Write-Host ""

$totalKoans = 9
$passed = 0
$failed = 0

$koans = @(
    "01_about_variables",
    "02_about_functions",
    "03_about_structs",
    "04_about_interfaces",
    "05_about_pointers",
    "06_about_errors",
    "07_about_goroutines",
    "08_about_channels",
    "09_about_packages"
)

foreach ($koan in $koans) {
    Write-Host "Testing $koan..." -NoNewline
    
    $result = go test "./$koan" -v 2>&1
    
    if ($LASTEXITCODE -eq 0) {
        Write-Host "  ✅ PASSED" -ForegroundColor Green
        $passed++
    } else {
        Write-Host "  ❌ FAILED" -ForegroundColor Red
        $failed++
    }
}

Write-Host ""
Write-Host "📊 Resumen:" -ForegroundColor Yellow
Write-Host "  ✅ Completados: $passed/$totalKoans" -ForegroundColor Green
Write-Host "  ❌ Pendientes: $failed/$totalKoans" -ForegroundColor Red
Write-Host ""

$percentage = [math]::Round(($passed / $totalKoans) * 100)
Write-Host "📈 Progreso: $percentage%" -ForegroundColor Cyan

if ($passed -eq $totalKoans) {
    Write-Host ""
    Write-Host "🎉 ¡Felicidades! ¡Has completado todos los koans!" -ForegroundColor Green
    Write-Host "🚀 Ahora estás listo para usar Go en proyectos reales." -ForegroundColor Cyan
}
