package main

import (
	"fmt"
	"os"

	"bitbucket.org/junglee_games/getsetgo/utils/files"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 10, 10)

	buildCreditNote(m)
	bytes, err := m.Output()
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}
	err = files.SaveFile("pdfs/", "credit"+".pdf", string(bytes.Bytes()))
	if err != nil {
		fmt.Println("⚠️  Could not save PDF:", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")
}

func buildCreditNote(m pdf.Maroto) {
	m.Row(20, func() {
		m.SetBorder(true)
		m.Col(12, func() {
			err := m.FileImage("images/xyz.png", props.Rect{
				Center:  false,
				Left:    5,
				Percent: 80,
				Top:     2,
			})

			if err != nil {
				fmt.Println("Image file was not loaded 😱 - ", err)
			}

			m.Text("XYZ", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
			})
			m.Text("Tel no. :", props.Text{
				Top:   10,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})
	addBreak(m)
	m.Row(15, func() {
		m.SetBackgroundColor(getTealColor())
		m.Col(12, func() {
			m.Text("Credit Note", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
				Size:  20,
			})
		})
		m.SetBackgroundColor(color.NewWhite())
	})
	m.Row(5, func() {
		addColLeft(m, 5, "Document No:")
		addColLeft(m, 7, "Against invoice:")
	})
	m.Row(5, func() {
		addColLeft(m, 5, "Date of Issue:")
		addColLeft(m, 7, "Date of Invoice:")
	})
	m.Row(5, func() {
		addColLeft(m, 3, "State:")
		addColLeft(m, 1, "Code")
		m.Col(1, func() {})
		m.Col(7, func() {})
	})
	addBreak(m)
	m.Row(5, func() {
		m.SetBackgroundColor(getTealColor())
		addColCenter(m, 12, "Detail of Receiver (Billed to)")
		m.SetBackgroundColor(color.NewWhite())
	})
	m.Row(5, func() {
		addColLeft(m, 12, "Name:")
	})
	m.Row(10, func() {
		addColLeft(m, 12, "Address:")
	})
	m.Row(5, func() {
		addColLeft(m, 12, "GSTIN:")
	})
	m.Row(5, func() {
		addColLeft(m, 8, "State:")
		addColLeft(m, 2, "Code")
		m.Col(2, func() {})
	})
	addBreak(m)
	m.Row(10, func() {
		m.SetBackgroundColor(getTealColor())
		addColCenterTable(m, 1, "S No.", false)
		addColCenterTable(m, 2, "Product Description", false)
		addColCenterTable(m, 1, "HSN/SAC Code", false)
		addColCenterTable(m, 1, "UOM", false)
		addColCenterTable(m, 1, "Amount", false)
		addColCenterTable(m, 1, "Discount", false)
		addColCenterTable(m, 1, "Taxable Value", false)
		addColCenterTable(m, 1, "CGST("+"18"+")", true)
		addColCenterTable(m, 1, "SGST("+"18"+")", true)
		addColCenterTable(m, 1, "IGST("+"18"+")", true)
		addColCenterTable(m, 1, "TOTAL", true)
		m.SetBackgroundColor(color.NewWhite())
	})
	m.Row(5, func() {
		m.SetBackgroundColor(getTealColor())
		addColCenter(m, 6, "Total amount in words")
		m.SetBackgroundColor(color.NewWhite())
		addColLeft(m, 4, "Total Amount before Tax")
		addColLeft(m, 2, "")
	})
	m.Row(20, func() {
		addColLeft(m, 6, "")
		m.Col(4, func() {
			m.Text("Add: CGST", props.Text{
				Align: consts.Left,
				Left:  1,
			})
			m.Text("Add: SGST", props.Text{
				Align: consts.Left,
				Left:  1,
				Top:   5,
			})
			m.Text("Total Tax Amount", props.Text{
				Align: consts.Left,
				Left:  1,
				Top:   10,
			})
			m.Text("Total Amount after Tax:", props.Text{
				Align: consts.Left,
				Left:  1,
				Top:   15,
			})
		})
		m.Col(2, func() {
			m.Text("", props.Text{
				Align: consts.Left,
				Left:  1,
			})
			m.Text("", props.Text{
				Align: consts.Left,
				Left:  1,
				Top:   5,
			})
			m.Text("", props.Text{
				Align: consts.Left,
				Left:  1,
				Top:   10,
			})
			m.Text("", props.Text{
				Align: consts.Left,
				Left:  1,
				Top:   15,
			})
		})
	})
	m.Row(5, func() {
		m.SetBackgroundColor(getTealColor())
		addColLeftBold(m, 6, "Bank Details")
		addColLeft(m, 6, "")
		m.SetBackgroundColor(color.NewWhite())
	})
	m.Row(40, func() {
		m.Col(6, func() {
			m.Text("Bank Name:", props.Text{
				Align: consts.Left,
				Left:  1,
			})
			m.Text("Bank A/C:", props.Text{
				Align: consts.Left,
				Left:  1,
				Top:   5,
			})
			m.Text("Bank IFSC:", props.Text{
				Align: consts.Left,
				Left:  1,
				Top:   10,
			})
			m.Text("Terms & conditions:-", props.Text{
				Align: consts.Left,
				Left:  1,
				Top:   15,
			})
		})
		m.Col(6, func() {
			m.Text("Ceritified that the particulars given above are true and correct", props.Text{
				Align: consts.Center,
				Left:  1,
				Size:  7,
			})
			m.Text("Junglee Games", props.Text{
				Align: consts.Center,
				Left:  1,
				Top:   4,
				Style: consts.Bold,
			})
			m.Text("Authorised signatory", props.Text{
				Align: consts.Center,
				Left:  1,
				Top:   35,
				Style: consts.Bold,
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Head office", props.Text{
				Align: consts.Center,
			})

			m.Text("CIN NO : ", props.Text{
				Align: consts.Center,
				Top:   5,
			})
		})
	})
}

func getTealColor() color.Color {
	return color.Color{
		Red:   194,
		Green: 214,
		Blue:  155,
	}
}

func addBreak(m pdf.Maroto) {
	m.Row(4, func() { m.Col(12, func() {}) })
}

func addColLeft(m pdf.Maroto, colWidth uint, text string) {
	m.Col(colWidth, func() {
		m.Text(text, props.Text{
			// Style: consts.Bold,
			Align: consts.Left,
			Left:  1,
		})
	})
}
func addColLeftBold(m pdf.Maroto, colWidth uint, text string) {
	m.Col(colWidth, func() {
		m.Text(text, props.Text{
			Style: consts.Bold,
			Align: consts.Left,
			Left:  1,
		})
	})
}

func addColCenter(m pdf.Maroto, colWidth uint, text string) {
	m.Col(colWidth, func() {
		m.Text(text, props.Text{
			// Style: consts.Bold,
			Align: consts.Center,
			Left:  1,
		})
	})
}
func addColCenterTable(m pdf.Maroto, colWidth uint, text string, bold bool) {
	m.Col(colWidth, func() {
		if bold {
			m.Text(text, props.Text{
				Style: consts.Bold,
				Align: consts.Center,
				Left:  1,
				Size:  7,
			})
		} else {
			m.Text(text, props.Text{
				// Style: consts.Bold,
				Align: consts.Center,
				Left:  1,
				Size:  7,
			})
		}
	})
}
