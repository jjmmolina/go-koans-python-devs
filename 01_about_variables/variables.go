// Package variables introduces basic Go variables and types
// Comparado con Python, Go es tipado estáticamente
package variables

// PASO 1: Arregla las declaraciones de variables
// En Python: nombre = "Juan"
// En Go: var nombre string = "Juan" o nombre := "Juan"

var (
	// TODO: Declara una variable string llamada 'saludo' con valor "Hola Go"
	saludo = "__FILL_ME__"

	// TODO: Declara una variable int llamada 'edad' con valor 25
	edad = 0

	// TODO: Declara una variable bool llamada 'esActivo' con valor true
	esActivo = false

	// TODO: Declara una variable float64 llamada 'precio' con valor 19.99
	precio = 0.0
)

// GetSaludo retorna el saludo
func GetSaludo() string {
	return saludo
}

// GetEdad retorna la edad
func GetEdad() int {
	return edad
}

// IsActivo retorna si está activo
func IsActivo() bool {
	return esActivo
}

// GetPrecio retorna el precio
func GetPrecio() float64 {
	return precio
}

// PASO 2: Declaración corta de variables
// En Go puedes usar := para declarar e inicializar en una línea
func DeclaracionCorta() (string, int, bool) {
	// TODO: Usa := para declarar y asignar:
	// - nombre con valor "Ana"
	// - puntos con valor 100
	// - completado con valor true

	nombre := "__FILL_ME__"
	puntos := 0
	completado := false

	return nombre, puntos, completado
}

// PASO 3: Constantes
// En Python: CONSTANT = "valor" (por convención)
// En Go: const CONSTANT = "valor" (real constante)
const (
	// TODO: Define una constante string PI_STRING con valor "3.14159"
	PI_STRING = "__FILL_ME__"

	// TODO: Define una constante int MAX_USUARIOS con valor 1000
	MAX_USUARIOS = 0
)

// GetConstants retorna las constantes
func GetConstants() (string, int) {
	return PI_STRING, MAX_USUARIOS
}

// PASO 4: Múltiple asignación
// En Python: a, b = 1, 2
// En Go: a, b := 1, 2
func MultiplesAsignaciones() (int, int, string) {
	// TODO: Asigna múltiples variables en una línea:
	// x = 10, y = 20, mensaje = "múltiple"
	x, y, mensaje := 0, 0, "__FILL_ME__"

	return x, y, mensaje
}

// PASO 5: Zero values (valores por defecto)
// Go inicializa automáticamente las variables con "zero values"
func ZeroValues() (int, string, bool, float64) {
	var (
		numero  int     // zero value de int es 0
		texto   string  // zero value de string es ""
		activo  bool    // zero value de bool es false
		decimal float64 // zero value de float64 es 0.0
	)

	// TODO: Retorna los zero values. No cambies nada aquí, solo observa.
	return numero, texto, activo, decimal
}

// PASO 6: Conversión de tipos
// En Python: int("123") o str(123)
// En Go: int(x) o string(x), pero más estricto
func ConversionTipos() (int, float64, string) {
	var entero int = 42

	// TODO: Convierte 'entero' a float64
	var decimal float64 = 0.0 // Hint: float64(entero)

	// TODO: Convierte 'entero' a string
	// Hint: necesitarás importar "strconv" y usar strconv.Itoa(entero)
	var texto string = "__FILL_ME__"

	return entero, decimal, texto
}
