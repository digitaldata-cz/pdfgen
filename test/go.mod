module github.com/digitaldata-cz/pdfgen/test

go 1.18

replace github.com/digitaldata-cz/pdfgen/proto/go => ../proto/go

require (
	github.com/digitaldata-cz/pdfgen/proto/go v0.0.0-20231217182316-02f3b48d11d6
	google.golang.org/grpc v1.60.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231212172506-995d672761c0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
