package api

import "net/http"

func GetCrypto(w http.ResponseWriter, r *http.Request) {

	key := r.FormValue("key")
	if key != "" {
		w.Write([]byte(key))
	}

}
