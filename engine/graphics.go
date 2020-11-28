package engine

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	ScreenWidth  = 500
	ScreenHeight = 700
)

type RenderLayer int

const (
	RenderLayerBackground       RenderLayer = iota
	RenderLayerUI               RenderLayer = iota
	RenderLayerForeground       RenderLayer = iota
	RenderLayerForegroundObject RenderLayer = iota
)

func LoadImageFromFile(filename string) (*ebiten.Image, error) {
	imgReader, err := ebitenutil.OpenFile(filename)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(imgReader)
	if err != nil {
		return nil, err
	}
	ebitenImage, err := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	return ebitenImage, nil
}

type SpriteInterface interface {
	Draw(screen *ebiten.Image, x, y, w, h, angle float64) error
	GetSize() (float64, float64)
}

type BasicSprite struct {
	image *ebiten.Image
	w, h  float64
}

func NewBasicSpriteFromPath(filename string) (SpriteInterface, error) {
	s := BasicSprite{}

	lImage, err := LoadImageFromFile(filename)
	if err != nil {
		return nil, err
	}
	s.image = lImage

	wInt, hInt := s.image.Size()
	s.w = float64(wInt)
	s.h = float64(hInt)

	return &s, nil
}

func (s *BasicSprite) Draw(screen *ebiten.Image, x, y, w, h, angle float64) error {
	drawOptions := ebiten.DrawImageOptions{}
	drawOptions.GeoM.Reset()
	drawOptions.GeoM.Scale(w/s.w, h/s.h)
	drawOptions.GeoM.Rotate(2 * math.Pi * angle)
	drawOptions.GeoM.Translate(x, y)
	return screen.DrawImage(s.image, &drawOptions)
}

func (s *BasicSprite) GetSize() (float64, float64) {
	return s.w, s.h
}

// generic component that's drawable
const ComponentTypeDrawable = "component-drawable"

type ComponentDrawable struct {
	Component
	sprite      SpriteInterface
	renderLayer RenderLayer
	active      bool
}

type ComponentDrawableInterface interface {
	ComponentInterface
	Draw(screen *ebiten.Image, renderLayer RenderLayer) error
	GetRenderLayer() RenderLayer

	GetActive() bool
	SetActive(active bool)
}

func NewComponentDrawable(parent ActorInterface, sprite SpriteInterface, renderLayer RenderLayer) (ComponentDrawableInterface, error) {
	component := ComponentDrawable{
		Component: Component{
			parent, ComponentTypeDrawable,
		},
		sprite:      sprite,
		renderLayer: renderLayer,
		active:      true,
	}
	return &component, nil
}

func (c *ComponentDrawable) Draw(screen *ebiten.Image, renderLayer RenderLayer) error {
	if !c.active {
		return nil
	}
	if c.renderLayer != renderLayer {
		return nil
	}
	wComp, err := c.parentActor.GetComponent(ComponentTypeWorldly)
	if err != nil {
		return err
	}
	w := wComp.(*ComponentWorldly)
	return c.sprite.Draw(screen, w.x, w.y, w.w, w.h, w.angle)
}

func (c *ComponentDrawable) GetRenderLayer() RenderLayer {
	return c.renderLayer
}

func (c *ComponentDrawable) GetActive() bool {
	return c.active
}

func (c *ComponentDrawable) SetActive(active bool) {
	c.active = active
}

const ActorTypeBackgroundImage = "actor-background-image"

func NewActorBackgroundImage(parentScene SceneInterface, id, filename string) (ActorInterface, error) {
	actor := Actor{
		parentScene: parentScene,
		actorType:   ActorTypeBackgroundImage,
		id:          id,
		components:  make([]ComponentInterface, 0),
	}

	sprite, err := NewBasicSpriteFromPath(filename)
	if err != nil {
		return nil, err
	}
	spriteComp, err := NewComponentDrawable(&actor, sprite, RenderLayerBackground)
	if err != nil {
		return nil, err
	}
	actor.components = append(actor.components, spriteComp)

	worldly, err := NewComponentWorldly(&actor, 0, 0, ScreenWidth, ScreenHeight, 0)
	if err != nil {
		return nil, err
	}
	actor.components = append(actor.components, worldly)

	return &actor, nil
}
