package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	pb "github.com/digitaldata-cz/pdfgen/proto/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPdfGenClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Generate(ctx, &pb.GenerateRequest{HtmlBody: `<html>
<body>
	<h1>Lorem ipsum...</h1>
	<hr>
	<p>Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Vivamus ac leo pretium faucibus.
	Duis ante orci, molestie vitae vehicula venenatis, tincidunt ac pede. Nulla est. Duis sapien
	nunc, commodo et, interdum suscipit, sollicitudin et, dolor. Mauris tincidunt sem sed arcu.
	Nam libero tempore, cum soluta nobis est eligendi optio cumque nihil impedit quo minus id
	quod maxime placeat facere possimus, omnis voluptas assumenda est, omnis dolor repellendus.
	Pellentesque pretium lectus id turpis. Phasellus et lorem id felis nonummy placerat. Maecenas
	ipsum velit, consectetuer eu lobortis ut, dictum at dui. Etiam commodo dui eget wisi. Duis
	pulvinar. Cras elementum. Nullam faucibus mi quis velit. Mauris suscipit, ligula sit amet
	pharetra semper, nibh ante cursus purus, vel sagittis velit mauris vel metus. Vestibulum
	fermentum tortor id mi. Morbi scelerisque luctus velit.</p>
</body>
</html>
`})
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	ioutil.WriteFile("test.pdf", r.Pdf, 0644)
}
