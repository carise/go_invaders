package main

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type SpaceInvaderScene struct{}

type Battlefield struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

func (*SpaceInvaderScene) Type() string { return "space invaders" }

func (*SpaceInvaderScene) Preload() {
	engo.Files.Load("sprites/aliensprite.png")
}

func (*SpaceInvaderScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)
	world.AddSystem(&common.RenderSystem{})
	battlefield := Battlefield{BasicEntity: ecs.NewBasic()}
	battlefield.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 10, Y: 10},
		Width: 468,
		Height: 39,
	}

	texture, err := common.LoadedSprite("sprites/aliensprite.png")
	log.Println("Load sprites/aliensprite.png")
	if err != nil {
		log.Println("Unable to load texture: " + err.Error())
	}
	battlefield.RenderComponent = common.RenderComponent{
		Drawable: texture,
		Scale: engo.Point{X: 1, Y: 1},
	}

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&battlefield.BasicEntity, &battlefield.RenderComponent, &battlefield.SpaceComponent)
		}
	}
}

func main() {
	opts := engo.RunOptions{
		Title:  "Space Invaders",
		Width:  400,
		Height: 400,
	}
	engo.Run(opts, &SpaceInvaderScene{})
}
