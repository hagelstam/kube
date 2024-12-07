#!/bin/bash

if [ -z "${API_URL}" ]; then
    echo "API_URL environment variable is required"
    exit 1
fi

apiUrl="${API_URL}" 

url=$(curl -Ls -o /dev/null -w %{url_effective} "https://en.wikipedia.org/wiki/Special:Random")

json_payload=$(jq -n --arg url "$url" '{text: ("Read " + $url)}')

curl -X  POST "${apiUrl}/todos" \
	-H "Content-Type: application/json" \
	-d "$json_payload"
