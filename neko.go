package main

import (
	_ "embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type NekoState int

const (
	NekoIdle NekoState = iota
	NekoSleep
	NekoFalling
	NekoDragging
)

const (
	clickThreshold = 5.0
	fallSpeed      = 7.0
)

//go:embed assets/catIdle.png
var catIdle []byte

//go:embed assets/catSleep.png
var catSleep []byte

//go:embed assets/catDown.png
var catDown []byte

//go:embed assets/catDrag.png
var catDrag []byte

// neko struct -> sprites n rendering
type Neko struct {
	textures     []rl.Texture2D
	currentState NekoState
	frame        int32
	frameSpeed   int
	animTimer    float32

	clickStart rl.Vector2
	isDragging bool
	isFalling  bool
}

func InitNeko() *Neko {
	textures := []rl.Texture2D{
		LoadTextureFrmBytes(catIdle),
		LoadTextureFrmBytes(catSleep),
		LoadTextureFrmBytes(catDown),
		LoadTextureFrmBytes(catDrag),
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
	//neko state -> tex
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

//---cat shit---

// click/drag
func (c *Neko) ClickNDrag() {
	mouse := rl.GetMousePosition()
	windowPos := rl.GetWindowPosition()

	//e
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		c.clickStart = mouse
		c.isDragging = false
	}

	//drag to where
	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		moveDist := rl.Vector2Length(rl.Vector2Subtract(mouse, c.clickStart))
		if moveDist >= clickThreshold {
			c.isDragging = true
		}
		if c.isDragging {
			newX := int(mouse.X) - int(c.clickStart.X) + int(windowPos.X)
			newY := int(mouse.Y) - int(c.clickStart.Y) + int(windowPos.Y)
			rl.SetWindowPosition(newX, newY)
		}
	}

	//click not drag
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		moveDist := rl.Vector2Length(rl.Vector2Subtract(mouse, c.clickStart))
		if !c.isDragging && moveDist < clickThreshold {
			c.ToggleIdleSleep()
		}
		c.isDragging = false
	}
}

// fall how
func (c *Neko) HandleFall() {
	winPos := rl.GetWindowPosition()
	monitorHeight := rl.GetMonitorHeight(0)
	c.isFalling = false
	if winPos.Y+float32(screenHeight) < float32(monitorHeight) {
		rl.SetWindowPosition(int(winPos.X), int(winPos.Y+fallSpeed))
		c.isFalling = true
	}
}

// fall/drag
func (c *Neko) FallNDrag() {
	if c.isDragging {
		c.SetState(NekoDragging)
	} else if c.isFalling {
		c.SetState(NekoFalling)
	} else { //idle when neitherr
		if c.GetState() == NekoFalling {
			c.SetState(NekoIdle)
		}
	}
}
