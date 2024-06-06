package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	piastrelle map[piastrella]colorazione
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
	piastrelleInPiano := make(map[piastrella]colorazione)
	return piano{piastrelleInPiano, &[]regolaSingola{}}
}

func esegui(p piano, s string) {
	if s == "q" {
		os.Exit(0)
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
			blocco(p, x, y)
		case "B":
			bloccoOmog(p, x, y)
		case "p":
			propaga(p, x, y)
		case "P":
			propagaBlocco(p, x, y)
		case "o":
			ordina(p)
			// fmt.Println(p.regole)
		default:
			return
		}
	}
}

func parseInput(argument string) (int, int, string, int) {
	if len(argument) <= 1 {
		return 0, 0, "", 0
	}

	args := strings.Split(argument, " ")

	if len(args) < 4 {
		x, _ := strconv.Atoi(args[0])
		y, _ := strconv.Atoi(args[1])

		return x, y, "", 0
	}

	x, _ := strconv.Atoi(args[0])
	y, _ := strconv.Atoi(args[1])
	colore := args[2]
	intensita, _ := strconv.Atoi(args[3])

	return x, y, colore, intensita
}

func colora(p piano, x int, y int, alpha string, i int) {
	p.piastrelle[piastrella{x, y}] = colorazione{i, alpha}
}

func spegni(p piano, x int, y int) {
	delete(p.piastrelle, piastrella{x, y})
}

func regola(p piano, r string) {
	args := strings.Split(r, " ")

	var nuovaRegola regolaSingola
	nuovaRegola.consumo = 0
	nuovaRegola.coloreFinale = args[0]
	addendoRegola := colorazione{}

	for i := 1; i < len(args); i++ {
		if i%2 != 0 { // è un valore intero
			addendoRegola.coefficiente, _ = strconv.Atoi(args[i])
		} else { // è un colore
			addendoRegola.colore = args[i]
			nuovaRegola.addendi = append(nuovaRegola.addendi, addendoRegola)
		}
	}

	*p.regole = append(*(p.regole), nuovaRegola)
}

func stato(p piano, x int, y int) (string, int) {
	val, ok := (p.piastrelle)[piastrella{x, y}]

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

func bloccoGenerico(p piano, x, y int, omogeneo bool) (int, []piastrella) {
	piastrelleInPiano := p.piastrelle
	start, ok := piastrelleInPiano[piastrella{x, y}]
	sommaIntensita := 0

	piastrelleBlocco := []piastrella{}
	piastrelleBlocco = append(piastrelleBlocco, piastrella{x, y})

	// piastrella spenta
	if !ok {
		return 0, nil
	}

	sommaIntensita += start.coefficiente

	// ricerca del blocco a partire dalle coordinate avute per argomento
	coda := queue{nil, nil}
	coda.enqueue(piastrella{x, y})

	visitate := make(map[piastrella]struct{})

	for !coda.isEmpty() {
		piastrella_ := coda.dequeue()
		visitate[piastrella_] = struct{}{} // utilizzo una mappa che contiene solo le chiavi (e come valori una struct vuota) per utilizzarla come set
		// utilizzo la struct vuota poiché in Go non vi sono tipi Set; avrei potuto usare un booleano come valore, ma con una struct vuota riesco a risparmiare spazio in memoria
		adiacenti := cercaAdiacenti(p, piastrella_)

		for i := 0; i < len(adiacenti); i++ {
			_, ok := visitate[adiacenti[i]]

			if !ok {
				visitate[adiacenti[i]] = struct{}{}
				val := piastrelleInPiano[adiacenti[i]]

				if !omogeneo || val.colore == start.colore {
					sommaIntensita += val.coefficiente
					piastrelleBlocco = append(piastrelleBlocco, adiacenti[i])
					coda.enqueue(adiacenti[i])
				}
			}

		}

	}

	return sommaIntensita, piastrelleBlocco
}

func cercaAdiacenti(p piano, piastrella_ piastrella) []piastrella {
	circonvicine := []piastrella{}
	piastrelleInPiano := p.piastrelle

	// le 8 combinazioni possibili per ogni piastrella:
	combX := []int{-1, 0, 0, 1, -1, -1, 1, 1}
	combY := []int{-1, -1, 1, -1, 1, 0, 0, 1}

	for i := 0; i < len(combX); i++ {
		_, ok := piastrelleInPiano[piastrella{piastrella_.x + combX[i], piastrella_.y + combY[i]}]

		if ok {
			circonvicine = append(circonvicine, piastrella{piastrella_.x + combX[i], piastrella_.y + combY[i]})
		}
	}

	return circonvicine
}

func blocco(p piano, x, y int) {
	sommaIntensita, _ := bloccoGenerico(p, x, y, false)

	fmt.Println(sommaIntensita)
}

func bloccoOmog(p piano, x, y int) {
	sommaIntensita, _ := bloccoGenerico(p, x, y, true)

	fmt.Println(sommaIntensita)
}

func propagaGenerico(p piano, x, y int) map[piastrella]regolaSingola {
	piastrelleInPiano := p.piastrelle
	piastrelleRegole := make(map[piastrella]regolaSingola)

	adiacenti := cercaAdiacenti(p, piastrella{x, y})
	regole := (*p.regole)

	for i := 0; i < len(regole); i++ {
		rispettata := true

		for j := 0; j < len(regole[i].addendi); j++ {
			coeffAddendo := regole[i].addendi[j].coefficiente

			for k := 0; k < len(adiacenti) && coeffAddendo > 0; k++ {

				if piastrelleInPiano[adiacenti[k]].colore == regole[i].addendi[j].colore {
					coeffAddendo--
				}

			}

			if coeffAddendo > 0 {
				rispettata = false

				break
			}
		}

		if rispettata {
			regole[i].consumo++

			piastrelleRegole[piastrella{x, y}] = regole[i]

			break
		}

	}

	return piastrelleRegole
}

func propaga(p piano, x, y int) {
	coloraPiastrelle(p, propagaGenerico(p, x, y))
}

func propagaBlocco(p piano, x, y int) {
	_, piastrelleBlocco := bloccoGenerico(p, x, y, false)
	sliceCambiamenti := make([]map[piastrella]regolaSingola, 0)

	for i := 0; i < len(piastrelleBlocco); i++ {
		piastrelleRegole := propagaGenerico(p, piastrelleBlocco[i].x, piastrelleBlocco[i].y)

		if len(piastrelleRegole) > 0 {
			sliceCambiamenti = append(sliceCambiamenti, piastrelleRegole)
		}
	}

	for j := 0; j < len(sliceCambiamenti); j++ {
		coloraPiastrelle(p, sliceCambiamenti[j])
	}
}

func coloraPiastrelle(p piano, piastrelleRegole map[piastrella]regolaSingola) {
	piastrelleInPiano := p.piastrelle

	for k, v := range piastrelleRegole {
		val, ok := piastrelleInPiano[k]

		if !ok {
			colora(p, k.x, k.y, v.coloreFinale, 1)
		} else {
			colora(p, k.x, k.y, v.coloreFinale, val.coefficiente)
		}
	}
}

func ordina(p piano) {
	regole := *p.regole

	slices.SortStableFunc(regole, func(a, b regolaSingola) int {
		return a.consumo - b.consumo
	})
}
