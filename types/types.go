package types

import (
	"engo.io/ecs"
	"engo.io/engo"
	"engo.io/engo/common"
	"image/color"
	"math/rand"
)

type VelocityComponent struct {
	Vel engo.Point
}

func (vc *VelocityComponent) GetVelocityComponent() *VelocityComponent {
	return vc
}

type DeathComponent struct {
	DeadTime float32
}

func (dc *DeathComponent) GetDeathComponent() *DeathComponent {
	return dc
}

type GameOb struct {
	ecs.BasicEntity
	common.SpaceComponent
	common.CollisionComponent
	common.RenderComponent
	DeathComponent
}

type MovingOb struct {
	ecs.BasicEntity
	common.SpaceComponent
	VelocityComponent
	common.RenderComponent
	common.CollisionComponent
}

func NewFrog() *GameOb {
	res := GameOb{BasicEntity: ecs.NewBasic()}
	res.SpaceComponent = common.SpaceComponent{Width: 50, Height: 50}
	res.RenderComponent = common.RenderComponent{
		Drawable: common.Triangle{},
		Color:    color.Black,
	}
	res.DeathComponent = DeathComponent{}
	res.CollisionComponent = common.CollisionComponent{Solid: false, Main: true}
	res.SetZIndex(4.5)

	return &res
}

func NewCar(loc, vel engo.Point) *MovingOb {
	res := MovingOb{BasicEntity: ecs.NewBasic()}
	res.SpaceComponent = common.SpaceComponent{Position: loc, Width: 100, Height: 50}
	res.VelocityComponent = VelocityComponent{vel}
	res.RenderComponent = common.RenderComponent{
		Drawable: common.Rectangle{},
		Color:    color.RGBA{uint8(rand.Intn(255)), 0, 255, 255},
	}
	res.CollisionComponent = common.CollisionComponent{Solid: true, Main: false}
	res.SetZIndex(3.5)
	return &res
}
