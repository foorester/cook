#!/bin/bash

# Usage
# ./scripts/curl/get_books.sh -h localhost -p 8080 -u "your_username"

# Default values
host="localhost"
port="8080"
username="johndoe"

# Command-line arguments
while getopts ":h:p:u:" opt; do
  case $opt in
    h) host="$OPTARG" ;;
    p) port="$OPTARG" ;;
    u) username="$OPTARG" ;;
    \?) echo "Invalid option -$OPTARG" >&2; exit 1 ;;
  esac
done

# URL
url="http://$host:$port/api/v1/$username/books"

# Get request using curl
curl -X GET "$url"
