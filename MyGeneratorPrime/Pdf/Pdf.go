package pdf

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

type PdfSaver struct {
	OutputDir string
}

func Createpdf(Name, Prime, Img string) {
	
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetTitle("Avis de recherche", true)
	// pageWidth, pageHeight := pdf.GetPageSize()

	pdf.Image("image/marron.jpeg", 0, 0, 300, 0, false, "", 0,"")
	pdf.Image("image/wantedVierge.png", 0, 0, 210, 0, false, "", 0,"")
	pdf.Image(Img, 25, 70, 160, 110, false, "",0,"")
	fmt.Println("TEST")
	

	pdf.SetFont("Arial", "B", 40)
	pdf.SetTextColor(61, 43, 12)
	pdf.SetXY(50, 231)
	pdf.WriteAligned(0, 0, Name, "C")

	pdf.SetFont("Arial", "B", 52)
	pdf.SetTextColor(61, 43, 12)
	pdf.SetXY(40, 260)
	pdf.WriteAligned(0, 0, Prime, "C")


	outputFile := Name + ".pdf"

	err := pdf.OutputFileAndClose(outputFile)

	if err != nil {
		fmt.Println("erreur lors de la création du PDF: %w", err)
	} else {
		fmt.Println("OK.")
	}
}
