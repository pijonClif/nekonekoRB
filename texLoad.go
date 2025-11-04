package main

import (
	_ "embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func LoadTextureFrmBytes(file []byte) rl.Texture2D {
	img := rl.LoadImageFromMemory(".png", file, int32(len(file)))
	texture := rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	return texture
}
