package netio

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

/*TCP server*/
func startTcpServer() {

	log.Println("Running TCP server")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err.Error())
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err.Error())
			break
		}
		mesg := fmt.Sprintf("Hi from Tcp server at %s\n", time.Now().String())
		conn.Write([]byte(mesg))
		conn.Close()
	}
	lis.Close()
}

/*TCP Server client*/
func dialTcpServer() {
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Println(err.Error())
		return
	}
	resp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Resp:= ", resp)
}

/*UDP server*/
func startUDPserver() {

	log.Println("Starting udp server")
	udpAdr, err := net.ResolveUDPAddr("udp4", ":8081")
	if err != nil {
		log.Println(err.Error())
		return
	}
	udplis, err := net.ListenUDP("udp", udpAdr)
	if err != nil {
		log.Println(err.Error())
		return
	}
	buffer := make([]byte, 1024)
	for {

		mesg := fmt.Sprintf("Hi from Udp server at %s\n", time.Now().String())
		_, addr, err := udplis.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err.Error())
		}
		udplis.WriteTo([]byte(mesg), addr)
	}
}

/*UDP client*/
func diaUDPServer() {

	time.Sleep(1 * time.Second)
	udpAdr, err := net.ResolveUDPAddr("udp4", ":8081")
	if err != nil {
		log.Println(err.Error())
		return
	}
	conn, err := net.DialUDP("udp", nil, udpAdr)
	if err != nil {
		log.Println(err.Error())
		return
	}
	conn.Write([]byte(""))
	buffer := make([]byte, 1024)
	_, _, err = conn.ReadFromUDP(buffer)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("resp:= ", string(buffer))
	defer conn.Close()
}

/*HTTP server*/
func startHttpSever() {

	log.Println("Starting HTTP server")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		mesg := fmt.Sprintf("Hi from HTTP server at %s\n", time.Now().String())
		rw.Write([]byte(mesg))

	})
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Println(err.Error())
		return
	}

}

/*HTTP server client*/
func dialHttpSever() {
	time.Sleep(1 * time.Second)
	resp, err := http.Get("http://localhost:8082/")
	if err != nil {
		log.Println(err.Error())
		return
	}
	io.Copy(os.Stdout, resp.Body)

}

/*Method to run all servers concurrently and
test with their respective client
*/
func RunNetoworkIO() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		startTcpServer()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		startUDPserver()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		startHttpSever()
	}()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dialTcpServer()
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			diaUDPServer()
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			dialHttpSever()
		}()
	}
	wg.Wait()
}
