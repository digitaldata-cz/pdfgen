package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strings"

	"github.com/digitaldata-cz/htmltopdf"
	pb "github.com/digitaldata-cz/pdfgen/proto"

	"google.golang.org/grpc"
)

var (
	run       = make(chan func())
	ipAddress = "0.0.0.0"
	port      = "50051"
)

type tServer struct {
	// TODO: Nastudovat k cemu je "UnimplementedPrinterServer"
	pb.UnimplementedPrinterServer
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
	if os.Getenv("PS_IP") != "" {
		ipAddress = os.Getenv("IP")
	}
	if os.Getenv("PS_PORT") != "" {
		port = os.Getenv("PORT")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", ipAddress, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)

	}
	s := grpc.NewServer()
	pb.RegisterPrinterServer(s, &tServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *tServer) PrintReport(ctx context.Context, in *pb.ReportRequest) (*pb.ReportResponse, error) {
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
		// TODO: vratit chybu na clienta (pridat do protoStruct Err)
		return &pb.ReportResponse{Report: nil}, nil
	}
	fmt.Println("gRPC JEDEEEEEE")
	return &pb.ReportResponse{Report: out.Bytes()}, nil
}
