package engine

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

type Actor struct {
	parentScene SceneInterface
	actorType   string
	id          string
	components  []ComponentInterface
}

type ActorInterface interface {
	Update() error
	Draw(screen *ebiten.Image, renderLayer RenderLayer) error
	GetComponent(componentType string) (ComponentInterface, error)
	GetActorType() string
	GetId() string
}

func (a *Actor) Update() error {
	for k := range a.components {
		if err := a.components[k].Update(); err != nil {
			return err
		}
	}
	return nil
}

func (a *Actor) Draw(screen *ebiten.Image, renderLayer RenderLayer) error {
	for k := range a.components {
		if a.components[k].GetComponentType() == ComponentTypeDrawable {
			if err := a.components[k].(ComponentDrawableInterface).Draw(screen, renderLayer); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *Actor) GetComponent(componentType string) (ComponentInterface, error) {
	for k := range a.components {
		if a.components[k].GetComponentType() == componentType {
			return a.components[k], nil
		}
	}
	return nil, fmt.Errorf("component of type %s on %s not found", componentType, a.id)
}

func (a *Actor) GetActorType() string {
	return a.actorType
}

func (a *Actor) GetId() string {
	return a.id
}
