#!/bin/bash
echo Hello Terminal!

echo "Script Name: $0"
echo "First Parameter of the script is $1"
echo "The second Parameter is $2"

echo "The complete list of arguments is $@"
echo "Total Number of Parameters: $#"
echo "The process ID is $$"
echo "Exit code for the script: $?"

echo "Go: "
go version
go run hello_world.go
go build hello_world.go

