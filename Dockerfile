FROM --platform=amd64 golang:1.19-buster AS builder
WORKDIR /build/
ADD *.go go.mod go.sum ./
ADD proto/go/*.go proto/go/go.mod proto/go/go.sum ./proto/go/
RUN go get
RUN GOOS=linux go build -trimpath -ldflags "-s -w" -a -o pdfgen .

FROM --platform=amd64 ubuntu:22.04
WORKDIR /
RUN apt update && apt upgrade && apt install -y --no-install-recommends \
    ca-certificates \
    fontconfig \
    libfreetype6 \
    libjpeg-turbo8 \
    libpng16-16 \
    libx11-6 \
    libxcb1 \
    libxext6 \
    libxrender1 \
    xfonts-75dpi \
    xfonts-base \
    wget \
    && rm -rf /var/lib/apt/lists/*
RUN wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6.1-2/wkhtmltox_0.12.6.1-2.jammy_amd64.deb \
    && dpkg -i wkhtmltox_0.12.6.1-2.jammy_amd64.deb \
    && rm -f wkhtmltox_0.12.6.1-2.jammy_amd64.deb
COPY --from=builder /build/pdfgen .
ENV IP=0.0.0.0
ENV PORT=50051
EXPOSE 50051
ENTRYPOINT ["/pdfgen"]