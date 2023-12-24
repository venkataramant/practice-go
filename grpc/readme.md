# Installing protoc
    brew install protobuf

    https://github.com/LinkedInLearning/grpc-in-go-4400430/blob/02_02b/gen.go

    protoc --go_out=. rides.proto

    https://protobuf.dev/reference/protobuf/google.protobuf/#google.protobuf.FieldMask

     protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto

Run the server
    lsof | grep 8282

Reflections

