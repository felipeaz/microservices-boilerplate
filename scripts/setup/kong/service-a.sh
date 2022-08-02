echo "Creating Service A"
curl --location --request POST 'http://localhost:8001/services' \
--form 'name="service-a"' \
--form 'protocol="http"' \
--form 'host="service-a"' \
--form 'port="8085"' \
--form 'path="/"'

echo "Creating Service A Routes"
curl --location --request POST 'http://localhost:8001/services/service-a/routes' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=service-a' \
--data-urlencode 'paths=/service-a' \
--data-urlencode 'methods=GET' \
--data-urlencode 'methods=POST' \
--data-urlencode 'methods=PUT' \
--data-urlencode 'methods=OPTIONS' \
--data-urlencode 'methods=DELETE'
