package main

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/carise/go_invaders/systems"
)

type SpaceInvaderScene struct{}

type Battlefield struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (*SpaceInvaderScene) Type() string { return "space invaders" }

func (*SpaceInvaderScene) Preload() {
	engo.Files.Load("sprites/alien1.png", "sprites/small_turret.png")

	engo.Input.RegisterButton("shoot", engo.KeySpace)
	engo.Input.RegisterAxis("updown", engo.AxisKeyPair{Min: engo.KeyW, Max: engo.KeyS})
	engo.Input.RegisterAxis("leftright", engo.AxisKeyPair{Min: engo.KeyA, Max: engo.KeyD})
}

func (*SpaceInvaderScene) Setup(u engo.Updater) {

	world, _ := u.(*ecs.World)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&systems.ControlSystem{})
	world.AddSystem(&systems.AlienSystem{})

	turret := systems.Turret{BasicEntity: ecs.NewBasic()}
	turret.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 200, Y: 350},
		Width:    50,
		Height:   50,
	}
	turretTexture, err := common.LoadedSprite("sprites/small_turret.png")
	log.Println("Load turret texture")
	if err != nil {
		log.Println("Unable to load turret texture")
	}
	turret.RenderComponent = common.RenderComponent{
		Drawable: turretTexture,
		Scale:    engo.Point{X: 1, Y: 1},
	}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&turret.BasicEntity, &turret.RenderComponent, &turret.SpaceComponent)
		case *systems.ControlSystem:
			sys.Add(&turret.BasicEntity, &turret.SpaceComponent)
		}
	}
}

func main() {
	opts := engo.RunOptions{
		Title:  "Space Invaders",
		Width:  600,
		Height: 600,
	}
	engo.Run(opts, &SpaceInvaderScene{})
}
