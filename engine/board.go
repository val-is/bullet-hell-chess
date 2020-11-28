package engine

const (
	BoardWidth        = 400.0
	BoardSpriteWidth  = 180.0
	BoardHeight       = 400.0
	BoardSpriteHeight = 180.0
	BoardPixelBorder  = 2.0

	// the first part is converting the border pixels to screen pixels, which we add to the original padding
	BoardConversionFactor = BoardWidth / BoardSpriteWidth
	BoardPieceOffsetX     = ((BoardWidth / BoardSpriteWidth) * BoardPixelBorder) + (ScreenWidth-BoardWidth)/2.0
	BoardPieceOffsetY     = ((BoardHeight / BoardSpriteHeight) * BoardPixelBorder) + (ScreenHeight-BoardHeight)/2.0

	BoardSpriteCellWidth  = 22.0
	BoardSpriteCellHeight = 22.0
	BoardCellWidth        = BoardSpriteCellWidth * BoardConversionFactor
	BoardCellHeight       = BoardSpriteCellHeight * BoardConversionFactor

	BoardCellPaddingWidth  = (BoardCellWidth - PieceWidth) / 2.0
	BoardCellPaddingHeight = (BoardCellHeight - PieceHeight) / 2.0

	// final piece offset is $BOARDOFFSET + $CELLOFFSET * CELLS + $BOARDCELLPADDING
)

const ActorTypeBoard = "actor-board"

func NewActorBoard(parentScene SceneInterface, id, imagePath string) (ActorInterface, error) {
	actor := Actor{
		parentScene: parentScene,
		actorType:   ActorTypeBoard,
		id:          id,
		components:  make([]ComponentInterface, 0),
	}

	sprite, err := NewBasicSpriteFromPath(imagePath)
	if err != nil {
		return nil, err
	}
	spriteComp, err := NewComponentDrawable(&actor, sprite, RenderLayerForeground)
	if err != nil {
		return nil, err
	}
	actor.components = append(actor.components, spriteComp)

	x := float64((ScreenWidth - BoardWidth) / 2)
	y := float64((ScreenHeight - BoardHeight) / 2)
	worldly, err := NewComponentWorldly(&actor, x, y, BoardWidth, BoardHeight, 0)
	if err != nil {
		return nil, err
	}
	actor.components = append(actor.components, worldly)

	return &actor, nil
}

// helpers for notation
func AlgebraicToNative(row, column int) BoardSquare {
	// algebraic is 1 indexed, bottom left -> 1,1
	return BoardSquare{row - 1, 8 - column}
}

func NativeToAlgebraic(square BoardSquare) (row, column int) {
	return square[0] + 1, 8 - square[1]
}
