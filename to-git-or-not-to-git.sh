URL="https://01.alem.school/assets/superhero/all.json"

curl -s $URL | jq -r '.[] | select(.id==170) | "\(.name)\n\(.powerstats.power)\n\(.appearance.gender)"'