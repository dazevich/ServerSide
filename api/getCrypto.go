package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CryptoAnswer struct {
	Key    string
	Input  string
	Output string
}

func Encrypt(text string, shift int, size int) string {
	var encrypted string
	var old_code, new_code int
	for _, ch := range text {
		old_code = int(ch)
		new_code = old_code + shift
		if ch >= 'a' && ch <= 'z' && new_code > int('z') ||
			ch >= 'A' && ch <= 'Z' && new_code > int('Z') {
			new_code -= size
		}

		encrypted += string(new_code)
	}
	return encrypted
}

func Decrypt(text string, shift int, size int) string {
	var decrypted string
	var old_code, new_code int
	for _, ch := range text {
		old_code = int(ch)
		new_code = old_code - shift
		if ch >= 'a' && ch <= 'z' && new_code < int('a') ||
			ch >= 'A' && ch <= 'Z' && new_code < int('A') {
			new_code += size
		}

		decrypted += string(new_code)
	}
	return decrypted
}

var (
	input  string
	output string
)

func GetCrypto(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	shift, err := strconv.Atoi(r.FormValue("key"))
	if err != nil {
		log.Fatal(err)
	}
	in := r.FormValue("input")
	out := r.FormValue("output")

	size := 33
	if shift > size {
		log.Fatalf("Cannot be more then %d", shift)
	}

	if out == "" && in != "" {
		output = Encrypt(in, shift, size)
		input = in
	}

	if out != "" && in == "" {
		input = Decrypt(out, shift, size)
		output = out
	}

	crypto := &CryptoAnswer{}

	crypto.Key = strconv.Itoa(shift)
	crypto.Input = input
	crypto.Output = output

	json, err := json.Marshal(crypto)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(json)
}
