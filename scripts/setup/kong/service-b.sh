echo "Creating Service B"
curl --location --request POST 'http://localhost:8001/services' \
--form 'name="service-b"' \
--form 'protocol="http"' \
--form 'host="service-b"' \
--form 'port="8086"' \
--form 'path="/"'

echo "Creating Service B Routes"
curl --location --request POST 'http://localhost:8001/services/service-b/routes' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=service-b' \
--data-urlencode 'paths=/service-b' \
--data-urlencode 'methods=GET' \
--data-urlencode 'methods=POST' \
--data-urlencode 'methods=PUT' \
--data-urlencode 'methods=OPTIONS' \
--data-urlencode 'methods=DELETE'
