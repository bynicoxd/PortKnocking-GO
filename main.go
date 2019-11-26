package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	ip:=""
	if len(os.Args)>1{
		ip= os.Args[1]
	}else{
		fmt.Println("Host?")
		os.Exit(0)
	}
	fmt.Println(" Prueba de Port Knock")
	var seq [3] string
	seq = [3]string{"444", "222", "333"}

	for _, x := range seq {
		address := ip + ":" + x
		fmt.Println("Tratando de conectarse a: " + address)
		net.DialTimeout("tcp", address, 1*time.Second)
		time.Sleep(500* time.Millisecond)
	}
}