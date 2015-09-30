#!/bin/sh


if [ "$NEODB_PATH" = "" ];
  then
	echo "environment var NEODB_PATH is not set, please see env.toSource file."
	exit -1
  else
	echo "Starting Neo4J instance..."
fi

echo "Downloading go dependencies...."
echo "go-json-rest from github.com/ant0ine/go-json-rest/rest"
go get github.com/ant0ine/go-json-rest/rest
echo "Go dependencies updated"



