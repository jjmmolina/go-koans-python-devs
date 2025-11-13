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

**🎓 Filosofía de Aprendizaje:**
- ❌ **NO** te damos las soluciones directamente
- ✅ Te damos **tests** que describen el comportamiento esperado
- ✅ Te damos **pistas y hints** sobre cómo resolverlo en Go
- ✅ Te mostramos **comparaciones con Python** para facilitar la comprensión
- ✅ Tú **implementas** la solución siguiendo los TODOs

**Este es un viaje de descubrimiento, no de copiar y pegar.** 🚀

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

## 🚀 Inicio Rápido

1. **Clona el repositorio**:
   ```bash
   git clone https://github.com/jjmmolina/go-koans-python-devs.git
   cd go-koans-python-devs
   ```

2. **Ejecuta todos los tests** para ver cuántos fallan:
   ```bash
   go test ./...
   ```
   Verás muchos tests rojos 🔴 - ¡esto es lo esperado!

3. **Empieza con el primer koan**:
   ```bash
   cd 01_about_variables
   go test -v
   ```

4. **Abre** `variables.go` y busca los TODOs:
   - Lee las comparaciones con Python
   - Sigue las pistas en los comentarios
   - Reemplaza `"__FILL_ME__"` y valores `0`, `false` por las soluciones correctas

5. **Ejecuta el test** hasta que pase 🟢:
   ```bash
   go test -v
   ```

6. **Repite** con cada koan siguiendo el orden numérico

7. **Lee la [GUIA.md](GUIA.md)** para un ejemplo paso a paso completo

## 📖 Cómo Usar los Koans

Cada archivo `.go` contiene:

```go
// PASO 1: Explicación del concepto
// En Python: ejemplo_python()
// En Go: ejemplo_go()

// TODO: Descripción clara de qué hacer
func MiFuncion() string {
    return "__FILL_ME__"  // ← Reemplaza esto
}
```

**Tu trabajo:**
1. Lee el comentario de comparación Python/Go
2. Lee el TODO
3. Consulta los hints si los hay
4. Implementa la solución
5. Ejecuta `go test` para verificar

**NO mires las soluciones en internet hasta que lo intentes primero!** 💪

## 📚 Estructura del Proyecto

```
go-koans/
├── 01_about_variables/      # Variables, tipos, constantes, conversiones
│   ├── variables.go         # ← Edita este archivo
│   └── variables_test.go    # ← Los tests que deben pasar
├── 02_about_functions/      # Funciones, parámetros, closures, errores
├── 03_about_structs/        # Structs, métodos, composición
├── 04_about_interfaces/     # Interfaces, duck typing, type assertions
├── 05_about_pointers/       # Punteros, referencias, memoria
├── 06_about_errors/         # Manejo de errores, panic/recover
├── 07_about_goroutines/     # Concurrencia, WaitGroups, Mutex
├── 08_about_channels/       # Channels, select, patrones de concurrencia
└── 09_about_packages/       # Organización, exports, imports
```

**Orden recomendado:** Sigue el orden numérico (01 → 09) ya que cada koan construye sobre los anteriores.

## 💡 Consejos para Desarrolladores Python

| Concepto | Python | Go |
|----------|--------|-----|
| **Clases** | `class Person:` | `type Person struct {}` |
| **Métodos** | `def method(self):` | `func (p Person) Method() {}` |
| **Herencia** | `class Child(Parent):` | Composición con embedding |
| **Excepciones** | `try/except` | Retornar `error` como segundo valor |
| **Async** | `async/await` | `goroutines` y `channels` |
| **None** | `None` | `nil` |
| **Duck Typing** | Implícito | Interfaces explícitas |
| **List** | `[1, 2, 3]` | `[]int{1, 2, 3}` (slices) |

**Diferencias clave:**
- 🔸 Go es **tipado estáticamente**: debes declarar tipos
- 🔸 Go no tiene **clases**, usa structs + métodos
- 🔸 Go maneja **errores como valores**, no excepciones
- 🔸 Go tiene **punteros explícitos**, Python tiene todo por referencia
- 🔸 Goroutines son **más ligeras** que threads de Python

## 📖 Documentación Adicional

- **[COMO_FUNCIONA.md](COMO_FUNCIONA.md)** - Explicación detallada de la metodología de koans
- **[GUIA.md](GUIA.md)** - Tutorial paso a paso con ejemplos
- **[CONTRIBUTING.md](CONTRIBUTING.md)** - Cómo contribuir al proyecto

## 🛠️ Comandos Útiles

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
