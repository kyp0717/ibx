package main

import (
	"crypto/tls"
	"fmt"
	"learn2/client"
	"net/http"
)

func setPolicy() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
}

func main() {
	setPolicy()
	acct := client.Default.Account
	list, err := acct.GetIserverAccounts(nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", list)
}
