package main

import (
	wg "main/widgets"
	"main/widgets/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetTraceLogLevel(rl.LogNone)

	rl.InitWindow(800, 600, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	wg.InitWidgetsContext()
	defer wg.DestroyWidgetsContext()

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(color.Background)

		rl.EndDrawing()
	}
}
