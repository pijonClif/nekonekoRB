//a good fricking tutorial -- https://www.youtube.com/watch?v=iWp-mCIQgMU&list=PLVotA8ycjnCsy30WQCwVU5RrZkt4lLgY5

package main

import (
	_ "embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ()

var (
	running bool = true
	neko    *Neko
)

func drawScene() {
	neko.DrawNeko(int32(screenWidth), int32(screenHeight))
	DrawPomoOverlay()
}

func render() {
	rl.BeginDrawing()
	//curses to opengl curses to last a thousand lifetimes ((i forgot to enable it))
	rl.ClearBackground(rl.NewColor(0, 0, 0, 0))
	drawScene()

	rl.EndDrawing()
}

func quit() {
	rl.CloseWindow()
}

func update() {
	running = !rl.WindowShouldClose()
}

func main() {

	rl.SetConfigFlags(rl.FlagWindowUndecorated | rl.FlagWindowTopmost | rl.FlagWindowTransparent)
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "NekoNekoRB")
	rl.SetExitKey(0)
	defer rl.CloseWindow()

	//inti cta and pomo
	neko = InitNeko()
	defer neko.UnloadNeko()

	InitPomo()
	defer UnloadPomo()

	for running {

		//updatee
		update()

		if rl.IsMouseButtonPressed(rl.MouseRightButton) && !IsPomoActive() {
			StartPomo()
		}
		UpdatePomo()

		neko.UpdateNeko()

		//nekoko methods
		neko.ClickNDrag()
		neko.HandleFall()
		neko.FallNDrag()

		render()
	}
	quit()
}
