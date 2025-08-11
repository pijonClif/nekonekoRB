# NekoNekoRB

Tiny desktop pet + Pomodoro timer.

Toy project to experiment with Golang & raylib-go.

## Features
- **Desktop pet**: idle, sleep, drag, and fall animations
- **Pomodoro overlay**: simple work/break cycle display
- **Frameless, always-on-top, transparent window**

## Run (Windows)
- Double-click `bin/NekoNekoRB.exe`.
- Keep the working directory as `bin/` so `assets/` can be found.

## Controls
- **Left click**: toggle idle/sleep
- **Drag**: hold left click and move to reposition
- **Right click**: start Pomodoro
- **Esc**: quit


Requirements: Go 1.24+, Windows.
Uses `github.com/gen2brain/raylib-go/raylib`.