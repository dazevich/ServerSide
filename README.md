# ServerSide

В файле main.go идет прослушка порта 9097 и /getCourses. При обращении к нему вызывается обработчик из пакета api apiserver. Этот обработчик получает xml файл, преобразует его в JSON и отдает

w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	k := r.FormValue("key")
	in := r.FormValue("input")
	out := r.FormValue("output")

	var input []byte
	var output []byte
	var key bhx.BoxSharedKey

	if k != "" {
		bytes, err := hex.DecodeString(k)
		if err != nil {
			log.Fatal("Error create key: ", err)
		}
		copy(key[:], bytes)
	} else {
		bytes, err := hex.DecodeString("edeef0")
		if err != nil {
			log.Fatal("Error create key: ", err)
		}
		copy(key[:], bytes)
	}

	var nnonce bhx.BoxNonce
	kh, err := hex.DecodeString("edeef0edeef0edeef0edeef0edeef0edeef0edeef0edeef0")
	if err != nil {
		log.Fatal(err)
	}
	copy(nnonce[:], kh[:bhx.BoxNonceLen])

	if out == "" && in != "" {
		output = Encrypt(in, key, nnonce)
		input = []byte(in)
	}

	if out != "" && in == "" {
		input = Decrypt([]byte(out), key, nnonce)
		output = []byte(out)
	}

	answer := &CryptoAnswer{}

	keyString := string(key[:len(key)])
	outputString := string(output[:len(output)])
	decInputString := string(input[:len(input)])

	answer.Key = keyString
	answer.Input = string(in)
	answer.Output = outputString
	answer.DecryptedInput = decInputString

	jsonAnsw, err := json.Marshal(answer)
	if err != nil {
		log.Fatal(err.Error())
	}

	w.Write(jsonAnsw)