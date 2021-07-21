package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Turret struct {
	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent
}

type Bullet struct {
	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent
}

type ControlSystem struct {
	entities []ControlEntity
}

type ControlEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
}

func (c *ControlSystem) New(*ecs.World) {
	log.Println("new control system")
}

func (c *ControlSystem) Add(basic *ecs.BasicEntity, space *common.SpaceComponent) {
	c.entities = append(c.entities, ControlEntity{basic, space})
}

func (*ControlSystem) Remove(ecs.BasicEntity) {}

func (c *ControlSystem) Update(dt float32) {
	for _, e := range c.entities {
		// for now, require that a bullet is shot only when the space is pressed
		// eventually the bullet speed needs to be rate limited. if the space is
		// held longer than a frame, it will have some max rate of shooting.
		if engo.Input.Button("shoot").JustPressed() {
			log.Println("Shoot bullet")
		}

		// some magic number
		speed := engo.GameWidth() / 2 * dt

		vert := engo.Input.Axis("updown")
		e.SpaceComponent.Position.Y += speed * vert.Value()
		if (e.SpaceComponent.Height + e.SpaceComponent.Position.Y) > engo.GameHeight() {
			e.SpaceComponent.Position.Y = engo.GameHeight() - e.SpaceComponent.Height
		} else if e.SpaceComponent.Position.Y < 0 {
			e.SpaceComponent.Position.Y = 0
		}

		horiz := engo.Input.Axis("leftright")
		e.SpaceComponent.Position.X += speed * horiz.Value()
		if (e.SpaceComponent.Width + e.SpaceComponent.Position.X) > engo.GameWidth() {
			e.SpaceComponent.Position.X = engo.GameWidth() - e.SpaceComponent.Width
		} else if e.SpaceComponent.Position.X < 0 {
			e.SpaceComponent.Position.X = 0
		}
	}
}
