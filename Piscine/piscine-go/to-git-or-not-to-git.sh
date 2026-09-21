#!/bin/bash
curl -s https://acad.learn2earn.org.ng/assets/superhero/all.json \ | jq '.[] | select(.id == 170) | "\(.name)\n\(.strength)\n\(.gender)"'
