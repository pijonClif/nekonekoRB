package main

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PomoPhase int

const (
	Work1 PomoPhase = iota
	Break1
	Work2
	Break2
	Work3
	FinalRest
	PhaseDone // not shown
)

var (
	pomoActive    = false
	pomoStartTime time.Time
	currentPhase  PomoPhase
	pomoTexture   rl.Texture2D
	screenWidth   = 128
	screenHeight  = 128
)

func InitPomo() {
	pomoTexture = rl.LoadTexture("assets/pomodoro.png")
}

func UnloadPomo() {
	rl.UnloadTexture(pomoTexture)
}

func StartPomo() {
	pomoActive = true
	currentPhase = Work1
	pomoStartTime = time.Now()
}

func UpdatePomo() {
	if !pomoActive {
		return
	}

	elapsed := time.Since(pomoStartTime)
	currentDuration := getPomoPhaseDuration(currentPhase)

	if elapsed >= currentDuration {
		if currentPhase == PhaseDone {
			pomoActive = false
		} else {
			currentPhase++
			pomoStartTime = time.Now()
		}
	}
}

func DrawPomoOverlay() {
	if !pomoActive {
		return
	}

	//draw overlay frame
	frameWidth := pomoTexture.Width / 7
	srcRect := rl.NewRectangle(float32(currentPhase)*float32(frameWidth), 0, float32(frameWidth), float32(pomoTexture.Height))
	dstRect := rl.NewRectangle(0, 0, float32(screenWidth), float32(screenHeight))

	rl.DrawTexturePro(pomoTexture, srcRect, dstRect, rl.NewVector2(0, 0), 0, rl.White)

	//remaining time @ southwest
	remaining := int(getPomoPhaseDuration(currentPhase).Seconds() - time.Since(pomoStartTime).Seconds())
	if remaining < 0 {
		remaining = 0
	}

	rl.DrawText(
		fmt.Sprintf("%dm %ds", remaining/60.0, remaining%60),
		7, int32(screenHeight-20),
		16,
		rl.White,
	)

}

func IsPomoActive() bool {
	return pomoActive
}

func getPomoPhaseDuration(phase PomoPhase) time.Duration {
	switch phase {
	case Work1, Work2, Work3:
		return 25 * time.Minute
	case Break1, Break2:
		return 5 * time.Minute
	case FinalRest:
		return 30 * time.Second
	case PhaseDone:
		return 30 * time.Second
	default:
		return 0
	}
}
