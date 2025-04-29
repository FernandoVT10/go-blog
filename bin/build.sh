#!/bin/sh -e

pnpm run build
go build -o $(pwd)/webapp ./app
