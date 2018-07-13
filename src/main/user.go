package main
import(
    "fmt"
	"os"
	"io/ioutil"
    "log"
	"strings"
	"bufio"
	"net/rpc"
)
func main(){
	var un string
	var pw string
	fmt.Println("Welcome Client")
	fmt.Println("Enter USERNAME")
	fmt.Scan(&un)
	fmt.Println("Enter PASSWORD")
	fmt.Scan(&pw)

	if((un == "puneet" || un == "rajat") && pw == "client"){
		fmt.Println(un,"is Logged in")
		fmt.Println("Enter 1 to Join group")
		fmt.Println("Enter 2 to close")
		var choice int
		fmt.Scan(&choice)
			switch choice {
			case 1:
				fmt.Println("Select Group to Join")
				files, err := ioutil.ReadDir("./grp/")
				if err != nil {
					log.Fatal(err)
				}

				for _, f := range files {
						fmt.Println(strings.Replace(f.Name(),".txt","",2))
				}
				fmt.Println("Enter name of the group you want to join")
				var gname string
				fmt.Scan(&gname)
				fmt.Println("You joined the group chat : ", gname)
				x:
				fmt.Println("Enter 1 to send message")
				fmt.Println("Enter 2 to see group chat")
				fmt.Println("Enter 3 to close")
				var choice2 int
				fmt.Scan(&choice2)
					switch choice2 {
						case 1:
							fmt.Println("Your message :")
							client, err := rpc.Dial("tcp", "localhost:42586")
							if err != nil {
								log.Fatal(err)
							}
							in := bufio.NewReader(os.Stdin)
							for i := 0; i < 2; i++{
								line, _:= in.ReadString('\n')
								name := []string{un,line}
								
								var reply1 bool
								err = client.Call("Listener.Getgrp", gname, &reply1)
								if err != nil {
									log.Fatal(err)
								}
								
								var reply2 bool
								err = client.Call("Listener.GetLine", strings.Join(name, " : "), &reply2)
								if err != nil {
									log.Fatal(err)
								}
							}
							goto x
							break
						case 2:
							name :=  []string{"**##",gname,"##**"}
							fmt.Println(strings.Join(name, ""))
							name2 :=  []string{"./grp/",gname,".txt"}
							 data, err := ioutil.ReadFile(strings.Join(name2, ""))
								if err != nil {
									log.Panicf("failed reading data from file: %s", err)
								}
								fmt.Printf("\n%s", data)
							goto x
							break
						case 3:
							os.Exit(2)
							break
					}
				break
			case 2:
				os.Exit(2)
			}
		fmt.Print("out of for loop")
	}else{
		fmt.Print("wrong credentials")
	}
	
}