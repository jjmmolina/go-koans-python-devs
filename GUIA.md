# Guía de Resolución - Go Koans

Esta guía te ayudará a empezar con los Go Koans. Recuerda: ¡el objetivo es aprender paso a paso!

## 🚀 Orden Recomendado

Completa los koans en este orden para una curva de aprendizaje óptima:

1. **01_about_variables** - Fundamentos básicos
2. **02_about_functions** - Funciones y métodos  
3. **03_about_structs** - Estructuras (las "clases" de Go)
4. **04_about_interfaces** - Polimorfismo en Go
5. **05_about_pointers** - Gestión de memoria
6. **06_about_errors** - Manejo de errores (muy importante!)
7. **07_about_goroutines** - Concurrencia básica
8. **08_about_channels** - Comunicación entre goroutines
9. **09_about_packages** - Organización de código

## 🔥 Cómo Empezar (Ejemplo Práctico)

### Paso 1: Ve al primer koan
```bash
cd 01_about_variables
```

### Paso 2: Ejecuta los tests para ver qué falla
```bash
go test -v
```

### Paso 3: Abre `variables.go` y busca los TODOs
Verás líneas como:
```go
// TODO: Declara una variable string llamada 'saludo' con valor "Hola Go"
saludo = "__FILL_ME__"
```

### Paso 4: Reemplaza `"__FILL_ME__"` y `0` con los valores correctos
```go
saludo = "Hola Go"  // ✅ Correcto
```

### Paso 5: Ejecuta los tests nuevamente
```bash
go test -v
```

### Paso 6: Continúa hasta que todos los tests pasen
¡Luego pasa al siguiente koan!

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
