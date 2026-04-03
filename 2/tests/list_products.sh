#/bin/sh

URL='https://localhost:8000'
if [[ $# -ge 1 ]]; then
	URL=$1
fi

curl --insecure "$URL/products"
