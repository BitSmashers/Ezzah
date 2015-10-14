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

