package interfaces

import (
	"strings"
	"testing"
)

func TestPerroHablar(t *testing.T) {
	perro := Perro{Nombre: "Rex"}
	resultado := perro.Hablar()
	
	if !strings.Contains(resultado, "Guau") || !strings.Contains(resultado, "Rex") {
		t.Errorf("perro.Hablar() = '%s', debe contener 'Guau' y 'Rex'", resultado)
	}
}

func TestGatoHablar(t *testing.T) {
	gato := Gato{Nombre: "Whiskers"}
	resultado := gato.Hablar()
	
	if !strings.Contains(resultado, "Miau") || !strings.Contains(resultado, "Whiskers") {
		t.Errorf("gato.Hablar() = '%s', debe contener 'Miau' y 'Whiskers'", resultado)
	}
}

func TestHacerHablar(t *testing.T) {
	perro := Perro{Nombre: "Buddy"}
	gato := Gato{Nombre: "Fluffy"}
	
	resultadoPerro := HacerHablar(perro)
	resultadoGato := HacerHablar(gato)
	
	if !strings.Contains(resultadoPerro, "Buddy") {
		t.Errorf("HacerHablar(perro) debe contener 'Buddy', obtuvo '%s'", resultadoPerro)
	}
	
	if !strings.Contains(resultadoGato, "Fluffy") {
		t.Errorf("HacerHablar(gato) debe contener 'Fluffy', obtuvo '%s'", resultadoGato)
	}
}

func TestDescribirTipo(t *testing.T) {
	tests := []struct {
		valor    any
		contiene string
	}{
		{"hola", "string"},
		{42, "int"},
		{true, "bool"},
		{3.14, "desconocido"},
	}
	
	for _, test := range tests {
		resultado := DescribirTipo(test.valor)
		if !strings.Contains(strings.ToLower(resultado), test.contiene) {
			t.Errorf("DescribirTipo(%v) = '%s', debe contener '%s'", test.valor, resultado, test.contiene)
		}
	}
}

func TestProgramador(t *testing.T) {
	prog := Programador{Lenguaje: "Go"}
	
	trabajo := prog.Trabajar()
	if !strings.Contains(trabajo, "Go") {
		t.Errorf("prog.Trabajar() = '%s', debe contener 'Go'", trabajo)
	}
	
	descanso := prog.DescansoAlmuerzo()
	if descanso == "__FILL_ME__" {
		t.Error("prog.DescansoAlmuerzo() no debe retornar '__FILL_ME__'")
	}
}

func TestEsPerro(t *testing.T) {
	perro := Perro{Nombre: "Max"}
	gato := Gato{Nombre: "Salem"}
	
	if !EsPerro(perro) {
		t.Error("EsPerro(perro) debe retornar true")
	}
	
	if EsPerro(gato) {
		t.Error("EsPerro(gato) debe retornar false")
	}
}

func TestRobot(t *testing.T) {
	robot := Robot{Modelo: "R2D2"}
	
	// Test como Hablador
	habla := robot.Hablar()
	if !strings.Contains(habla, "R2D2") {
		t.Errorf("robot.Hablar() = '%s', debe contener 'R2D2'", habla)
	}
	
	// Test como Trabajador
	trabajo := robot.Trabajar()
	if trabajo == "__FILL_ME__" {
		t.Error("robot.Trabajar() no debe retornar '__FILL_ME__'")
	}
	
	descanso := robot.DescansoAlmuerzo()
	if descanso == "__FILL_ME__" {
		t.Error("robot.DescansoAlmuerzo() no debe retornar '__FILL_ME__'")
	}
	
	// Test interface embebida
	var ht HabladorTrabajador = robot
	_ = ht.Hablar()
	_ = ht.Trabajar()
	_ = ht.DescansoAlmuerzo()
}

func TestProductoString(t *testing.T) {
	producto := Producto{Nombre: "Laptop", Precio: 999.99}
	resultado := producto.String()
	
	if !strings.Contains(resultado, "Laptop") {
		t.Errorf("producto.String() = '%s', debe contener 'Laptop'", resultado)
	}
	
	if !strings.Contains(resultado, "999.99") {
		t.Errorf("producto.String() = '%s', debe contener '999.99'", resultado)
	}
}
