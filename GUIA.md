# Guía de Resolución - Go Koans 📖

Esta guía te ayudará a navegar por los Go Koans. **Recuerda:** el objetivo es que **TÚ** implementes las soluciones, no copiarlas.

## 🎯 Filosofía de Aprendizaje

**Estos koans NO contienen las soluciones.** En su lugar encontrarás:
- ✅ Explicaciones de **cómo funciona Go** comparado con Python
- ✅ **Hints y pistas** sobre qué hacer
- ✅ **TODOs claros** que indican qué implementar
- ✅ **Tests** que describen el comportamiento esperado

**Tu trabajo es llenar los espacios en blanco** usando las pistas provistas.

## 🚀 Orden Recomendado

Completa los koans en este orden para una curva de aprendizaje óptima:

1. **01_about_variables** - Fundamentos: tipos, declaraciones, constantes
2. **02_about_functions** - Funciones, parámetros, retornos múltiples  
3. **03_about_structs** - Estructuras (las "clases" de Go)
4. **04_about_interfaces** - Polimorfismo en Go
5. **05_about_pointers** - Gestión de memoria y referencias
6. **06_about_errors** - Manejo de errores (muy importante en Go!)
7. **07_about_goroutines** - Concurrencia con goroutines
8. **08_about_channels** - Comunicación entre goroutines
9. **09_about_packages** - Organización y exports

## 🔥 Ejemplo Práctico Paso a Paso

### Paso 1: Navega al primer koan
```bash
cd 01_about_variables
```

### Paso 2: Ejecuta los tests para ver qué falla (esto es esperado!)
```bash
go test -v
```

Verás output como:
```
--- FAIL: TestVariables (0.00s)
    variables_test.go:12: Expected "Hola Go", got "__FILL_ME__"
```

### Paso 3: Abre `variables.go` y lee el contexto
Verás código como:
```go
// PASO 1: Arregla las declaraciones de variables
// En Python: nombre = "Juan"
// En Go: var nombre string = "Juan" o nombre := "Juan"

var (
	// TODO: Declara una variable string llamada 'saludo' con valor "Hola Go"
	saludo = "__FILL_ME__"
```

### Paso 4: Analiza qué hacer
1. Lee el comentario de comparación Python/Go
2. Lee el TODO
3. Entiende qué se pide: una variable string con valor "Hola Go"
4. En Go, puedes hacer: `saludo = "Hola Go"`

### Paso 5: Implementa la solución
```go
var (
	// TODO: Declara una variable string llamada 'saludo' con valor "Hola Go"
	saludo = "Hola Go"  // ← Reemplazaste __FILL_ME__
```

### Paso 6: Ejecuta los tests nuevamente
```bash
go test -v
```

Si pasa 🟢:
```
--- PASS: TestVariables (0.00s)
```

### Paso 7: Continúa con el siguiente TODO
Repite el proceso hasta que todos los tests del koan pasen.

### Paso 8: Pasa al siguiente koan
```bash
cd ../02_about_functions
go test -v
```

## ⚠️ Reglas Importantes

1. **NO busques soluciones en internet** hasta que lo hayas intentado
2. **Lee los comentarios** - contienen toda la información necesaria
3. **Usa los hints** - están ahí para ayudarte
4. **Los tests son tu guía** - describen exactamente qué debe hacer el código
5. **Experimenta** - no pasa nada si te equivocas, los tests te dirán qué está mal

## 💡 Consejos para Desarrolladores Python

### Variables y Tipos
```python
# Python
nombre = "Juan"        # Tipo inferido
edad = 25             # Tipo inferido

# Go
var nombre string = "Juan"  // Tipo explícito
nombre := "Juan"           // Tipo inferido (solo en funciones)
```

### Funciones
```python
# Python
def saludar(nombre):
    return f"Hola {nombre}"

# Go
func Saludar(nombre string) string {
    return fmt.Sprintf("Hola %s", nombre)
}
```

### Clases vs Structs
```python
# Python
class Persona:
    def __init__(self, nombre, edad):
        self.nombre = nombre
        self.edad = edad
    
    def saludar(self):
        return f"Hola, soy {self.nombre}"

# Go
type Persona struct {
    Nombre string
    Edad   int
}

func (p Persona) Saludar() string {
    return fmt.Sprintf("Hola, soy %s", p.Nombre)
}
```

### Manejo de Errores
```python
# Python
try:
    resultado = dividir(10, 0)
except ZeroDivisionError as e:
    print(f"Error: {e}")

# Go
resultado, err := dividir(10, 0)
if err != nil {
    fmt.Printf("Error: %v\n", err)
}
```

### Concurrencia
```python
# Python
import asyncio

async def tarea():
    await asyncio.sleep(1)
    return "completado"

# Go
func tarea() {
    time.Sleep(1 * time.Second)
    fmt.Println("completado")
}

go tarea() // ¡Ejecuta en paralelo!
```

## 🛠️ Comandos Útiles

```bash
# Ejecutar tests de un koan específico
go test ./01_about_variables -v

# Ejecutar todos los tests (ver progreso general)
go test ./... -v

# Formatear código automáticamente
go fmt ./...

# Verificar problemas en el código
go vet ./...

# Ver cobertura de tests
go test -cover ./...
```

## 🎯 Objetivos de Aprendizaje por Koan

- **Variables**: Declaración, tipos, constantes, zero values
- **Functions**: Parámetros, returns múltiples, closures, defer
- **Structs**: Composición, métodos, embedido
- **Interfaces**: Duck typing de Go, type assertions
- **Pointers**: Referencias vs valores, memory management
- **Errors**: Patrón error en Go, custom errors, wrapping
- **Goroutines**: Concurrencia, sync.WaitGroup, race conditions
- **Channels**: Comunicación, buffered, select, patterns
- **Packages**: Organización, visibilidad, imports

## 🏆 ¡Que disfrutes el viaje!

Recuerda: cada error es una oportunidad de aprender. Go tiene sus peculiaridades, pero una vez que las domines, amarás su simplicidad y poder.

**¡Feliz coding! 🚀**
