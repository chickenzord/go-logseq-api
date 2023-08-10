#!/bin/sh

base_url=localhost:8080
msg="$1"

curl -X POST "${base_url}/journals/today" -d "$msg"
