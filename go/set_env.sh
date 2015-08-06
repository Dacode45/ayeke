#GO environ variables

export PATH=$GOROOT/bin:$PATH
GOPATH=`pwd`:$GOPATH

#Environ variables for Go web
GOWEBEXP_DIR="src/github.com/dacode45/gowebexp"

export GOWEBEXP_STATIC = `pwd`/$GOWEBEXP_DIR"/web/static"
export GOWEBEXP_TEMPLATES=`pwd`/$GOWEBEXP_DIR"/web/templates"
