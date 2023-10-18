package services

import (
	"log"
	"path"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/tudemaha/ifassion-be/internal/export/dto"
)

func CreatePdf(ResponseData dto.ResponseData) (string, error) {
	pdf := gofpdf.New("P", "mm", "A5", "")

	pdf.AddUTF8Font("Poppins", "", "./files/fonts/Poppins-Regular.ttf")
	pdf.AddUTF8Font("Poppins", "B", "./files/fonts/Poppins-Bold.ttf")
	pdf.AddUTF8Font("Poppins", "I", "./files/fonts/Poppins-Italic.ttf")

	pdf.SetHeaderFunc(func() {
		pdf.SetFont("Poppins", "B", 40)
		width, _ := pdf.GetPageSize()
		pdf.Text(width-70, 15, "IFASSION")
		pdf.Line(0, 20, width, 20)
	})

	pdf.SetFooterFunc(func() {
		pdf.SetFont("Poppins", "", 8)
		_, height := pdf.GetPageSize()
		local, _ := time.LoadLocation("Asia/Makassar")
		pdf.Text(60, height-8, "Date Created: "+time.Now().In(local).Format("Monday, 02 January 2006 15:04:05")+" WITA")
	})

	pdf.AddPage()
	pdf.SetFillColor(47, 151, 148)

	pdf.SetFont("Poppins", "B", 14)
	pdf.Text(48, 27, "IFASSION Test Result")

	col1Width := 40.0
	col2Width := 88.0
	rowHeight := 8.0

	pdf.SetY(35)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Poppins", "B", 12)
	pdf.CellFormat(col1Width, rowHeight, "Test ID", "1", 0, "C", true, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Poppins", "", 12)
	pdf.CellFormat(col2Width, rowHeight, ResponseData.ID, "1", 0, "L", false, 0, "")
	pdf.Ln(-1)

	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Poppins", "B", 12)
	pdf.CellFormat(col1Width, rowHeight, "Test Date", "1", 0, "C", true, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Poppins", "", 12)
	pdf.CellFormat(col2Width, rowHeight, ResponseData.Time, "1", 0, "L", false, 0, "")
	pdf.Ln(-1)

	rowLen := len(ResponseData.Indicators)
	currentHeight := rowHeight * float64(rowLen)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Poppins", "B", 12)
	pdf.CellFormat(col1Width, currentHeight, "Indicators", "1", 0, "C", true, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Poppins", "", 10)
	for i, value := range ResponseData.Indicators {
		pdf.CellFormat(col2Width, rowHeight, value, "1", 0, "L", false, 0, "")
		if i < rowLen-1 {
			pdf.Ln(-1)
			pdf.CellFormat(col1Width, rowHeight, "", "", 0, "C", false, 0, "")
		}
	}
	pdf.Ln(-1)

	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("Poppins", "B", 12)
	pdf.CellFormat(col1Width, rowHeight+2, "Passion", "1", 0, "C", true, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Poppins", "B", 14)
	pdf.CellFormat(col2Width, rowHeight+2, ResponseData.Passion, "1", 0, "L", false, 0, "")
	pdf.Ln(-1)

	filename := ResponseData.ID + ".pdf"
	filePath := path.Join("files", "pdf", filename)
	err := pdf.OutputFileAndClose(filePath)
	if err != nil {
		log.Fatalf("ERROR GenerateReport fatal error:, %v", err)
		return "", err
	}
	return filename, nil
}
