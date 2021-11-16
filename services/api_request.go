package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func apiRequest(url string, key string, reqData interface{}) {
	jsonByte, err := json.Marshal(reqData)
	if err != nil {
		log.Println(err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonByte))
	if err != nil {
		//todo log
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", key)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(jsonByte))
	log.Println(string(body))
}
