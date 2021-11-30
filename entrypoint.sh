#!/bin/bash

sleep 10

migrate -path migrations -database "postgres://db/postgres?sslmode=disable&user=postgres&password=pass" up

exec "$@"