#!/usr/bin/sh

curl --request POST \
    --url http://localhost:1323/customers \
    --header 'content-type: application/json' \
    --data '{
    "cName"    : "akhil",
    "cTel"     : 9892934125,
    "cAddress" : "indiranagar, bangalore"
}'


curl --request POST \
    --url http://localhost:1323/customers \
    --header 'content-type: application/json' \
    --data '{
    "cName"    : "sharma",
    "cTel"     : 9892934125,
    "cAddress" : "indiranagar, bangalore"
}'


curl --request PUT \
    --url http://localhost:1323/customers/2 \
    --header 'content-type: application/json' \
    --data '{
    "cName"    : "akhilsharma",
    "cTel"     : 9892934125,
    "cAddress" : "indiranagar, bangalore"
}'


curl --request GET \
    --url http://localhost:1323/customers

curl --request GET \
    --url http://localhost:1323/report/1

curl --request GET \
    --url 'http://localhost:1323/customers?cName=ha'
