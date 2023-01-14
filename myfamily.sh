URL="https://01.alem.school/assets/superhero/all.json"

curl -s $URL | jq  ".[] | select (.id == $HERO_ID) | .connections.relatives" | tr -d '"'