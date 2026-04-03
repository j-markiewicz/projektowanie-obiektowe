#/bin/sh

URL='https://localhost:8000'
if [[ $# -ge 1 ]]; then
	URL=$1
fi

ID='0'
if [[ $# -ge 2 ]]; then
	ID=$2
fi

curl --insecure -X PUT \
	-H 'Content-Type: application/json' \
	--data '{"id":0,"name":"test","description":"test product","price":199}' \
	"$URL/products/$ID"
