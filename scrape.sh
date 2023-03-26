#!/bin/bash
set -euC

declare -a SCRAPING_URL_LIST=(
    "https://note.com/youth_waster/n/nfd7767496d21"
    "https://note.com/youth_waster/n/n3dbfaa732972"
)

for list in "${SCRAPING_URL_LIST[@]}"
do
    go run main.go "$list"
done
