// Package structs introduce los structs en Go
// En lugar de clases como Python, Go usa structs con métodos
package structs

// PASO 1: Definir un struct básico
// En Python: class Persona: def __init__(self, nombre, edad): ...
// En Go: type Persona struct { Nombre string; Edad int }

// TODO: Define un struct Persona con campos Nombre (string) y Edad (int)
type Persona struct {
	Nombre string
	Edad   int
}

// PASO 2: Constructor (función que retorna struct)
// En Python: def __init__(self, nombre, edad): ...
// En Go: func NuevaPersona(nombre string, edad int) Persona { ... }

// TODO: Implementa una función constructora
func NuevaPersona(nombre string, edad int) Persona {
	// TODO: Retorna un struct Persona inicializado
	return Persona{}
}

// PASO 3: Métodos en structs
// En Python: def saludar(self): return f"Hola, soy {self.nombre}"
// En Go: func (p Persona) Saludar() string { return fmt.Sprintf("Hola, soy %s", p.Nombre) }

// TODO: Implementa un método Saludar para Persona
func (p Persona) Saludar() string {
	// TODO: Retorna "Hola, soy " + p.Nombre
	return "__FILL_ME__"
}

// PASO 4: Métodos con receptor puntero (para modificar el struct)
// En Python: self.edad += 1
// En Go: func (p *Persona) CumplirAños() { p.Edad++ }

// TODO: Implementa un método que incremente la edad
func (p *Persona) CumplirAños() {
	// TODO: Incrementa p.Edad
}

// PASO 5: Structs anidados
// En Python: class Direccion: ...; class Persona: def __init__(self): self.direccion = Direccion()
type Direccion struct {
	Calle  string
	Ciudad string
}

type Empleado struct {
	Persona   Persona // Embebe Persona (composición)
	Direccion Direccion
	Salario   float64
}

// TODO: Implementa una función constructora para Empleado
func NuevoEmpleado(nombre string, edad int, calle, ciudad string, salario float64) Empleado {
	// TODO: Retorna un Empleado inicializado con todos sus campos
	return Empleado{}
}

// PASO 6: Embedded structs (herencia-like)
// Go no tiene herencia, pero tiene composición por embedding
type Estudiante struct {
	Persona // Embedded - Estudiante "hereda" campos y métodos de Persona
	Carrera string
}

// TODO: Implementa una función constructora para Estudiante
func NuevoEstudiante(nombre string, edad int, carrera string) Estudiante {
	// TODO: Retorna un Estudiante con Persona embebida y carrera
	return Estudiante{}
}

// PASO 7: Métodos que usan campos de structs embebidos
// TODO: Implementa un método que retorne información completa del estudiante
func (e Estudiante) InformacionCompleta() string {
	// TODO: Usa e.Saludar() (heredado) y agrega info de carrera
	// Retorna algo como "Hola, soy Juan y estudio Ingeniería"
	return "__FILL_ME__"
}

// PASO 8: Struct con métodos de diferentes tipos de receptor
type Contador struct {
	Valor int
}

// TODO: Método que retorna el valor actual (receptor valor)
func (c Contador) Obtener() int {
	// TODO: Retorna c.Valor
	return 0
}

// TODO: Método que incrementa el contador (receptor puntero)
func (c *Contador) Incrementar() {
	// TODO: Incrementa c.Valor
}

// TODO: Método que resetea el contador (receptor puntero)
func (c *Contador) Resetear() {
	// TODO: Asigna c.Valor = 0
}

// PASO 9: Struct literal y inicialización
func EjemplosInicializacion() (Persona, Persona, Persona) {
	// TODO: Crea tres personas de diferentes maneras:

	// 1. Con campos nombrados
	p1 := Persona{
		// TODO: Asigna Nombre: "Ana", Edad: 30
	}

	// 2. Con orden de campos (no recomendado pero posible)
	p2 := Persona{ /* TODO: "Luis", 25 */ }

	// 3. Zero value y asignación posterior
	var p3 Persona
	// TODO: Asigna p3.Nombre = "María" y p3.Edad = 35

	return p1, p2, p3
}

// PASO 10: Comparación de structs
// En Go, structs son comparables si todos sus campos son comparables
func CompararPersonas(p1, p2 Persona) bool {
	// TODO: Retorna true si p1 == p2
	return false
}
