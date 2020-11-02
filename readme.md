https://godoc.org/github.com/openshift/osin

https://www.oauth.com/oauth2-servers/pkce/

https://tonyxu-io.github.io/pkce-generator/

---------------------------------------
# Authorization Code Flow

```bash
curl -v 'http://localhost:14000/authorize?response_type=code&state=1&client_id=1234'

curl -v 'http://localhost:14000/token' \
-H "Authorization: Basic $(echo -ne '1234:aabbccdd' | base64)" \
-d 'grant_type=authorization_code&code=1RY7kUoaSq6fKR2BtUJoDw' | jq

curl -v 'http://localhost:14000/token' \
-H 'Authorization: Basic MTIzNDphYWJiY2NkZA==' \
-d 'grant_type=refresh_token&refresh_token=Gqp0TUdwTUmwFfZ7REd5mw' | jq

```
--------------------------------------------

# PKCE FLOW

CODE_VERIFIER=     dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk

HASH =             E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM

```bash
curl -v 'http://localhost:14000/authorize?response_type=code&client_id=public-client&code_challenge_method=S256&code_challenge=E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM'

curl -v 'http://localhost:14000/token' \
-H "Authorization: Basic $(echo -ne 'public-client:' | base64)" \
-d 'grant_type=authorization_code&code_verifier=dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk&code=OLH7Et5fT_iJ9PJ8vJJf9w'

curl -v 'http://localhost:14000/token' \
-H "Authorization: Basic $(echo -ne 'public-client:' | base64)" \
-d 'grant_type=refresh_token&refresh_token=Q2y_YrNgToWtXj8O-Geb7A'
```

