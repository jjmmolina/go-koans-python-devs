package channels

import (
	"strings"
	"testing"
	"time"
)

func TestEjemploCommunicacionBasica(t *testing.T) {
	resultado := EjemploCommunicacionBasica()

	if resultado != 42 {
		t.Errorf("EjemploCommunicacionBasica() = %d, esperado 42", resultado)
	}
}

func TestEjemploBufferedChannel(t *testing.T) {
	resultados := EjemploBufferedChannel()

	esperados := []int{1, 2, 3}
	if len(resultados) != len(esperados) {
		t.Errorf("len(resultados) = %d, esperado %d", len(resultados), len(esperados))
		return
	}

	for i, resultado := range resultados {
		if resultado != esperados[i] {
			t.Errorf("resultados[%d] = %d, esperado %d", i, resultado, esperados[i])
		}
	}
}

func TestEjemploProducirConsumir(t *testing.T) {
	resultados := EjemploProducirConsumir()

	esperados := []string{"a", "b", "c", "d", "e"}
	if len(resultados) != len(esperados) {
		t.Errorf("len(resultados) = %d, esperado %d", len(resultados), len(esperados))
		return
	}

	for i, resultado := range resultados {
		if resultado != esperados[i] {
			t.Errorf("resultados[%d] = '%s', esperado '%s'", i, resultado, esperados[i])
		}
	}
}

func TestEjemploSelect(t *testing.T) {
	// Ejecuta varias veces para probar que funciona
	for i := 0; i < 3; i++ {
		resultado := EjemploSelect()

		if resultado != "from ch1" && resultado != "from ch2" && resultado != "timeout" {
			t.Errorf("EjemploSelect() = '%s', debe ser 'from ch1', 'from ch2' o 'timeout'", resultado)
		}

		// Debería recibir "from ch1" la mayoría de las veces (es más rápido)
		if resultado == "from ch1" {
			break // Test exitoso
		}
	}
}

func TestEjemploSelectNonBlocking(t *testing.T) {
	resultado := EjemploSelectNonBlocking()

	if resultado != "no data available" {
		t.Errorf("EjemploSelectNonBlocking() = '%s', esperado 'no data available'", resultado)
	}
}

func TestEjemploFanOut(t *testing.T) {
	resultados := EjemploFanOut()

	if len(resultados) != 3 {
		t.Errorf("len(resultados) = %d, esperado 3", len(resultados))
		return
	}

	// Cada consumidor debería haber procesado algunos números
	for i, consumerResults := range resultados {
		if len(consumerResults) == 0 {
			t.Errorf("Consumer %d no procesó ningún dato", i)
		}

		for _, result := range consumerResults {
			if !strings.Contains(result, "processed") {
				t.Errorf("Resultado '%s' debe contener 'processed'", result)
			}
		}
	}
}

func TestEjemploPipeline(t *testing.T) {
	resultados := EjemploPipeline()

	if len(resultados) != 5 {
		t.Errorf("len(resultados) = %d, esperado 5", len(resultados))
		return
	}

	// Los números deberían ser multiplicados por 2 y formateados
	esperados := []string{"processed: 2", "processed: 4", "processed: 6", "processed: 8", "processed: 10"}

	for i, resultado := range resultados {
		if resultado != esperados[i] {
			t.Errorf("resultados[%d] = '%s', esperado '%s'", i, resultado, esperados[i])
		}
	}
}

func TestEjemploContextConChannel(t *testing.T) {
	start := time.Now()
	resultado := EjemploContextConChannel()
	duration := time.Since(start)

	// Debería cancelarse después de ~100ms
	if duration > 150*time.Millisecond {
		t.Errorf("Duración = %v, esperado ~100ms", duration)
	}

	if resultado != "tarea cancelada" {
		t.Errorf("resultado = '%s', esperado 'tarea cancelada'", resultado)
	}
}

func TestEjemploSemaforo(t *testing.T) {
	start := time.Now()
	resultados := EjemploSemaforo()
	duration := time.Since(start)

	if len(resultados) != 5 {
		t.Errorf("len(resultados) = %d, esperado 5", len(resultados))
	}

	// Con semáforo de 2, 5 tareas de 100ms cada una deberían tomar ~300ms
	// (primera ronda: 2 tareas en paralelo = 100ms, segunda: 2 tareas = 100ms, tercera: 1 tarea = 100ms)
	expectedMin := 250 * time.Millisecond
	if duration < expectedMin {
		t.Errorf("Duración = %v, esperado al menos %v (semáforo debería limitar concurrencia)", duration, expectedMin)
	}

	// Verifica que todas las tareas se completaron
	for i, resultado := range resultados {
		if !strings.Contains(resultado, "completada") {
			t.Errorf("resultados[%d] = '%s', debe contener 'completada'", i, resultado)
		}
	}
}
