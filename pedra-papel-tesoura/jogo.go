package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var ganhou = "você ganhou 🏆"
var perdeu = "você perdeu 🥈"

func main() {
	opcoesMaquina := [3]string{"pedra", "papel", "tesoura"}

	rand.Seed(time.Now().UnixNano())
	respostaMaquina := opcoesMaquina[rand.Intn(3)]

	getInput := bufio.NewScanner(os.Stdin)
	fmt.Print("> pedra, papel ou tesoura? escreva aqui -> ")
	getInput.Scan()
	respostaPessoa := getInput.Text()

	fmt.Println("outro jogador: ", respostaMaquina)

	if !(respostaPessoa == "pedra" || respostaPessoa == "papel" || respostaPessoa == "tesoura") {
		fmt.Println("> opção inválida.", respostaPessoa)
		fmt.Println("> as únicas opções aceitas são: pedra, papel ou tesoura.")
		fmt.Println("> por favor, tente novamente.")
	}

	if respostaPessoa == respostaMaquina {
		fmt.Println("> empate! tente novamente :D")
	}

	switch respostaPessoa {
	case "pedra":
		if respostaMaquina == "tesoura" {
			fmt.Println("pedra ganha de tesoura!", ganhou)
		} else if respostaMaquina == "papel" {
			fmt.Println("pedra perde pra tesoura!", perdeu)
		}
	case "papel":
		if respostaMaquina == "pedra" {
			fmt.Println("papel ganha de pedra!", ganhou)
		} else if respostaMaquina == "tesoura" {
			fmt.Println("papel perde pra tesoura!", perdeu)
		}
	case "tesoura":
		if respostaMaquina == "papel" {
			fmt.Println("tesoura ganha de papel!", ganhou)
		} else if respostaMaquina == "pedra" {
			fmt.Println("tesoura perde pra pedra!", perdeu)
		}
	}
}