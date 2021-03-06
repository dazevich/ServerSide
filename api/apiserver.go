package api

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

//Type Courses ...
type Courses struct {
	Course []struct {
		From      string  `xml:"from" json:"from"`
		To        string  `xml:"to" json:"to"`
		In        float64 `xml:"in" json:"in"`
		Out       float64 `xml:"out" json:"out"`
		Amount    int     `xml:"amount" json:"amount"`
		Minamount string  `xml:"minamount" json:"minamount"`
		Maxamount string  `xml:"maxamount" json:"maxamount"`
		Param     string  `xml:"param" json:"param"`
		City      string  `xml:"city" json:"city"`
	} `xml:"item" json:"courses"`
}

//APIServer ...
func APIServer(w http.ResponseWriter, r *http.Request) {

	byteCh := make(chan []byte)
	answerCh := make(chan []byte)

	go func() {

		resp, err := http.Get("https://test.cryptohonest.ru/request-exportxml.xml")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		byteCh <- body
	}()

	xmlCourses := &Courses{}

	err := xml.Unmarshal(<-byteCh, xmlCourses)
	if nil != err {
		log.Println(err)
	}

	go func() {
		answer, err := json.Marshal(xmlCourses)
		if err != nil {
			log.Println(err)
		}
		answerCh <- answer
	}()

	w.Write(<-answerCh)
}
