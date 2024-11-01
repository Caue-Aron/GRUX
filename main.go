package main

import (
	wg "main/widgets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetTraceLogLevel(rl.LogNone)

	rl.InitWindow(800, 600, "GRUX")
	defer rl.CloseWindow()

	wg.InitWidgetsContext()
	defer wg.DestroyWidgetsContext()

	panel := wg.NewPanel(wg.PanelStyleNormal | wg.PanelStyleDoubleBordered)
	panel.CalculateBounds(&rl.Rectangle{
		X: 100, Y: 100,
		Width: 600, Height: 250})
	panel.SetContainer(wg.NewSingleContainer(wg.NewPanel(wg.PanelStyleSunken), 0))

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {

		panel.Update()

		rl.BeginDrawing()
		rl.ClearBackground(wg.ColorBackground)

		panel.Draw()

		rl.EndDrawing()
	}
}
