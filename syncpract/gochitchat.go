package syncpract

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var (
	prefix = "Chat:Sever>"
	userdb = make(map[string]net.Conn)
	menu   = `
---Help---
/q : exit
/lu : list user
/cu: send message to user
----------
	`
)

func StartServer() {
	fmt.Println("Go chitchat")

	lis, err := net.Listen("tcp4", ":8080")
	log.Println("listening on port 8080")

	if err != nil {
		log.Fatal(err.Error())
	}
	for {

		conn, err := lis.Accept()
		if err != nil {
			log.Printf("Error in connectiton%s\n = %s",
				conn.RemoteAddr(), err.Error())
		}
		go handleConncection(conn)
	}
}

func handleConncection(conn net.Conn) {

mainConn:
	for {
		conn.Write([]byte(prefix + "Enter user name:- "))
		username, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println(err.Error())
			break mainConn
		}
		username = strings.TrimSpace(username)
		userdb[username] = conn
		userPrefix := fmt.Sprintf("\n\u001b[31mchat:%s>\u001b[0m", username)
		for {
			conn.Write([]byte(userPrefix))
			option, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Println(err.Error())
				break
			}
			switch strings.TrimSpace(option) {
			case "/q":
				delete(userdb, username)
				break mainConn
			case "/lu":
				list := ""
				for k, _ := range userdb {
					list = list + fmt.Sprintf("Username: %s\n", k)
				}
				conn.Write([]byte(list))
			case "/cu":
				conn.Write([]byte("user to connect: "))
				cuser, err := bufio.NewReader(conn).ReadString('\n')
				if err != nil {
					break
				}
				cuser = strings.TrimSpace(cuser)
				uc, ok := userdb[cuser]
				if !ok {
					conn.Write([]byte("user not found"))
					break
				}
				conn.Write([]byte("enter your message:- "))
				msg, _ := bufio.NewReader(conn).ReadString('\n')
				uc.Write([]byte(fmt.Sprintf("\n\u001b[32mfrom:%s:>%s\u001b[0m", username, msg)))
			case "/h":
				conn.Write([]byte(menu))
			default:
				conn.Write([]byte("please enter valid statment"))
			}

		}
	}
	conn.Close()
}
