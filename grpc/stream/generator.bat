@rem Generate the Go code for .proto files

setlocal

@rem enter this directory
cd /d %~dp0

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/route_guide.proto

endlocal