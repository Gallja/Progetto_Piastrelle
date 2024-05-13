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

type regolaSingola struct {
	addendi     []addendoRegola
	targetColor string
}

type piano struct {
	piastrelle *map[piastrella]addendoRegola
	regole     *[]regolaSingola
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

			x, y, colore := parametrizzaInput(argument)

			switch command {
			case "C":
				colora(p, x, y, colore)
			case "S":
				spegni(p, x, y)
			case "r":
				regola(p, argument)
			case "?":
				stato(p, x, y)
			case "s":
				stampa(p)
			default:
				return
			}
		}
	}
}

func creaPiano() piano {
	mappa := make(map[piastrella]addendoRegola)
	return piano{&mappa, &[]regolaSingola{}}
}

func parametrizzaInput(argument string) (int, int, string) {
	if len(argument) <= 1 {
		return 0, 0, ""
	}

	if len(argument) < 4 {
		args := strings.Split(argument, " ")
		x, _ := strconv.Atoi(args[0])
		y, _ := strconv.Atoi(args[1])

		return x, y, ""
	}

	args := strings.Split(argument, " ")
	x, _ := strconv.Atoi(args[0])
	y, _ := strconv.Atoi(args[1])
	colore := args[2]

	return x, y, colore
}

func colora(p piano, x int, y int, alpha string) {
	mappa := p.piastrelle
	(*mappa)[piastrella{x, y}] = addendoRegola{1, alpha}
}

func spegni(p piano, x int, y int) {
	delete(*p.piastrelle, piastrella{x, y})
}

func regola(p piano, r string) {
	args := strings.Split(r, " ")

	var newReg regolaSingola
	newReg.targetColor = args[0]
	addReg := addendoRegola{}

	for i := 1; i < len(args); i++ {
		if i%2 != 0 { // è un colore
			addReg.coefficiente, _ = strconv.Atoi(args[i])
		} else { // è un valore intero
			addReg.colore = args[i]
			newReg.addendi = append(newReg.addendi, addReg)
		}
	}

	*p.regole = append(*(p.regole), newReg)
}

func stato(p piano, x int, y int) (string, int) {
	val, ok := (*p.piastrelle)[piastrella{x, y}]

	if ok {
		fmt.Println(val.colore, val.coefficiente)
	}

	return val.colore, val.coefficiente
}

func stampa(p piano) {
	fmt.Println("(")

	for _, v := range *p.regole {
		stampaRegola(v)
	}

	fmt.Println(")")
}

func stampaRegola(r regolaSingola) {
	fmt.Print(r.targetColor)

	for i := 0; i < len(r.addendi); i++ {
		fmt.Print(" ", r.addendi[i].coefficiente, " ", r.addendi[i].colore)
	}

	fmt.Println()
}
