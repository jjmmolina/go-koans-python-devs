package variables

import (
	"testing"
)

// TestVariablesBasicas verifica las declaraciones básicas de variables
func TestVariablesBasicas(t *testing.T) {
	// Test saludo
	if GetSaludo() != "Hola Go" {
		t.Errorf("saludo debe ser 'Hola Go', pero obtuvimos '%s'", GetSaludo())
	}
	
	// Test edad
	if GetEdad() != 25 {
		t.Errorf("edad debe ser 25, pero obtuvimos %d", GetEdad())
	}
	
	// Test esActivo
	if !IsActivo() {
		t.Errorf("esActivo debe ser true, pero obtuvimos %t", IsActivo())
	}
	
	// Test precio
	if GetPrecio() != 19.99 {
		t.Errorf("precio debe ser 19.99, pero obtuvimos %f", GetPrecio())
	}
}

// TestDeclaracionCorta verifica el uso de :=
func TestDeclaracionCorta(t *testing.T) {
	nombre, puntos, completado := DeclaracionCorta()
	
	if nombre != "Ana" {
		t.Errorf("nombre debe ser 'Ana', pero obtuvimos '%s'", nombre)
	}
	
	if puntos != 100 {
		t.Errorf("puntos debe ser 100, pero obtuvimos %d", puntos)
	}
	
	if !completado {
		t.Errorf("completado debe ser true, pero obtuvimos %t", completado)
	}
}

// TestConstantes verifica las constantes
func TestConstantes(t *testing.T) {
	piString, maxUsuarios := GetConstants()
	
	if piString != "3.14159" {
		t.Errorf("PI_STRING debe ser '3.14159', pero obtuvimos '%s'", piString)
	}
	
	if maxUsuarios != 1000 {
		t.Errorf("MAX_USUARIOS debe ser 1000, pero obtuvimos %d", maxUsuarios)
	}
}

// TestMultiplesAsignaciones verifica asignación múltiple
func TestMultiplesAsignaciones(t *testing.T) {
	x, y, mensaje := MultiplesAsignaciones()
	
	if x != 10 {
		t.Errorf("x debe ser 10, pero obtuvimos %d", x)
	}
	
	if y != 20 {
		t.Errorf("y debe ser 20, pero obtuvimos %d", y)
	}
	
	if mensaje != "múltiple" {
		t.Errorf("mensaje debe ser 'múltiple', pero obtuvimos '%s'", mensaje)
	}
}

// TestZeroValues verifica los valores por defecto
func TestZeroValues(t *testing.T) {
	numero, texto, activo, decimal := ZeroValues()
	
	if numero != 0 {
		t.Errorf("zero value de int debe ser 0, pero obtuvimos %d", numero)
	}
	
	if texto != "" {
		t.Errorf("zero value de string debe ser '', pero obtuvimos '%s'", texto)
	}
	
	if activo != false {
		t.Errorf("zero value de bool debe ser false, pero obtuvimos %t", activo)
	}
	
	if decimal != 0.0 {
		t.Errorf("zero value de float64 debe ser 0.0, pero obtuvimos %f", decimal)
	}
}

// TestConversionTipos verifica la conversión entre tipos
func TestConversionTipos(t *testing.T) {
	entero, decimal, texto := ConversionTipos()
	
	if entero != 42 {
		t.Errorf("entero debe ser 42, pero obtuvimos %d", entero)
	}
	
	if decimal != 42.0 {
		t.Errorf("decimal debe ser 42.0, pero obtuvimos %f", decimal)
	}
	
	if texto != "42" {
		t.Errorf("texto debe ser '42', pero obtuvimos '%s'", texto)
	}
}
