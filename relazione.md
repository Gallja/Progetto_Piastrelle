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
- [Esempi di esecuzione](#esempi-di-esecuzione)


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
    colore       string
}
```

#### Le regole
Le **regole di propagazione** da applicare alle **piastrelle accese** nel **piano**, necessitano di 3 campi per poter essere rappresentate con una struttura: gli *addendi* che formano la regola, il *colore* che assume la **piastrella** dopo l'applicazione della regola ed il *consumo* (ovvero il numero di volte che la regola è stata applicata; questo campo permette di **ordinare** le regole in maniera **non decrescente**).  

```Go
type regolaSingola struct {
    addendi      []colorazione
    coloreFinale string
    consumo      int
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
- **Analisi del tempo**: Dal punto di vista del tempo, questa funzione è nell'ordine di **_O(1)_**, poiché tutte le operazioni che effettua (ovvero la restituzione di un valore della *mappa di piastrelle*, di un valore *bool* che indichi se quel valore esiste, il controllo prima della stmpa e il ritorno finale di **colore** e **intensità** della piastrella) impiegano tempo costante.  
- **Analisi dello spazio**: Anche lo spazio allocato, a livello di variabili dichiarate e memoria utilizzata, da parte di **stato** è nell'ordine di **_O(1)_**.


#### Stampa

```Go
func stampa(p piano) {
    // implementazione di "stampa"
}
```

La funzione **stampa** mostra tutte le **regole** del **piano** nel seguente formato:  
(    
*coloreFinale 1: coefficiente1 colore1 coefficiente2 colore2 ...*  
*coloreFinale 2: coefficiente1 colore1 coefficiente2 colore2 ...*  
.  
.  
.  
*coloreFinale n: coefficiente1 colore1 coefficiente2 colore2 ...*  
)  
Ciò che fa la funzione, a livello di codice, è *scorrere la slice di regole del **piano** e, per ognuna di essere scorrere gli addendi che la compongono stampando infine il coefficiente ed il colore dell'addendo* (separando opportunamente entrambi con uno spazio).

- **Analisi del tempo**: Questa funzione contiene 2 cicli: il primo scorre le **regole** nel **piano**, mentre il secondo scorre gli **addendi** di ogni regola. Il primo ciclo ha complessità **_O(n)_** (con *n = numero di regole nel piano*) ed il secondo effettua sempre, al più, 8 iterazioni (questo perché, per definizione del piano e dell'intorno di ogni piastrella con *max piastrelle circonvicine = 8*, **una regola di propagazione non può avere più di 8 addendi**).  
Di conseguenza, la complessità temporale totale della funzione **stampa** è pari a **_O(n) x O(8)_ = O(n)**.  
- **Analisi dello spazio**:   


#### Blocco

```Go
func blocco(p piano, x, y int) {
    // implementazione di "blocco"
}
```

La funzione **blocco** stampa la somma delle intensità delle piastrelle facenti parte del medesimo blocco; per poterlo fare con complessità spaziali e temporali contenute, è stato necessario partire dalle coordinate **_(x, y)_** di una piastrella avuta per argomento per poi _effettuare una **visita in ampiezza ("Breadth-First-Search")** ed avere a disposizione le piastrelle circonvicine del blocco_.  

La ricerca degli adiacenti o delle piastrelle circonvicine ad un'altra, le cui coordinate **_(x, y)_** sono passate per argomento ad un'apposita funzione **_"cercaAdiacenti"_**, non fa altro che _scorrere tutte le possibili **8 combinazioni** di coordinate di piastrelle circonvicine per poi restituirle all'interno di una slice di piastrelle_.

```Go
func cercaAdiacenti(p piano, piastrella_ piastrella) []piastrella {
    // le 8 combinazioni possibili per ogni piastrella:
    combX := []int{-1, 0, 0, 1, -1, -1, 1, 1}
    combY := []int{-1, -1, 1, -1, 1, 0, 0, 1}

    // implementazione di "cercaAdiacenti"
}
```

Per effettuare la *visita in ampiezza*, è stata inoltre utilizzata una **coda**, in cui vengono salvate temporaneamente le piastrelle visitate e dalle quali si andrà a visitarne le circonvicine.  
La struttura dati **coda**, con campi e funzioni scritte all'interno di un file a parte chiamato **_"queue.go"_**, è definita così:  

```Go
type queue struct {
	head *queueNode
	tail *queueNode
}

type queueNode struct {
	next  *queueNode
	value piastrella
}
```
La memorizzazione di **_tail_** è particolarmente utile all'interno della funzione di **_enqueue_**, poiché permette di _**NON** scorrere tutta la coda per aggiungere un elemento, ma di avere direttamente un puntatore all'ultimo nodo ed effettuare l'aggiunta risparmiando in termini di complessità temporale_.  

Verrà tenuto conto delle piastrelle già visitate salvandole permanentemente all'interno di una **mappa usata come _set_**.  
Dato che, però, all'interno del linguaggio **Go** _non esiste **"Set"** come vero e proprio tipo_, ecco come è stato realizzato:  

```Go
visitate := make(map[piastrella]struct{})
```

Questa mappa, **da piastrella a struct vuota**, permette di memorizzare solo le chiavi, in modo tale da trattare la struttura dati come un vero e proprio *set di piastrelle già visitate durante la BFS*.  
Viene utilizzata una *struct vuota* al posto di una *variabile di tipo bool* **per poter risparmiare ulteriormente spazio in memoria**.  

- **Analisi del tempo**: 
- **Analisi dello spazio**: 

#### Blocco Omogeneo

```Go
func bloccoOmog(p piano, x, y int) {
    // implementazione di "bloccoOmog"
}
```

La funzione **bloccoOmog** _stampa la somma delle intensità delle piastrelle circonvicine facenti parte dello stesso blocco, utilizzando lo stesso principio di funzionamento della funzione **blocco**_.  

Proprio allo scopo di _fattorizzare_ la parte di implementazione comune a **blocco** e **bloccoOmog**, entrambe le funzioni utilizzano una funzione "comune" chiamata **_"bloccoGenerico"_**, _alla quale viene passato un parametro di tipo **bool** (omogeneo = True/False) e che ha una condizione che valuta quando incrementare il valore della somma delle intensità_.  

```Go
func bloccoGenerico(p piano, x, y int, omogeneo bool) (int, []piastrella) {
    // implementazione di "bloccoGenerico"
}
```

A questo punto, è facile dedurre che le prestazioni riguardanti il *tempo*, che quelle riguardanti lo *spazio* non variano rispetto alla funzione **_blocco_**.
- **Analisi del tempo**: 
- **Analisi dello spazio**: 

#### Propaga

```Go
func propaga(p piano, x, y int) {
    // implementazione di "propaga"
}
```

La funzione **propaga** permette di applicare, ad una piastrella le cui coordinate **_(x,y)_** vengono passate per argomento, _la prima **regola di propagazione** disponibile dell'elenco di regole nel **piano**_.  
Ciò che viene fatto è 

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

La funzione **ordina** permette di _ordinare le **regole di propagazione** del **piano** in ordine **non decrescente** in base al consumo delle regole stesse_. Per fare l'ordinamento, è stata utilizzata la funzione di libreria di **Go** [SortStableFunc](https://pkg.go.dev/slices#SortStableFunc), che permette di ordinare **in maniera stabile** riscrivendo il **comparatore** per confrontare gli elementi di una slice allo stesso modo della funzione [SortFunc](https://pkg.go.dev/slices#SortFunc).  

- **Analisi del tempo**: L'ordinamento delle regole in base al loro consumo è basato su confronti e, nel caso peggiore, non si può scendere al di sotto dell'ordine di **_O(n log n)_**.
- **Analisi dello spazio**: Essendo un algoritmo di ordinamento **_in-place_**, non utilizza spazio ulteriore per la creazione di copie di slice, di conseguenza la funzione **ordina** utilizza solo un **puntatore alla slice da ordinare** ed è nell'ordine di **_O(1)_**.  

### Esempi di esecuzione
Per testare il corretto funzionamento del programma e le sue prestazioni, sono stati scritti ulteriori _file di input_ con i relativi _file di output_.  
Per questi esempi è stata inoltre creata una griglia per avere una **visualizzazione grafica** del piano per capire come viene modificato a fronte dei comandi in input.