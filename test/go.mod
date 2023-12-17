module github.com/digitaldata-cz/pdfgen/test

go 1.18

replace github.com/digitaldata-cz/pdfgen/proto/go => ../proto/go

require (
	github.com/digitaldata-cz/pdfgen/proto/go v0.0.0
	google.golang.org/grpc v1.56.3
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
