package main

import (
	"bufio"
	"fmt"
	"os"
)

type piastrella struct {
	x      int
	y      int
	colore string
	accesa bool
}

type addendoRegola struct {
	coefficiente int
	colore       string
}

type piano struct {
	piastrelle []piastrella
	regole     []addendoRegola
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	parseInput(scanner)

	fmt.Println("ciao mondo")
}

func parseInput(scanner *bufio.Scanner) {
	for scanner.Scan() {
		input := scanner.Text()

		if input == "q" {
			break
		}

		if len(input) > 0 {
			command := input[:1]
			// argument := strings.TrimSpace(input[1:])

			switch command {
			case "C":
				fmt.Println("colora")
			case "S":
				fmt.Println("spegni")
			case "r":
				fmt.Println("regola")
			case "?":
				fmt.Println("stato")
			case "s":
				fmt.Println("stampa")
			default:
				fmt.Println("comando non valido")
			}
		}
	}
}
