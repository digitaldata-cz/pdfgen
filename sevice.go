package main

import (
	"context"
	"net"
	"os"
	"strings"

	htmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
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
	logger.Infof("Starting pdfgen server version %s", appVersion)
	listener, err := net.Listen("tcp", net.JoinHostPort(config.Address, config.Port))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	s := grpc.NewServer(grpc.MaxRecvMsgSize(1024 * 1024 * 100))
	pb.RegisterPdfGenServer(s, &tGrpcServer{})
	logger.Infof("Server listening at %s", listener.Addr().String())
	if err := s.Serve(listener); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func (s *tGrpcServer) Generate(ctx context.Context, in *pb.GenerateRequest) (*pb.GenerateResponse, error) {
	/*
		startTime := time.Now()
		peer, _ := peer.FromContext(ctx)
			logger.Infof("Generate request from %s started", peer.Addr.String())
			defer func() {
				logger.Infof("Generate request from %s finished aster %s", peer.Addr.String(), time.Since(startTime))
			}()
	*/

	pdfg, err := htmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	pdfg.Dpi.Set(uint(in.GetDpi()))
	pdfg.Orientation.Set(in.GetOrientation())
	pdfg.Grayscale.Set(in.GetGrayscale())
	pdfg.PageSize.Set(in.GetPageSize())
	pdfg.MarginTop.Set(uint(in.GetMarginTop()))
	pdfg.MarginBottom.Set(uint(in.GetMarginBottom()))
	pdfg.MarginLeft.Set(uint(in.GetMarginLeft()))
	pdfg.MarginRight.Set(uint(in.GetMarginRight()))

	pdfPages := htmltopdf.NewPageReader(strings.NewReader(in.GetHtmlBody()))
	pdfPages.DisableJavascript.Set(false)
	pdfPages.Zoom.Set(in.GetZoom())

	if in.GetHtmlHeader() != "" {
		headerFile, err := templateToTempFile(in.GetHtmlHeader())
		if err != nil {
			return nil, err
		}
		defer func() {
			headerFile.Close()
			os.Remove(headerFile.Name())
		}()
		pdfPages.HeaderHTML.Set(headerFile.Name())
	}
	if in.GetHtmlFooter() != "" {
		footerFile, err := templateToTempFile(in.GetHtmlFooter())
		if err != nil {
			return nil, err
		}
		defer func() {
			footerFile.Close()
			os.Remove(footerFile.Name())
		}()
		pdfPages.FooterHTML.Set(footerFile.Name())
	}

	pdfg.AddPage(pdfPages)

	if err := pdfg.Create(); err != nil {
		return nil, err
	}
	return &pb.GenerateResponse{Pdf: pdfg.Bytes()}, nil
}

func templateToTempFile(templateData string) (*os.File, error) {
	if templateData == "" {
		return nil, nil
	}
	file, err := os.CreateTemp("", "template-*.html")
	if err != nil {
		return nil, err
	}
	if _, err := file.Write([]byte(templateData)); err != nil {
		return nil, err
	}
	return file, nil
}
