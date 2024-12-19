#!/bin/bash 

PATH_FROM=./
PATH_TO=./

clear
protoc --go_out=${PATH_TO} --go_opt=paths=import --go-grpc_out=${PATH_TO} --go-grpc_opt=paths=import ${PATH_FROM}postgres_migrate.proto