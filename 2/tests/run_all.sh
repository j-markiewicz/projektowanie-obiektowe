#/bin/sh

URL='https://localhost:8000'
if [[ $# -ge 1 ]]; then
	URL=$1
fi

ID='0'
if [[ $# -ge 2 ]]; then
	ID=$2
fi

./list_products.sh "$URL"
echo
./create_product.sh "$URL" "$ID"
echo
./get_product.sh "$URL" "$ID"
echo
./update_product.sh "$URL" "$ID"
echo
./get_product.sh "$URL" "$ID"
echo
./list_products.sh "$URL"
echo
./delete_product.sh "$URL" "$ID"
echo
./list_products.sh "$URL"
echo
