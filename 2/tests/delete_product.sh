#/bin/sh

URL='https://localhost:8000'
if [[ $# -ge 1 ]]; then
	URL=$1
fi

ID='0'
if [[ $# -ge 2 ]]; then
	ID=$2
fi

curl --insecure -X DELETE "$URL/products/$ID"
