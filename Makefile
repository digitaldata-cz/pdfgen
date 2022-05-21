all:
	docker build -t digitaldata/pdfgen .

push:
	docker push digitaldata/pdfgen