/*
Grafiquitos con colores y posibilidad de ver algunas funciones en consola
*/

//.......| STACK
package main

import (
	"fmt"
	"math"
	"time"
)

// ..tamaño de la pantalla
const (
	wid = 50 // ancho
	hei = 20 // alto
)

// ..estructura para definir los colores
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	White  = "\033[97m"
)

// .función para limpiar la pantalla
func clr() {
	fmt.Print("\033[H\033[2J")
}

// inicializamos la pantalla
func ini() [][]string {
	pix := make([][]string, hei)
	for i := range pix {
		pix[i] = make([]string, wid)
		for j := range pix[i] {
			pix[i][j] = " " // espacio vacío
		}
	}
	return pix
}

// función para dibujar la pantalla con la gráfica
func dib(pix [][]string) {
	for _, fila := range pix {
		for _, col := range fila {
			fmt.Print(col)
		}
		fmt.Println()
	}
}

// función para agregar el eje
func eje(pix [][]string) {
	// dibujamos el eje X
	for i := 0; i < wid; i++ {
		pix[hei/2][i] = White + "-" // eje X en blanco
	}
	// dibujamos el eje Y
	for i := 0; i < hei; i++ {
		pix[i][wid/2] = White + "|" // eje Y en blanco
	}
	pix[hei/2][wid/2] = White + "+" // punto de origen
}

// graficamos una función matemática
func graficarFuncion(pix [][]string, f func(float64) float64, color string, xMin, xMax, step float64) {
	for x := xMin; x < xMax; x += step {
		// convertimos a coordenadas de píxeles
		screenX := int((x - xMin) * float64(wid) / (xMax - xMin))
		screenY := int((f(x) - -2) * float64(hei) / 4) // Limites de y de -2 a 2 (ajustar según la función)

		if screenX >= 0 && screenX < wid && screenY >= 0 && screenY < hei {
			pix[hei-1-screenY][screenX] = color + "[]"
		}
	}
}

// función para mostrar la tabla de valores
func tabla(f func(float64) float64, xMin, xMax, step float64) {
	fmt.Println("Tabla de valores:")
	for x := xMin; x < xMax; x += step {
		y := f(x)
		fmt.Printf("x: %.2f, y: %.2f\n", x, y)
	}
}

// funciones matemáticas
func seno(x float64) float64 {
	return math.Sin(x)
}

func coseno(x float64) float64 {
	return math.Cos(x)
}

func tangente(x float64) float64 {
	return math.Tan(x)
}

func exponencial(x float64) float64 {
	return math.Exp(x)
}

func logaritmo(x float64) float64 {
	return math.Log(x)
}

// función para mostrar el menú de selección de funciones
func menu() int {
	fmt.Println("Elija una función para graficar:")
	fmt.Println("1 - Seno")
	fmt.Println("2 - Coseno")
	fmt.Println("3 - Tangente")
	fmt.Println("4 - Exponencial")
	fmt.Println("5 - Logaritmo")
	fmt.Println("6 - Función Lineal")
	fmt.Println("7 - Función Cuadrática")
	fmt.Print("Seleccione una opción: ")

	var opcion int
	fmt.Scan(&opcion)
	return opcion
}

func main() {
	// inicializamos la pantalla
	pix := ini()

	// limpiamos la pantalla
	clr()

	// elige la función a graficar
	opcion := menu()

	// establecemos los parámetros para la gráfica
	var f func(float64) float64
	var color string
	var xMin, xMax, step float64

	switch opcion {
	case 1:
		f = seno
		color = Blue
		xMin, xMax, step = -3.0, 3.0, 0.1
	case 2:
		f = coseno
		color = Green
		xMin, xMax, step = -3.0, 3.0, 0.1
	case 3:
		f = tangente
		color = Yellow
		xMin, xMax, step = -3.0, 3.0, 0.1
	case 4:
		f = exponencial
		color = Red
		xMin, xMax, step = 0.0, 3.0, 0.1
	case 5:
		f = logaritmo
		color = Blue
		xMin, xMax, step = 0.1, 3.0, 0.1
	case 6:
		f = func(x float64) float64 { return 2*x + 1 } // ejemplo de función lineal
		color = Green
		xMin, xMax, step = -3.0, 3.0, 0.1
	case 7:
		f = func(x float64) float64 { return x*x - 2*x + 1 } // ejemplo de función cuadrática
		color = Yellow
		xMin, xMax, step = -3.0, 3.0, 0.1
	default:
		fmt.Println("Opción inválida.")
		return
	}

	// .mostramos la tabla de valores
	tabla(f, xMin, xMax, step)

	// .dibujamos los ejes
	eje(pix)

	// .graficamos la función seleccionada
	graficarFuncion(pix, f, color, xMin, xMax, step)

	// .mostramos la pantalla con la gráfica
	dib(pix)

	// .animación simple
	for i := 0; i < 5; i++ {
		clr()
		fmt.Println("Animando...")
		graficarFuncion(pix, f, color, xMin, xMax, step)
		dib(pix)
		time.Sleep(500 * time.Millisecond) // .espera de 500ms antes de redibujar
	}
}
//.....................................| FIN/THE END
