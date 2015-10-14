source ./env.toSource
if [ -e ./env.toSource.local ]
  then
	source ./env.toSource.local
fi

echo $GOPATH

cd ui/
bower update

cd ../

