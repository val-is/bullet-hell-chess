package engine

func NewMainScene() (SceneInterface, error) {
	baseScene, err := NewScene()
	if err != nil {
		return nil, err
	}

	bgActor, err := NewActorBackgroundImage(baseScene, "scene-background", "assets/sprites/chessboard/chess_green/bg.png")
	if err != nil {
		return nil, err
	}
	baseScene.AddActor(bgActor)

	return baseScene, nil
}
