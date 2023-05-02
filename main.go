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
			err := m.FileImage("images/rummy.png", props.Rect{
				Center:  false,
				Left:    5,
				Percent: 80,
				Top:     2,
			})

			if err != nil {
				fmt.Println("Image file was not loaded 😱 - ", err)
			}

			m.Text("Junglee Games", props.Text{
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
	addGap(m)
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
		m.Col(5, func() {
			m.Text("Document No:", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
		m.Col(7, func() {
			m.Text("Against invoice:", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
	})
	m.Row(5, func() {
		m.Col(5, func() {
			m.Text("Date of Issue:", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
		m.Col(7, func() {
			m.Text("Date of Invoice:", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
	})
	m.Row(5, func() {
		m.Col(3, func() {
			m.Text("State:", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
		m.Col(1, func() {
			m.Text("Code", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
		m.Col(1, func() {})
		m.Col(7, func() {})
	})
	addGap(m)
	m.Row(5, func() {
		m.SetBackgroundColor(getTealColor())
		m.Col(5, func() {
			m.Text("Detail of Receiver (Billed to)", props.Text{
				Style: consts.Bold,
				Align: consts.Center,
				Left:  1,
			})
		})
		m.Col(7, func() {
			m.Text("Detail of Consignee (Shipped to)", props.Text{
				Style: consts.Bold,
				Align: consts.Center,
				Left:  1,
			})
		})
		m.SetBackgroundColor(color.NewWhite())
	})
	m.Row(5, func() {
		m.Col(5, func() {
			m.Text("Name:", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
		m.Col(7, func() {
			m.Text("Name:", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
	})
	m.Row(5, func() {
		m.Col(5, func() {
			m.Text("Address:", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
		m.Col(7, func() {
			m.Text("Address:", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
	})
	m.Row(5, func() {
		m.Col(5, func() {
			m.Text("GSTIN:", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
		m.Col(7, func() {
			m.Text("GSTIN:", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
	})
	m.Row(5, func() {
		m.Col(3, func() {
			m.Text("State:", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
		m.Col(1, func() {
			m.Text("Code", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
		m.Col(1, func() {})
		m.Col(3, func() {
			m.Text("State:", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
		m.Col(1, func() {})
		m.Col(1, func() {})
		m.Col(1, func() {
			m.Text("Code", props.Text{
				Style: consts.Bold,
				Align: consts.Left,
				Left:  1,
			})
		})
		m.Col(1, func() {})
	})
	addGap(m)
	m.Row(15, func() {

	})

}

func getTealColor() color.Color {
	return color.Color{
		Red:   194,
		Green: 214,
		Blue:  155,
	}
}

func addGap(m pdf.Maroto) {
	m.Row(4, func() { m.Col(12, func() {}) })
}
