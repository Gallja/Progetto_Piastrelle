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

type colorazione struct {
	coefficiente int
	colore       string
}

type regolaSingola struct {
	addendi      []colorazione
	coloreFinale string
}

type piano struct {
	piastrelle *map[piastrella]colorazione
	regole     *[]regolaSingola
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	p := creaPiano()

	for scanner.Scan() {
		input := scanner.Text()
		esegui(p, input)
	}
}

func creaPiano() piano {
	mappa := make(map[piastrella]colorazione)
	return piano{&mappa, &[]regolaSingola{}}
}

func esegui(p piano, s string) {
	if s == "q" {
		return
	}

	if len(s) > 0 {
		command := s[:1]
		argument := strings.TrimSpace(s[1:])

		x, y, colore := parseInput(argument)

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

func parseInput(argument string) (int, int, string) {
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
	(*mappa)[piastrella{x, y}] = colorazione{1, alpha}
}

func spegni(p piano, x int, y int) {
	delete(*p.piastrelle, piastrella{x, y})
}

func regola(p piano, r string) {
	args := strings.Split(r, " ")

	var newReg regolaSingola
	newReg.coloreFinale = args[0]
	addReg := colorazione{}

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
	fmt.Print(r.coloreFinale)

	for i := 0; i < len(r.addendi); i++ {
		fmt.Print(" ", r.addendi[i].coefficiente, " ", r.addendi[i].colore)
	}

	fmt.Println()
}
