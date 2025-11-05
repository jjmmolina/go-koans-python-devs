# Go Koans - Aprende Go con TDD 🚀

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

¡Bienvenido a Go Koans! Un tutorial interactivo para aprender Go usando **Test-Driven Development (TDD)**, especialmente diseñado para **desarrolladores que ya conocen Python**.

## ✨ ¿Qué son los Koans?

Los Koans son ejercicios de programación que siguen el ciclo TDD:

```
🔴 Red → 🟢 Green → 🔵 Refactor
```

1. **🔴 Red**: Un test falla (punto de partida)
2. **🟢 Green**: Escribes el código mínimo para que pase
3. **🔵 Refactor**: Mejoras el código manteniendo los tests verdes

## 🎯 ¿Para quién es esto?

✅ Desarrolladores Python que quieren aprender Go  
✅ Personas que prefieren aprender haciendo  
✅ Quienes buscan entender Go mediante comparaciones con Python  
✅ Desarrolladores que quieren dominar TDD en Go  

## 📋 Requisitos

- Go 1.21 o superior ([Descargar](https://go.dev/dl/))
- VS Code (recomendado) con extensión de Go
- Conocimientos básicos de Python
- Ganas de aprender 😊

## 🚀 Inicio Rápido

1. **Ejecuta todos los tests** para ver los fallos:
   ```bash
   go test ./...
   ```

2. **Empieza con el primer koan**:
   ```bash
   go test ./01_about_variables -v
   ```

3. **Edita** `01_about_variables/variables.go` y reemplaza `"__FILL_ME__"` con los valores correctos

4. **Re-ejecuta** hasta que los tests pasen, luego continúa con el siguiente

5. **Lee la [GUIA.md](GUIA.md)** para consejos detallados y ejemplos

## Cómo empezar

1. Ejecuta todos los tests para ver los fallos:
   ```bash
   go test ./...
   ```

2. Ve al primer koan (`01_about_variables`) y ejecuta sus tests:
   ```bash
   go test ./01_about_variables
   ```

3. Edita el archivo para hacer que el primer test pase
4. Continúa con el siguiente test que falle
5. Repite hasta completar todos los koans

## Estructura del proyecto

```
go-koans/
├── 01_about_variables/      # Variables y tipos básicos ✅
├── 02_about_functions/      # Funciones y parámetros ✅
├── 03_about_structs/        # Estructuras y métodos ✅
├── 04_about_interfaces/     # Interfaces y polimorfismo ✅
├── 05_about_pointers/       # Punteros y gestión de memoria ✅
├── 06_about_errors/         # Manejo de errores ✅
├── 07_about_goroutines/     # Concurrencia con goroutines ✅
├── 08_about_channels/       # Comunicación con channels ✅
└── 09_about_packages/       # Paquetes y módulos ✅
```

## Consejos para desarrolladores Python

- En Go no hay clases, usa `structs` e `interfaces`
- Go es tipado estáticamente, declara tipos explícitamente
- Go maneja memoria automáticamente, pero puedes usar punteros
- Las goroutines son como async/await pero más poderosas
- Los channels son como las queues de Python pero integradas en el lenguaje

## Comandos útiles

```bash
# Ejecutar todos los tests
go test ./...

# Ejecutar tests de un koan específico
go test ./01_about_variables

# Ver coverage
go test -cover ./...

# Formatear código
go fmt ./...

# Verificar código
go vet ./...
```

## Usando VS Code

Este proyecto incluye tareas de VS Code configuradas:
- **Ctrl+Shift+P** → "Tasks: Run Task" → "Go: Test All Koans"
- La extensión de Go te ayudará con autocompletado y diagnósticos
- Usa **F5** para depurar código Go

## 🤝 Contribuir

¡Las contribuciones son bienvenidas! Lee [CONTRIBUTING.md](CONTRIBUTING.md) para más detalles.

## 📄 Licencia

Este proyecto está bajo la Licencia MIT - ver [LICENSE](LICENSE) para más detalles.

## 🌟 Agradecimientos

Inspirado por:
- [Ruby Koans](https://github.com/edgecase/ruby_koans)
- [Go by Example](https://gobyexample.com/)
- La increíble comunidad de Go

---

**¿Te resultó útil? ¡Dale una ⭐ al proyecto!**

¡Que disfrutes aprendiendo Go! 🚀
