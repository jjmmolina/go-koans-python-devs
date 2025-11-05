package goroutines

import (
	"strings"
	"testing"
	"time"
)

func TestEjecutarGoroutineSimple(t *testing.T) {
	resultado = "" // Reset
	EjecutarGoroutineSimple()

	if resultado != "tarea completada" {
		t.Errorf("resultado = '%s', esperado 'tarea completada'", resultado)
	}
}

func TestEjecutarVariasGoroutines(t *testing.T) {
	resultados := EjecutarVariasGoroutines()

	if len(resultados) != 3 {
		t.Errorf("len(resultados) = %d, esperado 3", len(resultados))
		return
	}

	for i, resultado := range resultados {
		expected := "tarea 0 completada"
		if i == 1 {
			expected = "tarea 1 completada"
		} else if i == 2 {
			expected = "tarea 2 completada"
		}

		if resultado != expected {
			t.Errorf("resultados[%d] = '%s', esperado '%s'", i, resultado, expected)
		}
	}
}

func TestContadorSeguro(t *testing.T) {
	resultado := IncrementarConcurrente()

	if resultado != 1000 {
		t.Errorf("IncrementarConcurrente() = %d, esperado 1000", resultado)
	}
}

func TestContadorInseguro(t *testing.T) {
	// Ejecutamos múltiples veces para aumentar probabilidad de race condition
	resultadosIncorrectos := 0
	for i := 0; i < 5; i++ {
		resultado := IncrementarInseguro()
		if resultado != 1000 {
			resultadosIncorrectos++
		}
	}

	// No necesariamente falla, pero debería fallar algunas veces
	// Este test demuestra el problema de race conditions
	t.Logf("Race conditions detectadas en %d/5 ejecuciones", resultadosIncorrectos)
}

func TestCache(t *testing.T) {
	cache := NuevoCache()

	// Test escribir y leer
	cache.Escribir("clave1", "valor1")
	valor, existe := cache.Leer("clave1")

	if !existe {
		t.Error("cache.Leer('clave1') debería encontrar la clave")
	}

	if valor != "valor1" {
		t.Errorf("cache.Leer('clave1') = '%s', esperado 'valor1'", valor)
	}

	// Test clave inexistente
	_, existe = cache.Leer("clave_inexistente")
	if existe {
		t.Error("cache.Leer('clave_inexistente') no debería encontrar la clave")
	}
}

func TestEjecutarConCancelacion(t *testing.T) {
	start := time.Now()
	resultados := EjecutarConCancelacion()
	duration := time.Since(start)

	// Debería tomar aproximadamente 250ms (timeout)
	if duration > 400*time.Millisecond {
		t.Errorf("Duración = %v, esperado ~250ms", duration)
	}

	// Al menos una tarea debería ser cancelada
	canceladas := 0
	for _, resultado := range resultados {
		if strings.Contains(resultado, "cancelada") {
			canceladas++
		}
	}

	if canceladas == 0 {
		t.Error("Al menos una tarea debería haber sido cancelada")
	}
}

func TestEjecutarWorkerPool(t *testing.T) {
	resultados := EjecutarWorkerPool()

	if len(resultados) != 5 {
		t.Errorf("len(resultados) = %d, esperado 5", len(resultados))
		return
	}

	// Verifica que todos los trabajos fueron procesados
	for i, resultado := range resultados {
		if !strings.Contains(resultado, "procesó trabajo") {
			t.Errorf("resultados[%d] = '%s', debe contener 'procesó trabajo'", i, resultado)
		}
	}
}

func TestRateLimiter(t *testing.T) {
	start := time.Now()
	resultados := RateLimiter()
	duration := time.Since(start)

	// 5 operaciones a 2 por segundo = mínimo 2.5 segundos
	expectedMin := 2 * time.Second
	if duration < expectedMin {
		t.Errorf("Duración = %v, esperado al menos %v", duration, expectedMin)
	}

	if len(resultados) != 5 {
		t.Errorf("len(resultados) = %d, esperado 5", len(resultados))
	}

	// Verifica que cada resultado contiene información de timestamp
	for i, resultado := range resultados {
		if !strings.Contains(resultado, "operación") {
			t.Errorf("resultados[%d] = '%s', debe contener 'operación'", i, resultado)
		}
	}
}
