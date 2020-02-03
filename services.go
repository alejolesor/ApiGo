package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
)

var url = "http://gateway.marvel.com/v1/"

var apikeyPublic = "4767bd441b9c524b7c9db29e8b2c3b16"
var apikeyPrivate = "edcae0e40a18c8d19b91454202a3b17aeaa5b049"
var ts = "1"
var hash string

func digestString(t string, publicKey string, privateKey string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(t+privateKey+publicKey)))
}

//Services ...
func Services(methodComics string) *http.Response {
	hash = digestString(ts, apikeyPublic, apikeyPrivate)
	urlConcat := url + methodComics + "apikey=" + apikeyPublic + "&" + "ts=" + ts + "&" + "hash=" + hash
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlConcat, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", `*/*`)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}
