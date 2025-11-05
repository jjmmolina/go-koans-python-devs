// Package interfaces introduce las interfaces en Go
// En Python usas duck typing, en Go uses interfaces explícitas
package interfaces

import "fmt"

// PASO 1: Definir una interface básica
// En Python: No necesitas declarar interfaces, solo implementas métodos
// En Go: type Speaker interface { Speak() string }

// TODO: Define una interface Hablador con método Hablar() string
type Hablador interface {
	Hablar() string
}

// PASO 2: Implementar interface con structs
// En Go, implementas interfaces implícitamente (no hay palabra clave 'implements')

type Perro struct {
	Nombre string
}

// TODO: Implementa el método Hablar para Perro
func (p Perro) Hablar() string {
	// TODO: Retorna algo como "Guau! Soy " + p.Nombre
	return "__FILL_ME__"
}

type Gato struct {
	Nombre string
}

// TODO: Implementa el método Hablar para Gato
func (g Gato) Hablar() string {
	// TODO: Retorna algo como "Miau! Soy " + g.Nombre
	return "__FILL_ME__"
}

// PASO 3: Usar interfaces como parámetros
// TODO: Función que acepta cualquier Hablador
func HacerHablar(h Hablador) string {
	// TODO: Retorna h.Hablar()
	return "__FILL_ME__"
}

// PASO 4: Interface vacía
// En Python: cualquier object
// En Go: interface{} (en Go 1.18+ es 'any')

// TODO: Función que acepta cualquier tipo
func DescribirTipo(valor any) string {
	// TODO: Usa type assertion o switch para determinar el tipo
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

// PASO 5: Interface con múltiples métodos
type Trabajador interface {
	Trabajar() string
	DescansoAlmuerzo() string
}

type Programador struct {
	Lenguaje string
}

// TODO: Implementa Trabajar para Programador
func (p Programador) Trabajar() string {
	// TODO: Retorna "Programando en " + p.Lenguaje
	return "__FILL_ME__"
}

// TODO: Implementa DescansoAlmuerzo para Programador
func (p Programador) DescansoAlmuerzo() string {
	// TODO: Retorna "Debugeando mientras como"
	return "__FILL_ME__"
}

// PASO 6: Type assertion
// TODO: Función que verifica si un Hablador es específicamente un Perro
func EsPerro(h Hablador) bool {
	// TODO: Usa type assertion para verificar si h es *Perro o Perro
	_, esPerro := h.(Perro)
	return esPerro
}

// PASO 7: Interface embebida
type HabladorTrabajador interface {
	Hablador   // Interface embebida
	Trabajador // Interface embebida
}

type Robot struct {
	Modelo string
}

// TODO: Implementa todos los métodos necesarios para que Robot satisfaga HabladorTrabajador
func (r Robot) Hablar() string {
	return "__FILL_ME__" // "Bip boop! Soy " + r.Modelo
}

func (r Robot) Trabajar() string {
	return "__FILL_ME__" // "Procesando datos..."
}

func (r Robot) DescansoAlmuerzo() string {
	return "__FILL_ME__" // "Recargando baterías"
}

// PASO 8: Stringer interface (muy común en Go)
// Equivalente a __str__ en Python
type Producto struct {
	Nombre string
	Precio float64
}

// TODO: Implementa String() para que Producto implemente fmt.Stringer
func (p Producto) String() string {
	// TODO: Retorna formato como "Producto: Laptop ($999.99)"
	return "__FILL_ME__"
}
