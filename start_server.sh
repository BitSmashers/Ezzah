#!/bin/sh

source ./env.toSource
if [ -e ./env.toSource.local ]
  then
	source ./env.toSource.local
fi

cd core

# Go dependencies
echo "Downloading go dependencies...."

echo "go-json-rest from github.com/ant0ine/go-json-rest/rest"
go get -v github.com/ant0ine/go-json-rest/rest

echo "httpgzip from github.com/daaku/go.httpgzip"
go get -v github.com/daaku/go.httpgzip

echo "sqlx from github.com/jmoiron/sqlx"
go get -v github.com/jmoiron/sqlx

echo "cypher driver (neoism) from github.com/jmcvetta/neoism"
go get -v github.com/jmcvetta/neoism

echo "Go dependencies updated"

# Check environment
if [ "$NEODB_PATH" = "" ];
  then
	echo "environment var NEODB_PATH is not set, please see env.toSource file."
	exit -1
  else
	echo "Starting Neo4J instance... $NEODB_PATH"
fi

go install github.com/BitSmashers/Ezzah/main
./bin/main

