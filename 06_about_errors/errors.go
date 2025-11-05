// Package errors introduce el manejo de errores en Go
// En Python usas try/except, en Go usas valores de retorno múltiples
package errors

import (
	"errors"
	"fmt"
	"strconv"
)

// PASO 1: Error básico
// En Python: raise ValueError("mensaje")
// En Go: return value, error

// TODO: Función que divide pero retorna error si divisor es 0
func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		// TODO: Retorna 0 y un error usando errors.New("división por cero")
		return 0, nil
	}
	return a / b, nil
}

// PASO 2: Crear errores personalizados
// En Python: class MiError(Exception): pass
// En Go: type MiError struct{}; func (e MiError) Error() string

type ErrorEdadInvalida struct {
	Edad int
}

// TODO: Implementa el método Error() para ErrorEdadInvalida
func (e ErrorEdadInvalida) Error() string {
	// TODO: Retorna mensaje como "edad inválida: 150"
	return "__FILL_ME__"
}

// TODO: Función que valide edad (0-120)
func ValidarEdad(edad int) error {
	if edad < 0 || edad > 120 {
		// TODO: Retorna ErrorEdadInvalida{Edad: edad}
		return nil
	}
	return nil
}

// PASO 3: Verificación de errores
// En Python: try...except
// En Go: if err != nil

// TODO: Función que use Dividir y maneje el error
func DividirSeguro(a, b float64) string {
	// TODO: Llama a Dividir y maneja el error
	// resultado, err := Dividir(a, b)
	// if err != nil {
	//     return "Error: " + err.Error()
	// }
	// return fmt.Sprintf("%g", resultado)
	return "__FILL_ME__"
}

// PASO 4: Propagación de errores
// En Python: re-raise exception
// En Go: return err

func ConvertirYDividir(aStr, bStr string) (float64, error) {
	// TODO: Convierte aStr a float64 usando strconv.ParseFloat
	// TODO: Si hay error, retorna 0 y el error
	// TODO: Convierte bStr a float64
	// TODO: Si hay error, retorna 0 y el error
	// TODO: Llama a Dividir y retorna el resultado

	// a, err := strconv.ParseFloat(aStr, 64)
	// if err != nil {
	//     return 0, err
	// }
	// b, err := strconv.ParseFloat(bStr, 64)
	// if err != nil {
	//     return 0, err
	// }
	// return Dividir(a, b)

	return 0, nil
} // PASO 5: Errores con contexto usando fmt.Errorf
func ConvertirConContexto(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		// TODO: Retorna error con contexto usando fmt.Errorf
		// "no se pudo convertir '%s' a entero: %w", str, err
		return 0, nil
	}
	return num, nil
}

// PASO 6: Múltiples tipos de error
type ErrorCalculadora struct {
	Operacion string
	Motivo    string
}

func (e ErrorCalculadora) Error() string {
	return fmt.Sprintf("error en %s: %s", e.Operacion, e.Motivo)
}

// TODO: Función calculadora que puede tener diferentes errores
func Calcular(operacion string, a, b float64) (float64, error) {
	switch operacion {
	case "suma":
		return a + b, nil
	case "resta":
		return a - b, nil
	case "multiplicacion":
		return a * b, nil
	case "division":
		if b == 0 {
			// TODO: Retorna ErrorCalculadora{"division", "división por cero"}
			return 0, nil
		}
		return a / b, nil
	default:
		// TODO: Retorna ErrorCalculadora{operacion, "operación no soportada"}
		return 0, nil
	}
}

// PASO 7: Defer para cleanup
// En Python: try/finally o context managers
// En Go: defer

func ProcesarArchivo(nombre string) error {
	// Simulamos abrir un archivo
	fmt.Printf("Abriendo archivo: %s\n", nombre)

	// TODO: Usa defer para simular cerrar el archivo
	// defer fmt.Printf("Cerrando archivo: %s\n", nombre)

	if nombre == "archivo_malo.txt" {
		return errors.New("archivo corrupto")
	}

	fmt.Println("Procesando archivo...")
	return nil
}

// PASO 8: Panic y recover (equivalente a excepciones no manejadas)
// En Python: uncaught exception
// En Go: panic (para errores irrecuperables)

func FuncionPeligrosa(panico bool) (resultado string) {
	// TODO: Usa defer con recover para manejar panic
	defer func() {
		if r := recover(); r != nil {
			// TODO: Asigna resultado = "se recuperó del pánico"
		}
	}()

	if panico {
		panic("¡algo salió muy mal!")
	}

	return "todo bien"
}

// PASO 9: Verificar tipo de error
// En Python: isinstance(e, MiError)
// En Go: errors.Is() y errors.As()

func ManejarErrorEspecifico(err error) string {
	// TODO: Verifica si err es de tipo ErrorEdadInvalida
	var edadErr ErrorEdadInvalida
	if errors.As(err, &edadErr) {
		// TODO: Retorna "Error de edad: " + err.Error()
		return "__FILL_ME__"
	}

	// TODO: Para otros errores, retorna "Error genérico: " + err.Error()
	return "__FILL_ME__"
}

// PASO 10: Sentinel errors (errores predefinidos)
// En Python: usar constantes de error
// En Go: var ErrAlgo = errors.New("algo")

var (
	ErrSaldoInsuficiente = errors.New("saldo insuficiente")
	ErrCuentaInexistente = errors.New("cuenta inexistente")
)

type Cuenta struct {
	Saldo float64
}

func (c *Cuenta) Retirar(cantidad float64) error {
	if cantidad > c.Saldo {
		// TODO: Retorna ErrSaldoInsuficiente
		return nil
	}
	c.Saldo -= cantidad
	return nil
}

func VerificarTipoError(err error) string {
	// TODO: Usa errors.Is() para verificar si err es ErrSaldoInsuficiente
	if errors.Is(err, ErrSaldoInsuficiente) {
		return "saldo insuficiente"
	}
	// TODO: Verifica ErrCuentaInexistente
	if errors.Is(err, ErrCuentaInexistente) {
		return "cuenta inexistente"
	}
	return "error desconocido"
}
