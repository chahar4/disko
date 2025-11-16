#!/bin/bash


BASE_URL="http://localhost:8080/api/v1"

echo $BASE_URL

RESPONSE=$(curl -s -X POST "$BASE_URL/auth/login" \
    -H "Content-Type: application/json" \
    -d '{"email":"user@email.com","password" : "user1234"}')

echo $RESPONSE
TOKEN=$(echo $RESPONSE | jq -r '.token')



echo $TOKEN
