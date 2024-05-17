package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type piastrella struct {
	x int
	y int
}

type colorazione struct {
	coefficiente int    // intensità
	colore       string // {a, b, ..., z} --> colore
}

type regolaSingola struct {
	addendi      []colorazione
	coloreFinale string
	consumo      int
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

		x, y, colore, intensita := parseInput(argument)

		switch command {
		case "C":
			colora(p, x, y, colore, intensita)
		case "S":
			spegni(p, x, y)
		case "r":
			regola(p, argument)
		case "?":
			stato(p, x, y)
		case "s":
			stampa(p)
		case "b":
			sommaIntensita := blocco(p, x, y)
			fmt.Println(sommaIntensita)
		case "B":
			sommaIntensita := bloccoOmog(p, x, y)
			fmt.Println(sommaIntensita)
		case "p":
			propaga(p, x, y)
		case "o":
			ordina(p)
		default:
			return
		}
	}
}

func parseInput(argument string) (int, int, string, int) {
	if len(argument) <= 1 {
		return 0, 0, "", 0
	}

	if len(argument) < 4 {
		args := strings.Split(argument, " ")
		x, _ := strconv.Atoi(args[0])
		y, _ := strconv.Atoi(args[1])

		return x, y, "", 0
	}

	args := strings.Split(argument, " ")
	x, _ := strconv.Atoi(args[0])
	y, _ := strconv.Atoi(args[1])
	colore := args[2]
	intensita, _ := strconv.Atoi(args[3])

	return x, y, colore, intensita
}

func colora(p piano, x int, y int, alpha string, i int) {
	mappa := p.piastrelle
	(*mappa)[piastrella{x, y}] = colorazione{i, alpha}
}

func spegni(p piano, x int, y int) {
	delete(*p.piastrelle, piastrella{x, y})
}

func regola(p piano, r string) {
	args := strings.Split(r, " ")

	var newReg regolaSingola
	newReg.consumo = 0
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
	fmt.Print(r.coloreFinale, ":")

	for i := 0; i < len(r.addendi); i++ {
		fmt.Print(" ", r.addendi[i].coefficiente, " ", r.addendi[i].colore)
	}

	fmt.Println()
}

func bloccoGenerico(p piano, x, y int, omogeneo bool) int {
	mappa := p.piastrelle
	start, ok := (*mappa)[piastrella{x, y}]
	sommaIntensita := 0

	// piastrella spenta
	if !ok {
		return 0
	}

	sommaIntensita += start.coefficiente

	// ricerca del blocco a partire dalle coordinate avute per argomento
	coda := queue{nil}
	coda.enqueue(piastrella{x, y})

	visitate := make(map[piastrella]struct{})

	for !coda.isEmpty() {
		piastrella_ := coda.dequeue()
		visitate[piastrella_] = struct{}{} // utilizzo una mappa che contiene solo le chiavi (e come valori una struct vuota) per utilizzarla come set
		adiacenti := cercaAdiacenti(p, piastrella_)

		for i := 0; i < len(adiacenti); i++ {
			_, ok := visitate[adiacenti[i]]

			if !ok {
				visitate[adiacenti[i]] = struct{}{}
				val := (*mappa)[adiacenti[i]]

				if !omogeneo || val.colore == start.colore {
					sommaIntensita += val.coefficiente
					coda.enqueue(adiacenti[i])
				}
			}

		}

	}

	return sommaIntensita
}

func cercaAdiacenti(p piano, piastrella_ piastrella) []piastrella {
	sliceRet := []piastrella{}
	mappa := p.piastrelle

	// le 8 combinazioni possibili per ogni piastrella:
	diffX := []int{-1, 0, 0, 1, -1, -1, 1, 1}
	diffY := []int{-1, -1, 1, -1, 1, 0, 0, 1}

	for i := 0; i < len(diffX); i++ {
		_, ok := (*mappa)[piastrella{piastrella_.x + diffX[i], piastrella_.y + diffY[i]}]

		if ok {
			sliceRet = append(sliceRet, piastrella{piastrella_.x + diffX[i], piastrella_.y + diffY[i]})
		}
	}

	return sliceRet
}

func blocco(p piano, x, y int) int {
	return bloccoGenerico(p, x, y, false)
}

func bloccoOmog(p piano, x, y int) int {
	return bloccoGenerico(p, x, y, true)
}

func propaga(p piano, x, y int) {
	mappa := p.piastrelle

	val := (*mappa)[piastrella{x, y}]

	adiacenti := cercaAdiacenti(p, piastrella{x, y})
	regole := (*p.regole)

	for i := 0; i < len(regole); i++ {
		rispettata := true

		for j := 0; j < len(regole[i].addendi); j++ {
			coeffBkcp := regole[i].addendi[j].coefficiente

			for k := 0; k < len(adiacenti) && coeffBkcp > 0; k++ {

				if (*mappa)[adiacenti[k]].colore == regole[i].addendi[j].colore {
					coeffBkcp--
				}

			}

			if coeffBkcp > 0 {
				rispettata = false
				break
			}
		}

		if rispettata {
			(*mappa)[piastrella{x, y}] = colorazione{val.coefficiente, regole[i].coloreFinale}

			regole[i].consumo++

			break
		}

	}
}

func ordina(p piano) {
	regole := *p.regole

	sort.SliceStable(regole, func(i, j int) bool {
		return regole[i].consumo < regole[j].consumo
	})
}
