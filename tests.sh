source ./env.toSource
if [ -e ./env.toSource.local ]
  then
	source ./env.toSource.local
fi

cd core/

gb test api
