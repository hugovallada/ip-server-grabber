package main

import (
	"fmt"
	"ip-server/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ip_Server Grabber")
	aplicacao := app.Gerar()
	
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}

}
