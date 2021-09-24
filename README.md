# MockyMcMockface

A GRPC mock server

## Prerequisites

Protocol Buffers

### TODO
- Add runner, so the app can be built into an image and ran
- Add full ability to create responses

protoc --plugin=mockymock=build\protoc-gen-mockymock.exe --mockymock_out=test -I test\basic\protobuf test\basic\protobuf\basic.proto --go-grpc_out=test\basic\grpc --go_out=test\basic\grpc


protoc --plugin=mockymock=build/protoc-gen-mockymock --mockymock_out=test -I test/basic/protobuf test/basic/protobuf/basic.proto --go-grpc_out=test/basic/grpc --go_out=test/basic/grpc