package main

import (
	"gocpf/cpfgen"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	mask = "***.***.***-**"
)

func main() {
	gocpf := app.New()
	window := gocpf.NewWindow("Go CPF Generator")
	window.SetTitle("Go CPF Generator")

	cpfLabel := widget.NewLabel("CPF:")
	cpf := widget.NewLabel(mask)

	generate := widget.NewButton("Generate", func() {
		gen, ok := cpfgen.GenCPF()
		if !ok {
			cpf.SetText(mask)
			return
		}
		cpf.SetText(cpfgen.CPFFormat(gen))
	})

	copy := widget.NewButton("Copy", func() {
		window.Clipboard().SetContent(cpf.Text)
	})

	horizontalCPF := container.NewHBox(cpfLabel, cpf)
	horizontalButtons := container.NewHBox(generate, copy)
	vertical := container.NewVBox(horizontalCPF, horizontalButtons)

	window.SetContent(container.New(layout.NewCenterLayout(), vertical))
	window.Resize(fyne.NewSize(270, 100))
	window.ShowAndRun()
}
