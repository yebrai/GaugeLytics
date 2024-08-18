#!/bin/bash

# Define the URL to check
URL=http://localhost:8080/health

# Make a request to the health check endpoint
HTTP_RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" $URL)

# Check if the response code is 200
if [ "$HTTP_RESPONSE" -eq 200 ]; then
  echo "Health-check passed: Service is up and running."
  exit 0
else
  echo "Health-check failed: Service responded with status code $HTTP_RESPONSE."
  exit 1
fi
