package tcp

import (
	"bufio"
	"io"
	"log"
	"net"
	"strconv"
)

type TcpServer struct{}

func (srv *TcpServer) Serve(hst string, prt int) error {

	lsn, err := net.Listen("tcp4", hst+":"+strconv.Itoa(prt))
	defer lsn.Close()

	if err != nil {

		log.Fatalf("Tcp Server was failed to run %s:%s\n", hst+":"+strconv.Itoa(prt), err)
		return err

	}

	log.Printf("Tcp Server Running %s\n", hst+":"+strconv.Itoa(prt))

	for {

		con, err := lsn.Accept()
		defer con.Close()

		if err != nil {

			log.Fatalf("TCP Server was failed to accepte client:%s\n", err)
			continue

		}

		log.Printf("TCP Server was accepted client\n")

		go func(con net.Conn) error {

			res := "pong"
			buf := make([]byte, 1024)

			rdr := bufio.NewReader(con)
			wtr := bufio.NewWriter(con)

			CONNECTION: for {

				len, err := rdr.Read(buf)

				switch err {

					case io.EOF: {
						
						break CONNECTION

					}

						

					case nil: {
						
						req := string(buf[:len])
						log.Printf("TCP Server was received %s", req)

						wtr.Write([]byte(res))
						wtr.Flush()

						log.Printf("TCP Server was sent %s", req)

					}

						

					default: {
						
						log.Fatalf("TCP Server was failed to receive:%s", err)
						return err

					}

				}

			}

			return nil

		}(con)

	}

}
