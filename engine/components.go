package engine

import (
	"github.com/hajimehoshi/ebiten"
)

// BASE COMPONENTS

// basic component, nothing special
const ComponentTypeBasic = "component-basic"

type Component struct {
	parentActor   ActorInterface
	componentType string
}

type ComponentInterface interface {
	Update() error
	GetComponentType() string
}

func NewComponent(parent ActorInterface, componentType string) (ComponentInterface, error) {
	return &Component{
		parent, componentType,
	}, nil
}

func (c *Component) Update() error {
	return nil
}

func (c *Component) GetComponentType() string {
	return c.componentType
}

// component that exists in the world (i.e. has position)
const ComponentTypeWorldly = "component-worldly"

type ComponentWorldly struct {
	Component
	x, y  float64
	w, h  float64
	angle float64
}

type ComponentWorldlyInterface interface {
	ComponentInterface

	GetPosition() (x, y float64)
	SetPosition(x, y float64)
	GetScale() (w, h float64)
	SetScale(w, h float64)
	GetAngle() float64
	SetAngle(angle float64)
}

func NewComponentWorldly(parent ActorInterface, x, y, w, h, angle float64) (ComponentWorldlyInterface, error) {
	return &ComponentWorldly{
		Component{parent, ComponentTypeWorldly},
		x, y, w, h, angle,
	}, nil
}

func (c *ComponentWorldly) GetPosition() (x, y float64) {
	return c.x, c.y
}

func (c *ComponentWorldly) SetPosition(x, y float64) {
	c.x = x
	c.y = y
}

func (c *ComponentWorldly) GetScale() (w, h float64) {
	return c.x, c.y
}

func (c *ComponentWorldly) SetScale(w, h float64) {
	c.w = w
	c.h = h
}

func (c *ComponentWorldly) GetAngle() float64 {
	return c.angle
}

func (c *ComponentWorldly) SetAngle(angle float64) {
	c.angle = angle
}

// GRAPHICAL COMPONENTS

// generic component that's drawable
const ComponentTypeDrawable = "component-drawable"

type ComponentDrawable struct {
	Component
	sprite SpriteInterface
}

type ComponentDrawableInterface interface {
	ComponentInterface
	Draw(screen *ebiten.Image) error
}

func NewComponentDrawable(parent ActorInterface, sprite SpriteInterface) (ComponentDrawableInterface, error) {
	component := ComponentDrawable{
		Component: Component{
			parent, ComponentTypeDrawable,
		},
		sprite: sprite,
	}
	return &component, nil
}

func (c *ComponentDrawable) Draw(screen *ebiten.Image) error {
	wComp, err := c.parentActor.GetComponent(ComponentTypeWorldly)
	if err != nil {
		return err
	}
	w := wComp.(*ComponentWorldly)
	return c.sprite.Draw(screen, w.x, w.y, w.w, w.h, w.angle)
}
