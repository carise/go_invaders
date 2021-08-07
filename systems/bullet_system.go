package systems

/** There are 2 kinds of bullets (for now), the alien bullets and the turret bullets. */
import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type Bullet struct {
	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent
	common.CollisionComponent

	BulletSystem
}

type BulletSystem struct {
	entities []BulletEntity
	world    *ecs.World
}

type BulletEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
}

type BulletMessage struct {
	direction int
	x         float32
	y         float32
}

func (BulletMessage) Type() string { return "BulletMessage" }

func (b *BulletSystem) New(w *ecs.World) {
	b.world = w
	log.Println("new bullet system")

	engo.Mailbox.Listen("CollisionMessage", func(message engo.Message) {
		log.Println("bullet collision")

		collision, isCollision := message.(common.CollisionMessage)
		if isCollision {
			for _, e := range b.entities {
				if e.ID() == collision.Entity.BasicEntity.ID() {
					log.Println(e.ID())
				}
			}
		}
	})

	engo.Mailbox.Listen("BulletMessage", func(message engo.Message) {
		log.Println("bullet message")

		bulletMessage, isBullet := message.(BulletMessage)
		if isBullet {
			b.AddBullet(bulletMessage.x, bulletMessage.y, bulletMessage.direction)
		}
	})
}

func (b *BulletSystem) Add(basic *ecs.BasicEntity, space *common.SpaceComponent) {
	b.entities = append(b.entities, BulletEntity{basic, space})
}

func (*BulletSystem) Remove(ecs.BasicEntity) {}

func (b *BulletSystem) Update(dt float32) {}

func (bs *BulletSystem) AddBullet(x float32, y float32, direction int) {
	b := Bullet{BasicEntity: ecs.NewBasic()}
	b.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{X: x - 1, Y: y - 4},
		Width:    2,
		Height:   4,
	}
	bulletTexture, err := common.LoadedSprite("sprites/bullet.png")
	log.Println("Load bullet texture")
	if err != nil {
		log.Println("Unable to load bullet texture")
	}
	b.RenderComponent = common.RenderComponent{
		Drawable: bulletTexture,
		Scale:    engo.Point{X: 1, Y: 1},
	}
	for _, system := range bs.world.Systems() {
		switch s := system.(type) {
		case *common.RenderSystem:
			log.Println("Render bullet", b.Position)
			s.Add(&b.BasicEntity, &b.RenderComponent, &b.SpaceComponent)
		case *common.CollisionSystem:
			break
		}
	}
}
