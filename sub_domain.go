package main

import (
	"fmt"
	"os"
	//"strings"
	"net/http"
	"net/url"
	"bufio"
	"os/exec"
)

func main() {
	var Domain,Filename string
	
	fmt.Print("Enter the Domain name: ")
	fmt.Scanln(&Domain)
	fmt.Print("Enter the file name with(.txt) :")
	fmt.Scanln(&Filename)
	_,err :=exec.Command("sublist3r","-d",Domain,"-o",Filename).Output()

	if err != nil{
		println(err)
	}
	file,_:= os.Open("sub.txt")
	sub_domain:= bufio.NewScanner(file)
	sub_domain.Split(bufio.ScanLines)
	for sub_domain.Scan(){
		fmt.Println(http_s(sub_domain.Text()))
	}

 }

func http_s(domain string) (w *url.URL) {
	url,_:= http.Get("http://"+domain)
	return url.Request.URL
}
