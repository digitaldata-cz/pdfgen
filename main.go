package main

import (
	"bytes"
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"github.com/digitaldata-cz/htmltopdf"
	pb "github.com/digitaldata-cz/pdfgen/proto/go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

var run = make(chan func())

type tGrpcServer struct {
	// TODO: Nastudovat k cemu je "UnimplementedPrinterServer"
	pb.UnimplementedPdfGenServer
}

func init() {
	// Set main function to run on the main thread.
	runtime.LockOSThread()

	// Initialize library.
	if err := htmltopdf.Init(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer htmltopdf.Destroy()

	go startServer()

	// Listen for functions that need to run on the main thread.
	var quit = make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	for {
		select {
		case f := <-run:
			f()
		case <-quit:
			log.Println("shutting down")
			return
		}
	}
}

// callFunc calls the provided function on the main thread.
func callFunc(f func() error) error {
	err := make(chan error)
	run <- func() {
		err <- f()
	}
	return <-err
}

func startServer() {
	var (
		ipAddress = "::"
		port      = "50051"
	)

	if s := os.Getenv("PS_IP"); s != "" {
		ipAddress = s
	}

	if s := os.Getenv("PS_PORT"); s != "" {
		port = s
	}

	listener, err := net.Listen("tcp", net.JoinHostPort(ipAddress, port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterPdfGenServer(s, &tGrpcServer{})
	log.Printf("server listening at %s", listener.Addr().String())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}

func (s *tGrpcServer) Generate(ctx context.Context, in *pb.GenerateRequest) (*pb.GenerateResponse, error) {
	startTime := time.Now()
	defer func() {
		p, _ := peer.FromContext(ctx)
		log.Printf("[%s] Generate request %s", p.Addr.String(), time.Since(startTime))
	}()
	out := bytes.NewBuffer(nil)
	if err := callFunc(func() error {
		tmpl, err := htmltopdf.NewObjectFromReader(strings.NewReader(in.GetHtmlBody()))
		if err != nil {
			return err
		}
		converter, err := htmltopdf.NewConverter()
		if err != nil {
			log.Fatal(err)
		}
		defer converter.Destroy()
		converter.Add(tmpl)

		colorMode := "Color"
		if in.GetGrayscale() {
			colorMode = "Grayscale"
		}

		headerFile, err := templateToTempFile(in.GetHtmlHeader())
		if err != nil {
			return err
		}
		defer func() {
			headerFile.Close()
			os.Remove(headerFile.Name())
		}()
		footerFile, err := templateToTempFile(in.GetHtmlFooter())
		if err != nil {
			return err
		}
		defer func() {
			footerFile.Close()
			os.Remove(footerFile.Name())
		}()

		tmpl.Header.CustomLocation = headerFile.Name()
		tmpl.Footer.CustomLocation = footerFile.Name()
		tmpl.Zoom = in.GetZoom()
		converter.DPI = in.GetDpi()
		converter.PaperSize = htmltopdf.PaperSize(in.GetPageSize())
		converter.Orientation = htmltopdf.Orientation(in.GetOrientation())
		converter.Colorspace = htmltopdf.Colorspace(colorMode)
		converter.MarginLeft = in.GetMarginLeft()
		converter.MarginRight = in.GetMarginRight()
		converter.MarginTop = in.GetMarginTop()
		converter.MarginBottom = in.GetMarginBottom()
		converter.UseCompression = true
		return converter.Run(out)
	}); err != nil {
		return &pb.GenerateResponse{Pdf: nil, Error: err.Error()}, nil
	}
	return &pb.GenerateResponse{Pdf: out.Bytes()}, nil
}

func templateToTempFile(templateData string) (*os.File, error) {
	if templateData == "" {
		return nil, nil
	}
	file, err := os.CreateTemp("", "portunusTmpl-*.html")
	if err != nil {
		return nil, err
	}
	if _, err := file.Write([]byte(templateData)); err != nil {
		return nil, err
	}
	return file, nil
}
