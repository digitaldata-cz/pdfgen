module github.com/digitaldata-cz/pdfgen

go 1.18

replace github.com/digitaldata-cz/pdfgen/proto/go => ./proto/go

require (
	github.com/digitaldata-cz/htmltopdf v0.0.0-20220509182308-3449e5e90b74
	github.com/digitaldata-cz/pdfgen/proto/go v0.0.0-00000000000000-000000000000
	google.golang.org/grpc v1.46.2
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
