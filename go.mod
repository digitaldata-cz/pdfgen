module github.com/digitaldata-cz/pdfgen

go 1.18

replace github.com/digitaldata-cz/pdfgen/proto/go => ./proto/go

require (
	github.com/SebastiaanKlippert/go-wkhtmltopdf v1.9.2
	github.com/digitaldata-cz/pdfgen/proto/go v0.0.0
	github.com/kardianos/service v1.2.2
	google.golang.org/grpc v1.60.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/stretchr/testify v1.8.3 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231212172506-995d672761c0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
