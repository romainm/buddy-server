#!/bin/bash
printf "\nRegenerating gqlgen files\n"
rm -f internal/buddy-server/generated.go \
    internal/buddy-server/models/generated.go \
    internal/buddy-server/resolvers/generated.go
time go run -v github.com/99designs/gqlgen $1
printf "\nDone.\n\n"
