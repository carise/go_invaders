package systems

import (
	"fmt"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
	"github.com/Noofbiz/engoBox2dSystem"
)

type AlienSystem struct{}

type Alien struct {
	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent
	engoBox2dSystem.Box2dComponent
}

func (*AlienSystem) Remove(ecs.BasicEntity) {}

func (*AlienSystem) Update(dt float32) {}

func (*AlienSystem) New(*ecs.World) {
	fmt.Println("new AlienSystem added to scene")
}
