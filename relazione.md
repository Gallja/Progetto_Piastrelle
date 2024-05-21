# GALLIANO ANDREA 05460A - RELAZIONE PROGETTO DI LABORATORIO DI ALGORITMI E STRUTTURE DATI

### Indice
- [Introduzione](#introduzione)
- [Strutture utilizzate](#strutture-utilizzate)  
    - [Piano](#il-piano)
    - [Piastrelle](#le-piastrelle)
    - [Colorazione](#la-colorazione)
    - [Regole di propagazione](#le-regole)


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
Come abbiamo visto per la prima struttura, ogni **piastrella** accesa facente parte del **piano**, è necessario avere a disposizione 2 dati oltre le sue coordinate: l'*intensità* con cui è accesa nel **piano** ed il *colore*.  

```Go
type colorazione struct {
    coefficiente int
    colore string
}
```

#### Le regole
Le **regole di propagazione** da applicare alle **piastrelle accese** nel **piano**, necessitano di 3 campi per poter essere rappresentate con una struttura: gli *addendi* che formano la regola, il *colore* che assume la **piastrella** dopo l'applicazione della regola ed il *consumo* (ovvero il numero di volte che la regola è stata applicata; questo campo permette di ordinare le regole in maniera non decrescente).  

```Go
type regolaSingola struct {
    addendi []colorazione
    coloreFinale string
    consumo int
}
```

