package main

import (
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
	engo.Files.Load("sprites/small_alien1.png", "sprites/small_turret.png", "sprites/bullet.png")

	engo.Input.RegisterButton("shoot", engo.KeySpace)
	engo.Input.RegisterAxis("updown", engo.AxisKeyPair{Min: engo.KeyW, Max: engo.KeyS})
	engo.Input.RegisterAxis("leftright", engo.AxisKeyPair{Min: engo.KeyA, Max: engo.KeyD})
}

func (*SpaceInvaderScene) Setup(u engo.Updater) {
	world, _ := u.(*ecs.World)
	as := systems.AlienSystem{}
	cs := systems.ControlSystem{}
	world.AddSystem(&as)
	world.AddSystem(&cs)
	world.AddSystem(&systems.BulletSystem{})
	world.AddSystem(&common.RenderSystem{})

	as.AddAliens(48)
	cs.AddTurret(world)
}

func main() {
	opts := engo.RunOptions{
		Title:  "Space Invaders",
		Width:  600,
		Height: 600,
	}
	engo.Run(opts, &SpaceInvaderScene{})
}
