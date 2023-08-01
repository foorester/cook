#!/bin/bash

# Usage
# ./scripts/curl/create_book.sh -h localhost -p 8080 -u "your_username" -n "Recipe Book One" -d "Favorite Recipes"

# Default values
host="localhost"
port="8080"
username="johndoe"
default_username=$(uuidgen)  # Generate a random UUID as the default username

# Command-line arguments
while getopts ":h:p:u:n:d:" opt; do
  case $opt in
    h) host="$OPTARG" ;;
    p) port="$OPTARG" ;;
    u) username="$OPTARG" ;;
    n) name="$OPTARG" ;;
    d) description="$OPTARG" ;;
    \?) echo "Invalid option -$OPTARG" >&2; exit 1 ;;
  esac
done

# Check if the username is empty, then assign the default username
if [ -z "$username" ]; then
  username="$default_username"
fi

# URL
url="http://$host:$port/api/v1/$username/books"

# Payload
payload="{ \"Name\": \"$name\", \"Description\": \"$description\" }"

# Post request using curl
curl -X POST -H "Content-Type: application/json" -d "$payload" "$url"
