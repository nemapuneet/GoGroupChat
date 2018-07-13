package main
import(
    "fmt"
	"os"
	"log"
	"strings"
	"net"
	"net/rpc"
)

type Listener int
var grpnm string
func (l *Listener) Getgrp(grp string, ack *bool) error {
	grpnm = grp
	return nil
}

func (l *Listener) GetLine(line string, ack *bool) error {
	if(strings.HasSuffix(line,": \n")){
		return nil
	}else{
		//fmt.Println(strings.Replace(line,"\n","",2))
		//fmt.Println(grpnm)
		name :=  []string{"./grp/",grpnm,".txt"}
		file, err := os.OpenFile(strings.Join(name, ""), os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
		defer file.Close()
 
		len, err := file.WriteString(line)
		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}
		len++

		client, err := rpc.Dial("tcp", "localhost:8080")
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i<1 ;i++ {
			//var gname string = "abcd"
			var reply bool
			err = client.Call("Listener.GetGrpn", grpnm, &reply)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	}
}

func main(){
	var un string
	var pw string
	fmt.Println("Welcome Admin")
	fmt.Println("Enter USERNAME")
	fmt.Scan(&un)
	fmt.Println("Enter PASSWORD")
	fmt.Scan(&pw)

	if(un == "admin" && pw == "admin"){
		fmt.Println("Logged in")
		x:
		fmt.Println("Enter 1 to add group")
		fmt.Println("Enter 2 to start chat server")
		fmt.Println("Enter 3 to close")
		var choice int
		fmt.Scan(&choice)
			switch choice {
			case 1:
				//fmt.Println("code for start group")
				var gn string
				fmt.Println("Enter Groupname :")
				fmt.Scan(&gn)
				name :=  []string{"./grp/",gn,".txt"}
				file, err := os.Create(strings.Join(name, ""))  
				   if err != nil {  
					  log.Fatal(err)  
				   }  
				file.Close() 
				fmt.Println("You have successfully created group",gn)
				goto x
				break
			case 2:
				fmt.Println("server started")

				con, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
				if err != nil {
					log.Fatal(err)
				}

				inbound, err := net.ListenTCP("tcp", con)
				if err != nil {
					log.Fatal(err)
				}

				listener := new(Listener)
				rpc.Register(listener)
				rpc.Accept(inbound)
				break
			case 3:
				os.RemoveAll("./grp/")
				os.MkdirAll("./grp/",0777)
				os.Exit(2)
			}
		fmt.Print("out of for loop")
	}else{
		fmt.Print("wrong credentials")
	}
	
}