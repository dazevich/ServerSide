package api

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
)

// Структура для преобразования в JSON
type CryptoAnswer struct {
	Key    string
	Input  string
	Output string
}

var (
	input  string
	output string
)

// GetKey... Пользовательский ключ складывается с шаблоном так, чтобы итоговый ключ состоял из 64 символов
func GetKey(key string) ([]byte, error) {
	userKey := []byte(key)
	keyTmp := []byte("6368616e176520746869732070617373776f726420746f206120736563726574")
	var newKey []byte
	if len(userKey) < len(keyTmp) {
		addTmp := make([]byte, len(keyTmp)-len(userKey))
		copy(addTmp, keyTmp)
		newKey = append(newKey, userKey...)
		newKey = append(newKey, addTmp...)
	}
	if len(userKey) > len(keyTmp) {
		addTmp := make([]byte, len(keyTmp))
		copy(addTmp, userKey)
		newKey = append(newKey, addTmp...)
	}
	result, err := hex.DecodeString(string(newKey))
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return result, nil
}

// GetNonce... Получение номера
func GetNonce() []byte {
	result, _ := hex.DecodeString("64a9433eae7ccceee2fc0eda")
	return result
}

// Encrypt... Шифрование текста с помощью полученного ключа и номера
func Encrypt(text string, key []byte, nonce []byte) []byte {
	textByte := []byte(text)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, textByte, nil)
	return ciphertext
}

// Decrypt... Дешифровка зашифрованного текста с помощью ключа и номера
func Decrypt(ciphertext []byte, key []byte, nonce []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func GetCrypto(w http.ResponseWriter, r *http.Request) {

	// Разрешаю подключение с различных доменов
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Получаю ключ, текст и шифротекст
	userKey := r.FormValue("key")
	in := r.FormValue("input")
	out := r.FormValue("output")

	if userKey == "" {
		userKey = "edeef0"
	}

	// Привожу ключ к нужному виду
	key, err := GetKey(userKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Получаю номер
	nonce := GetNonce()

	// Если дан только текст, шифрую его
	if in == "" && out != "" {
		decodedStr, err := hex.DecodeString(out)
		if err != nil {
			log.Fatal(err)
		}
		input = string(Decrypt(decodedStr, key, nonce))
		output = out
	}

	// Если дан только шифротекст, дешифрую его
	if in != "" && out == "" {
		output = hex.EncodeToString(Encrypt(in, key, nonce))
		input = in
	}

	// Возвращаю ответ в формате JSON
	answer := &CryptoAnswer{}
	answer.Key = userKey
	answer.Input = input
	answer.Output = output

	json, _ := json.Marshal(answer)

	w.Write(json)

}
