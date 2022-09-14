package main

import (
	"context"
	"log"
	"os"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.Generate(ctx, &pb.GenerateRequest{
		Name:         "test",
		Dpi:          96,
		Zoom:         1,
		PageSize:     "A4",
		Grayscale:    false,
		Orientation:  "Landscape",
		MarginLeft:   10,
		MarginRight:  10,
		MarginTop:    10,
		MarginBottom: 10,
		HtmlBody: `
<!DOCTYPE html>
<html>
<body>
	<h1 style="color: red">Lorem ipsum...</h1>
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
`,
		HtmlHeader: "<!DOCTYPE html> <h1>Header</h1>",
		HtmlFooter: `
<!DOCTYPE html>  
<html>    
<body onload="getPdfInfo()">      
    <p style="width: 30mm; display: inline-block;">FOOTER - Page:</p>      
    <p style="width: 5mm; display: inline-block;" id="pdfkit_page_current"></p>     
    <p style="width: 3mm; display: inline-block;">/</p>
    <p style="width: 8mm; display: inline-block;" id="pdfkit_page_count"></p>
</body>
</html>
<script>
var pdfInfo = {};
var x = document.location.search.substring(1).split('&');
for (var i in x) {
    var z = x[i].split('=', 2);
    pdfInfo[z[0]] = unescape(z[1]);
}
function getPdfInfo() {      
    var page = pdfInfo.page || 1;
    var pageCount = pdfInfo.topage || 1;
    document.getElementById('pdfkit_page_current').textContent = page;
    document.getElementById('pdfkit_page_count').textContent = pageCount;
}  
</script>
		`,
	})
	if err != nil {
		log.Fatalf("Error1: %s", err.Error())
	}
	if r.Error != "" {
		log.Fatalf("Error2: %s", r.Error)
	}
	os.WriteFile("test.pdf", r.Pdf, 0644)
}
