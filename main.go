package main

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Fish struct {
	position fyne.Position
	icon     *canvas.Image
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Water Filling Effect with Fish")

	// Set up the water rectangle
	water := canvas.NewRectangle(theme.PrimaryColor())
	water.FillColor = theme.PrimaryColor()
	water.Resize(fyne.NewSize(400, 0))
	water.Move(fyne.NewPos(0, 400))

	// Container to hold the water rectangle
	waterContainer := container.NewWithoutLayout(water)

	startButton := widget.NewButton("Start Filling", func() {
		go fillWater(water, waterContainer, myWindow)
	})

	content := container.NewVBox(
		startButton,
		waterContainer,
	)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.ShowAndRun()
}

func fillWater(water *canvas.Rectangle, waterContainer *fyne.Container, window fyne.Window) {
	rand.Seed(time.Now().UnixNano())
	fishList := []*Fish{}
	fishCount := 0

	for i := 0; i <= 400; i++ {
		time.Sleep(10 * time.Millisecond)
		water.Resize(fyne.NewSize(400, float32(i)))
		water.Move(fyne.NewPos(0, float32(400-i)))
		water.Refresh()

		// Add fish as water level increases
		if i%40 == 0 && fishCount < 10 { // Adjust the condition and count as needed
			fishCount++
			newFish := createFish(i)
			fishList = append(fishList, newFish)
			waterContainer.Add(newFish.icon)
		}

		// Update fish positions
		for _, fish := range fishList {
			fish.position.Y = float32(400 - i + rand.Intn(20))
			fish.icon.Move(fish.position)
			fish.icon.Refresh()
		}

		window.Canvas().Refresh(waterContainer)
	}
}

func createFish(waterLevel int) *Fish {
	x := float32(rand.Intn(380) + 10) // Random x position
	y := float32(400 - waterLevel)    // Initial y position based on water level
	fishIcon := canvas.NewImageFromFile("fish.png")
	fishIcon.FillMode = canvas.ImageFillContain
	fishIcon.Resize(fyne.NewSize(30, 30)) // Adjust the size as needed
	fish := &Fish{
		position: fyne.NewPos(x, y),
		icon:     fishIcon,
	}
	return fish
}
