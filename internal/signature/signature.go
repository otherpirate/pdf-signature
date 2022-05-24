package signature

import (
	"fmt"
	"strings"

	"github.com/gosimple/slug"
	"github.com/jung-kurt/gofpdf"
)

func Signature(sourceImage, outputDir string, signPosX, signPosY float64, names []string) error {
	for _, name := range names {
		pdf := buildBasePDF(sourceImage)
		pdf.Text(signPosX, signPosY, strings.Title(name))
		filename := fmt.Sprintf("%s/%s.pdf", outputDir, slug.Make(name))
		if err := pdf.OutputFileAndClose(filename); err != nil {
			return fmt.Errorf("couldn't save pdf %s reason %+v", filename, err)
		}
		fmt.Printf("Exported %s\n", filename)
	}
	return nil
}

func buildBasePDF(sourceImage string) *gofpdf.Fpdf {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetMargins(0.1, 0.1, 0.1)
	pdf.AddPage()
	height, width := pdf.GetPageSize()
	pdf.Image(sourceImage, 0.0, 0.0, height, width, false, "", 0, "")
	pdf.SetFont("Arial", "", 16)
	return pdf
}
