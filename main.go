package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// func init() {
// 	chave := make([]byte, 64)

// 	if _, erro := rand.Read(chave); erro != nil {
// 		log.Fatal(erro)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	fmt.Println(stringBase64)
// }

func main() {

	config.Carregar()

	fmt.Printf("Escutando a porta %d", config.PortaAPI)

	r := router.GerarRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PortaAPI), r))

}
