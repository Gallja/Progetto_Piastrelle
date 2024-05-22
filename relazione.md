# GALLIANO ANDREA 05460A - RELAZIONE PROGETTO DI LABORATORIO DI ALGORITMI E STRUTTURE DATI

### Indice
- [Introduzione](#introduzione)
- [Strutture utilizzate](#strutture-utilizzate)  
    - [Piano](#il-piano)
    - [Piastrelle](#le-piastrelle)
    - [Colorazione](#la-colorazione)
    - [Regole di propagazione](#le-regole)  
- [Le funzioni principali](#le-funzioni-principali)
    - [Colora](#colora)
    - [Spegni](#spegni)
    - [Regola](#regola)
    - [Stato](#stato)
    - [Stampa](#stampa)
    - [Blocco](#blocco)
    - [Blocco Omogeneo](#blocco-omogeneo)
    - [Propaga](#propaga)
    - [Propaga Blocco](#propaga-blocco)
    - [Ordina](#ordina)


## Introduzione
Per poter affrontare ragionevolmente il problmema, sono state utilizzate apposite strutture che risolvessero tutti i punti richiesti e che rappresentassero fedelmente il piano descritto all'interno della traccia.  

### Strutture utilizzate

#### Il piano

Per poter rappresentare fedelmente il **piano** contenente le piastrelle a cui poter applicare le **regole di propagazione**, è stato necessario utilizzare una struttura che, prima di tutto, avesse un campo che mettesse in relazione le coordinate **_(x, y)_** di una piastrella e i dati relativi all'*intensità con cui è accesa* ed il *colore*.  
Per questo motivo, il primo campo del **piano** è il *puntatore all'indirizzo di memoria di una mappa dalla piastrella alla corrispondente colorazione*.  
Il secondo campo della struttura è invece il *puntatore all'indirizzo di una slice di regole*, che torna utile nel momento in cui si decide di applicare una **regola di propagazione** a una o più piastrelle.

```Go
type piano struct {
    piastrelle *map[piastrella]colorazione
	regole     *[]regolaSingola
}
```

#### Le piastrelle
Le **piastrelle** sono state pensate come una struttura i cui campi sono 2 interi rappresentati le coordinate **(x, y)** della stessa all'interno del piano.  

```Go
type piastrella struct {
    x int
    y int
}
```

#### La colorazione
Come abbiamo visto per la prima struttura, per ogni **piastrella** accesa facente parte del **piano**, è necessario avere a disposizione 2 dati oltre le sue coordinate: l'*intensità* con cui è accesa nel **piano** ed il *colore*.  

```Go
type colorazione struct {
    coefficiente int
    colore string
}
```

#### Le regole
Le **regole di propagazione** da applicare alle **piastrelle accese** nel **piano**, necessitano di 3 campi per poter essere rappresentate con una struttura: gli *addendi* che formano la regola, il *colore* che assume la **piastrella** dopo l'applicazione della regola ed il *consumo* (ovvero il numero di volte che la regola è stata applicata; questo campo permette di **ordinare** le regole in maniera **non decrescente**).  

```Go
type regolaSingola struct {
    addendi []colorazione
    coloreFinale string
    consumo int
}
```

### Le funzioni principali
Le funzioni implementate all'interno del programma, a fronte di un apposito input con i giusti comandi, permettono di modificare il piano e prestando particolare attenzioni all'uso delle risorse sia spaziali che temporali.  

#### Colora

```Go
func colora(p piano, x int, y int, alpha string, i int) {
    // implementazione di "colora"
}
```

La funzione **_colora_** riceve come parametri il **piano**, le coordinate intere **x** e **y**, un **colore** e l'**intensità** con cui si intende colorare la *piastrella*.  
Per effettuare l'operazione di *colorazione*, viene assegnata alla mappa contenente le *piastrelle* nel **piano** il valore della corrispondente **colorazione**.  
- **Analisi del tempo**: l'accesso alla mappa ha costo **_O(1)_** in termini di tempo. 
- **Analisi dello spazio**: non viene allocato alcuno spazio, di conseguenza il costo in termini di spazio è costante e nell'ordine di **_O(1)_**.

#### Spegni

```Go
func spegni(p piano, x int, y int) {
    // implementazione di "spegni"
}
```

La funzione **_spegni_** permette di spegnere una piastrella che, al momento, si trova (accesa) all'interno del **piano** con intensità ≥ 1.  
Per eseguire da codice questa operazione, ciò che viene fatto è un'*eliminazione della piastrella avuta per argomento tramite coordinate*.  
- **Analisi del tempo**: Anche l'operazione di *delete* dalla mappa impiega tempo costante, di conseguenza la complessità temporale è nell'ordine di **_O(1)_**.  
- **Analisi dello spazio**: Come per la complessità temporale, anche l'uso dello spazio è costante: **_O(1)_**.  

#### Regola

```Go
func regola(p piano, r string) {
    // implementazione di "regola"
}
```

La funzione **_regola_** permette, dati in ingresso il **piano** ed una **stringa**, di aggiungere una nuova regola all'interno del sistema.  
Per poterlo fare, è necessario, in primo luogo, effettuare un _parsing_ della stringa avuta per argomento, successivamente creare la regola (composta dai suoi 3 campi analizzati in precedenza) e, infine, *aggiungere la regola appena creata alla slice di regole facenti già parti del piano*.  
- **Analisi del tempo**:   
- **Analisi dello spazio**:  

#### Stato

```Go
func stato(p piano, x int, y int) (string, int) {
    // implementazione di "stato"
}
```
La funzione **stato** *restituisce e stampa i valori relativi al colore e l'intensità della piastrella delle coordinate avute per argomento*.  
Per farlo, assegno ad una variabile il valore della mappa contenente le piastrelle del piano e un'altra, di tipo *bool*, per stampare (e, conseguentemente, anche ritornare) **se e solo se quella piastrella esiste nel piano**.  
- **Analisi del tempo**:   
- **Analisi dello spazio**: 


#### Stampa

```Go
func stampa(p piano) {
    // implementazione di "stampa"
}
```
La funzione **stampa** mostra tutte le **regole** del **piano** nel seguente formato:  
(  
*coloreFinale: coefficiente1 colore1 coefficiente2 colore2 ...*  
*coloreFinale: coefficiente1 colore1 coefficiente2 colore2 ...*  
.  
.  
.  
)  
- **Analisi del tempo**:  
- **Analisi dello spazio**:  


#### Blocco

```Go
func blocco(p piano, x, y int) {
    // implementazione di "blocco"
}
```

#### Blocco Omogeneo

```Go
func bloccoOmog(p piano, x, y int) {
    // implementazione di "bloccoOmog"
}
```

#### Propaga

```Go
func propaga(p piano, x, y int) {
    // implementazione di "propaga"
}
```

#### Propaga Blocco

```Go
func propagaBlocco(p piano, x, y int) {
    // implementazione di "propagaBlocco"
}
```

#### Ordina

```Go
func ordina(p piano) {
    // implementazione di "ordina"
}
```