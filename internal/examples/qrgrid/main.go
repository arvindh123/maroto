package main

import (
	"fmt"
	"github.com/arvindh123/maroto/pkg/consts"
	"github.com/arvindh123/maroto/pkg/pdf"
	"github.com/arvindh123/maroto/pkg/props"
	"os"
	"time"
)

func main() {
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.SetBorder(true)

	m.Row(40, func() {
		m.Col(2, func() {
			m.QrCode("https://github.com/arvindh123/maroto", props.Rect{
				Percent: 50,
			})
		})
		m.Col(4, func() {
			m.QrCode("https://github.com/arvindh123/maroto", props.Rect{
				Percent: 75,
			})
		})
		m.Col(6, func() {
			m.QrCode("https://github.com/arvindh123/maroto", props.Rect{
				Percent: 100,
			})
		})
	})

	m.Row(40, func() {
		m.Col(2, func() {
			m.QrCode("https://github.com/arvindh123/maroto", props.Rect{
				Center:  true,
				Percent: 50,
			})
		})
		m.Col(4, func() {
			m.QrCode("https://github.com/arvindh123/maroto", props.Rect{
				Center:  true,
				Percent: 75,
			})
		})
		m.Col(6, func() {
			m.QrCode("https://github.com/arvindh123/maroto", props.Rect{
				Center:  true,
				Percent: 100,
			})
		})
	})

	m.Row(40, func() {
		m.Col(6, func() {
			m.QrCode("https://github.com/arvindh123/maroto", props.Rect{
				Percent: 50,
			})
		})
		m.Col(4, func() {
			m.QrCode("https://github.com/arvindh123/maroto", props.Rect{
				Percent: 75,
			})
		})
		m.Col(2, func() {
			m.QrCode("https://github.com/arvindh123/maroto", props.Rect{
				Percent: 100,
			})
		})
	})

	m.Row(40, func() {
		m.Col(6, func() {
			m.QrCode("https://github.com/arvindh123/maroto", props.Rect{
				Center:  true,
				Percent: 50,
			})
		})
		m.Col(4, func() {
			m.QrCode("https://github.com/arvindh123/maroto", props.Rect{
				Center:  true,
				Percent: 75,
			})
		})
		m.Col(2, func() {
			m.QrCode("https://github.com/arvindh123/maroto", props.Rect{
				Center:  true,
				Percent: 100,
			})
		})
	})

	err := m.OutputFileAndClose("internal/examples/pdfs/qrgrid.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}
