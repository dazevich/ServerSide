package api

import (
	"ServerSide/api/bhx"
	"fmt"
	"log"
	"net/http"
)

type CryptoAnser struct {
}

func GetCrypto(wr http.ResponseWriter, r *http.Request) {

	key := r.FormValue("key")
	text := []byte(r.FormValue("text"))

	var keyBx bhx.BoxSharedKey
	copy(keyBx[:], bhx.Sha256H([]byte(key)).Bytes())

	nnonce := bhx.GetKeyNonce(keyBx)

	result, err := bhx.Encrypt(text, keyBx, nnonce)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result))

	output, err := bhx.Decrypt(result, keyBx, nnonce)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(output))

}
