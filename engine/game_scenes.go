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

	// set up standard board
	pieceSpriteDir := "assets/sprites/chessboard/chess_green/"

	addPieces := func(pieceType ChessPiece, row int, columns ...int) error {
		for _, col := range columns {
			// white
			piece, err := NewActorChessPiece(baseScene, BoardSideWhite, pieceType, AlgebraicToNative(col, row), pieceSpriteDir)
			if err != nil {
				return err
			}
			baseScene.AddActor(piece)
			// black (mirrored)
			piece, err = NewActorChessPiece(baseScene, BoardSideBlack, pieceType, AlgebraicToNative(col, 9-row), pieceSpriteDir)
			if err != nil {
				return err
			}
			baseScene.AddActor(piece)
		}
		return nil
	}

	// pawns
	if err := addPieces(PiecePawn, 2, 1, 2, 3, 4, 5, 6, 7, 8); err != nil {
		return nil, err
	}
	// rooks
	if err := addPieces(PieceRook, 1, 1, 8); err != nil {
		return nil, err
	}
	// knights
	if err := addPieces(PieceKnight, 1, 2, 7); err != nil {
		return nil, err
	}
	// bishops
	if err := addPieces(PieceBishop, 1, 3, 6); err != nil {
		return nil, err
	}
	// queens
	if err := addPieces(PieceQueen, 1, 4); err != nil {
		return nil, err
	}
	// kings
	if err := addPieces(PieceKing, 1, 5); err != nil {
		return nil, err
	}

	testBoardActor, err := NewActorBoard(baseScene, "board-actor", "assets/sprites/chessboard/chess_green/board.png")
	if err != nil {
		return nil, err
	}
	baseScene.AddActor(testBoardActor)

	return baseScene, nil
}
