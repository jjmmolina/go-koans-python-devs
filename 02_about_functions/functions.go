// Package functions introduce las funciones en Go
// Comparado con Python, Go requiere tipos explícitos para parámetros y retorno
package functions

// PASO 1: Función básica
// En Python: def saludar(nombre): return f"Hola {nombre}"
// En Go: func saludar(nombre string) string { return "Hola " + nombre }

// TODO: Implementa una función que tome un string y retorne "Hola " + nombre
func Saludar(nombre string) string {
	return "__FILL_ME__"
}

// PASO 2: Función con múltiples parámetros
// En Python: def sumar(a, b): return a + b
// En Go: func sumar(a int, b int) int { return a + b }

// TODO: Implementa una función que sume dos enteros
func Sumar(a int, b int) int {
	return 0
}

// PASO 3: Función con parámetros del mismo tipo (sintaxis corta)
// En Go puedes escribir: func multiplicar(a, b int) int
func Multiplicar(a, b int) int {
	// TODO: Retorna a * b
	return 0
}

// PASO 4: Función con múltiples valores de retorno
// En Python: return a, b (tupla)
// En Go: return a, b (múltiples valores)

// TODO: Implementa una función que retorne el cociente y el residuo de una división
func DividirConResto(dividendo, divisor int) (int, int) {
	// TODO: Retorna dividendo/divisor, dividendo%divisor
	return 0, 0
}

// PASO 5: Valores de retorno nombrados
// Go permite nombrar los valores de retorno
func Coordenadas() (x, y int) {
	// TODO: Asigna x = 10, y = 20
	// Con valores nombrados, puedes solo hacer 'return' al final
	x = 0
	y = 0
	return // return desnudo, retorna x e y
}

// PASO 6: Función variádica (argumentos variables)
// En Python: def suma(*args): return sum(args)
// En Go: func suma(nums ...int) int

// TODO: Implementa una función que sume todos los números pasados
func SumarTodos(nums ...int) int {
	total := 0
	// TODO: Itera sobre nums y suma todos los valores
	// Hint: usa 'for _, num := range nums'
	return total
}

// PASO 7: Función como valor
// En Go las funciones son valores de primera clase
var OperacionMatematica func(int, int) int

// TODO: Asigna una función que multiplique dos números
func InicializarOperacion() {
	OperacionMatematica = func(a, b int) int {
		// TODO: Retorna a * b
		return 0
	}
}

// PASO 8: Función que retorna función (closure)
// En Python: def multiplicador(factor): return lambda x: x * factor
func CrearMultiplicador(factor int) func(int) int {
	// TODO: Retorna una función que multiplique su parámetro por 'factor'
	return func(x int) int {
		return 0 // TODO: Retorna x * factor
	}
}

// PASO 9: Defer - ejecuta al final de la función
// Único de Go, no existe en Python
func EjemploDefer() string {
	resultado := "inicio"

	// TODO: Usa defer para cambiar resultado a "final" al terminar la función
	// defer func() { resultado = "final" }()

	resultado = "medio"
	return resultado
}

// PASO 10: Manejo de errores - patrón común en Go
// En Python: raise Exception("error")
// En Go: return value, error

// TODO: Implementa una función que retorne error si el divisor es 0
func DividirSeguro(a, b int) (int, error) {
	if b == 0 {
		// TODO: Retorna 0 y un error
		// Hint: necesitarás importar "fmt" y usar fmt.Errorf("división por cero")
		return 0, nil
	}
	return a / b, nil
}
