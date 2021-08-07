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
	common.CollisionComponent
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
	engo.Mailbox.Listen("CollisionMessage", func(message engo.Message) {
		log.Println("turret collision")

		collision, isCollision := message.(common.CollisionMessage)
		if isCollision {
			for _, e := range c.entities {
				if e.ID() == collision.Entity.BasicEntity.ID() {
					log.Println(e.ID())
				}
			}
		}
	})
}

func (c *ControlSystem) Add(basic *ecs.BasicEntity, space *common.SpaceComponent) {
	c.entities = append(c.entities, ControlEntity{basic, space})
}

func (*ControlSystem) Remove(ecs.BasicEntity) {
	log.Println("removed entity")
}

func (c *ControlSystem) Update(dt float32) {
	for _, e := range c.entities {
		// for now, require that a bullet is shot only when the space is pressed
		// eventually the bullet speed needs to be rate limited. if the space is
		// held longer than a frame, it will have some max rate of shooting.
		if engo.Input.Button("shoot").JustPressed() {
			bX := e.SpaceComponent.Position.X + e.SpaceComponent.Width/2
			bY := e.SpaceComponent.Position.Y
			log.Println("Shoot bullet from (x, y)", bX, ",", bY)
			engo.Mailbox.Dispatch(BulletMessage{direction: -1, x: bX, y: bY})
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

func (*ControlSystem) AddTurret(world *ecs.World) {
	turret := Turret{BasicEntity: ecs.NewBasic()}
	turret.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: 200, Y: 350},
		Width:    30,
		Height:   19,
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
		switch s := system.(type) {
		case *common.RenderSystem:
			s.Add(&turret.BasicEntity, &turret.RenderComponent, &turret.SpaceComponent)
		case *ControlSystem:
			s.Add(&turret.BasicEntity, &turret.SpaceComponent)
		}
	}
}
