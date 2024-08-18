#!/bin/bash

# Definir la URL del health check
URL="${HEALTH_CHECK_URL:-http://localhost:8080/health}"

# Realizar la solicitud al endpoint de health-check
response=$(curl --write-out "%{http_code}" --silent --output /dev/null "$URL")

# Verificar el código de estado HTTP
if [ "$response" -eq 200 ]; then
  echo "Health check passed."
  exit 0
else
  echo "Health check failed with status code $response."
  exit 1
fi
