#!/bin/bash

# Usage
# ./scripts/curl/create_book.sh -h localhost -p 8080 -n "Recipe Book One" -d "Favorite Recipes"

# Default values
host="localhost"
port="8080"

# Command-line arguments
while getopts ":h:p:n:d:" opt; do
  case $opt in
    h) host="$OPTARG" ;;
    p) port="$OPTARG" ;;
    n) name="$OPTARG" ;;
    d) description="$OPTARG" ;;
    \?) echo "Invalid option -$OPTARG" >&2; exit 1 ;;
  esac
done

# URL
url="http://$host:$port/api/v1/books"

# Payload
payload="{ \"Name\": \"$name\", \"Description\": \"$description\" }"

# Post request using curl
curl -X POST -H "Content-Type: application/json" -d "$payload" "$url"
