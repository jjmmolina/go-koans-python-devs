package errors

import (
	"errors"
	"strings"
	"testing"
)

func TestDividir(t *testing.T) {
	// Test división normal
	resultado, err := Dividir(10, 2)
	if err != nil {
		t.Errorf("Dividir(10, 2) no debería retornar error: %v", err)
	}
	if resultado != 5.0 {
		t.Errorf("Dividir(10, 2) = %f, esperado 5.0", resultado)
	}

	// Test división por cero
	_, err = Dividir(10, 0)
	if err == nil {
		t.Error("Dividir(10, 0) debería retornar un error")
	}
}

func TestErrorEdadInvalida(t *testing.T) {
	err := ErrorEdadInvalida{Edad: 150}
	mensaje := err.Error()

	if !strings.Contains(mensaje, "150") {
		t.Errorf("Error() = '%s', debe contener '150'", mensaje)
	}
}

func TestValidarEdad(t *testing.T) {
	// Test edad válida
	err := ValidarEdad(25)
	if err != nil {
		t.Errorf("ValidarEdad(25) no debería retornar error: %v", err)
	}

	// Test edad inválida (negativa)
	err = ValidarEdad(-5)
	if err == nil {
		t.Error("ValidarEdad(-5) debería retornar un error")
	}

	// Test edad inválida (muy alta)
	err = ValidarEdad(150)
	if err == nil {
		t.Error("ValidarEdad(150) debería retornar un error")
	}
}

func TestDividirSeguro(t *testing.T) {
	// Test división exitosa
	resultado := DividirSeguro(10, 2)
	if resultado != "5" {
		t.Errorf("DividirSeguro(10, 2) = '%s', esperado '5'", resultado)
	}

	// Test división por cero
	resultado = DividirSeguro(10, 0)
	if !strings.Contains(resultado, "Error") {
		t.Errorf("DividirSeguro(10, 0) = '%s', debe contener 'Error'", resultado)
	}
}

func TestConvertirYDividir(t *testing.T) {
	// Test conversión y división exitosa
	resultado, err := ConvertirYDividir("10", "2")
	if err != nil {
		t.Errorf("ConvertirYDividir('10', '2') no debería retornar error: %v", err)
	}
	if resultado != 5.0 {
		t.Errorf("ConvertirYDividir('10', '2') = %f, esperado 5.0", resultado)
	}

	// Test error de conversión
	_, err = ConvertirYDividir("abc", "2")
	if err == nil {
		t.Error("ConvertirYDividir('abc', '2') debería retornar un error")
	}
}

func TestConvertirConContexto(t *testing.T) {
	// Test conversión exitosa
	num, err := ConvertirConContexto("42")
	if err != nil {
		t.Errorf("ConvertirConContexto('42') no debería retornar error: %v", err)
	}
	if num != 42 {
		t.Errorf("ConvertirConContexto('42') = %d, esperado 42", num)
	}

	// Test error con contexto
	_, err = ConvertirConContexto("abc")
	if err == nil {
		t.Error("ConvertirConContexto('abc') debería retornar un error")
	}
}

func TestCalcular(t *testing.T) {
	// Test suma
	resultado, err := Calcular("suma", 3, 4)
	if err != nil {
		t.Errorf("Calcular('suma', 3, 4) no debería retornar error: %v", err)
	}
	if resultado != 7 {
		t.Errorf("Calcular('suma', 3, 4) = %f, esperado 7", resultado)
	}

	// Test división por cero
	_, err = Calcular("division", 10, 0)
	if err == nil {
		t.Error("Calcular('division', 10, 0) debería retornar un error")
	}

	// Test operación inválida
	_, err = Calcular("raiz", 4, 0)
	if err == nil {
		t.Error("Calcular('raiz', 4, 0) debería retornar un error")
	}
}

func TestProcesarArchivo(t *testing.T) {
	// Test archivo bueno
	err := ProcesarArchivo("archivo_bueno.txt")
	if err != nil {
		t.Errorf("ProcesarArchivo('archivo_bueno.txt') no debería retornar error: %v", err)
	}

	// Test archivo malo
	err = ProcesarArchivo("archivo_malo.txt")
	if err == nil {
		t.Error("ProcesarArchivo('archivo_malo.txt') debería retornar un error")
	}
}

func TestFuncionPeligrosa(t *testing.T) {
	// Test sin pánico
	resultado := FuncionPeligrosa(false)
	if resultado != "todo bien" {
		t.Errorf("FuncionPeligrosa(false) = '%s', esperado 'todo bien'", resultado)
	}

	// Test con pánico (debería recuperarse)
	resultado = FuncionPeligrosa(true)
	if !strings.Contains(resultado, "recuperó") {
		t.Errorf("FuncionPeligrosa(true) = '%s', debe contener 'recuperó'", resultado)
	}
}

func TestManejarErrorEspecifico(t *testing.T) {
	// Test con ErrorEdadInvalida
	err := ErrorEdadInvalida{Edad: 200}
	resultado := ManejarErrorEspecifico(err)
	if !strings.Contains(resultado, "Error de edad") {
		t.Errorf("ManejarErrorEspecifico(ErrorEdadInvalida) = '%s', debe contener 'Error de edad'", resultado)
	}

	// Test con error genérico
	err2 := errors.New("error genérico")
	resultado = ManejarErrorEspecifico(err2)
	if !strings.Contains(resultado, "Error genérico") {
		t.Errorf("ManejarErrorEspecifico(error genérico) = '%s', debe contener 'Error genérico'", resultado)
	}
}

func TestCuentaRetirar(t *testing.T) {
	cuenta := &Cuenta{Saldo: 100}

	// Test retiro exitoso
	err := cuenta.Retirar(50)
	if err != nil {
		t.Errorf("Retirar(50) no debería retornar error: %v", err)
	}
	if cuenta.Saldo != 50 {
		t.Errorf("Saldo después de retirar 50 = %f, esperado 50", cuenta.Saldo)
	}

	// Test saldo insuficiente
	err = cuenta.Retirar(100)
	if err == nil {
		t.Error("Retirar(100) debería retornar error de saldo insuficiente")
	}
}

func TestVerificarTipoError(t *testing.T) {
	// Test ErrSaldoInsuficiente
	resultado := VerificarTipoError(ErrSaldoInsuficiente)
	if resultado != "saldo insuficiente" {
		t.Errorf("VerificarTipoError(ErrSaldoInsuficiente) = '%s', esperado 'saldo insuficiente'", resultado)
	}

	// Test ErrCuentaInexistente
	resultado = VerificarTipoError(ErrCuentaInexistente)
	if resultado != "cuenta inexistente" {
		t.Errorf("VerificarTipoError(ErrCuentaInexistente) = '%s', esperado 'cuenta inexistente'", resultado)
	}

	// Test error desconocido
	resultado = VerificarTipoError(errors.New("otro error"))
	if resultado != "error desconocido" {
		t.Errorf("VerificarTipoError(otro error) = '%s', esperado 'error desconocido'", resultado)
	}
}
