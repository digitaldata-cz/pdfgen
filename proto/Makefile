rebuild-go:
	protoc \
	--go_out=go \
	--go_opt=paths=source_relative \
	--go-grpc_out=go \
	--go-grpc_opt=paths=source_relative \
	pdfgen.proto

rebuild-js:
	protoc \
	--js_out=import_style=commonjs,binary:js \
	pdfgen.proto


install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
