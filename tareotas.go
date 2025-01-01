//.........----..........//
//Guardar y revisar tareas//
//.......................//

//......Stack:
package main

import (
	"fmt"
	"os"
	"strings"
)

//...............................................................|Comienza código
// ...func principal donde se ejecuta el programa
func main() {
	// creamos una lista de tareas
	tks := []string{"leer", "escribir", "codificar", "revisar", "debuggear"}
	
	// mostramos las tareas iniciales
	fmt.Println("tareas iniciales:")
	mostrar(tks)

	// pedimos al usuario que ingrese una nueva tarea
	fmt.Println("\n¿qué tarea quieres agregar?")
	var nt string
	fmt.Scanln(&nt)
	tks = agregar(nt, tks)

	// mostramos las tareas luego de agregar una nueva
	fmt.Println("\ntareas después de agregar:")
	mostrar(tks)

	// eliminamos una tarea
	fmt.Println("\n¿qué tarea quieres eliminar?")
	var dt string
	fmt.Scanln(&dt)
	tks = eliminar(dt, tks)

	// mostramos las tareas después de eliminar
	fmt.Println("\ntareas después de eliminar:")
	mostrar(tks)

	// contamos el número de tareas
	cont := contar(tks)
	fmt.Printf("\nNúmero de tareas restantes: %d\n", cont)

	// guardamos las tareas en un archivo
	err := guardar(tks)
	if err != nil {
		fmt.Println("error al guardar las tareas:", err)
		return
	}

	// cargamos las tareas desde el archivo
	fmt.Println("\nCargando tareas guardadas...")
	tksC, err := cargar()
	if err != nil {
		fmt.Println("error al cargar las tareas:", err)
		return
	}
	fmt.Println("\ntareas cargadas:")
	mostrar(tksC)
}

// func que muestra las tareas
func mostrar(tks []string) {
	for i, t := range tks {
		fmt.Printf("%d. %s\n", i+1, t)
	}
}

// func para agregar una tarea
func agregar(tk string, tks []string) []string {
	tks = append(tks, tk)
	return tks
}

// func para eliminar una tarea
func eliminar(tk string, tks []string) []string {
	var idx int
	for i, t := range tks {
		if t == tk {
			idx = i
			break
		}
	}
	tks = append(tks[:idx], tks[idx+1:]...)
	return tks
}

// func que cuenta las tareas
func contar(tks []string) int {
	return len(tks)
}

// func que guarda las tareas en un archivo
func guardar(tks []string) error {
	file, err := os.Create("tareas.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	for _, t := range tks {
		_, err := file.WriteString(t + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

// func que carga las tareas desde un archivo
func cargar() ([]string, error) {
	file, err := os.Open("tareas.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var tks []string
	var linea string
	for {
		_, err := fmt.Fscanf(file, "%s\n", &linea)
		if err != nil {
			break
		}
		tks = append(tks, linea)
	}
	return tks, nil
}

//................................|Final
