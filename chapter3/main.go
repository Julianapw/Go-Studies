package main

import (
	"fmt"
)

func main() {
	fmt.Println("--- Exercício 5: Divisão ---")
	res, err := Ex5(10, 2)
	fmt.Printf("10/2 = %d (Erro: %v)\n", res, err)
	_, err = Ex5(10, 0)
	fmt.Printf("10/0 erro esperado: %v\n", err)

	fmt.Println("\n--- Exercício 6 e 7: Busca de Usuário ---")
	Err() // A função Err já chama FindUser e imprime o resultado internamente

	fmt.Println("\n--- Exercício 8: Error Wrapping (%w) ---")
	_, err = Err2("10") // ID diferente de "12" retorna erro com wrapping
	fmt.Printf("Erro com wrapping: %v\n", err)

	fmt.Println("\n--- Exercício 10 e 11: Injeção de Dependência ---")
	storage := NewInMemoryStorage()
	userService := NewUserService(storage)

	err = userService.CreateUser("Juliana")
	fmt.Printf("Usuário criado: %v\n", err)

	fmt.Println("\n--- Exercício 14 e 15: Custom Errors e Validation ---")
	// Testando o erro de validação (campo vazio)
	userService.UserCustomError("")

	fmt.Println("\n--- Exercício 2: Interfaces de Printer ---")
	var p Printer
	p = ConsolePrinter{}
	p.Print("Olá do Console")

	p = FilePrinter{}
	p.Print("Olá do Arquivo")

	fmt.Println("\n--- Testes ---")
	fmt.Println("Dica: Para rodar os testes dos exercícios 16 a 19, use o comando: go test ./...")

	
}
