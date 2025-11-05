// Package pointers introduce los punteros en Go
// En Python todo son referencias, en Go puedes elegir valor o puntero
package pointers

import "fmt"

// PASO 1: Entender punteros básicos
// En Python: todo son referencias (a = [1,2]; b = a; b.append(3) modifica a)
// En Go: puedes elegir pasar por valor o por referencia (puntero)

// TODO: Función que incremente un valor usando puntero
func IncrementarConPuntero(x *int) {
	// TODO: Incrementa el valor al que apunta x
	// Hint: *x++
}

// TODO: Función que trate de incrementar pero falle (paso por valor)
func IncrementarSinPuntero(x int) {
	// TODO: Incrementa x (no funcionará para el caller)
	x++
}

// PASO 2: Obtener dirección de memoria
// En Python: id(objeto) te da algo similar
// En Go: &variable te da la dirección

func ObtenerDireccion(x int) *int {
	// TODO: Retorna la dirección de x usando &
	return nil
}

// PASO 3: Dereferencia de punteros
// En Python: no necesitas hacer esto explícitamente
// En Go: *puntero para obtener el valor

func DereferenciarPuntero(ptr *int) int {
	// TODO: Retorna el valor al que apunta ptr
	return 0
}

// PASO 4: Punteros con structs
type Persona struct {
	Nombre string
	Edad   int
}

// TODO: Método que modifica la persona usando receptor puntero
func (p *Persona) CumplirAños() {
	// TODO: Incrementa p.Edad
}

// TODO: Método que NO modifica (receptor valor)
func (p Persona) ObtenerInfo() string {
	// TODO: Retorna información de la persona sin modificarla
	return fmt.Sprintf("%s tiene %d años", p.Nombre, p.Edad)
}

// PASO 5: Función constructora que retorna puntero
func NuevaPersona(nombre string, edad int) *Persona {
	// TODO: Retorna un puntero a una nueva Persona
	// Puedes usar &Persona{...} o crear variable y retornar &variable
	return nil
}

// PASO 6: Punteros a slices (arrays dinámicos)
// En Python: lista = [1,2,3]; modificar_lista(lista) modifica la original
// En Go: los slices ya son referencias, pero puedes usar punteros para más control

func ModificarSlice(numeros *[]int) {
	// TODO: Agrega el número 4 al slice
	// Hint: *numeros = append(*numeros, 4)
}

// PASO 7: Comparación de punteros
func CompararPunteros(p1, p2 *int) bool {
	// TODO: Compara si p1 y p2 apuntan a la misma dirección
	return false
}

// PASO 8: Puntero nulo (nil)
// En Python: None
// En Go: nil

func EsPunteroNulo(ptr *int) bool {
	// TODO: Verifica si ptr es nil
	return false
}

// PASO 9: Función que puede retornar nil
func BuscarPersona(nombre string, personas []*Persona) *Persona {
	// TODO: Busca una persona por nombre, retorna nil si no la encuentra
	for _, persona := range personas {
		if persona.Nombre == nombre {
			return persona
		}
	}
	return nil
}

// PASO 10: Aritmética de punteros (limitada en Go)
// En C puedes hacer ptr++, en Go es más seguro
type Contador struct {
	valor *int
}

func NuevoContador() *Contador {
	// TODO: Crea un contador con valor inicial 0
	// Hint: crea una variable int, toma su dirección
	inicial := 0
	return &Contador{valor: &inicial}
}

func (c *Contador) Incrementar() {
	// TODO: Incrementa el valor usando el puntero
}

func (c *Contador) ObtenerValor() int {
	// TODO: Retorna el valor actual dereferenciando el puntero
	return 0
}
