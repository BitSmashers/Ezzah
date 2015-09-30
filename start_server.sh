#!/bin/sh

source env.toSource

# Check environment
if [ "$NEODB_PATH" = "" ];
  then
	echo "environment var NEODB_PATH is not set, please see env.toSource file."
	exit -1
  else
	echo "Starting Neo4J instance... $NEODB_PATH"
fi

# Go dependencies
echo "Downloading go dependencies...."
echo "go-json-rest from github.com/ant0ine/go-json-rest/rest"
go get github.com/ant0ine/go-json-rest/rest
echo "httpgzip from github.com/daaku/go.httpgzip"
go get github.com/daaku/go.httpgzip
echo "Go dependencies updated"

# Run Ezzah binary
go run core/src/main/main.go
