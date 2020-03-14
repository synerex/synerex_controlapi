package synerex_nodeserv_controlapi

// dummy go file for generating pb.go
//go:generate protoc -I . -I ../../nodeapi  --go_out=paths=source_relative,plugins=grpc:. nodeserv_controlapi.proto
