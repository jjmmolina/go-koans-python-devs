// Package channels introduce los channels en Go
// En Python usas queue.Queue o asyncio.Queue, en Go los channels están integrados
package channels

import (
	"context"
	"fmt"
	"time"
)

// PASO 1: Channel básico
// En Python: import queue; q = queue.Queue(); q.put(item); item = q.get()
// En Go: ch := make(chan int); ch <- valor; valor := <-ch

// TODO: Función que envía un valor por un channel
func EnviarPorChannel(ch chan int, valor int) {
	// TODO: Envía valor por el channel
	// ch <- valor
}

// TODO: Función que recibe un valor de un channel
func RecibirDeChannel(ch chan int) int {
	// TODO: Recibe y retorna valor del channel
	// return <-ch
	return 0
}

func EjemploCommunicacionBasica() int {
	ch := make(chan int)

	// Enviar en una goroutine
	go EnviarPorChannel(ch, 42)

	// Recibir en la goroutine principal
	return RecibirDeChannel(ch)
}

// PASO 2: Buffered channels
// En Python: queue.Queue(maxsize=n)
// En Go: make(chan int, capacity)

func EjemploBufferedChannel() []int {
	// TODO: Crea un channel con buffer de tamaño 3 y envía/recibe 3 valores
	// ch := make(chan int, 3)
	// ch <- 1
	// ch <- 2
	// ch <- 3

	// var resultados []int
	// resultados = append(resultados, <-ch)
	// resultados = append(resultados, <-ch)
	// resultados = append(resultados, <-ch)

	return []int{}
} // PASO 3: Channel como parámetro (direccional)
// En Go puedes especificar si un channel es solo para envío o recepción

// TODO: Función que solo puede enviar por el channel
func ProducirDatos(ch chan<- string, datos []string) {
	// TODO: Itera sobre datos, envía por channel y cierra
	// for _, dato := range datos {
	//     ch <- dato
	// }
	// close(ch)
}

// TODO: Función que solo puede recibir del channel
func ConsumirDatos(ch <-chan string) []string {
	var resultados []string
	// TODO: Recibe todos los datos hasta que se cierre el channel
	// for dato := range ch {
	//     resultados = append(resultados, dato)
	// }
	return resultados
}

func EjemploProducirConsumir() []string {
	ch := make(chan string, 5)
	datos := []string{"a", "b", "c", "d", "e"}

	go ProducirDatos(ch, datos)
	return ConsumirDatos(ch)
}

// PASO 4: Select statement
// En Python: no hay equivalente directo, usarías asyncio.wait
// En Go: select permite esperar en múltiples channels

func EjemploSelect() string {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	// Envía datos después de diferentes delays
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	// TODO: Usa select para recibir del primer channel que esté listo
	select {
	case msg1 := <-ch1:
		return msg1
	case msg2 := <-ch2:
		return msg2
	case <-time.After(300 * time.Millisecond):
		return "timeout"
	}
}

// PASO 5: Select con default (non-blocking)
func EjemploSelectNonBlocking() string {
	ch := make(chan string)

	// TODO: Usa select con default para no bloquear
	select {
	case msg := <-ch:
		return msg
	default:
		// TODO: Retorna "no data available"
		return "__FILL_ME__"
	}
}

// PASO 6: Fan-out pattern (un productor, múltiples consumidores)
func FanOutProducer(input chan int, outputs []chan int) {
	// TODO: Envía cada número del input a todos los outputs
	// for num := range input {
	//     for _, output := range outputs {
	//         output <- num
	//     }
	// }
	// for _, output := range outputs {
	//     close(output)
	// }
}

func FanOutConsumer(id int, input chan int) []string {
	var resultados []string
	// TODO: Procesa números del canal
	// for num := range input {
	//     resultados = append(resultados, fmt.Sprintf("consumer %d processed %d", id, num))
	// }
	return resultados
}

func EjemploFanOut() [][]string {
	input := make(chan int, 5)
	outputs := make([]chan int, 3)

	// Crea channels de salida
	for i := range outputs {
		outputs[i] = make(chan int, 5)
	}

	// Produce datos
	go func() {
		for i := 1; i <= 5; i++ {
			input <- i
		}
		close(input)
	}()

	// Fan-out
	go FanOutProducer(input, outputs)

	// Consume con múltiples consumidores
	resultados := make([][]string, 3)
	done := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go func(consumerID int) {
			resultados[consumerID] = FanOutConsumer(consumerID, outputs[consumerID])
			done <- true
		}(i)
	}

	// Espera que todos terminen
	for i := 0; i < 3; i++ {
		<-done
	}

	return resultados
}

// PASO 7: Pipeline pattern
func Pipeline1(input chan int) chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		// TODO: Multiplica cada número por 2
		// for num := range input {
		//     output <- num * 2
		// }
	}()
	return output
}

func Pipeline2(input chan int) chan string {
	output := make(chan string)
	go func() {
		defer close(output)
		// TODO: Convierte a string con formato
		// for num := range input {
		//     output <- fmt.Sprintf("processed: %d", num)
		// }
	}()
	return output
}

func EjemploPipeline() []string {
	// Crea el pipeline
	input := make(chan int)
	stage1 := Pipeline1(input)
	stage2 := Pipeline2(stage1)

	// Envía datos
	go func() {
		defer close(input)
		for i := 1; i <= 5; i++ {
			input <- i
		}
	}()

	// Recolecta resultados
	var resultados []string
	for resultado := range stage2 {
		resultados = append(resultados, resultado)
	}

	return resultados
}

// PASO 8: Context con channels

func TareaConContext(ctx context.Context, result chan string) {
	select {
	case <-time.After(200 * time.Millisecond):
		// TODO: Envía "tarea completada" si no hay cancelación
		// result <- "tarea completada"
	case <-ctx.Done():
		// TODO: Envía "tarea cancelada" si hay cancelación
		// result <- "tarea cancelada"
	}
}

func EjemploContextConChannel() string {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	result := make(chan string, 1)
	go TareaConContext(ctx, result)

	return <-result
}

// PASO 9: Semáforo usando channels
type Semaforo chan struct{}

func NuevoSemaforo(capacity int) Semaforo {
	// TODO: Crea un channel buffered de struct{} con la capacidad dada
	return make(chan struct{}, capacity)
}

func (s Semaforo) Acquire() {
	// TODO: Envía un struct{} para "tomar" un permit
	// s <- struct{}{}
}

func (s Semaforo) Release() {
	// TODO: Recibe un struct{} para "liberar" un permit
	// <-s
}

func TareaConSemaforo(id int, sem Semaforo, resultados chan string) {
	sem.Acquire()
	defer sem.Release()

	// Simula trabajo
	time.Sleep(100 * time.Millisecond)
	resultados <- fmt.Sprintf("tarea %d completada", id)
}

func EjemploSemaforo() []string {
	sem := NuevoSemaforo(2) // Solo 2 tareas concurrentes
	resultados := make(chan string, 5)

	// Lanza 5 tareas
	for i := 1; i <= 5; i++ {
		go TareaConSemaforo(i, sem, resultados)
	}

	// Recolecta resultados
	var output []string
	for i := 0; i < 5; i++ {
		output = append(output, <-resultados)
	}

	return output
}
