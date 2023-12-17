module github.com/digitaldata-cz/pdfgen/test

go 1.18

replace github.com/digitaldata-cz/pdfgen/proto/go => ../proto/go

require (
	github.com/digitaldata-cz/pdfgen/proto/go v0.0.0-20230112121713-d734e008a729
	google.golang.org/grpc v1.52.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
