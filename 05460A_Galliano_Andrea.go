package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type piastrella struct {
	x int
	y int
}

type addendoRegola struct {
	coefficiente int
	colore       string
}

type piano struct {
	piastrelle map[piastrella]string
	regole     []addendoRegola
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	parseInput(scanner)
}

func parseInput(scanner *bufio.Scanner) {
	p := creaPiano()

	for scanner.Scan() {
		input := scanner.Text()

		if input == "q" {
			break
		}

		if len(input) > 0 {
			command := input[:1]
			argument := strings.TrimSpace(input[1:])

			switch command {
			case "C":
				args := strings.Split(argument, " ")
				x, _ := strconv.Atoi(args[0])
				y, _ := strconv.Atoi(args[1])
				colora(p, x, y, args[2])
			case "S":
				args := strings.Split(argument, " ")
				x, _ := strconv.Atoi(args[0])
				y, _ := strconv.Atoi(args[1])
				spegni(p, x, y)
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

func creaPiano() piano {
	return piano{}
}

func colora(p piano, x int, y int, alpha string) {
	p.piastrelle[piastrella{x, y}] = alpha
}

func spegni(p piano, x int, y int) {
	delete(p.piastrelle, piastrella{x, y})
}

func regola(p piano, r string) {

}

func stato(p piano, x int, y int) (string, int) {
	return "", 0
}

func stampa(p piano) {

}
