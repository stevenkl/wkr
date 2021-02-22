#!/usr/bin/env sh

METHOD=$1
ROUTE=$2
CONTENT=$3

RESPONSE=$(curl -X "${METHOD:-GET}" \
	-H 'Content-Type: application/json' \
	-d "${CONTENT:-\{\}}" \
	http://127.0.0.1:8000${ROUTE:-/} 2> NUL)


echo "Response: $RESPONSE"
