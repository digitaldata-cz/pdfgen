all:
	docker build -t digitaldata/pdfgen .

run:
	docker run --rm -it -p 50051:50051 --platform linux/amd64 digitaldata/pdfgen

push:
	docker push digitaldata/pdfgen