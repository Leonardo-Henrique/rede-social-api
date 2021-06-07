package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	PortaAPI           = 0
	SecretKey          []byte
)

func Carregar() {

	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	PortaAPI, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		PortaAPI = 9000
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

}
