#/bin/sh

URL='https://localhost:8000'
if [[ $# -ge 1 ]]; then
	URL=$1
fi

curl --insecure -X POST \
	-H 'Content-Type: application/json' \
	--data '{"id":0,"name":"test","description":"test product","price":99}' \
	"$URL/products"
