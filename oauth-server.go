package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RangelReale/osin"
	"github.com/bmalecki/go-oauth-server/genjwt"
	ex "github.com/bmalecki/go-oauth-server/storage"
)

func main() {
	config := osin.NewServerConfig()
	config.RequirePKCEForPublicClients = true
	config.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{osin.CODE}
	config.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.REFRESH_TOKEN}

	server := osin.NewServer(config, ex.NewTestStorage())

	var err error
	var privatekeyPEM, publickeyPEM []byte

	if privatekeyPEM, err = ioutil.ReadFile("cert/jwt.key"); err != nil {
		panic(err)
	}

	if publickeyPEM, err = ioutil.ReadFile("cert/jwt.key.pub"); err != nil {
		panic(err)
	}

	if server.AccessTokenGen, err = genjwt.NewAccessTokenGenJWT(privatekeyPEM, publickeyPEM); err != nil {
		panic(err)
	}

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()

		if ar := server.HandleAuthorizeRequest(resp, r); ar != nil {

			// HANDLE LOGIN PAGE HERE

			ar.Authorized = true
			server.FinishAuthorizeRequest(resp, r, ar)
		}
		osin.OutputJSON(resp, w, r)
	})

	// Access token endpoint
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()

		if ar := server.HandleAccessRequest(resp, r); ar != nil {
			ar.Authorized = true
			server.FinishAccessRequest(resp, r, ar)
		}
		osin.OutputJSON(resp, w, r)
	})

	fmt.Println("Start listening on port 14000")
	http.ListenAndServe(":14000", nil)
}
