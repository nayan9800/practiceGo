package syncpract

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

/* Go ChitChat Sever*/
/*Simple Tcp chat server in golang*/

var (
	prefix = "Chat:Sever>"             //chat prefix
	userdb = make(map[string]net.Conn) //map to store username and connection
	//help menu
	menu = `
---Help---
/q : exit
/lu : list user
/cu: send message to user
----------
	`
)

func Startserver() {
	fmt.Println("Go chitchat")

	lis, err := net.Listen("tcp4", ":8080") //create tcp listener on port 8080s
	log.Println("listening on port 8080")
	if err != nil { //check if there is any error
		log.Fatal(err.Error())
	}
	for { //accecpt conncections

		conn, err := lis.Accept()
		if err != nil { //check if there is a error
			log.Printf("Error in connectiton%s\n = %s",
				conn.RemoteAddr(), err.Error())
			continue
		}
		go handleConncection(conn) //run connection handller in goroutine
	}
}

/*handleConnection function is used to handle connection*/
func handleConncection(conn net.Conn) {

mainConn: //label for main for loop
	for {
		conn.Write([]byte(prefix + "Enter user name:- "))       //write to user to ask username
		username, err := bufio.NewReader(conn).ReadString('\n') //get user input using bufio
		if err != nil {
			log.Println(err.Error())
			break mainConn
		}
		username = strings.TrimSpace(username) //trim spaces in username using trimspace in strings
		_, ok := userdb[username]              //check in map if username already exsits
		if ok {
			conn.Write([]byte("username already exists\n"))
			continue
		}
		userdb[username] = conn                                              //put username in map in which username as key and net.conn as value
		userPrefix := fmt.Sprintf("\n\u001b[31mchat:%s>\u001b[0m", username) //create user prefix
		for {
			conn.Write([]byte(userPrefix))                        //print user prefix
			option, err := bufio.NewReader(conn).ReadString('\n') //get user input
			if err != nil {
				log.Println(err.Error())
				break
			}
			switch strings.TrimSpace(option) {
			case "/q": //quit connection and delete username from map
				delete(userdb, username)
				break mainConn
			case "/lu": //list the active users
				list := ""
				for k := range userdb {
					list = list + fmt.Sprintf("Username: %s\n", k)
				}
				conn.Write([]byte(list))
			case "/cu": //connect the user
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
			case "/h": //print help menu
				conn.Write([]byte(menu))
			default: //if option in not in menu
				conn.Write([]byte("please enter valid statment"))
			}

		}
	}
	conn.Close() //close connection
}
