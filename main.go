package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type sentence struct {
	Word string
	pos  map[int]rl.Vector2
}

func main() {

	rl.InitWindow(800, 800, "Matrix")
	characters := "abcdefghijklmnopqrstuvwxyz1234567890[]{}:/"
	var allWords [2][96]sentence
	rl.SetWindowState(rl.FlagWindowResizable)
	rl.SetTargetFPS(60)
	// framesCounter := 0
	// colors := []rl.Color{rl.Blue, rl.Green, rl.Red, rl.Beige, rl.Brown, rl.DarkBlue, rl.DarkPurple, rl.DarkBrown, rl.Yellow}
	for y := range allWords {
		for x := range allWords[y] {
			wordLen := rl.GetRandomValue(6, 8)
			m := make(map[int]rl.Vector2)
			for i := 0; i < int(wordLen); i++ {
				allWords[y][x].Word += string(characters[rl.GetRandomValue(0, rl.GetRandomValue(0, int32(len(characters)-1)))])
				m[i] = rl.Vector2{X: float32(x * 20), Y: float32(i * 20)}
			}
			allWords[y][x].pos = m
		}
	}
	for !rl.WindowShouldClose() {
		for y := range allWords {
			if y == 1 {
				continue
			}
			for x := range allWords[y] {
				for k, _ := range allWords[y][x].pos {
					fmt.Println(allWords[y][x].pos[k].Y)
				}
			}
		}

		rl.BeginDrawing()
		for y := range allWords {
			if y == 1 {
				continue

			}
			for x, v := range allWords[y] {
				for key, pos := range v.pos {
					rl.DrawText(string(allWords[y][x].Word[key]), int32(pos.X), int32(pos.Y), 20, rl.Green)
				}
			}
		}

		// rl.DrawFPS(0, 0)
		rl.EndDrawing()

	}

	rl.CloseWindow()
}
