package packages

import (
	"strings"
	"testing"
)

func TestVariablePublica(t *testing.T) {
	if VariablePublica == "__FILL_ME__" {
		t.Error("VariablePublica debe ser asignada con un valor real")
	}
}

func TestFuncionPublica(t *testing.T) {
	resultado := FuncionPublica()

	if resultado != "función pública" {
		t.Errorf("FuncionPublica() = '%s', esperado 'función pública'", resultado)
	}
}

func TestUsarFuncionPrivada(t *testing.T) {
	resultado := UsarFuncionPrivada()

	if resultado != "función privada" {
		t.Errorf("UsarFuncionPrivada() = '%s', esperado 'función privada'", resultado)
	}
}

func TestPersonaPublica(t *testing.T) {
	persona := NuevaPersonaPublica("Ana", 25)

	if persona == nil {
		t.Error("NuevaPersonaPublica() no debería retornar nil")
		return
	}

	if persona.Nombre != "Ana" {
		t.Errorf("persona.Nombre = '%s', esperado 'Ana'", persona.Nombre)
	}

	if persona.ObtenerEdad() != 25 {
		t.Errorf("persona.ObtenerEdad() = %d, esperado 25", persona.ObtenerEdad())
	}

	// Test modificar edad
	persona.EstablecerEdad(30)
	if persona.ObtenerEdad() != 30 {
		t.Errorf("Después de EstablecerEdad(30), ObtenerEdad() = %d, esperado 30", persona.ObtenerEdad())
	}
}

func TestEjemplosPackagesMath(t *testing.T) {
	resultados := EjemplosPackagesMath()

	if len(resultados) != 3 {
		t.Errorf("len(resultados) = %d, esperado 3", len(resultados))
		return
	}

	// sqrt(16) = 4
	if resultados[0] != 4.0 {
		t.Errorf("resultados[0] = %f, esperado 4.0", resultados[0])
	}

	// pow(2, 3) = 8
	if resultados[1] != 8.0 {
		t.Errorf("resultados[1] = %f, esperado 8.0", resultados[1])
	}

	// abs(-5) = 5
	if resultados[2] != 5.0 {
		t.Errorf("resultados[2] = %f, esperado 5.0", resultados[2])
	}
}

func TestEjemplosPackagesStrings(t *testing.T) {
	resultados := EjemplosPackagesStrings()

	if len(resultados) < 2 {
		t.Errorf("len(resultados) = %d, esperado al menos 2", len(resultados))
		return
	}

	// Verifica que contiene las transformaciones esperadas
	encontradoMayuscula := false
	encontradoContiene := false

	for _, resultado := range resultados {
		if strings.Contains(resultado, "HOLA MUNDO GO") {
			encontradoMayuscula = true
		}
		if resultado == "contiene Go" {
			encontradoContiene = true
		}
	}

	if !encontradoMayuscula {
		t.Error("No se encontró el texto en mayúsculas")
	}
	if !encontradoContiene {
		t.Error("No se encontró la verificación de 'contiene Go'")
	}
}

func TestEjemplosPackagesTime(t *testing.T) {
	resultados := EjemplosPackagesTime()

	if len(resultados) == 0 {
		t.Error("EjemplosPackagesTime() debería retornar al menos un resultado")
	}

	// Verifica que hay algún contenido relacionado con tiempo
	for _, resultado := range resultados {
		if resultado != "" {
			return // Al menos un resultado no vacío es suficiente para el test
		}
	}

	t.Error("Todos los resultados están vacíos")
}

func TestEjemplosAlias(t *testing.T) {
	resultados := EjemplosAlias()

	if len(resultados) < 1 {
		t.Error("EjemplosAlias() debería retornar al menos un resultado")
		return
	}

	// Verifica que se usaron los alias
	for _, resultado := range resultados {
		if resultado != "" && resultado != "__FILL_ME__" {
			return // Test pasó si hay al menos un resultado válido
		}
	}
}

func TestPackageInicializado(t *testing.T) {
	if !PackageInicializado() {
		t.Error("El paquete debería estar inicializado (función init() debería ejecutarse)")
	}
}

func TestCrearConfiguracion(t *testing.T) {
	// Test configuración válida
	config, err := CrearConfiguracion("localhost", 8080)
	if err != nil {
		t.Errorf("CrearConfiguracion('localhost', 8080) no debería retornar error: %v", err)
	}
	if config == nil {
		t.Error("CrearConfiguracion() no debería retornar nil")
		return
	}
	if config.Host != "localhost" {
		t.Errorf("config.Host = '%s', esperado 'localhost'", config.Host)
	}
	if config.Puerto != 8080 {
		t.Errorf("config.Puerto = %d, esperado 8080", config.Puerto)
	}

	// Test configuración inválida
	_, err = CrearConfiguracion("", 0)
	if err == nil {
		t.Error("CrearConfiguracion('', 0) debería retornar un error")
	}
}

func TestCalcularAreaCirculo(t *testing.T) {
	// Test radio válido
	area, err := CalcularAreaCirculo(5.0)
	if err != nil {
		t.Errorf("CalcularAreaCirculo(5.0) no debería retornar error: %v", err)
	}
	// π * 5² ≈ 78.54
	if area < 78.0 || area > 79.0 {
		t.Errorf("CalcularAreaCirculo(5.0) = %f, esperado ~78.54", area)
	}

	// Test radio negativo
	_, err = CalcularAreaCirculo(-1.0)
	if err == nil {
		t.Error("CalcularAreaCirculo(-1.0) debería retornar un error")
	}
}

func TestObtenerVersion(t *testing.T) {
	version := ObtenerVersion()

	if version == "0.0.0" {
		t.Error("Las constantes de versión deben ser actualizadas")
	}

	if version != "1.2.3" {
		t.Errorf("ObtenerVersion() = '%s', esperado '1.2.3'", version)
	}
}

func TestObtenerConfiguracion(t *testing.T) {
	env := ObtenerConfiguracion("env")
	if env == "__FILL_ME__" {
		t.Error("ObtenerConfiguracion() debe implementarse")
	}

	if env != "development" {
		t.Errorf("ObtenerConfiguracion('env') = '%s', esperado 'development'", env)
	}

	// Test clave inexistente
	inexistente := ObtenerConfiguracion("inexistente")
	if inexistente != "" {
		t.Errorf("ObtenerConfiguracion('inexistente') = '%s', esperado ''", inexistente)
	}
}

func TestObtenerClavesConfiguracion(t *testing.T) {
	claves := ObtenerClavesConfiguracion()

	if len(claves) != 3 {
		t.Errorf("len(claves) = %d, esperado 3", len(claves))
	}

	// Verifica que contiene las claves esperadas
	esperadas := map[string]bool{"env": false, "debug": false, "timeout": false}
	for _, clave := range claves {
		if _, existe := esperadas[clave]; existe {
			esperadas[clave] = true
		}
	}

	for clave, encontrada := range esperadas {
		if !encontrada {
			t.Errorf("Clave '%s' no encontrada en las claves de configuración", clave)
		}
	}
}

func TestInicializarPaqueteYObtenerInstancia(t *testing.T) {
	// Test inicialización
	err := InicializarPaquete("testhost", 9000)
	if err != nil {
		t.Errorf("InicializarPaquete('testhost', 9000) no debería retornar error: %v", err)
	}

	// Test obtener instancia
	instancia, err := ObtenerInstancia()
	if err != nil {
		t.Errorf("ObtenerInstancia() no debería retornar error: %v", err)
	}
	if instancia == nil {
		t.Error("ObtenerInstancia() no debería retornar nil")
		return
	}
	if instancia.Host != "testhost" {
		t.Errorf("instancia.Host = '%s', esperado 'testhost'", instancia.Host)
	}

	// Test doble inicialización
	err = InicializarPaquete("otrohost", 8000)
	if err == nil {
		t.Error("Segunda llamada a InicializarPaquete() debería retornar error")
	}
}
