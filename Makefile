all:
	docker build -t digitaldata/pdfgen -t digitaldata/pdfgen:1.1 .

run:
	docker run --rm -it -p 50051:50051 --platform linux/amd64 digitaldata/pdfgen

publish:
	docker push digitaldata/pdfgen:1.1
	docker push digitaldata/pdfgen:latest