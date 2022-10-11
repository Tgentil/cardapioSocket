package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func conexaoServidor() {

	// iniciando a conexão tcp
	fmt.Println("Servidor aguardando conexões...")

	ln, erro1 := net.Listen("tcp", ":8081")
	if erro1 != nil {
		fmt.Println(erro1)

		os.Exit(3)
	}

	// aguarde de conexão do cliente
	conexao, erro2 := ln.Accept()
	if erro2 != nil {
		fmt.Println(erro2)
		os.Exit(3)
	}

	defer ln.Close() // fecha a conexão no fim

	fmt.Println("Conexão aceita...")

	for { // Laço eterno
		// Recebe informações no buffer de leitura

		mensagem, erro3 := bufio.NewReader(conexao).ReadString('\n')
		if erro3 != nil {
			fmt.Println(erro3)
			os.Exit(3)
		}

		fmt.Print("Mensagem recebida:", string(mensagem))
		novamensagem := teste(string(mensagem))

		conexao.Write([]byte(novamensagem + "\n"))

	}
}

func teste(mensagem string) string {
	// eliminando a quebra de parágrafo e de linha
	msg := mensagem[0:(len(mensagem) - 2)]
	// itens do cardapio
	switch msg {

	case "1":

		fmt.Println("R$ 10,00")
		return "R$ 10,00"

	case "2":

		fmt.Println("R$ 15,00")
		return "R$ 15,00"

	case "3":

		fmt.Println("R$ 18,00")
		return "R$ 18,00"

	case "4":

		fmt.Println("R$ 30,00")
		return "R$ 30,00"

	case "5":

		fmt.Println("R$ 35,00")
		return "R$ 35,00"

	case "6":

		fmt.Println("R$ 37,00")
		return "R$ 37,00"

	case "7":

		fmt.Println("R$ 55,00")
		return "R$ 55,00"

	case "8":

		fmt.Println("R$ 60,00")
		return "R$ 60,00"

	case "9":

		fmt.Println("62,00")
		return "R$ 62,00"

	case "10":

		fmt.Println("R$ 100,00")
		return "R$ 100,00"

	case "0":

		fmt.Println("saindo...")
		os.Exit(0)
		return "saindo..."

	default:

		fmt.Println("valor não indentificado")
		return "Não identificado"

	}

}

// chama a função de conexão
func main() {
	conexaoServidor()
}
