package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
    "os/exec"
    "runtime"
    "time"
	"io/ioutil"
	"strings"
)

type Listener int

func (l *Listener) GetGrpn(grp string, ack *bool) error {
	//fmt.Println(grp)
	
	time.Sleep(1 * time.Second)
    CallClear()

	name :=  []string{"**##",grp,"##**"}
	fmt.Println(strings.Join(name, ""))
	name2 :=  []string{"./grp/",grp,".txt"}
		data, err := ioutil.ReadFile(strings.Join(name2, ""))
		if err != nil {
			log.Panicf("failed reading data from file: %s", err)
		}
	fmt.Printf("\n%s", data)
	return nil
}


var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}

func main() {
	con, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8080")
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

}