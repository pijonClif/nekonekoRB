package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	const screenWidth, screenHeight = 128, 128

	rl.SetConfigFlags(rl.FlagWindowUndecorated | rl.FlagWindowTopmost | rl.FlagWindowTransparent)
	rl.InitWindow(screenWidth, screenHeight, "NekoNekoRB")
	defer rl.CloseWindow()

	//inti cta and pomo
	neko := InitNeko()
	defer neko.UnloadNeko()

	InitPomo()
	defer UnloadPomo()

	//e
	var clickStart rl.Vector2
	var isDragging bool
	var isFalling bool

	const clickThreshold = 5.0
	const fallSpeed = 7.0

	//e
	for !rl.WindowShouldClose() {

		neko.UpdateNeko()

		//neko
		mouse := rl.GetMousePosition()
		windowPos := rl.GetWindowPosition()

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			clickStart = mouse
			isDragging = false
		}

		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			moveDist := rl.Vector2Length(rl.Vector2Subtract(mouse, clickStart))
			if moveDist >= clickThreshold {
				isDragging = true
			}
			if isDragging {
				newX := int(mouse.X) - int(clickStart.X) + int(windowPos.X)
				newY := int(mouse.Y) - int(clickStart.Y) + int(windowPos.Y)
				rl.SetWindowPosition(newX, newY)
			}
		}

		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			moveDist := rl.Vector2Length(rl.Vector2Subtract(mouse, clickStart))
			if !isDragging && moveDist < clickThreshold {
				neko.ToggleIdleSleep()
			}
			isDragging = false
		}

		//fall logic(?)
		winPos := rl.GetWindowPosition()
		monitorHeight := rl.GetMonitorHeight(0)
		isFalling = false
		if winPos.Y+screenHeight < float32(monitorHeight) {
			rl.SetWindowPosition(int(winPos.X), int(winPos.Y+fallSpeed))
			isFalling = true
		}

		//fall vs drag anim
		if isDragging {
			neko.SetState(NekoDragging)
		} else if isFalling {
			neko.SetState(NekoFalling)
		} else {
			//idle when neitherr
			if neko.GetState() == NekoFalling {
				neko.SetState(NekoIdle)
			}
		}

		//pomo
		if rl.IsMouseButtonPressed(rl.MouseRightButton) && !IsPomoActive() {
			StartPomo()
		}
		UpdatePomo()

		rl.BeginDrawing()
		//curses to opengl curses to last a thousand lifetimes
		rl.ClearBackground(rl.NewColor(0, 0, 0, 0))

		neko.DrawNeko(screenWidth, screenHeight)
		DrawPomoOverlay()

		rl.EndDrawing()
	}
}
