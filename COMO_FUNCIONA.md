# ¿Cómo Funcionan Estos Koans? 🤔

## Concepto Principal

Este proyecto sigue el **verdadero patrón Koan**: no se te dan las respuestas, sino que **tú las descubres** con ayuda de:
- Tests que describen el comportamiento esperado
- Comentarios que explican cómo funciona Go
- Comparaciones con Python para facilitar la comprensión
- Hints específicos cuando necesitas usar funciones o patrones especiales

## Anatomía de un Koan

### 1. El Archivo de Código (`*.go`)

```go
// Package variables introduces basic Go variables and types
// Comparado con Python, Go es tipado estáticamente
package variables

// PASO 1: Arregla las declaraciones de variables
// En Python: nombre = "Juan"
// En Go: var nombre string = "Juan" o nombre := "Juan"

var (
	// TODO: Declara una variable string llamada 'saludo' con valor "Hola Go"
	saludo = "__FILL_ME__"
```

**Componentes:**
- **Comentario de comparación**: Muestra equivalente Python vs Go
- **TODO**: Instrucción clara de qué implementar
- **Placeholder**: `"__FILL_ME__"`, `0`, `false`, `nil` - valores que debes reemplazar

### 2. El Archivo de Tests (`*_test.go`)

```go
func TestVariablesBasicas(t *testing.T) {
	if saludo != "Hola Go" {
		t.Errorf("saludo debe ser 'Hola Go', pero obtuvimos '%s'", saludo)
	}
}
```

**El test te dice:**
- ✅ Qué valor se espera (`"Hola Go"`)
- ✅ Qué obtuviste realmente (`'__FILL_ME__'`)
- ✅ Si tu código es correcto (PASS) o no (FAIL)

## Ciclo de Trabajo TDD

```
┌─────────────────────────────────────┐
│  1. Ejecutar Tests (go test -v)    │
│     ↓                               │
│  🔴 Ver qué falla                   │
└─────────────────────────────────────┘
           ↓
┌─────────────────────────────────────┐
│  2. Leer el código                  │
│     - Comentario Python/Go          │
│     - TODO                          │
│     - Hints si los hay              │
└─────────────────────────────────────┘
           ↓
┌─────────────────────────────────────┐
│  3. Pensar en la solución           │
│     - ¿Cómo se haría en Python?    │
│     - ¿Cómo dice que se hace en Go? │
│     - ¿Qué pide el TODO?            │
└─────────────────────────────────────┘
           ↓
┌─────────────────────────────────────┐
│  4. Implementar                     │
│     Reemplazar placeholders         │
└─────────────────────────────────────┘
           ↓
┌─────────────────────────────────────┐
│  5. Ejecutar Tests de nuevo         │
│     ↓                               │
│  🟢 ¿Pasa? → Siguiente TODO        │
│  🔴 ¿Falla? → Revisar código       │
└─────────────────────────────────────┘
```

## Ejemplos de Tipos de TODOs

### Tipo 1: Valores Simples
```go
// TODO: Declara una variable string llamada 'saludo' con valor "Hola Go"
saludo = "__FILL_ME__"

// Solución:
saludo = "Hola Go"
```

### Tipo 2: Expresiones
```go
// TODO: Retorna a * b
func Multiplicar(a, b int) int {
	return 0
}

// Solución:
func Multiplicar(a, b int) int {
	return a * b
}
```

### Tipo 3: Con Hints
```go
// TODO: Convierte 'entero' a string
// Hint: necesitarás importar "strconv" y usar strconv.Itoa(entero)
var texto string = "__FILL_ME__"

// Solución:
import "strconv"
var texto string = strconv.Itoa(entero)
```

### Tipo 4: Código Completo (Comentado)
```go
// TODO: Usa defer para cambiar resultado a "final" al terminar la función
// defer func() { resultado = "final" }()

// Solución: Descomenta y coloca en el lugar correcto
defer func() { resultado = "final" }()
```

## Niveles de Dificultad

### 🟢 Fácil (01-03)
- Reemplazar valores directos
- Sintaxis básica
- Conceptos fundamentales

### 🟡 Medio (04-06)
- Implementar funciones completas
- Entender nuevos conceptos (interfaces, punteros, errores)
- Pensar en la lógica

### 🔴 Avanzado (07-09)
- Patrones de concurrencia
- Múltiples conceptos combinados
- Diseño de soluciones más complejas

## Consejos para el Éxito

1. **Lee TODO el código antes de empezar**: Los comentarios tienen información valiosa

2. **Los tests son tu amigo**: Te dicen exactamente qué está mal

3. **Usa los hints**: Si un TODO menciona un paquete o función, úsalo

4. **Compara con Python**: Usa tu conocimiento existente como base

5. **Experimenta**: No pasa nada si te equivocas, los tests te guiarán

6. **NO busques soluciones en internet**: Intenta resolver con la información provista

7. **Si te atascas**: 
   - Relee los comentarios
   - Revisa el error del test
   - Prueba en un playground de Go
   - Consulta la documentación oficial de Go
   - Como último recurso, consulta [SOLUCIONES.md](SOLUCIONES.md) (pero intenta al menos 30 minutos primero)

## Verificar tu Progreso

```bash
# Ver cuántos tests pasan actualmente
./check_progress.ps1  # Windows PowerShell
# o
./check_progress.sh   # Linux/Mac

# Ejecutar un koan específico
cd 01_about_variables
go test -v

# Ejecutar todos los koans
go test ./...
```

## Cuando Termines un Koan

✅ Todos los tests en verde
✅ Entiendes por qué funciona tu código
✅ Puedes explicar la diferencia con Python

→ **¡Pasa al siguiente!**

---

**Recuerda**: El objetivo no es completar los koans rápido, sino **entender Go profundamente**. Tómate tu tiempo. 🧘
