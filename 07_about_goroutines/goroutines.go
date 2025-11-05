// Package goroutines introduce la concurrencia en Go
// En Python usas async/await o threading, en Go usas goroutines
package goroutines

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// PASO 1: Goroutine básica
// En Python: import asyncio; asyncio.create_task(funcion())
// En Go: go funcion()

var resultado string

// TODO: Función que será ejecutada como goroutine
func TareaSimple() {
	time.Sleep(100 * time.Millisecond)
	// TODO: Asigna resultado = "tarea completada"
}

func EjecutarGoroutineSimple() {
	// TODO: Ejecuta TareaSimple como goroutine usando 'go'
	// go TareaSimple()

	// Esperar un poco para que termine
	time.Sleep(200 * time.Millisecond)
}

// PASO 2: WaitGroup - esperar que terminen las goroutines
// En Python: await asyncio.gather(*tasks)
// En Go: sync.WaitGroup

func TareaConWaitGroup(id int, wg *sync.WaitGroup, resultados []string) {
	// TODO: Marca que esta goroutine terminó al final
	// defer wg.Done()

	time.Sleep(50 * time.Millisecond)
	// TODO: Asigna resultados[id] = fmt.Sprintf("tarea %d completada", id)
}

func EjecutarVariasGoroutines() []string {
	// TODO: Crea un WaitGroup y un slice para resultados
	// var wg sync.WaitGroup
	// resultados := make([]string, 3)

	// TODO: Itera 3 veces, incrementa WaitGroup y ejecuta goroutines
	// for i := 0; i < 3; i++ {
	//     wg.Add(1)
	//     go TareaConWaitGroup(i, &wg, resultados)
	// }

	// TODO: Espera que terminen todas las goroutines
	// wg.Wait()

	return make([]string, 3)
} // PASO 3: Mutex - proteger datos compartidos
// En Python: import threading; lock = threading.Lock(); with lock: ...
// En Go: sync.Mutex

type ContadorSeguro struct {
	mutex sync.Mutex
	valor int
}

// TODO: Método para incrementar de forma segura
func (c *ContadorSeguro) Incrementar() {
	// TODO: Bloquea el mutex, incrementa valor, desbloquea
	// c.mutex.Lock()
	// defer c.mutex.Unlock()
	// c.valor++
}

// TODO: Método para obtener valor de forma segura
func (c *ContadorSeguro) ObtenerValor() int {
	// TODO: Bloquea el mutex, lee valor, desbloquea y retorna
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.valor
}

func IncrementarConcurrente() int {
	contador := &ContadorSeguro{}
	var wg sync.WaitGroup

	// 1000 goroutines incrementando
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			contador.Incrementar()
		}()
	}

	wg.Wait()
	return contador.ObtenerValor()
}

// PASO 4: Race condition (para demostrar por qué necesitamos mutex)
type ContadorInseguro struct {
	valor int
}

func (c *ContadorInseguro) Incrementar() {
	// TODO: Incrementa valor sin protección (race condition)
	c.valor++
}

func (c *ContadorInseguro) ObtenerValor() int {
	return c.valor
}

func IncrementarInseguro() int {
	contador := &ContadorInseguro{}
	var wg sync.WaitGroup

	// 1000 goroutines incrementando sin protección
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			contador.Incrementar()
		}()
	}

	wg.Wait()
	return contador.ObtenerValor()
}

// PASO 5: RWMutex - lectores múltiples, escritor único
// En Python: threading.RLock() es similar pero no igual
type Cache struct {
	mutex sync.RWMutex
	datos map[string]string
}

func NuevoCache() *Cache {
	return &Cache{
		datos: make(map[string]string),
	}
}

// TODO: Método para escribir (necesita lock exclusivo)
func (c *Cache) Escribir(clave, valor string) {
	// TODO: Usa c.mutex.Lock() y defer c.mutex.Unlock()
	// c.datos[clave] = valor
}

// TODO: Método para leer (puede ser concurrente)
func (c *Cache) Leer(clave string) (string, bool) {
	// TODO: Usa c.mutex.RLock() y defer c.mutex.RUnlock()
	// valor, existe := c.datos[clave]
	// return valor, existe
	return "", false
}

// PASO 6: Context para cancelación
// En Python: asyncio.CancelledError
// En Go: context.Context

func TareaLarga(ctx context.Context, id int) string {
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			// TODO: Retorna mensaje de cancelación
			return fmt.Sprintf("tarea %d cancelada", id)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
	// TODO: Retorna mensaje de completado
	return fmt.Sprintf("tarea %d completada", id)
}

func EjecutarConCancelacion() []string {
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()

	var wg sync.WaitGroup
	resultados := make([]string, 3)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			resultados[id] = TareaLarga(ctx, id)
		}(i)
	}

	wg.Wait()
	return resultados
}

// PASO 7: Worker pool pattern
// En Python: concurrent.futures.ThreadPoolExecutor
func Worker(id int, trabajos <-chan int, resultados chan<- string) {
	// TODO: Itera sobre el canal de trabajos
	// for trabajo := range trabajos {
	//     time.Sleep(50 * time.Millisecond)
	//     resultados <- fmt.Sprintf("worker %d procesó trabajo %d", id, trabajo)
	// }
}

func EjecutarWorkerPool() []string {
	const numWorkers = 3
	const numTrabajos = 5

	// TODO: Crea canales de trabajos y resultados
	// trabajos := make(chan int, numTrabajos)
	// resultados := make(chan string, numTrabajos)

	// TODO: Inicia workers
	// for w := 1; w <= numWorkers; w++ {
	//     go Worker(w, trabajos, resultados)
	// }

	// TODO: Envía trabajos y cierra el canal
	// for j := 1; j <= numTrabajos; j++ {
	//     trabajos <- j
	// }
	// close(trabajos)

	// TODO: Recolecta resultados
	// var output []string
	// for a := 1; a <= numTrabajos; a++ {
	//     output = append(output, <-resultados)
	// }

	return []string{}
}

// PASO 8: Rate limiting
func RateLimiter() []string {
	// TODO: Implementa rate limiting
	// rate := time.Second / 2
	// limiter := time.NewTicker(rate)
	// defer limiter.Stop()

	// var resultados []string
	// for i := 0; i < 5; i++ {
	//     <-limiter.C
	//     resultados = append(resultados, fmt.Sprintf("operación %d a las %v", i, time.Now().Format("15:04:05")))
	// }

	return []string{}
}
