#!/bin/bash
# Script para verificar progreso en los Go Koans

echo "🔍 Verificando progreso en Go Koans..."
echo ""

total_koans=9
passed=0
failed=0

for i in $(seq 1 9); do
    dir=$(printf "%02d_about_" $i)
    
    case $i in
        1) dir="${dir}variables" ;;
        2) dir="${dir}functions" ;;
        3) dir="${dir}structs" ;;
        4) dir="${dir}interfaces" ;;
        5) dir="${dir}pointers" ;;
        6) dir="${dir}errors" ;;
        7) dir="${dir}goroutines" ;;
        8) dir="${dir}channels" ;;
        9) dir="${dir}packages" ;;
    esac
    
    echo "Testing $dir..."
    
    if go test ./$dir -v > /dev/null 2>&1; then
        echo "  ✅ $dir - PASSED"
        ((passed++))
    else
        echo "  ❌ $dir - FAILED"
        ((failed++))
    fi
done

echo ""
echo "📊 Resumen:"
echo "  ✅ Completados: $passed/$total_koans"
echo "  ❌ Pendientes: $failed/$total_koans"
echo ""

percentage=$((passed * 100 / total_koans))
echo "📈 Progreso: $percentage%"

if [ $passed -eq $total_koans ]; then
    echo ""
    echo "🎉 ¡Felicidades! ¡Has completado todos los koans!"
    echo "🚀 Ahora estás listo para usar Go en proyectos reales."
fi
