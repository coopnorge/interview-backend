module github.com/coopnorge/interview-backend/src/client

go 1.21

replace github.com/coopnorge/interview-backend/src/generated => ../generated

require (
	github.com/brianvoe/gofakeit/v6 v6.28.0
	github.com/coopnorge/interview-backend/src/generated v0.0.0-20240125094456-5b766eadd0e1
	github.com/google/wire v0.6.0
	google.golang.org/grpc v1.62.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20240304212257-790db918fca8 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240304212257-790db918fca8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240304212257-790db918fca8 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)
