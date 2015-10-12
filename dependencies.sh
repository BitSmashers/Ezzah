source ./env.toSource
if [ -e ./env.toSource.local ]
  then
	source ./env.toSource.local
fi

echo $GOPATH

cd ui/
bower update

cd ../

cd core/

go get github.com/constabulary/gb/...
sleep 1

# Go dependencies
echo "Downloading go dependencies...."

echo "go-json-rest from github.com/ant0ine/go-json-rest/rest"
gb vendor fetch github.com/ant0ine/go-json-rest/rest

echo "httpgzip from github.com/daaku/go.httpgzip"
gb vendor fetch github.com/daaku/go.httpgzip

echo "sqlx from github.com/jmoiron/sqlx"
gb vendor fetch github.com/jmoiron/sqlx

echo "cypher driver (neoism) from github.com/jmcvetta/neoism"
gb vendor fetch github.com/jmcvetta/neoism

echo "Go dependencies updated"
