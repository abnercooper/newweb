package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func reLaunch() {
	cmd := exec.Command("sh", "./deploy.sh")  //exec是go用来调用命令 ?./deploy.sh
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func firstPage(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "<h1>hello, this is my deploy server!</h1>")
	n :=n+1
	fmt.Print("io-OK/n*1*%s",n)
	reLaunch()
	m :=m+1
	fmt.Print("reLaunch-OK/n***%s***",m)
}

func main(){
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":5000",nil)
}