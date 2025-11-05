package structs

import (
	"strings"
	"testing"
)

func TestPersonaStruct(t *testing.T) {
	p := NuevaPersona("Juan", 30)
	
	if p.Nombre != "Juan" {
		t.Errorf("p.Nombre = '%s', esperado 'Juan'", p.Nombre)
	}
	
	if p.Edad != 30 {
		t.Errorf("p.Edad = %d, esperado 30", p.Edad)
	}
}

func TestSaludar(t *testing.T) {
	p := NuevaPersona("Ana", 25)
	resultado := p.Saludar()
	esperado := "Hola, soy Ana"
	
	if resultado != esperado {
		t.Errorf("p.Saludar() = '%s', esperado '%s'", resultado, esperado)
	}
}

func TestCumplirAños(t *testing.T) {
	p := NuevaPersona("Luis", 20)
	edadInicial := p.Edad
	
	p.CumplirAños()
	
	if p.Edad != edadInicial+1 {
		t.Errorf("Después de CumplirAños(), edad = %d, esperado %d", p.Edad, edadInicial+1)
	}
}

func TestEmpleado(t *testing.T) {
	e := NuevoEmpleado("María", 28, "Av. Principal", "Madrid", 45000.0)
	
	if e.Persona.Nombre != "María" {
		t.Errorf("e.Persona.Nombre = '%s', esperado 'María'", e.Persona.Nombre)
	}
	
	if e.Persona.Edad != 28 {
		t.Errorf("e.Persona.Edad = %d, esperado 28", e.Persona.Edad)
	}
	
	if e.Direccion.Calle != "Av. Principal" {
		t.Errorf("e.Direccion.Calle = '%s', esperado 'Av. Principal'", e.Direccion.Calle)
	}
	
	if e.Direccion.Ciudad != "Madrid" {
		t.Errorf("e.Direccion.Ciudad = '%s', esperado 'Madrid'", e.Direccion.Ciudad)
	}
	
	if e.Salario != 45000.0 {
		t.Errorf("e.Salario = %f, esperado 45000.0", e.Salario)
	}
}

func TestEstudiante(t *testing.T) {
	est := NuevoEstudiante("Carlos", 22, "Ingeniería")
	
	if est.Nombre != "Carlos" {
		t.Errorf("est.Nombre = '%s', esperado 'Carlos'", est.Nombre)
	}
	
	if est.Edad != 22 {
		t.Errorf("est.Edad = %d, esperado 22", est.Edad)
	}
	
	if est.Carrera != "Ingeniería" {
		t.Errorf("est.Carrera = '%s', esperado 'Ingeniería'", est.Carrera)
	}
	
	// Test método heredado
	saludo := est.Saludar()
	if saludo != "Hola, soy Carlos" {
		t.Errorf("est.Saludar() = '%s', esperado 'Hola, soy Carlos'", saludo)
	}
}

func TestInformacionCompleta(t *testing.T) {
	est := NuevoEstudiante("Elena", 21, "Medicina")
	info := est.InformacionCompleta()
	
	if !strings.Contains(info, "Elena") {
		t.Errorf("InformacionCompleta() debe contener 'Elena', pero obtuvo '%s'", info)
	}
	
	if !strings.Contains(info, "Medicina") {
		t.Errorf("InformacionCompleta() debe contener 'Medicina', pero obtuvo '%s'", info)
	}
}

func TestContador(t *testing.T) {
	var c Contador
	
	// Test valor inicial
	if c.Obtener() != 0 {
		t.Errorf("c.Obtener() = %d, esperado 0", c.Obtener())
	}
	
	// Test incrementar
	c.Incrementar()
	if c.Obtener() != 1 {
		t.Errorf("Después de Incrementar(), c.Obtener() = %d, esperado 1", c.Obtener())
	}
	
	// Test incrementar múltiples veces
	c.Incrementar()
	c.Incrementar()
	if c.Obtener() != 3 {
		t.Errorf("Después de 3 incrementos, c.Obtener() = %d, esperado 3", c.Obtener())
	}
	
	// Test resetear
	c.Resetear()
	if c.Obtener() != 0 {
		t.Errorf("Después de Resetear(), c.Obtener() = %d, esperado 0", c.Obtener())
	}
}

func TestEjemplosInicializacion(t *testing.T) {
	p1, p2, p3 := EjemplosInicializacion()
	
	// Test p1
	if p1.Nombre != "Ana" || p1.Edad != 30 {
		t.Errorf("p1 = {%s, %d}, esperado {Ana, 30}", p1.Nombre, p1.Edad)
	}
	
	// Test p2
	if p2.Nombre != "Luis" || p2.Edad != 25 {
		t.Errorf("p2 = {%s, %d}, esperado {Luis, 25}", p2.Nombre, p2.Edad)
	}
	
	// Test p3
	if p3.Nombre != "María" || p3.Edad != 35 {
		t.Errorf("p3 = {%s, %d}, esperado {María, 35}", p3.Nombre, p3.Edad)
	}
}

func TestCompararPersonas(t *testing.T) {
	p1 := NuevaPersona("Juan", 30)
	p2 := NuevaPersona("Juan", 30)
	p3 := NuevaPersona("Ana", 25)
	
	if !CompararPersonas(p1, p2) {
		t.Error("Personas idénticas deberían ser iguales")
	}
	
	if CompararPersonas(p1, p3) {
		t.Error("Personas diferentes no deberían ser iguales")
	}
}
