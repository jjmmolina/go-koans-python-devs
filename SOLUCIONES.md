# 🔓 Soluciones Completas de los Go Koans

> ⚠️ **ADVERTENCIA**: Este documento contiene las soluciones completas de todos los koans.
> 
> **Se recomienda fuertemente** intentar resolver los ejercicios por ti mismo antes de consultar estas soluciones.
> El aprendizaje real viene de la lucha y el descubrimiento, no de copiar respuestas.
>
> **Usa este documento solo cuando:**
> - Has intentado resolver un problema múltiples veces
> - Has consultado la documentación de Go
> - Quieres verificar tu solución
> - Estás completamente atascado después de 30+ minutos

---

## 📖 Índice

1. [01 - Variables](#01---variables)
2. [02 - Funciones](#02---funciones)
3. [03 - Structs](#03---structs)
4. [04 - Interfaces](#04---interfaces)
5. [05 - Punteros](#05---punteros)
6. [06 - Errores](#06---errores)
7. [07 - Goroutines](#07---goroutines)
8. [08 - Channels](#08---channels)
9. [09 - Packages](#09---packages)

---

## 01 - Variables

### 📚 Conceptos Clave

**Variables en Go:**
- Go es tipado estáticamente - debes declarar tipos explícitamente
- Tres formas de declarar variables:
  1. `var nombre tipo = valor` - forma completa
  2. `var nombre = valor` - tipo inferido
  3. `nombre := valor` - declaración corta (solo en funciones)
- Zero values: `0` para números, `""` para strings, `false` para bool

**Constantes:**
- Se declaran con `const`
- No pueden cambiar después de la declaración
- Pueden ser tipadas o no tipadas

### ✅ Soluciones

#### PASO 1: Variables Básicas

```go
var (
	saludo = "Hola Go"
	edad = 25
	esActivo = true
	precio = 19.99
)
```

**Explicación:**
- `saludo` es un string, necesita comillas dobles
- `edad` es un int, número entero sin decimales
- `esActivo` es bool, valores `true` o `false`
- `precio` es float64, número con decimales

#### PASO 2: Declaración Corta

```go
func DeclaracionCorta() (string, int, bool) {
	nombre := "Ana"
	puntos := 100
	completado := true
	
	return nombre, puntos, completado
}
```

**Explicación:**
- El operador `:=` declara e inicializa en una línea
- Solo funciona dentro de funciones
- El tipo se infiere automáticamente

#### PASO 3: Constantes

```go
const (
	PI_STRING = "3.14159"
	MAX_USUARIOS = 1000
)
```

**Explicación:**
- Las constantes usan `const` en lugar de `var`
- No pueden ser modificadas después de la declaración
- Por convención, las constantes se escriben en MAYÚSCULAS

#### PASO 4: Múltiples Asignaciones

```go
func MultiplesAsignaciones() (int, int, string) {
	x, y, mensaje := 10, 20, "múltiple"
	return x, y, mensaje
}
```

**Explicación:**
- Puedes declarar múltiples variables en una sola línea
- Similar a Python: `x, y, z = 1, 2, 3`

#### PASO 5: Zero Values

```go
func ZeroValues() (int, string, bool, float64) {
	var (
		numero  int     // 0
		texto   string  // ""
		activo  bool    // false
		decimal float64 // 0.0
	)
	
	return numero, texto, activo, decimal
}
```

**Explicación:**
- No necesitas cambiar nada aquí
- Go inicializa automáticamente las variables con "zero values"
- Diferente de Python donde variables no inicializadas causan error

#### PASO 6: Conversión de Tipos

```go
import "strconv"

func ConversionTipos() (int, float64, string) {
	var entero int = 42
	var decimal float64 = float64(entero)
	var texto string = strconv.Itoa(entero)
	
	return entero, decimal, texto
}
```

**Explicación:**
- `float64(entero)` convierte int a float64
- `strconv.Itoa(entero)` convierte int a string
- En Python: `float(42)` y `str(42)`
- Go requiere conversiones explícitas, no automáticas

---

## 02 - Funciones

### 📚 Conceptos Clave

**Funciones en Go:**
- Sintaxis: `func nombre(param tipo) tipoRetorno { ... }`
- Pueden retornar múltiples valores
- Valores de retorno pueden ser nombrados
- Las funciones son valores de primera clase (pueden asignarse a variables)

**Closures:**
- Funciones anónimas que capturan variables del scope exterior
- Similar a lambdas en Python

**Defer:**
- Ejecuta una función al final de la función contenedora
- Útil para cleanup (cerrar archivos, unlock mutex, etc.)

### ✅ Soluciones

#### PASO 1: Función Básica

```go
func Saludar(nombre string) string {
	return "Hola " + nombre
}
```

**Explicación:**
- Concatenación simple de strings
- Equivalente Python: `return f"Hola {nombre}"`

#### PASO 2: Múltiples Parámetros

```go
func Sumar(a int, b int) int {
	return a + b
}
```

**Explicación:**
- Cada parámetro necesita su tipo especificado
- El tipo de retorno se especifica después de los parámetros

#### PASO 3: Parámetros del Mismo Tipo

```go
func Multiplicar(a, b int) int {
	return a * b
}
```

**Explicación:**
- Cuando múltiples parámetros son del mismo tipo, puedes especificar el tipo una vez
- `(a, b int)` es equivalente a `(a int, b int)`

#### PASO 4: Múltiples Valores de Retorno

```go
func DividirConResto(dividendo, divisor int) (int, int) {
	return dividendo / divisor, dividendo % divisor
}
```

**Explicación:**
- Go puede retornar múltiples valores directamente
- `/` es división entera, `%` es módulo (resto)
- Equivalente Python: `return dividendo // divisor, dividendo % divisor`

#### PASO 5: Valores de Retorno Nombrados

```go
func Coordenadas() (x, y int) {
	x = 10
	y = 20
	return // return desnudo
}
```

**Explicación:**
- Puedes nombrar los valores de retorno en la firma
- Con `return` sin valores, retorna automáticamente las variables nombradas
- Útil para funciones con múltiples retornos

#### PASO 6: Función Variádica

```go
func SumarTodos(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
```

**Explicación:**
- `...int` permite pasar número variable de argumentos
- `range` itera sobre el slice
- `_` ignora el índice (solo queremos el valor)
- Equivalente Python: `def suma(*args): return sum(args)`

#### PASO 7: Función como Valor

```go
func InicializarOperacion() {
	OperacionMatematica = func(a, b int) int {
		return a * b
	}
}
```

**Explicación:**
- Las funciones son valores de primera clase
- Puedes asignar funciones anónimas a variables

#### PASO 8: Closures

```go
func CrearMultiplicador(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
```

**Explicación:**
- La función retornada "captura" el valor de `factor`
- Cada closure mantiene su propia copia de las variables capturadas
- Equivalente Python: `lambda x: x * factor`

#### PASO 9: Defer

```go
func EjemploDefer() string {
	resultado := "inicio"
	
	defer func() {
		resultado = "final"
	}()
	
	resultado = "medio"
	return resultado
}
```

**Explicación:**
- `defer` pospone la ejecución hasta que la función retorne
- Se ejecuta DESPUÉS del `return` pero ANTES de salir de la función
- Útil para cleanup, pero aquí modifica el valor de retorno

#### PASO 10: Manejo de Errores

```go
import (
	"fmt"
)

func DividirSeguro(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("división por cero")
	}
	return a / b, nil
}
```

**Explicación:**
- En Go, los errores son valores que se retornan
- Convencion: el error es el último valor de retorno
- `fmt.Errorf` crea un nuevo error con un mensaje
- `nil` significa "sin error"

---

## 03 - Structs

### 📚 Conceptos Clave

**Structs en Go:**
- Reemplazan a las clases de Python
- Agrupan datos relacionados
- Los métodos se definen fuera del struct

**Métodos:**
- Funciones asociadas a un tipo
- Receptor por valor: `(p Persona)` - no modifica el original
- Receptor por puntero: `(p *Persona)` - puede modificar el original

**Composición:**
- Go no tiene herencia, usa composición
- Embedding: incluir un struct dentro de otro

### ✅ Soluciones

#### PASO 1: Struct Básico

```go
type Persona struct {
	Nombre string
	Edad   int
}
```

**Explicación:**
- `type` define un nuevo tipo
- Los campos con mayúscula inicial son públicos (exportados)
- Equivalente Python: `class Persona: def __init__(self, nombre, edad): ...`

#### PASO 2: Constructor

```go
func NuevaPersona(nombre string, edad int) Persona {
	return Persona{
		Nombre: nombre,
		Edad:   edad,
	}
}
```

**Explicación:**
- Go no tiene constructores especiales como `__init__`
- Por convención, se crean funciones `Nueva...` o `New...`
- Retorna un struct inicializado

#### PASO 3: Métodos

```go
func (p Persona) Saludar() string {
	return "Hola, soy " + p.Nombre
}
```

**Explicación:**
- `(p Persona)` es el receptor - como `self` en Python
- El método se define fuera del struct
- Equivalente Python: `def saludar(self): return f"Hola, soy {self.nombre}"`

#### PASO 4: Métodos con Receptor Puntero

```go
func (p *Persona) CumplirAños() {
	p.Edad++
}
```

**Explicación:**
- `*Persona` es un puntero al struct
- Permite modificar el struct original
- Sin el `*`, solo modificarías una copia

#### PASO 5: Structs Anidados

```go
func NuevoEmpleado(nombre string, edad int, calle, ciudad string, salario float64) Empleado {
	return Empleado{
		Persona: Persona{
			Nombre: nombre,
			Edad:   edad,
		},
		Direccion: Direccion{
			Calle:  calle,
			Ciudad: ciudad,
		},
		Salario: salario,
	}
}
```

**Explicación:**
- Puedes incluir structs dentro de otros structs
- Cada struct anidado se inicializa por separado

#### PASO 6: Embedded Structs

```go
func NuevoEstudiante(nombre string, edad int, carrera string) Estudiante {
	return Estudiante{
		Persona: Persona{
			Nombre: nombre,
			Edad:   edad,
		},
		Carrera: carrera,
	}
}
```

**Explicación:**
- Embedding permite "heredar" campos y métodos
- `Estudiante` tiene acceso directo a `Nombre` y `Edad`
- También hereda el método `Saludar()`

#### PASO 7: Métodos de Structs Embebidos

```go
func (e Estudiante) InformacionCompleta() string {
	return e.Saludar() + " y estudio " + e.Carrera
}
```

**Explicación:**
- Puede usar `e.Saludar()` heredado de Persona
- También accede a `e.Carrera` del propio Estudiante
- Combina funcionalidad heredada con nueva

#### PASO 8: Receptores Mixtos

```go
func (c Contador) Obtener() int {
	return c.Valor
}

func (c *Contador) Incrementar() {
	c.Valor++
}

func (c *Contador) Resetear() {
	c.Valor = 0
}
```

**Explicación:**
- `Obtener()` usa receptor valor (solo lectura)
- `Incrementar()` y `Resetear()` usan receptor puntero (modifican)
- Regla: si algún método necesita puntero, úsalo en todos por consistencia

#### PASO 9: Inicialización

```go
func EjemplosInicializacion() (Persona, Persona, Persona) {
	// 1. Con campos nombrados (recomendado)
	p1 := Persona{
		Nombre: "Ana",
		Edad:   30,
	}
	
	// 2. Con orden de campos
	p2 := Persona{"Luis", 25}
	
	// 3. Zero value y asignación
	var p3 Persona
	p3.Nombre = "María"
	p3.Edad = 35
	
	return p1, p2, p3
}
```

**Explicación:**
- Tres formas de inicializar structs
- Forma 1 es la más clara y mantenible
- Forma 2 es más corta pero frágil (cambiar orden de campos rompe código)
- Forma 3 útil cuando inicializas condicionalmente

#### PASO 10: Comparación de Structs

```go
func CompararPersonas(p1, p2 Persona) bool {
	return p1 == p2
}
```

**Explicación:**
- Structs son comparables si todos sus campos son comparables
- `==` compara campo por campo
- No funciona con slices, maps o funciones

---

## 04 - Interfaces

### 📚 Conceptos Clave

**Interfaces en Go:**
- Definen comportamiento (conjunto de métodos)
- Implementación implícita (no hay palabra `implements`)
- Duck typing: "si camina como pato y grazna como pato, es un pato"

**Empty Interface:**
- `interface{}` o `any` (Go 1.18+)
- Acepta cualquier tipo
- Equivalente a `object` en Python

**Type Assertions:**
- Verificar y extraer el tipo concreto de una interface
- `valor.(Tipo)` o `valor, ok := interfaz.(Tipo)`

### ✅ Soluciones

#### PASO 1: Definir Interface

```go
type Hablador interface {
	Hablar() string
}
```

**Explicación:**
- Una interface es un contrato de métodos
- Cualquier tipo que implemente `Hablar()` satisface la interface
- No necesitas declarar explícitamente que implementas la interface

#### PASO 2: Implementar Interface

```go
func (p Perro) Hablar() string {
	return "Guau! Soy " + p.Nombre
}

func (g Gato) Hablar() string {
	return "Miau! Soy " + g.Nombre
}
```

**Explicación:**
- Solo necesitas implementar el método para satisfacer la interface
- Perro y Gato ahora son `Hablador` automáticamente
- Polimorfismo sin herencia

#### PASO 3: Usar Interfaces

```go
func HacerHablar(h Hablador) string {
	return h.Hablar()
}
```

**Explicación:**
- Puedes pasar cualquier tipo que implemente `Hablador`
- No importa si es Perro, Gato, o cualquier otro tipo
- Polimorfismo en acción

#### PASO 4: Interface Vacía

```go
func DescribirTipo(valor any) string {
	switch v := valor.(type) {
	case string:
		return fmt.Sprintf("Es un string: %s", v)
	case int:
		return fmt.Sprintf("Es un int: %d", v)
	case bool:
		return fmt.Sprintf("Es un bool: %t", v)
	default:
		return "Tipo desconocido"
	}
}
```

**Explicación:**
- `any` es alias de `interface{}` (Go 1.18+)
- Type switch permite verificar el tipo en runtime
- `.(type)` solo funciona en switch statements

#### PASO 5: Interface con Múltiples Métodos

```go
func (p Programador) Trabajar() string {
	return "Programando en " + p.Lenguaje
}

func (p Programador) DescansoAlmuerzo() string {
	return "Debugeando mientras como"
}
```

**Explicación:**
- Una interface puede requerir múltiples métodos
- Debes implementar TODOS los métodos para satisfacer la interface
- Programador implementa `Trabajador` porque tiene ambos métodos

#### PASO 6: Type Assertion

```go
func EsPerro(h Hablador) bool {
	_, esPerro := h.(Perro)
	return esPerro
}
```

**Explicación:**
- `h.(Perro)` intenta convertir la interface a Perro
- Retorna dos valores: el valor convertido y un bool
- Si es exitoso, `esPerro` es `true`
- Forma segura: `valor, ok := interfaz.(Tipo)`

#### PASO 7: Interface Embebida

```go
func (r Robot) Hablar() string {
	return "Bip boop! Soy " + r.Modelo
}

func (r Robot) Trabajar() string {
	return "Procesando datos..."
}

func (r Robot) DescansoAlmuerzo() string {
	return "Recargando baterías"
}
```

**Explicación:**
- `HabladorTrabajador` combina dos interfaces
- Robot debe implementar TODOS los métodos de ambas
- Total: 3 métodos (1 de Hablador + 2 de Trabajador)

#### PASO 8: Stringer Interface

```go
func (p Producto) String() string {
	return fmt.Sprintf("Producto: %s ($%.2f)", p.Nombre, p.Precio)
}
```

**Explicación:**
- `String()` es como `__str__` en Python
- `fmt.Stringer` es una interface estándar de Go
- `fmt.Println(producto)` llamará automáticamente a `String()`
- `%.2f` formatea float con 2 decimales

---

## 05 - Punteros

### 📚 Conceptos Clave

**Punteros en Go:**
- `&variable` obtiene la dirección de memoria
- `*puntero` dereferencia (obtiene el valor)
- `*Tipo` en parámetros indica un puntero
- Go tiene garbage collection - no necesitas liberar memoria

**¿Cuándo usar punteros?**
- Para modificar el valor original
- Para evitar copiar structs grandes
- Para compartir datos entre funciones

### ✅ Soluciones

#### PASO 1: Incrementar con Puntero

```go
func IncrementarConPuntero(x *int) {
	*x++
}
```

**Explicación:**
- `*x` dereferencia el puntero para obtener/modificar el valor
- Modifica la variable original, no una copia
- Equivalente Python: todo es por referencia por defecto

#### PASO 2: Incrementar sin Puntero (Falla)

```go
func IncrementarSinPuntero(x int) {
	x++
}
```

**Explicación:**
- Recibe una COPIA del valor
- Modificar `x` solo afecta la copia local
- El valor original no cambia
- Esto demuestra por qué necesitas punteros

#### PASO 3: Obtener Dirección

```go
func ObtenerDireccion(x int) *int {
	return &x
}
```

**Explicación:**
- `&x` retorna la dirección de memoria de `x`
- El tipo de retorno es `*int` (puntero a int)
- En Python no tienes acceso directo a direcciones de memoria

#### PASO 4: Dereferencia

```go
func DereferenciarPuntero(ptr *int) int {
	return *ptr
}
```

**Explicación:**
- `*ptr` obtiene el valor al que apunta el puntero
- Convierte `*int` (puntero) a `int` (valor)

#### PASO 5: Punteros con Structs

```go
func (p *Persona) CumplirAños() {
	p.Edad++
}

func (p Persona) ObtenerInfo() string {
	return fmt.Sprintf("%s tiene %d años", p.Nombre, p.Edad)
}
```

**Explicación:**
- Receptor puntero `*Persona` puede modificar
- Receptor valor `Persona` no puede modificar
- Go automáticamente dereferencia: `p.Edad` en lugar de `(*p).Edad`

#### PASO 6: Constructor con Puntero

```go
func NuevaPersona(nombre string, edad int) *Persona {
	return &Persona{
		Nombre: nombre,
		Edad:   edad,
	}
}
```

**Explicación:**
- `&Persona{...}` crea un struct y retorna su dirección
- Común en Go para "constructores"
- La memoria se gestiona automáticamente (garbage collection)

#### PASO 7: Modificar Slice

```go
func ModificarSlice(numeros *[]int) {
	*numeros = append(*numeros, 4)
}
```

**Explicación:**
- Slices ya son referencias, pero a veces necesitas el puntero
- `*numeros` dereferencia para acceder al slice
- Útil cuando quieres reasignar el slice completo

#### PASO 8: Comparar Punteros

```go
func CompararPunteros(p1, p2 *int) bool {
	return p1 == p2
}
```

**Explicación:**
- Compara las direcciones de memoria
- `true` solo si apuntan a la MISMA ubicación
- Para comparar valores: `*p1 == *p2`

#### PASO 9: Puntero Nulo

```go
func EsPunteroNulo(ptr *int) bool {
	return ptr == nil
}
```

**Explicación:**
- `nil` es el valor cero de los punteros
- Equivalente a `None` en Python
- Siempre verifica `nil` antes de dereferencia para evitar panic

#### PASO 10: Búsqueda con Punteros

```go
func BuscarPersona(nombre string, personas []*Persona) *Persona {
	for _, persona := range personas {
		if persona.Nombre == nombre {
			return persona
		}
	}
	return nil
}
```

**Explicación:**
- Retorna puntero o `nil` si no encuentra
- Patrón común en Go
- Verificar `nil` antes de usar el resultado

#### PASO 11: Contador con Puntero

```go
func NuevoContador() *Contador {
	inicial := 0
	return &Contador{valor: &inicial}
}

func (c *Contador) Incrementar() {
	*c.valor++
}

func (c *Contador) ObtenerValor() int {
	return *c.valor
}
```

**Explicación:**
- Contador guarda un puntero a un int
- `*c.valor` dereferencia para acceder/modificar
- Demuestra punteros dentro de structs

---

## 06 - Errores

### 📚 Conceptos Clave

**Manejo de Errores en Go:**
- Los errores son valores, no excepciones
- Patrón: retornar error como último valor
- Siempre verificar: `if err != nil`
- Errores personalizados implementan `Error() string`

**Panic/Recover:**
- `panic`: para errores irrecuperables (como excepciones)
- `recover`: captura panics (como try/except)
- Usar solo en situaciones excepcionales

### ✅ Soluciones

#### PASO 1: Error Básico

```go
func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("división por cero")
	}
	return a / b, nil
}
```

**Explicación:**
- `errors.New()` crea un error simple
- Retornar `nil` significa "sin error"
- Convencion: error es el último valor de retorno

#### PASO 2: Error Personalizado

```go
func (e ErrorEdadInvalida) Error() string {
	return fmt.Sprintf("edad inválida: %d", e.Edad)
}
```

**Explicación:**
- Implementar `Error()` hace que tu struct sea un `error`
- Puedes incluir datos adicionales (como `Edad`)
- Similar a crear una clase de excepción en Python

#### PASO 3: Validar con Error Personalizado

```go
func ValidarEdad(edad int) error {
	if edad < 0 || edad > 120 {
		return ErrorEdadInvalida{Edad: edad}
	}
	return nil
}
```

**Explicación:**
- Retorna tu error personalizado cuando falla la validación
- `nil` indica validación exitosa
- El caller puede verificar el tipo de error

#### PASO 4: Verificación de Errores

```go
func DividirSeguro(a, b float64) string {
	resultado, err := Dividir(a, b)
	if err != nil {
		return "Error: " + err.Error()
	}
	return fmt.Sprintf("%g", resultado)
}
```

**Explicación:**
- Siempre verificar `if err != nil` después de llamadas que retornan error
- `err.Error()` obtiene el mensaje de error
- Equivalente Python: try/except

#### PASO 5: Propagación de Errores

```go
func ConvertirYDividir(aStr, bStr string) (float64, error) {
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		return 0, err
	}
	
	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		return 0, err
	}
	
	return Dividir(a, b)
}
```

**Explicación:**
- Patrón común: verificar error, si existe, retornarlo inmediatamente
- Los errores se propagan hacia arriba
- Similar a re-raise en Python

#### PASO 6: Errores con Contexto

```go
func ConvertirConContexto(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("no se pudo convertir '%s' a entero: %w", str, err)
	}
	return num, nil
}
```

**Explicación:**
- `fmt.Errorf` con `%w` envuelve el error original
- Agrega contexto sin perder el error original
- `errors.Unwrap()` puede extraer el error envuelto

#### PASO 7: Error de Calculadora

```go
func Calcular(operacion string, a, b float64) (float64, error) {
	switch operacion {
	case "suma":
		return a + b, nil
	case "resta":
		return a - b, nil
	case "multiplicacion":
		return a * b, nil
	case "division":
		if b == 0 {
			return 0, ErrorCalculadora{"division", "división por cero"}
		}
		return a / b, nil
	default:
		return 0, ErrorCalculadora{operacion, "operación no soportada"}
	}
}
```

**Explicación:**
- Diferentes tipos de errores según el caso
- El struct `ErrorCalculadora` contiene información específica
- El caller puede inspeccionar el error

#### PASO 8: Defer para Cleanup

```go
func ProcesarArchivo(nombre string) error {
	fmt.Printf("Abriendo archivo: %s\n", nombre)
	
	defer fmt.Printf("Cerrando archivo: %s\n", nombre)
	
	if nombre == "archivo_malo.txt" {
		return errors.New("archivo corrupto")
	}
	
	fmt.Println("Procesando archivo...")
	return nil
}
```

**Explicación:**
- `defer` garantiza que el archivo se "cierre" incluso si hay error
- Similar a `try/finally` o context managers en Python
- El defer se ejecuta antes de retornar

#### PASO 9: Panic y Recover

```go
func FuncionPeligrosa(panico bool) (resultado string) {
	defer func() {
		if r := recover(); r != nil {
			resultado = "se recuperó del pánico"
		}
	}()
	
	if panico {
		panic("¡algo salió muy mal!")
	}
	
	return "todo bien"
}
```

**Explicación:**
- `panic` detiene la ejecución normal (como exception)
- `recover` captura el panic (como except)
- Solo funciona dentro de `defer`
- Valor de retorno nombrado permite modificarlo en defer

#### PASO 10: Verificar Tipo de Error

```go
func ManejarErrorEspecifico(err error) string {
	var edadErr ErrorEdadInvalida
	if errors.As(err, &edadErr) {
		return "Error de edad: " + err.Error()
	}
	
	return "Error genérico: " + err.Error()
}
```

**Explicación:**
- `errors.As` verifica si el error es de un tipo específico
- Funciona con errores envueltos
- Similar a `isinstance()` en Python

#### PASO 11: Sentinel Errors

```go
func (c *Cuenta) Retirar(cantidad float64) error {
	if cantidad > c.Saldo {
		return ErrSaldoInsuficiente
	}
	c.Saldo -= cantidad
	return nil
}

func VerificarTipoError(err error) string {
	if errors.Is(err, ErrSaldoInsuficiente) {
		return "saldo insuficiente"
	}
	if errors.Is(err, ErrCuentaInexistente) {
		return "cuenta inexistente"
	}
	return "error desconocido"
}
```

**Explicación:**
- Sentinel errors son variables de error predefinidas
- `errors.Is` compara errores (funciona con wrapping)
- Patrón común en Go para errores conocidos

---

## 07 - Goroutines

### 📚 Conceptos Clave

**Goroutines:**
- Concurrencia ligera (más ligera que threads)
- Se lanzan con `go funcion()`
- No bloquean la ejecución
- Necesitas sincronización (WaitGroup, Mutex, Channels)

**Sincronización:**
- `sync.WaitGroup`: esperar que terminen goroutines
- `sync.Mutex`: proteger datos compartidos
- `context.Context`: cancelación y timeouts

### ✅ Soluciones

#### PASO 1: Goroutine Básica

```go
func TareaSimple() {
	time.Sleep(100 * time.Millisecond)
	resultado = "tarea completada"
}

func EjecutarGoroutineSimple() {
	go TareaSimple()
	time.Sleep(200 * time.Millisecond)
}
```

**Explicación:**
- `go TareaSimple()` ejecuta la función concurrentemente
- La función principal continúa sin esperar
- `time.Sleep` es una forma primitiva de esperar (no recomendado en producción)

#### PASO 2: WaitGroup

```go
func TareaConWaitGroup(id int, wg *sync.WaitGroup, resultados []string) {
	defer wg.Done()
	
	time.Sleep(50 * time.Millisecond)
	resultados[id] = fmt.Sprintf("tarea %d completada", id)
}

func EjecutarVariasGoroutines() []string {
	var wg sync.WaitGroup
	resultados := make([]string, 3)
	
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go TareaConWaitGroup(i, &wg, resultados)
	}
	
	wg.Wait()
	return resultados
}
```

**Explicación:**
- `wg.Add(1)` incrementa el contador
- `defer wg.Done()` decrementa cuando termina
- `wg.Wait()` bloquea hasta que el contador llegue a 0
- Forma correcta de esperar goroutines

#### PASO 3: Mutex para Datos Compartidos

```go
func (c *ContadorSeguro) Incrementar() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.valor++
}

func (c *ContadorSeguro) ObtenerValor() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.valor
}
```

**Explicación:**
- `Lock()` adquiere el mutex (bloquea)
- `Unlock()` libera el mutex
- `defer Unlock()` garantiza que se libere incluso si hay panic
- Previene race conditions

#### PASO 4: Race Condition (Demostración)

```go
func (c *ContadorInseguro) Incrementar() {
	c.valor++
}
```

**Explicación:**
- SIN mutex, múltiples goroutines acceden al mismo tiempo
- Causa race condition: el resultado es impredecible
- Podrías obtener 873 en lugar de 1000
- Ejecuta con `go test -race` para detectar race conditions

#### PASO 5: RWMutex

```go
func (c *Cache) Escribir(clave, valor string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.datos[clave] = valor
}

func (c *Cache) Leer(clave string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	valor, existe := c.datos[clave]
	return valor, existe
}
```

**Explicación:**
- `RWMutex` permite múltiples lectores o un escritor
- `Lock()` bloquea para escritura (exclusivo)
- `RLock()` bloquea para lectura (compartido)
- Más eficiente cuando hay muchas lecturas

#### PASO 6: Context para Cancelación

```go
func TareaLarga(ctx context.Context, id int) string {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			return fmt.Sprintf("tarea %d cancelada", id)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
	return fmt.Sprintf("tarea %d completada", id)
}
```

**Explicación:**
- `ctx.Done()` es un channel que se cierra cuando se cancela
- `select` permite verificar cancelación sin bloquear
- `WithTimeout` crea un context que se cancela automáticamente

#### PASO 7: Worker Pool

```go
func Worker(id int, trabajos <-chan int, resultados chan<- string) {
	for trabajo := range trabajos {
		time.Sleep(50 * time.Millisecond)
		resultados <- fmt.Sprintf("worker %d procesó trabajo %d", id, trabajo)
	}
}

func EjecutarWorkerPool() []string {
	const numWorkers = 3
	const numTrabajos = 5
	
	trabajos := make(chan int, numTrabajos)
	resultados := make(chan string, numTrabajos)
	
	// Inicia workers
	for w := 1; w <= numWorkers; w++ {
		go Worker(w, trabajos, resultados)
	}
	
	// Envía trabajos
	for j := 1; j <= numTrabajos; j++ {
		trabajos <- j
	}
	close(trabajos)
	
	// Recolecta resultados
	var output []string
	for a := 1; a <= numTrabajos; a++ {
		output = append(output, <-resultados)
	}
	
	return output
}
```

**Explicación:**
- Patrón común: pool de workers procesando tareas
- Múltiples workers comparten el canal de trabajos
- `close(trabajos)` indica que no hay más trabajos
- Workers terminan cuando el canal se vacía y cierra

#### PASO 8: Rate Limiting

```go
func RateLimiter() []string {
	rate := time.Second / 2
	limiter := time.NewTicker(rate)
	defer limiter.Stop()
	
	var resultados []string
	for i := 0; i < 5; i++ {
		<-limiter.C
		resultados = append(resultados, fmt.Sprintf("operación %d a las %v", i, time.Now().Format("15:04:05")))
	}
	
	return resultados
}
```

**Explicación:**
- `Ticker` envía un tick a intervalos regulares
- `<-limiter.C` bloquea hasta el próximo tick
- Limita operaciones a 2 por segundo
- Útil para rate limiting de APIs

---

## 08 - Channels

### 📚 Conceptos Clave

**Channels:**
- Forma de comunicación entre goroutines
- Tipados: `chan int`, `chan string`, etc.
- Pueden ser buffered o unbuffered
- Se cierran con `close(ch)`

**Direccionalidad:**
- `chan<- T`: solo envío
- `<-chan T`: solo recepción
- Previene errores en tiempo de compilación

### ✅ Soluciones

#### PASO 1: Channel Básico

```go
func EnviarPorChannel(ch chan int, valor int) {
	ch <- valor
}

func RecibirDeChannel(ch chan int) int {
	return <-ch
}
```

**Explicación:**
- `ch <- valor` envía valor al channel
- `<-ch` recibe valor del channel
- Por defecto, ambas operaciones bloquean

#### PASO 2: Buffered Channels

```go
func EjemploBufferedChannel() []int {
	ch := make(chan int, 3)
	
	ch <- 1
	ch <- 2
	ch <- 3
	
	var resultados []int
	resultados = append(resultados, <-ch)
	resultados = append(resultados, <-ch)
	resultados = append(resultados, <-ch)
	
	return resultados
}
```

**Explicación:**
- `make(chan int, 3)` crea channel con buffer de 3
- Puedes enviar 3 valores sin bloquear
- Se bloquea cuando el buffer está lleno
- Útil para desacoplar productor y consumidor

#### PASO 3: Channels Direccionales

```go
func ProducirDatos(ch chan<- string, datos []string) {
	for _, dato := range datos {
		ch <- dato
	}
	close(ch)
}

func ConsumirDatos(ch <-chan string) []string {
	var resultados []string
	for dato := range ch {
		resultados = append(resultados, dato)
	}
	return resultados
}
```

**Explicación:**
- `chan<- string`: solo puede enviar
- `<-chan string`: solo puede recibir
- `close(ch)` cierra el channel
- `range` itera hasta que se cierre

#### PASO 4: Select Statement

```go
func EjemploSelect() string {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from ch1"
	}()
	
	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "from ch2"
	}()
	
	select {
	case msg1 := <-ch1:
		return msg1
	case msg2 := <-ch2:
		return msg2
	case <-time.After(300 * time.Millisecond):
		return "timeout"
	}
}
```

**Explicación:**
- `select` espera en múltiples channels
- Ejecuta el primer case que esté listo
- `time.After` crea timeout
- Similar a `select` en Unix

#### PASO 5: Select Non-Blocking

```go
func EjemploSelectNonBlocking() string {
	ch := make(chan string)
	
	select {
	case msg := <-ch:
		return msg
	default:
		return "no data available"
	}
}
```

**Explicación:**
- `default` hace que select no bloquee
- Se ejecuta inmediatamente si no hay datos
- Útil para polling

#### PASO 6: Fan-Out Pattern

```go
func FanOutProducer(input chan int, outputs []chan int) {
	for num := range input {
		for _, output := range outputs {
			output <- num
		}
	}
	for _, output := range outputs {
		close(output)
	}
}

func FanOutConsumer(id int, input chan int) []string {
	var resultados []string
	for num := range input {
		resultados = append(resultados, fmt.Sprintf("consumer %d processed %d", id, num))
	}
	return resultados
}
```

**Explicación:**
- Un productor envía a múltiples consumidores
- Cada consumidor recibe TODOS los datos
- Útil para distribuir trabajo

#### PASO 7: Pipeline Pattern

```go
func Pipeline1(input chan int) chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			output <- num * 2
		}
	}()
	return output
}

func Pipeline2(input chan int) chan string {
	output := make(chan string)
	go func() {
		defer close(output)
		for num := range input {
			output <- fmt.Sprintf("processed: %d", num)
		}
	}()
	return output
}
```

**Explicación:**
- Cadena de transformaciones
- Cada etapa lee de un channel y escribe a otro
- `defer close` garantiza que se cierre
- Composición de funciones concurrentes

#### PASO 8: Context con Channels

```go
func TareaConContext(ctx context.Context, result chan string) {
	select {
	case <-time.After(200 * time.Millisecond):
		result <- "tarea completada"
	case <-ctx.Done():
		result <- "tarea cancelada"
	}
}
```

**Explicación:**
- Combina channels con context
- Permite cancelación o timeout
- `select` espera lo que ocurra primero

#### PASO 9: Semáforo con Channels

```go
func NuevoSemaforo(capacity int) Semaforo {
	return make(chan struct{}, capacity)
}

func (s Semaforo) Acquire() {
	s <- struct{}{}
}

func (s Semaforo) Release() {
	<-s
}
```

**Explicación:**
- Channel buffered como semáforo
- `struct{}` no ocupa memoria
- Limita concurrencia a N goroutines
- `Acquire` bloquea si está lleno

---

## 09 - Packages

### 📚 Conceptos Clave

**Visibilidad en Go:**
- Mayúscula inicial = público (exportado)
- Minúscula inicial = privado (no exportado)
- Se aplica a: variables, funciones, tipos, campos de struct

**Organización:**
- Un directorio = un package
- Múltiples archivos pueden estar en el mismo package
- `package main` es especial (punto de entrada)

### ✅ Soluciones

#### PASO 1: Exportación

```go
var VariablePublica = "valor público"
var variablePrivada = "secreto"

func FuncionPublica() string {
	return "función pública"
}

func funcionPrivada() string {
	return "función privada"
}

func UsarFuncionPrivada() string {
	return funcionPrivada()
}
```

**Explicación:**
- Mayúscula = exportado (accesible fuera del package)
- Minúscula = no exportado (solo dentro del package)
- Funciones públicas pueden llamar a privadas

#### PASO 2: Structs Exportados

```go
func NuevaPersonaPublica(nombre string, edad int) *PersonaPublica {
	return &PersonaPublica{
		Nombre: nombre,
		edad:   edad,
	}
}

func (p *PersonaPublica) ObtenerEdad() int {
	return p.edad
}

func (p *PersonaPublica) EstablecerEdad(nuevaEdad int) {
	p.edad = nuevaEdad
}
```

**Explicación:**
- `Nombre` es público (puede leerlo/modificarlo cualquiera)
- `edad` es privado (solo accesible vía métodos)
- Patrón de encapsulación en Go

#### PASO 3: Paquete Math

```go
import "math"

func EjemplosPackagesMath() []float64 {
	var resultados []float64
	
	resultados = append(resultados, math.Sqrt(16))    // 4.0
	resultados = append(resultados, math.Pow(2, 3))   // 8.0
	resultados = append(resultados, math.Abs(-5))     // 5.0
	
	return resultados
}
```

**Explicación:**
- Importar paquetes estándar con `import`
- Usar nombre del paquete: `math.Sqrt`
- Todas las funciones exportadas empiezan con mayúscula

#### PASO 4: Paquete Strings

```go
func EjemplosPackagesStrings() []string {
	texto := "Hola Mundo Go"
	var resultados []string
	
	resultados = append(resultados, strings.ToUpper(texto))
	
	if strings.Contains(texto, "Go") {
		resultados = append(resultados, "contiene Go")
	}
	
	resultados = append(resultados, strings.Replace(texto, "Mundo", "Universe", 1))
	
	return resultados
}
```

**Explicación:**
- `strings.ToUpper`: convierte a mayúsculas
- `strings.Contains`: verifica si contiene substring
- `strings.Replace`: reemplaza N ocurrencias

#### PASO 5: Paquete Time

```go
import "time"

func EjemplosPackagesTime() []string {
	var resultados []string
	
	ahora := time.Now()
	resultados = append(resultados, ahora.Format("2006-01-02"))
	
	duracion := 5 * time.Second
	resultados = append(resultados, duracion.String())
	
	resultados = append(resultados, time.Now().Weekday().String())
	
	return resultados
}
```

**Explicación:**
- `time.Now()` obtiene tiempo actual
- `Format("2006-01-02")`: formato específico de Go (fecha de referencia)
- `time.Second`, `time.Minute`, etc. son constantes

#### PASO 6: Alias de Paquetes

```go
import (
	str "strings"
	m "math"
)

func EjemplosAlias() []string {
	texto := "hello world"
	var resultados []string
	
	resultados = append(resultados, str.Title(texto))
	resultados = append(resultados, fmt.Sprintf("%.0f", m.Ceil(3.7)))
	
	return resultados
}
```

**Explicación:**
- `str "strings"`: importa strings con alias `str`
- Útil para evitar conflictos de nombres
- `str.Title` en lugar de `strings.Title`

#### PASO 7: Configuración

```go
func CrearConfiguracion(host string, puerto int) (*Configuracion, error) {
	config := &Configuracion{Host: host, Puerto: puerto}
	
	if !validarConfiguracion(config) {
		return nil, fmt.Errorf("configuración inválida")
	}
	
	return config, nil
}
```

**Explicación:**
- Función pública usa función privada
- Encapsulación: validación es detalle interno
- Patrón constructor con validación

#### PASO 8: Documentación

```go
import "math"

func CalcularAreaCirculo(radio float64) (float64, error) {
	if radio < 0 {
		return 0, fmt.Errorf("el radio no puede ser negativo")
	}
	
	return math.Pi * radio * radio, nil
}
```

**Explicación:**
- Comentario antes de la función se convierte en documentación
- `godoc` generará documentación automáticamente
- Formato: nombre de función + descripción

#### PASO 9: Constantes Exportadas

```go
const (
	VersionMajor = 1
	VersionMinor = 2
	VersionPatch = 3
)

func ObtenerVersion() string {
	return fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
}
```

**Explicación:**
- Constantes exportadas para configuración
- Agrupadas en bloque `const ()`
- Accesibles desde fuera del package

#### PASO 10: Getters de Solo Lectura

```go
func ObtenerConfiguracion(clave string) string {
	return configuracionInterna[clave]
}
```

**Explicación:**
- Variable privada con función pública de acceso
- Solo lectura - no hay setter
- Encapsulación de datos internos

#### PASO 11: Patrón Singleton

```go
func InicializarPaquete(host string, puerto int) error {
	if inicializado {
		return fmt.Errorf("paquete ya inicializado")
	}
	
	config, err := CrearConfiguracion(host, puerto)
	if err != nil {
		return err
	}
	
	instanciaSingleton = config
	inicializado = true
	return nil
}

func ObtenerInstancia() (*Configuracion, error) {
	if !inicializado {
		return nil, fmt.Errorf("paquete no inicializado")
	}
	
	return instanciaSingleton, nil
}
```

**Explicación:**
- Singleton: una sola instancia global
- Inicialización explícita con validación
- Getter retorna error si no está inicializado
- Patrón común en Go para configuración global

---

## 🎓 Consejos Finales

### Verificar Soluciones

```bash
# Ejecutar todos los tests
go test ./...

# Ejecutar un koan específico con verbose
go test ./01_about_variables -v

# Ejecutar con detección de race conditions
go test -race ./...
```

### Próximos Pasos

1. **Completa todos los koans en orden**
2. **Experimenta modificando el código**
3. **Lee la documentación oficial**: https://go.dev/doc/
4. **Practica con proyectos reales**
5. **Aprende sobre:**
   - Testing avanzado
   - Benchmarking
   - Generics (Go 1.18+)
   - Context patterns
   - gRPC y APIs

### Recursos Adicionales

- **Go Tour**: https://go.dev/tour/
- **Go by Example**: https://gobyexample.com/
- **Effective Go**: https://go.dev/doc/effective_go
- **Go Blog**: https://go.dev/blog/

---

**¡Felicidades por completar los Go Koans!** 🎉

Ahora tienes una base sólida en Go. Recuerda: el mejor aprendizaje viene de construir proyectos reales.

**Keep coding!** 🚀
