package main

// O pacote net/http é usado para criar servidores web e manipular requisições HTTP em Go.
import (
	"fmt"
	"net/http"
)

// A função http.HandleFunc associa a URL "/" a uma função que será chamada
// sempre que essa URL for acessada. Neste caso, a função anônima passada
// como argumento escreve "Olá mundo!" na resposta HTTP.
// A função http.ListenAndServe inicia um servidor web na porta 8080 e
// aguarda por requisições. Quando uma requisição é recebida na URL "/",
// a função associada é executada.
// Nota: Não é possível iniciar o servidor na porta 80 sem permissões de superusuário.
// Em sistemas Unix, portas abaixo de 1024 são reservadas para processos privilegiados.
func main() {
	// http.HandleFunc associa a URL "/" a uma função que será chamada sempre que essa URL for acessada.
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Olá mundo!")
		},
	)
	fmt.Println("Rodando servidor. Pressione Ctrl+C para sair...")
	// inicia o servidor na porta 8080
	http.ListenAndServe(":8080", nil)
}
