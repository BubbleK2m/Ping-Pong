package tcp

import (
	"bufio"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type TcpClient struct{}

func (cli *TcpClient) Connect(hst string, prt int) error {

	log.Printf("Tcp Client try to connect %s\n", hst+":"+strconv.Itoa(prt))

	con, err := net.Dial("tcp", hst+":"+strconv.Itoa(prt))
	defer con.Close()

	if err != nil {

		log.Fatalf("Tcp Client try failed to connect %s:%s\n", hst+":"+strconv.Itoa(prt), err)
		return err

	}

	log.Printf("Tcp Client was connected %s\n", hst+":"+strconv.Itoa(prt))

	req := "ping"
	buf := make([]byte, 1024)

	rdr := bufio.NewReader(con)
	wtr := bufio.NewWriter(con)

	CONNECTION: for {

		time.Sleep(time.Second * 1)
		log.Printf("Tcp Client was sent message %s\n", req)
		
		wtr.Write([]byte(req))
		wtr.Flush()

		len, err := rdr.Read(buf)

		switch err {

			case io.EOF: {
				
				break CONNECTION

			}

			case nil: {

				res := string(buf[:len])
				log.Printf("Tcp Client was received message %s\n", res)

			}
				
			default: {

				log.Fatalf("Tcp Client was failed to receive message %s\n", err)
				return err

			}

		}

	}

	log.Printf("Tcp Client finished connection %s\n", hst+":"+strconv.Itoa(prt))
	return nil

}
