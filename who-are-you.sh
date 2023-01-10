URL="https://01.alem.school/assets/superhero/all.json"

curl -s $URL | jq '.[] | select(.id==70) | "\(.name)"'

