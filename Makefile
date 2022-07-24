all:
	docker build -t digitaldata/pdfgen .

run:
	docker run -it -p 50051:50051 --platform linux/amd64 digitaldata/pdfgen

push:
	docker push digitaldata/pdfgen