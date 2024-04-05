package main

import (
	"fmt"
	"math/rand"
)

func show_tablero(tablero []string) {
	fmt.Println("\n                    Tablero")
	fmt.Printf("               ┌─────┐  ┌─────┐\n")
	fmt.Printf("               │  %s  │  │  %s  │\n", tablero[0], tablero[1])
	fmt.Printf("               └─────┘  └─────┘\n")
	fmt.Printf("               ┌─────┐  ┌─────┐\n")
	fmt.Printf("               │  %s  │  │  %s  │\n", tablero[2], tablero[3])
	fmt.Printf("               └─────┘  └─────┘\n")
}

func letra_azar() string {
	letras := []string{"A", "B", "C", "D"}

	indice := rand.Intn(len(letras))

	return letras[indice]

}

func stringNoEncontrado(array []string, valor string) bool {
	for _, elemento := range array {
		if elemento == valor {
			return false
		}
	}
	return true
}
func verificar_input(tab []string, input string) bool {
	for _, v := range tab {
		if v == input {
			return true
		}
	}
	fmt.Println("Entrada inválida. Inténtalo de nuevo.")
	return false
}

func verificar_pos(tablero []string, posicion string, letra_intento string) bool {

	if posicion == letra_intento {
		for i := range tablero {
			if tablero[i] == letra_intento {
				tablero[i] = "0"
			}
		}

		show_tablero(tablero)
		fmt.Print("\n HAY UN GANADOR !!! \n")
		return true
	} else {
		for i := range tablero {
			if tablero[i] == letra_intento {
				tablero[i] = "X"
			}
		}

		show_tablero(tablero)
		fmt.Print("\n Rayos! \n")
		return false
	}
}
func lore() {
	fmt.Print("\nBienvenid@s al juego del pirata, veamos quién puede sobrevivir en altamar\n")
	fmt.Print("\n")
	fmt.Print("Dame tu ubicación, sólo puedes contestar con: A, B, C o D: ")
	fmt.Print("\n")

}

func mov_cliente(tab []string, ind string) bool {
	var input string
	fmt.Println("Tu turno")
	fmt.Println("Adivina donde estoy!")
	fmt.Println("Escoge una letra:")

	for {
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error al leer la entrada:", err)
		} else if verificar_input(tab, input) {
			break
		}
	}

	return verificar_pos(tab, ind, input)
}

func mov_server(tabla []string, pos string) bool {
	fmt.Println("Es mi turno")
	x := letra_azar()

	for {
		if verificar_input(tabla, x) {
			break
		}
		x = letra_azar()
	}

	return verificar_pos(tabla, pos, x)
}

// func main() {
// 	lore()

// 	tablarrayS := []string{"A", "B", "C", "D"}
// 	tablarrayC := []string{"A", "B", "C", "D"}
// 	show_tablero(tablarrayS)
// 	fmt.Print("\n")

// 	var input string
// 	for {
// 		_, err := fmt.Scanln(&input)
// 		if err != nil {
// 			fmt.Println("Error al leer la entrada:", err)
// 		} else if verificar_input(tablarrayC, input) {
// 			break
// 		}
// 	}

// 	pos_server := letra_azar()

// 	for {

// 		if mov_cliente(tablarrayS, pos_server) {
// 			break
// 		}

// 		if mov_server(tablarrayC, input) {
// 			break
// 		}
// 	}
// }
