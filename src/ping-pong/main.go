package main

import (
	"log"
	"os"
	"ping-pong/tcp"
	"strconv"
)

func main() {

	beh, hst := os.Args[1], os.Args[2]
	prt, _ := strconv.Atoi(os.Args[3])

	log.Printf("%s %s %s\n", beh, hst, strconv.Itoa(prt))

	switch beh {

		case "client": {
			
			cli := new(tcp.TcpClient)
			cli.Connect(hst, prt)

		}

		case "server": {
			
			srv := new(tcp.TcpServer)
			srv.Serve(hst, prt)

		}

		default: {
			log.Fatal("Usage : [--behaviour] [--host] [--port]\n")
		}

	}

}
