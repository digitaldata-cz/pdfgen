rebuild:
	protoc \
	--go-grpc_out=go \
	--go-grpc_opt=paths=source_relative \
	--go_out=go \
	--go_opt=paths=source_relative \
	--js_out=import_style=commonjs,binary:js \
	pdfgen.proto

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2