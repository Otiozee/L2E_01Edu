#!/bin/bash

# Ensure HERO_ID is set to 1
if [ -z "$HERO_ID" ] || [ "$HERO_ID" -ne 1]; then
  echo "Error: HERO_ID environment variable is not set to 1."
  exit 1
fi

# Fetch relatives based on HERO_ID
relatives=$(curl -s https://acad.learn2earn.org.ng/assets/superhero/all.json | \
jq -r ".[] | select(.id == $HERO_ID) | .connections.relatives")

# Convert actual newlines to literal "\n"
relatives_escaped=$(echo "$relatives" | sed ':a;N;$!ba;s/\n/\\n/g')

# Check if relatives is empty or equals "null"
if [ -z "$relatives" ] || [ "$relatives" = "null" ]; then
  echo "No relatives found for HERO_ID=$HERO_ID"
else
  echo "$relatives_escaped"
fi

export HERO_ID=1