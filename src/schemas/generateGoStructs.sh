echo "Remove old models..."
set -e

#rm ../models/*
cd "$(dirname "$0")"

echo "Starting proto to struct..."
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
#protoc -I=. -I=$GOPATH/src --go_out=.. --gorm_out=engine=postgres:.. *.proto

protoc \
  --proto_path=. \
  --proto_path=${GOPATH}/src \
  --proto_path=${GOPATH}/src/github.com/mwitkow/go-proto-validators \
  --proto_path=${GOPATH}/src/github.com/infobloxopen/protoc-gen-gorm/options \
  --go_out=.. \
  --gorm_out="engine=postgres:.." \
  *.proto

# Remove omitempty option
# Credit: https://stackoverflow.com/a/37335452
ls ../models/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'

echo "Completed proto to struct..."
