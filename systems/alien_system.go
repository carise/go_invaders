package systems

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type AlienSystem struct {
	entities []Alien
	world    *ecs.World
}

type Alien struct {
	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent
}

func (a *AlienSystem) Add(alien Alien) {
	a.entities = append(a.entities, alien)
}

func (*AlienSystem) Remove(ecs.BasicEntity) {}

func (*AlienSystem) Update(dt float32) {}

func (a *AlienSystem) New(w *ecs.World) {
	a.world = w
	log.Println("new AlienSystem added to scene")

	engo.Mailbox.Listen("CollisionMessage", func(message engo.Message) {
		log.Println("alien collision")

		collision, isCollision := message.(common.CollisionMessage)
		if isCollision {
			for _, e := range a.entities {
				if e.ID() == collision.Entity.BasicEntity.ID() {
					log.Println(e.ID())
				}
			}
		}
	})
}

func (a *AlienSystem) AddAliens(numAliens int) {
	currentX := float32(60.0)
	currentY := float32(100.0)

	alien1Texture, err := common.LoadedSprite("sprites/small_alien1.png")
	log.Println("Load alien1 texture")
	if err != nil {
		log.Println("Unable to load alien1 texture", err)
	}

	for i := 0; i < numAliens; i++ {
		alien := Alien{BasicEntity: ecs.NewBasic()}
		if i > 0 {
			if i%12 == 0 {
				currentX = float32(60.0)
				currentY = currentY + 40
			} else {
				currentX += 40
			}
		}
		alien.SpaceComponent = common.SpaceComponent{
			Position: engo.Point{X: currentX, Y: currentY},
			Width:    30,
			Height:   22,
		}
		alien.RenderComponent = common.RenderComponent{
			Drawable: alien1Texture,
			Scale:    engo.Point{X: 1, Y: 1},
		}

		for _, system := range a.world.Systems() {
			switch s := system.(type) {
			case *common.RenderSystem:
				s.Add(&alien.BasicEntity, &alien.RenderComponent, &alien.SpaceComponent)
			case *AlienSystem:
				s.Add(alien)
			}
		}
	}
}
