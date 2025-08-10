package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// CatState represents the current state of the cat
type NekoState int

const (
	NekoIdle NekoState = iota
	NekoSleep
	NekoFalling
	NekoDragging
)

// neko manages the cat sprite and its rendering
type Neko struct {
	textures     []rl.Texture2D
	currentState NekoState
	frame        int32
	frameSpeed   int
	animTimer    float32
}

func InitNeko() *Neko {
	textures := []rl.Texture2D{
		rl.LoadTexture("assets/catIdle.png"),
		rl.LoadTexture("assets/catSleep.png"),
		rl.LoadTexture("assets/catDown.png"),
		rl.LoadTexture("assets/catDrag.png"),
	}

	return &Neko{
		textures:     textures,
		currentState: NekoIdle,
		frame:        0,
		frameSpeed:   4,
		animTimer:    0,
	}
}

func (c *Neko) UnloadNeko() {
	for _, tex := range c.textures {
		rl.UnloadTexture(tex)
	}
}

// neko animation
func (c *Neko) UpdateNeko() {
	c.animTimer += rl.GetFrameTime()
	if c.animTimer >= 1.0/float32(c.frameSpeed) {
		c.animTimer -= 1.0 / float32(c.frameSpeed)
		c.frame = (c.frame + 1) % 4
	}
}

func (c *Neko) SetState(state NekoState) {
	c.currentState = state
}

func (c *Neko) GetState() NekoState {
	return c.currentState
}

// sleep/idel
func (c *Neko) ToggleIdleSleep() {
	if c.currentState == NekoIdle {
		c.currentState = NekoSleep
	} else if c.currentState == NekoSleep {
		c.currentState = NekoIdle
	}
}

func (c *Neko) DrawNeko(screenWidth, screenHeight int32) {
	//neko state => tex
	textureIndex := 0
	switch c.currentState {
	case NekoIdle:
		textureIndex = 0
	case NekoSleep:
		textureIndex = 1
	case NekoFalling:
		textureIndex = 2
	case NekoDragging:
		textureIndex = 3
	}

	drawTexture := c.textures[textureIndex]

	//draw spirte
	frameWidth := drawTexture.Width / 4
	srcRect := rl.NewRectangle(float32(c.frame*frameWidth), 0, float32(frameWidth), float32(drawTexture.Height))
	dstRect := rl.NewRectangle(
		float32((screenWidth-frameWidth)/2),
		float32((screenHeight-drawTexture.Height)/2),
		float32(frameWidth),
		float32(drawTexture.Height),
	)

	rl.DrawTexturePro(drawTexture, srcRect, dstRect, rl.NewVector2(0, 0), 0, rl.White)
}
