package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func conexaoCliente() {
	// conectando na porta 3030 via protocolo tcp/ip na máquina local
	conexao, erro1 := net.Dial("tcp", "127.0.0.1:8081")
	if erro1 != nil {
		fmt.Println(erro1)
		os.Exit(3)
	}

	defer conexao.Close()

	requisicoes := 11

	for i := 0; i <= requisicoes; i++ {
		fmt.Println("Cliente Conectado ao Servidor")

		for {
			fmt.Println("Por favor, informe o número da bebida que você deseja consultar o valor :\n" +
				"1: Cachaça\n" +
				"2: Cerveja\n" +
				"3: Champanhe\n" +
				"4: Conhaque\n" +
				"5: Gin\n" +
				"6: Licor\n" +
				"7: Tequila\n" +
				"8: Vinho\n" +
				"9: Vodka\n" +
				"10: Whisky\n" +
				"0: Sair")
			fmt.Print("Opção escolhida: ")
			leitor := bufio.NewReader(os.Stdin)
			texto, erro2 := leitor.ReadString('\n')
			if erro2 != nil {
				fmt.Println(erro2)
				os.Exit(3)
			}

			fmt.Fprintf(conexao, texto+"\n")

			// ouvindo a resposta do servidor (eco)
			mensagem, err3 := bufio.NewReader(conexao).ReadString('\n')
			if err3 != nil {
				fmt.Println(err3)
				os.Exit(3)
			}
			// escrevendo a resposta do servidor no terminal
			fmt.Print("Resposta do servidor: " + mensagem)
			time.Sleep(time.Second * 2)
		}
	}
}

type vertex struct {
	dia int
	mes int
	ano int
}

func routine(str string) {
	for i := 0; i < 1; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}
func main() {

	routine("Seja bem vidno ao Cardapio 51")

	go routine("Siga os produtos a seguir")

	fmt.Println(vertex{
		dia: 0,
		mes: 00,
		ano: 000,
	})

	conexaoCliente()

}
