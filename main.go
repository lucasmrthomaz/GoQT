package main

import (
	"os"

	"github.com/therecipe/qt/core"

	"github.com/therecipe/qt/widgets"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(400, 200)
	window.SetWindowTitle("Exemplo Qt em Go")

	centralWidget := widgets.NewQWidget(nil, 0)
	centralWidget.SetLayout(widgets.NewQVBoxLayout())

	label := widgets.NewQLabel2("Olá, Mundo!", nil, 0)
	label.SetAlignment(core.Qt__AlignCenter)

	button := widgets.NewQPushButton2("Clique em mim", nil)
	button.ConnectClicked(func(bool) {
		label.SetText("Botão Clicado!")
	})

	centralWidget.Layout().AddWidget(label)
	centralWidget.Layout().AddWidget(button)

	window.SetCentralWidget(centralWidget)

	window.Show()

	app.Exec()
}
