package middleware

import (
	"fmt"
	"net/http"
)

func HandlerConv(handlerF func(http.ResponseWriter, *http.Request)) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("url request is ", req.URL.RequestURI())
		 rw.Header().Set("Access-Control-Allow-Origin", "*")
         rw.Header().Set("Access-Control-Allow-Credentials", "true")
         rw.Header().Set("Access-Control-Allow-Methods", "POST")
         rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		rw.Header().Set("Content-Type", "application/json")
		handlerF(rw, req)
	})
}
