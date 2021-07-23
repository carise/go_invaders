package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type AlienSystem struct {
	entities []Alien
}

type Alien struct {
	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent
}

func (a *AlienSystem) Add(alien Alien) {
	a.entities = append(a.entities, alien)
}

func (*AlienSystem) Remove(ecs.BasicEntity) {}

func (*AlienSystem) Update(dt float32) {}

func (*AlienSystem) New(*ecs.World) {
	log.Println("new AlienSystem added to scene")
}
