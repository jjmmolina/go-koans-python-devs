package pointers

import "testing"

func TestIncrementarConPuntero(t *testing.T) {
	x := 5
	IncrementarConPuntero(&x)

	if x != 6 {
		t.Errorf("IncrementarConPuntero(&x) debería incrementar x a 6, pero obtuvo %d", x)
	}
}

func TestIncrementarSinPuntero(t *testing.T) {
	x := 5
	IncrementarSinPuntero(x)

	// x no debería cambiar porque se pasa por valor
	if x != 5 {
		t.Errorf("IncrementarSinPuntero(x) no debería modificar x, pero cambió a %d", x)
	}
}

func TestObtenerDireccion(t *testing.T) {
	x := 42
	ptr := ObtenerDireccion(x)

	if ptr == nil {
		t.Error("ObtenerDireccion(x) no debería retornar nil")
	}

	if *ptr != x {
		t.Errorf("*ObtenerDireccion(x) = %d, esperado %d", *ptr, x)
	}
}

func TestDereferenciarPuntero(t *testing.T) {
	x := 100
	valor := DereferenciarPuntero(&x)

	if valor != x {
		t.Errorf("DereferenciarPuntero(&x) = %d, esperado %d", valor, x)
	}
}

func TestPersonaCumplirAños(t *testing.T) {
	p := &Persona{Nombre: "Ana", Edad: 25}
	edadInicial := p.Edad

	p.CumplirAños()

	if p.Edad != edadInicial+1 {
		t.Errorf("Después de CumplirAños(), edad = %d, esperado %d", p.Edad, edadInicial+1)
	}
}

func TestPersonaObtenerInfo(t *testing.T) {
	p := Persona{Nombre: "Luis", Edad: 30}
	info := p.ObtenerInfo()

	if info != "Luis tiene 30 años" {
		t.Errorf("ObtenerInfo() = '%s', esperado 'Luis tiene 30 años'", info)
	}
}

func TestNuevaPersona(t *testing.T) {
	p := NuevaPersona("María", 28)

	if p == nil {
		t.Error("NuevaPersona() no debería retornar nil")
		return
	}

	if p.Nombre != "María" {
		t.Errorf("p.Nombre = '%s', esperado 'María'", p.Nombre)
	}

	if p.Edad != 28 {
		t.Errorf("p.Edad = %d, esperado 28", p.Edad)
	}
}

func TestModificarSlice(t *testing.T) {
	numeros := []int{1, 2, 3}
	longitudInicial := len(numeros)

	ModificarSlice(&numeros)

	if len(numeros) != longitudInicial+1 {
		t.Errorf("ModificarSlice() debería agregar un elemento, longitud = %d, esperado %d", len(numeros), longitudInicial+1)
	}

	if numeros[len(numeros)-1] != 4 {
		t.Errorf("Último elemento = %d, esperado 4", numeros[len(numeros)-1])
	}
}

func TestCompararPunteros(t *testing.T) {
	x := 10
	y := 10

	// Mismo puntero
	ptr1 := &x
	ptr2 := &x
	if !CompararPunteros(ptr1, ptr2) {
		t.Error("Punteros al mismo objeto deberían ser iguales")
	}

	// Diferentes punteros
	ptr3 := &y
	if CompararPunteros(ptr1, ptr3) {
		t.Error("Punteros a diferentes objetos no deberían ser iguales")
	}
}

func TestEsPunteroNulo(t *testing.T) {
	var ptr *int = nil

	if !EsPunteroNulo(ptr) {
		t.Error("EsPunteroNulo(nil) debería retornar true")
	}

	x := 42
	if EsPunteroNulo(&x) {
		t.Error("EsPunteroNulo(&x) debería retornar false")
	}
}

func TestBuscarPersona(t *testing.T) {
	personas := []*Persona{
		{Nombre: "Ana", Edad: 25},
		{Nombre: "Luis", Edad: 30},
		{Nombre: "María", Edad: 28},
	}

	// Buscar persona existente
	encontrada := BuscarPersona("Luis", personas)
	if encontrada == nil {
		t.Error("BuscarPersona('Luis') no debería retornar nil")
	} else if encontrada.Nombre != "Luis" {
		t.Errorf("encontrada.Nombre = '%s', esperado 'Luis'", encontrada.Nombre)
	}

	// Buscar persona inexistente
	noEncontrada := BuscarPersona("Pedro", personas)
	if noEncontrada != nil {
		t.Error("BuscarPersona('Pedro') debería retornar nil")
	}
}

func TestContador(t *testing.T) {
	c := NuevoContador()

	if c == nil {
		t.Error("NuevoContador() no debería retornar nil")
		return
	}

	// Test valor inicial
	if c.ObtenerValor() != 0 {
		t.Errorf("Valor inicial = %d, esperado 0", c.ObtenerValor())
	}

	// Test incrementar
	c.Incrementar()
	if c.ObtenerValor() != 1 {
		t.Errorf("Después de Incrementar(), valor = %d, esperado 1", c.ObtenerValor())
	}

	// Test múltiples incrementos
	c.Incrementar()
	c.Incrementar()
	if c.ObtenerValor() != 3 {
		t.Errorf("Después de 3 incrementos, valor = %d, esperado 3", c.ObtenerValor())
	}
}
