https://godoc.org/github.com/openshift/osin
https://www.oauth.com/oauth2-servers/pkce/
https://tonyxu-io.github.io/pkce-generator/

---------------------------------------
curl -v 'http://localhost:14000/authorize?response_type=code&state=1&client_id=1234'

LJaQR9BLRqepNP4f4_u-jA

curl -v 'http://localhost:14000/token' \
-H "Authorization: Basic $(echo -ne '1234:aabbccdd' | base64)" \
-d 'grant_type=authorization_code&code=Sl5yfJ7kRtGXhef0nxSP7w'

{"access_token":"OB3vIlP5QcCrz261-A-CCA","expires_in":3600,"refresh_token":"3hYkZyL7T8GRYoWGCyMcQQ","token_type":"Bearer"}


curl -v 'http://localhost:14000/token' \
-H 'Authorization: Basic MTIzNDphYWJiY2NkZA==' \
-d 'grant_type=refresh_token&refresh_token=3hYkZyL7T8GRYoWGCyMcQQ'

--------------------------------------------

CODE_VERIFIER=     dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk
HASH =             E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM

curl -v 'http://localhost:14000/authorize?response_type=code&client_id=public-client&code_challenge_method=S256&code_challenge=E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM'

curl -v 'http://localhost:14000/token' \
-H "Authorization: Basic $(echo -ne 'public-client:' | base64)" \
-d 'grant_type=authorization_code&code_verifier=dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk&code=OLH7Et5fT_iJ9PJ8vJJf9w'

{"access_token":"u-xpX1lTTHao8ZZWSfTslA","expires_in":3600,"refresh_token":"Q2y_YrNgToWtXj8O-Geb7A","token_type":"Bearer"}

curl -v 'http://localhost:14000/token' \
-H "Authorization: Basic $(echo -ne 'public-client:' | base64)" \
-d 'grant_type=refresh_token&refresh_token=Q2y_YrNgToWtXj8O-Geb7A'
