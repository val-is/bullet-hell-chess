package engine

import (
	"os"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	sceneManager SceneMachineInterface
}

func NewGameInstance() (ebiten.Game, error) {
	sceneMachine, err := NewSceneMachine()
	if err != nil {
		return nil, err
	}

	sceneMachine.AddScene(StartSceneId, NewMainScene)
	if err := sceneMachine.RunScene(StartSceneId); err != nil {
		return nil, err
	}

	g := Game{
		sceneManager: sceneMachine,
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Bullet Hell Chess")

	return &g, nil
}

func (g *Game) Update(screen *ebiten.Image) error {
	if g.sceneManager.GetCurrentScene().GetId() == StopSceneId {
		os.Exit(0)
	}
	return g.sceneManager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	if err := g.sceneManager.Draw(screen); err != nil {
		panic(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
