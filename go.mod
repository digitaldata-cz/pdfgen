module github.com/digitaldata-cz/pdfgen

go 1.18

replace github.com/digitaldata-cz/pdfgen/proto/go => ./proto/go

require (
	github.com/SebastiaanKlippert/go-wkhtmltopdf v1.7.2
	github.com/digitaldata-cz/pdfgen/proto/go v0.0.0
	github.com/kardianos/service v1.2.2-0.20220428125717-29f8c79c511b
	google.golang.org/grpc v1.49.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/kr/text v0.2.0 // indirect
	golang.org/x/net v0.0.0-20220930213112-107f3e3c3b0b // indirect
	golang.org/x/sys v0.0.0-20220928140112-f11e5e49a4ec // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220930163606-c98284e70a91 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
