package engine

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

type SceneGenerator func() (SceneInterface, error)

const (
	StartSceneId = "scene-start"
	StopSceneId  = "scene-stop"
)

type SceneMachine struct {
	activeSceneId string
	activeScene   SceneInterface
	scenes        map[string]SceneGenerator
}

type SceneMachineInterface interface {
	RunScene(sceneId string) error
	Update() error
	Draw(screen *ebiten.Image) error
	GetCurrentScene() SceneInterface
	AddScene(sceneId string, generator SceneGenerator)
}

func NewSceneMachine() (SceneMachineInterface, error) {
	return &SceneMachine{
		activeSceneId: StartSceneId,
		activeScene:   nil,
		scenes:        make(map[string]SceneGenerator),
	}, nil
}

func (s *SceneMachine) RunScene(sceneId string) error {
	s.activeSceneId = sceneId
	scene, err := s.scenes[sceneId]()
	if err != nil {
		return err
	}
	s.activeScene = scene
	return nil
}

func (s *SceneMachine) Update() error {
	return s.activeScene.Update()
}

func (s *SceneMachine) Draw(screen *ebiten.Image) error {
	for _, layer := range []RenderLayer{
		RenderLayerBackground, RenderLayerForeground,
		RenderLayerForegroundObject, RenderLayerUI} {
		if err := s.activeScene.Draw(screen, layer); err != nil {
			return err
		}
	}
	return nil
}

func (s *SceneMachine) GetCurrentScene() SceneInterface {
	return s.activeScene
}

func (s *SceneMachine) AddScene(sceneId string, generator SceneGenerator) {
	s.scenes[sceneId] = generator
}

type Scene struct {
	id     string
	actors []ActorInterface
}

type SceneInterface interface {
	Update() error
	Draw(screen *ebiten.Image, renderLayer RenderLayer) error
	GetActorsType(actorType string) []ActorInterface
	GetActorId(actorId string) (ActorInterface, error)
	AddActor(actor ActorInterface)
	GetId() string
}

func NewScene() (SceneInterface, error) {
	s := Scene{
		actors: make([]ActorInterface, 0),
	}

	return &s, nil
}

func (s *Scene) Update() error {
	for k := range s.actors {
		if err := s.actors[k].Update(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Scene) Draw(screen *ebiten.Image, renderLayer RenderLayer) error {
	for k := range s.actors {
		if err := s.actors[k].Draw(screen, renderLayer); err != nil {
			return err
		}
	}
	return nil
}

func (s *Scene) GetActorsType(actorType string) []ActorInterface {
	actorsFound := make([]ActorInterface, 0)
	for k := range s.actors {
		if s.actors[k].GetActorType() == actorType {
			actorsFound = append(actorsFound, s.actors[k])
		}
	}
	return actorsFound
}

func (s *Scene) GetActorId(actorId string) (ActorInterface, error) {
	for k := range s.actors {
		if s.actors[k].GetId() == actorId {
			return s.actors[k], nil
		}
	}
	return nil, fmt.Errorf("actor %s not found in scene %s", actorId, s.id)
}

func (s *Scene) AddActor(actor ActorInterface) {
	s.actors = append(s.actors, actor)
}

func (s *Scene) GetId() string {
	return s.id
}
