// Package packages introduce el sistema de paquetes en Go
// En Python usas import, from...import, en Go es más simple y estricto
package packages

import (
	"fmt" // Alias para strings
	// Alias para math
	"strings"
)

// PASO 1: Exportación (visibility)
// En Python: todo es público, _private por convención
// En Go: Mayúscula = público, minúscula = privado

// TODO: Variable pública (exportada)
var VariablePublica = "__FILL_ME__" // Debe empezar con mayúscula

// TODO: Variable privada (no exportada)
var variablePrivada = "secreto"

// TODO: Función pública
func FuncionPublica() string {
	// TODO: Retorna "función pública"
	return "__FILL_ME__"
}

// TODO: Función privada
func funcionPrivada() string {
	return "función privada"
}

// TODO: Función que use la función privada
func UsarFuncionPrivada() string {
	// TODO: Llama funcionPrivada() y retorna su resultado
	return funcionPrivada()
}

// PASO 2: Structs y métodos exportados
type PersonaPublica struct {
	Nombre string // Público
	edad   int    // Privado
}

// TODO: Constructor público
func NuevaPersonaPublica(nombre string, edad int) *PersonaPublica {
	// TODO: Retorna nueva PersonaPublica
	return &PersonaPublica{Nombre: nombre, edad: edad}
}

// TODO: Método público para obtener edad privada
func (p *PersonaPublica) ObtenerEdad() int {
	// TODO: Retorna p.edad
	return 0
}

// TODO: Método público para modificar edad privada
func (p *PersonaPublica) EstablecerEdad(nuevaEdad int) {
	// TODO: Asigna p.edad = nuevaEdad
}

// PASO 3: Uso de paquetes estándar
// En Python: import math; math.sqrt(16)
// En Go: import "math"; math.Sqrt(16)

func EjemplosPackagesMath() []float64 {
	// TODO: Usa funciones del paquete math
	var resultados []float64

	// TODO: Agrega math.Sqrt(16)
	// resultados = append(resultados, math.Sqrt(16))

	// TODO: Agrega math.Pow(2, 3)
	// resultados = append(resultados, math.Pow(2, 3))

	// TODO: Agrega math.Abs(-5)
	// resultados = append(resultados, math.Abs(-5))

	return resultados
}

func EjemplosPackagesStrings() []string {
	texto := "Hola Mundo Go"
	var resultados []string

	// TODO: Usa funciones del paquete strings

	// TODO: Agrega strings.ToUpper(texto)
	// resultados = append(resultados, strings.ToUpper(texto))

	// TODO: Agrega strings.Contains(texto, "Go")
	if strings.Contains(texto, "Go") {
		resultados = append(resultados, "contiene Go")
	}

	// TODO: Agrega strings.Replace(texto, "Mundo", "Universe", 1)
	// resultados = append(resultados, strings.Replace(texto, "Mundo", "Universe", 1))

	return resultados
}

func EjemplosPackagesTime() []string {
	var resultados []string

	// TODO: Usa funciones del paquete time

	// TODO: Obtén tiempo actual y formatealo
	// ahora := time.Now()
	// resultados = append(resultados, ahora.Format("2006-01-02"))

	// TODO: Crea una duración y conviértela a string
	// duracion := 5 * time.Second
	// resultados = append(resultados, duracion.String())

	// TODO: Agrega día de la semana
	// resultados = append(resultados, time.Now().Weekday().String())

	return resultados
}

// PASO 4: Alias de paquetes
// En Python: import numpy as np
// En Go: import np "some/package"

func EjemplosAlias() []string {
	// TODO: Usa los alias str y m definidos en los imports
	// texto := "hello world"
	// var resultados []string
	// resultados = append(resultados, str.Title(texto))
	// resultados = append(resultados, fmt.Sprintf("%.0f", m.Ceil(3.7)))
	return []string{}
} // PASO 5: Blank import (solo para efectos secundarios)
// En Python: import module  # solo para ejecutar código
// En Go: import _ "package"

// En este ejemplo simularemos el concepto
var inicializacionCompleta bool

func init() {
	// init() se ejecuta automáticamente al importar el paquete
	inicializacionCompleta = true
	fmt.Println("Paquete inicializado")
}

func PackageInicializado() bool {
	return inicializacionCompleta
}

// PASO 6: Organización de archivos (simulado)
// En un proyecto real tendrías múltiples archivos en el mismo paquete

// types.go (simulado)
type Configuracion struct {
	Host   string
	Puerto int
}

// utils.go (simulado)
func validarConfiguracion(config *Configuracion) bool {
	return config.Host != "" && config.Puerto > 0
}

// main.go (simulado)
func CrearConfiguracion(host string, puerto int) (*Configuracion, error) {
	config := &Configuracion{Host: host, Puerto: puerto}

	if !validarConfiguracion(config) {
		return nil, fmt.Errorf("configuración inválida")
	}

	return config, nil
}

// PASO 7: Documentación con comentarios
// En Python: docstrings """Documentación"""
// En Go: comentarios antes de la declaración

// CalcularAreaCirculo calcula el área de un círculo dado su radio.
// Parámetros:
//   - radio: el radio del círculo (debe ser positivo)
//
// Retorna:
//   - el área del círculo
//   - error si el radio es negativo
func CalcularAreaCirculo(radio float64) (float64, error) {
	if radio < 0 {
		return 0, fmt.Errorf("el radio no puede ser negativo")
	}

	// TODO: Retorna math.Pi * radio * radio, nil
	return 0, nil
}

// PASO 8: Constantes exportadas
// En Python: CONSTANT = value (por convención)
// En Go: const Constant = value (real constante)

const (
	// TODO: Define constantes exportadas
	VersionMajor = 0 // Cambia por 1
	VersionMinor = 0 // Cambia por 2
	VersionPatch = 0 // Cambia por 3
)

func ObtenerVersion() string {
	return fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
}

// PASO 9: Variables exportadas de solo lectura (usando funciones)
// En Python: no hay const real, se simula con propiedades
// En Go: puedes usar funciones para "getters" de solo lectura

var configuracionInterna = map[string]string{
	"env":     "development",
	"debug":   "true",
	"timeout": "30s",
}

// TODO: Función para obtener configuración (solo lectura)
func ObtenerConfiguracion(clave string) string {
	// TODO: Retorna configuracionInterna[clave]
	return "__FILL_ME__"
}

// TODO: Función para obtener todas las claves de configuración
func ObtenerClavesConfiguracion() []string {
	var claves []string
	// TODO: Itera sobre configuracionInterna y agrega claves
	for clave := range configuracionInterna {
		claves = append(claves, clave)
	}
	return claves
}

// PASO 10: Patrón de inicialización de paquete
var (
	// Variables que se inicializan una vez
	instanciaSingleton *Configuracion
	inicializado       bool
)

// TODO: Función de inicialización (debe llamarse antes de usar el paquete)
func InicializarPaquete(host string, puerto int) error {
	if inicializado {
		return fmt.Errorf("paquete ya inicializado")
	}

	// TODO: Crea configuración y asigna a singleton
	// config, err := CrearConfiguracion(host, puerto)
	// if err != nil {
	//     return err
	// }
	// instanciaSingleton = config
	// inicializado = true

	return nil
} // TODO: Función para obtener la instancia singleton
func ObtenerInstancia() (*Configuracion, error) {
	if !inicializado {
		return nil, fmt.Errorf("paquete no inicializado")
	}

	// TODO: Retorna instanciaSingleton, nil
	return nil, nil
}
