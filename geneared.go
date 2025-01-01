/*
Pa visualizar un árbol genealógico.
*/
package main

//............................| Stack
import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

//............................| Código
// ..estructura de una persona en el árbol genealógico
type Psn struct {
	ID       string
	Nombre   string
	Padres   []string
	Hijos    []string
	Genero   string
}

var fam map[string]Psn // mapa para almacenar los miembros de la familia

// lee el CSV y carga los datos en el mapa
func cargarCSV(archivo string) error {
	file, err := os.Open(archivo)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // omitir cabecera
	if err != nil {
		return err
	}

	// lee todas las filas
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// procesa cada fila
	for _, row := range rows {
		id := row[0]
		nom := row[1]
		gen := row[2]
		padres := strings.Split(row[3], ",")
		hijos := strings.Split(row[4], ",")

		fam[id] = Psn{ID: id, Nombre: nom, Padres: padres, Hijos: hijos, Genero: gen}
	}

	return nil
}

// agrega un nuevo miembro al árbol
func agregarMiembro(id, nom, gen string, padres, hijos []string) {
	fam[id] = Psn{ID: id, Nombre: nom, Genero: gen, Padres: padres, Hijos: hijos}
}

// muestra el árbol genealógico
func mostrarArbol() {
	for _, psn := range fam {
		fmt.Printf("ID: %s | Nombre: %s | Género: %s | Padres: %v | Hijos: %v\n", psn.ID, psn.Nombre, psn.Genero, psn.Padres, psn.Hijos)
	}
}

func main() {
	fam = make(map[string]Psn) // inicializa el mapa de familia

	// carga el archivo CSV
	err := cargarCSV("familia.csv")
	if err != nil {
		fmt.Println("Error al cargar archivo CSV:", err)
		return
	}

	// muestra el árbol genealógico
	mostrarArbol()

	// agregar nuevo miembro
	agregarMiembro("5", "Carlos", "Masculino", []string{"2", "3"}, []string{"6", "7"})

	// muestra el árbol después de agregar al nuevo miembro
	fmt.Println("\nÁrbol actualizado:")
	mostrarArbol()
}
