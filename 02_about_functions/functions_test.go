package functions

import (
	"testing"
)

func TestSaludar(t *testing.T) {
	resultado := Saludar("Mundo")
	esperado := "Hola Mundo"
	
	if resultado != esperado {
		t.Errorf("Saludar('Mundo') = '%s', esperado '%s'", resultado, esperado)
	}
}

func TestSumar(t *testing.T) {
	resultado := Sumar(3, 5)
	esperado := 8
	
	if resultado != esperado {
		t.Errorf("Sumar(3, 5) = %d, esperado %d", resultado, esperado)
	}
}

func TestMultiplicar(t *testing.T) {
	resultado := Multiplicar(4, 6)
	esperado := 24
	
	if resultado != esperado {
		t.Errorf("Multiplicar(4, 6) = %d, esperado %d", resultado, esperado)
	}
}

func TestDividirConResto(t *testing.T) {
	cociente, resto := DividirConResto(17, 5)
	
	if cociente != 3 {
		t.Errorf("DividirConResto(17, 5) cociente = %d, esperado 3", cociente)
	}
	
	if resto != 2 {
		t.Errorf("DividirConResto(17, 5) resto = %d, esperado 2", resto)
	}
}

func TestCoordenadas(t *testing.T) {
	x, y := Coordenadas()
	
	if x != 10 {
		t.Errorf("Coordenadas() x = %d, esperado 10", x)
	}
	
	if y != 20 {
		t.Errorf("Coordenadas() y = %d, esperado 20", y)
	}
}

func TestSumarTodos(t *testing.T) {
	resultado := SumarTodos(1, 2, 3, 4, 5)
	esperado := 15
	
	if resultado != esperado {
		t.Errorf("SumarTodos(1, 2, 3, 4, 5) = %d, esperado %d", resultado, esperado)
	}
	
	// Test con cero argumentos
	resultado = SumarTodos()
	esperado = 0
	
	if resultado != esperado {
		t.Errorf("SumarTodos() = %d, esperado %d", resultado, esperado)
	}
}

func TestOperacionMatematica(t *testing.T) {
	InicializarOperacion()
	
	if OperacionMatematica == nil {
		t.Error("OperacionMatematica no debe ser nil después de InicializarOperacion()")
		return
	}
	
	resultado := OperacionMatematica(6, 7)
	esperado := 42
	
	if resultado != esperado {
		t.Errorf("OperacionMatematica(6, 7) = %d, esperado %d", resultado, esperado)
	}
}

func TestCrearMultiplicador(t *testing.T) {
	multiplicarPor3 := CrearMultiplicador(3)
	resultado := multiplicarPor3(10)
	esperado := 30
	
	if resultado != esperado {
		t.Errorf("CrearMultiplicador(3)(10) = %d, esperado %d", resultado, esperado)
	}
}

func TestEjemploDefer(t *testing.T) {
	resultado := EjemploDefer()
	esperado := "final"
	
	if resultado != esperado {
		t.Errorf("EjemploDefer() = '%s', esperado '%s'", resultado, esperado)
	}
}

func TestDividirSeguro(t *testing.T) {
	// Test división normal
	resultado, err := DividirSeguro(10, 2)
	if err != nil {
		t.Errorf("DividirSeguro(10, 2) no debería retornar error, pero obtuvo: %v", err)
	}
	if resultado != 5 {
		t.Errorf("DividirSeguro(10, 2) = %d, esperado 5", resultado)
	}
	
	// Test división por cero
	_, err = DividirSeguro(10, 0)
	if err == nil {
		t.Error("DividirSeguro(10, 0) debería retornar un error")
	}
}
