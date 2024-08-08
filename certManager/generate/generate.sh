#!/bin/bash

# Set environment variables
export VAULT_ADDR='http://192.168.0.100:8200'
export VAULT_TOKEN='root'

# Generate JWT token
JWT_TOKEN=$(go run generate_token.go | grep "Generated Token:" | awk '{print $3}')

# Check if JWT token generation was successful
if [ -z "$JWT_TOKEN" ]; then
  echo "Failed to generate JWT token"
  exit 1
fi

# Make request to the /request_cert endpoint
RESPONSE=$(curl -s -X POST -H "Authorization: Bearer $JWT_TOKEN" -H "Content-Type: application/json" -d '{"common_name": "test.surexsend.com"}' http://192.168.0.100:8080/request_cert)

# Print the response
echo "Response from server: $RESPONSE"
