#!/bin/sh

source ./env.toSource
if [ -e ./env.toSource.local ]
  then
	source ./env.toSource.local
fi

cd core

# Check environment
if [ "$NEODB_PATH" = "" ];
  then
	echo "environment var NEODB_PATH is not set, please see env.toSource file."
	exit -1
  else
	echo "Starting Neo4J instance... $NEODB_PATH"
fi

gb build all && ./bin/main
