package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	var username string = "identitymanager"
	var passwd string = "S1ta@123"
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://dev-persondb.uksouth.cloudapp.azure.com:9200/", nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	fmt.Println(s)
}
