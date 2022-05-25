protoc ^
	--go-grpc_out=go ^
	--go-grpc_opt=paths=source_relative ^
	--go_out=go ^
	--go_opt=paths=source_relative ^
	--js_out=import_style=commonjs,binary:js ^
	pdfgen.proto
