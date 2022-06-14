set -e

# For compiling issues
# https://www.notion.so/Protobuf-d5efc374891a452798e3f3a414722eec

# This could be run to nuke it all
#echo "Remove old models..."
#rm ../models/*
cd "$(dirname "$0")"

echo "Starting proto to struct..."
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
#protoc -I=. -I=$GOPATH/src --go_out=.. --gorm_out=engine=postgres:.. *.proto

# Enable extended glob
shopt -s extglob
protoc \
  --proto_path=. \
  --proto_path=${GOPATH}/src \
  --proto_path=${GOPATH}/src/github.com/mwitkow/go-proto-validators \
  --proto_path=${GOPATH}/src/github.com/infobloxopen/protoc-gen-gorm/options \
  --go_out=.. \
  --gorm_out="engine=postgres:.." \
  !(*contract_processed).proto

# For contract_processed.proto which has an optional field
protoc \
  --proto_path=. \
  --proto_path=${GOPATH}/src \
  --proto_path=${GOPATH}/src/github.com/mwitkow/go-proto-validators \
  --go_out=.. \
  contract_processed.proto

# Remove omitempty option
# Credit: https://stackoverflow.com/a/37335452
ls ../models/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'

echo "Completed proto to struct..."
