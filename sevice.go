package main

import (
	"bytes"
	"context"
	"net"
	"os"
	"strings"

	"github.com/digitaldata-cz/htmltopdf"
	pb "github.com/digitaldata-cz/pdfgen/proto/go"
	"google.golang.org/grpc"
)

type tGrpcServer struct {
	pb.UnimplementedPdfGenServer
}

func (p *tProgram) run() {
	p.loadConfig()
	go startServer(p.config)
}

func startServer(config *tConfig) {
	listener, err := net.Listen("tcp", net.JoinHostPort(config.Address, config.Port))
	if err != nil {
		logger.Errorf("errStartServer: %s", err.Error())
		os.Exit(1)
	}
	s := grpc.NewServer(grpc.MaxMsgSize(1024 * 1024 * 100))
	pb.RegisterPdfGenServer(s, &tGrpcServer{})
	logger.Infof("server listening at %s", listener.Addr().String())
	if err := s.Serve(listener); err != nil {
		logger.Errorf("errStartServer: %s", err.Error())
		os.Exit(1)
	}
}

func (s *tGrpcServer) Generate(ctx context.Context, in *pb.GenerateRequest) (*pb.GenerateResponse, error) {
	/* For testing purposes only
	startTime := time.Now()
	defer func() {
		p, _ := peer.FromContext(ctx)
		log.Printf("[%s] Generate request %s", p.Addr.String(), time.Since(startTime))
	}()
	*/
	out := bytes.NewBuffer(nil)
	if err := callFunc(func() error {
		tmpl, err := htmltopdf.NewObjectFromReader(strings.NewReader(in.GetHtmlBody()))
		if err != nil {
			return err
		}
		converter, err := htmltopdf.NewConverter()
		if err != nil {
			return err
		}
		defer converter.Destroy()
		converter.Add(tmpl)

		colorMode := "Color"
		if in.GetGrayscale() {
			colorMode = "Grayscale"
		}
		if in.GetHtmlHeader() != "" {
			headerFile, err := templateToTempFile(in.GetHtmlHeader())
			if err != nil {
				return err
			}
			defer func() {
				headerFile.Close()
				os.Remove(headerFile.Name())
			}()
			tmpl.Header.CustomLocation = headerFile.Name()
		}
		if in.GetHtmlFooter() != "" {
			footerFile, err := templateToTempFile(in.GetHtmlFooter())
			if err != nil {
				return err
			}
			defer func() {
				footerFile.Close()
				os.Remove(footerFile.Name())
			}()
			tmpl.Footer.CustomLocation = footerFile.Name()
		}

		tmpl.EnableJavascript = true
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
