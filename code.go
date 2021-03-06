package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type jogador struct {
	saldo string
}
type Lotes struct {
	Lotes []Lote `json:"lotes"`
}
type Lote struct {
	Nome   string `json:"nome"`
	Custo  int    `json:"custo"`
	Espaco int    `json:"espaco"`
}
type Brinquedos struct {
	Brinquedos []Brinquedo `json:"brinquedos"`
}
type Brinquedo struct {
	Nome          string  `json:"nome"`
	Custo         int     `json:"custo"`
	Popularidade  float64 `json:"popularidade"`
	Ingresso      int     `json:"ingresso"`
	Espaco        int     `json:"espaco"`
	QuantidadeMax int     `json:"quantidademax"`
	Code          int     `json:"code"`
}
type Quantidade struct {
	aux int
}

var fileScanner *bufio.Scanner
var ToysList []Brinquedo
var LotesList []Lote

func main() {
	quantidades := []int{}
	rand.Seed(time.Now().UnixNano())
	jsonFile, err := os.Open("brinquedos.json")
	jsonFile2, err2 := os.Open("lotes.json")
	if err != nil {
		log.Fatalf("Can't %s", err)
	}
	if err2 != nil {
		log.Fatalf("Can't %s", err2)
	}
	textoFile, _ := ioutil.ReadAll(jsonFile)
	textoFile2, _ := ioutil.ReadAll(jsonFile2)
	var brinquedos Brinquedos
	var lotes Lotes
	for e := 0; e < len(brinquedos.Brinquedos); e++ {
		quantidade := Quantidade{
			aux: brinquedos.Brinquedos[e].Code,
		}
		quantidades = append(quantidades, quantidade.aux)
	}
	json.Unmarshal(textoFile2, &lotes)
	json.Unmarshal(textoFile, &brinquedos)
	for e := 1; e < len(brinquedos.Brinquedos); e++ {
		brinquedo := Brinquedo{
			Nome:          brinquedos.Brinquedos[e].Nome,
			Custo:         brinquedos.Brinquedos[e].Custo,
			Popularidade:  brinquedos.Brinquedos[e].Popularidade,
			Ingresso:      brinquedos.Brinquedos[e].Ingresso,
			Espaco:        brinquedos.Brinquedos[e].Espaco,
			QuantidadeMax: brinquedos.Brinquedos[e].QuantidadeMax,
			Code:          brinquedos.Brinquedos[e].Code,
		}
		ToysList = append(ToysList, brinquedo)
	}
	for i := 0; i < len(lotes.Lotes); i++ {
		lote := Lote{
			Nome:   lotes.Lotes[i].Nome,
			Custo:  lotes.Lotes[i].Custo,
			Espaco: lotes.Lotes[i].Espaco,
		}
		LotesList = append(LotesList, lote)
	}
	player := jogador{
		saldo: "1000000",
	}
	var (
		Introdu????o        string
		comando           string
		comecou           bool
		vezes             int
		Opcoes            string
		MenuDeCompras     string
		Espaco_disponivel int
		MenuDeVendas      string
	)
	Opcoes = "1 - Menu de Compras \n2 - Menu de Vendas\n3 - Passar um dia\n4- Passar uma semana"
	MenuDeCompras = "O que deseja comprar?\n1- Lotes\n2-Brinquedos"
	MenuDeVendas = "O que deseja vender?\n1- Lotes\n2- Brinquedos"
	Introdu????o = "Bem Vinda(o) ao MyAmusementPark\n Neste jogo voc?? pode criar o seu propio parque de divers??es! \n Deseja Come??ar?\n 1-Sim\n 2-N??o"
	fmt.Println(Introdu????o)
	fmt.Scanf("%s", &comando)

	if comando == "s" || comando == "sim" || comando == "ss" || comando == "Sim" || comando == "1" {
		comecou = true
	} else {
		fmt.Printf("Ok, ent??o volte sempre que quiser\n")
	}
	if comecou {
		Espaco_disponivel = 1000
		saldo_jogador, err := strconv.Atoi(player.saldo)
		if err != nil {
			log.Fatal(err)
		}
		index := 1
		i := 0
		z := 0
		vezes = 0
		dia := 1

		myToys := []Brinquedo{}
		meusLotes := []Lote{}
		pessoas := []int{}
		dias := []int{}
		sas1 := []int{0, 0, 0, 00, 00, 00, 0, 00, 0, 0, 0, 0, 0, 0}
		sas2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		array := [][]int{sas1, sas2}
		//quantidades := []int{}
		//quantidades = append(quantidades,Quantidade)
		brinquedo2 := brinquedos.Brinquedos[0]
		var p int
		p = int(brinquedo2.Popularidade)
		array[0][0] = p
		array[1][0] = brinquedo2.Ingresso
		myToys = append(myToys, brinquedos.Brinquedos[0])
		pessoas = append(pessoas, population(brinquedo2.Popularidade))
		for comando != "Q" {
			z = 0
			fmt.Printf("Dia %d\n\n", dia)
			fmt.Printf("Espa??o dispon??vel: %d\n", Espaco_disponivel)
			fmt.Printf("Saldo: %d\n", saldo_jogador)
			fmt.Printf("Brinquedos Atuais: \n\n")
			for z != len(myToys) {
				fmt.Printf("%d - \nNome: %s\nPopularidade: %.0f\nCusto Mensal: %d\n", z+1, myToys[z].Nome, myToys[z].Popularidade, myToys[z].Custo)
				z++
			}
			if vezes > 1 {
				fmt.Printf("\n\nOque deseja fazer agora?\n")
			} else {
				fmt.Printf("\n\nOque deseja fazer logo de inicio?\n")
			}
			fmt.Println(Opcoes)
			fmt.Scanf("%s", &comando)
			switch comando {
			case "1":
				fmt.Println(MenuDeCompras)
				fmt.Scanf("%s", &comando)
				switch comando {
				case "1":
					listaLotes()
					fmt.Scanf("%d", &index)
					LoteComprado := LotesList[index-1]
					if LoteComprado.Custo > saldo_jogador {
						fmt.Printf("Voc?? n??o tem dinheiro o suficiente\n")
					}
					fmt.Printf("Parab??ns voce comprou o(a) %s\n", LoteComprado.Nome)
					meusLotes = append(meusLotes, LoteComprado)
					Espaco_disponivel += LoteComprado.Espaco
				case "2":
					listaBrinquedos()
					fmt.Scanf("%d", &index)
					brinquedoComprado := ToysList[index-1]
					if brinquedoComprado.Custo > saldo_jogador {
						fmt.Printf("Voc?? n??o tem dinheiro o suficiente\n")
					} else {
						if brinquedoComprado.Espaco > Espaco_disponivel {
							fmt.Printf("Voc?? n??o tem espa??o suficiente\n")
						} else {
							if quantidades[brinquedoComprado.Code] > brinquedoComprado.QuantidadeMax {
								fmt.Printf("Voc?? atingiu a maxima quantidade desse brinquedo\n")
							} else {
								if quantidades[brinquedoComprado.Code] == brinquedoComprado.QuantidadeMax {
									ToysList = remove2(ToysList, brinquedoComprado.Code)
								} else {
									m := 1
									saldo_jogador -= brinquedoComprado.Custo
									fmt.Printf("Parab??ns voce comprou o(a) %s\n", brinquedoComprado.Nome)
									myToys = append(myToys, brinquedoComprado)
									dias = append(dias, dia)
									pessoas = append(pessoas, population(brinquedoComprado.Popularidade))
									var j int = int(brinquedoComprado.Popularidade)
									array[0][m] = j
									array[1][m] = brinquedoComprado.Ingresso
									m++
								}
							}
						}
					}
				}
			case "2":
				fmt.Println(MenuDeVendas)
				fmt.Scanf("%s", &comando)
				switch comando {
				case "1":
				}
				for i != len(myToys) {
					fmt.Printf("%d - %s\n Valor: ", i+1, myToys[i].Nome)
					i++
				}
				fmt.Scanf("%d", &index)
				myToys = remove2(myToys, index-1)
			case "3":
				Total := reduceSoma(pessoas)
				RendaDeHoje := exe(array, len(sas1))
				fmt.Printf("Hoje vieram %d pessoas ao parque\n", Total)
				fmt.Printf("Lucro de hoje: %d\n", RendaDeHoje)
				timer := time.NewTimer(3 * time.Second)
				fmt.Printf("Passando o dia (Aguarde 3 segundos)\n")
				<-timer.C
				saldo_jogador += RendaDeHoje
				for i := 0; i < len(dias); i++ {
					dias[i] += 1
					if dias[i] == 30 {
						saldo_jogador = myToys[i].Custo
					}
				}
			case "4":
				var semana int
				for i := 0; i < 7; i++ {
					semana += exe(array, len(sas1))
				}
				timer2 := time.NewTimer(7 * time.Second)
				Total := reduceSoma(pessoas)
				fmt.Printf("Essa semana vieram %d pessoas ao parque\n", Total)
				fmt.Printf("Lucro dessa semana: %d\n", semana)
				fmt.Printf("Passando a semana(Aguarde 7 segundos)\n")
				<-timer2.C
				for i := 0; i < len(dias); i++ {
					dias[i] += 7
					if dias[i] == 30 {
						saldo_jogador = myToys[i].Custo
					}
				}
				vezes++
			}
		}
	}
}
func listaBrinquedos() {
	i := 1
	for i != len(ToysList) {
		brinquedo := ToysList[i]
		fmt.Printf("\n%d- %s\nCusto: %d\nPopularidade: %0.f\nEspa??o: %d\nIngresso: %d\n", i, brinquedo.Nome, brinquedo.Custo, math.Round(brinquedo.Popularidade), brinquedo.Espaco, brinquedo.Ingresso)
		i++
	}
}
func listaLotes() {
	i := 0
	j := 0
	for i != len(LotesList) {
		lote := LotesList[i]
		fmt.Printf("\n%d- %s\nCusto: %d\n", j, lote.Nome, lote.Custo)
		i++
		j++
	}
}
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
func remove2(slice []Brinquedo, s int) []Brinquedo {
	return append(slice[:s], slice[s+1:]...)
}
func renda(s int, r int) int {
	s *= 10
	n := rand.Intn(100)
	v := rand.Intn(100)
	l := rand.Intn(100)
	q := (s + n + v + l) / 4
	q /= 5
	w := r * q
	return w
}
func population(s float64) int {
	s = math.Round(s)
	var y int = int(s)
	return y*rand.Intn(100) + 400
}
func reduceSoma(s []int) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}
func exe(array [][]int, size int) int {
	var rend int
	for i := 0; i < size; i++ {
		rend += renda(population2(array[0][i]), array[1][i])
	}
	return rend
}
func population2(s int) int {
	return s*rand.Intn(100) + 400
}
