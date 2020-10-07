protoc --proto_path=. --proto_path=third_party --go_out=plugins=grpc:. pingpong.proto

protoc --proto_path=. --proto_path=third_party --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. pingpong.proto