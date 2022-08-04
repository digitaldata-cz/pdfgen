module github.com/digitaldata-cz/pdfgen

go 1.18

replace github.com/digitaldata-cz/pdfgen/proto/go => ./proto/go

require (
	github.com/SebastiaanKlippert/go-wkhtmltopdf v1.7.2
	github.com/digitaldata-cz/pdfgen/proto/go v0.0.0
	github.com/kardianos/service v1.2.2-0.20220428125717-29f8c79c511b
	google.golang.org/grpc v1.48.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/kr/text v0.2.0 // indirect
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b // indirect
	golang.org/x/sys v0.0.0-20220722155257-8c9f86f7a55f // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220722212130-b98a9ff5e252 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
