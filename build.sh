echo "Building"
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
BINFILE=$(basename $DIR)
BINFILENAME=$DIR/bin/$BINFILE
PROJSRCPATH=$DIR:$DIR/src
export GOPATH=$PROJSRCPATH
go build -v -o $BINFILENAME main && echo " -> ./bin/$BINFILE"