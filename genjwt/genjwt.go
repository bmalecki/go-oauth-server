package genjwt

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"

	"github.com/RangelReale/osin"
	"github.com/dgrijalva/jwt-go"
	"github.com/pborman/uuid"
)

// JWT access token generator
type AccessTokenGenJWT struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewAccessTokenGenJWT(privatekeyPEM, publickeyPEM []byte) (osin.AccessTokenGen, error) {
	var err error
	var accessTokenGenJWT AccessTokenGenJWT

	if accessTokenGenJWT.privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privatekeyPEM); err != nil {
		return nil, err
	}

	if accessTokenGenJWT.publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publickeyPEM); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return nil, err
	}

	return &accessTokenGenJWT, nil
}

func (c *AccessTokenGenJWT) GenerateAccessToken(data *osin.AccessData, generaterefresh bool) (
	accesstoken string, refreshtoken string, err error) {

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"cid": data.Client.GetId(),
		"exp": data.ExpireAt().Unix(),
	})

	accesstoken, err = token.SignedString(c.privateKey)
	if err != nil {
		return "", "", err
	}

	if generaterefresh {
		rtoken := uuid.NewRandom()
		refreshtoken = base64.RawURLEncoding.EncodeToString([]byte(rtoken))
	}

	return
}
