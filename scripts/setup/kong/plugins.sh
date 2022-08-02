echo "Configuring JWT Authentication"
curl --location --request POST 'http://localhost:8001/plugins' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=jwt' \
--data-urlencode 'config.secret_is_base64=false' \
--data-urlencode 'config.claims_to_verify=exp' \
--data-urlencode 'config.run_on_preflight=true'

echo "Enabling CORS"
curl --location --request POST 'http://localhost:8001/plugins' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=cors' \
--data-urlencode 'config.origins=*' \
--data-urlencode 'config.methods=GET' \
--data-urlencode 'config.methods=POST' \
--data-urlencode 'config.methods=PUT' \
--data-urlencode 'config.methods=DELETE' \
--data-urlencode 'config.methods=OPTIONS' \
--data-urlencode 'config.headers=Content-Type' \
--data-urlencode 'config.headers=Access-Control-Allow-Origin' \
--data-urlencode 'config.exposed_headers=Access-Control-Allow-Origin' \
--data-urlencode 'config.credentials=true' \
--data-urlencode 'config.preflight_continue=true' \
--data-urlencode 'config.max_age=3600'
