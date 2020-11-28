package engine

const (
	PieceSpriteWidth  = 18.0
	PieceSpriteHeight = 18.0

	PieceWidth  = PieceSpriteWidth * BoardConversionFactor
	PieceHeight = PieceSpriteHeight * BoardConversionFactor
)

type BoardSide string

const (
	BoardSideBlack BoardSide = "black"
	BoardSideWhite BoardSide = "white"
)

type ChessPiece string

const (
	PiecePawn   ChessPiece = "pawn"
	PieceRook   ChessPiece = "rook"
	PieceKnight ChessPiece = "knight"
	PieceBishop ChessPiece = "bishop"
	PieceQueen  ChessPiece = "queen"
	PieceKing   ChessPiece = "king"
)

type BoardSquare [2]int

const ComponentTypeChessPiece = "component-chess-piece"

type ComponentChessPiece struct {
	Component
	color     BoardSide
	pieceType ChessPiece
	position  BoardSquare
}

type ComponentChessPieceInterface interface {
	ComponentInterface
	GetColor() BoardSide
	GetPieceType() ChessPiece
	GetPosition() BoardSquare

	SetPosition(square BoardSquare) bool
	GetAvailableMoves() []BoardSquare

	LockToGrid() error
}

func NewComponentChessPiece(parent ActorInterface, color BoardSide,
	pieceType ChessPiece, position BoardSquare) (ComponentChessPieceInterface, error) {

	component := ComponentChessPiece{
		Component: Component{
			parent, ComponentTypeChessPiece,
		},
		color:     color,
		pieceType: pieceType,
		position:  position,
	}

	return &component, nil
}

func (c *ComponentChessPiece) GetColor() BoardSide {
	return c.color
}

func (c *ComponentChessPiece) GetPieceType() ChessPiece {
	return c.pieceType
}

func (c *ComponentChessPiece) GetPosition() BoardSquare {
	return c.position
}

func (c *ComponentChessPiece) SetPosition(square BoardSquare) bool {
	// TODO replace with check for valid move
	validSpaces := c.GetAvailableMoves()
	movePresent := false
	for _, space := range validSpaces {
		if space[0] == square[0] && space[1] == square[1] {
			movePresent = true
			break
		}
	}
	if !movePresent {
		return false
	}
	// TODO handle piece-piece interactions
	c.position = square
	return true
}

func (c *ComponentChessPiece) GetAvailableMoves() []BoardSquare {
	// i hate this _so_ inexplicably much but i guess it works
	availMoves := make([]BoardSquare, 0)
	pos := c.position
	switch c.pieceType {
	case PiecePawn:
		if c.color == BoardSideWhite {
			availMoves = append(availMoves,
				BoardSquare{pos[0], pos[1] + 1},
				BoardSquare{pos[0], pos[1] + 2},
				BoardSquare{pos[0] - 1, pos[1] + 1},
				BoardSquare{pos[0] + 1, pos[1] + 1},
			)
		} else if c.color == BoardSideBlack {
			availMoves = append(availMoves,
				BoardSquare{pos[0], pos[1] - 1},
				BoardSquare{pos[0], pos[1] - 2},
				BoardSquare{pos[0] - 1, pos[1] - 1},
				BoardSquare{pos[0] + 1, pos[1] - 1},
			)
		}
	case PieceRook:
		for i := 0; i < 8; i++ {
			// move along row
			availMoves = append(availMoves, BoardSquare{i, pos[1]})
			// move along col
			availMoves = append(availMoves, BoardSquare{pos[0], i})
		}
	case PieceBishop:
		for i := 0; i < 8; i++ {
			// towards top right
			availMoves = append(availMoves, BoardSquare{pos[0] + i, pos[1] + i})
			// towards top left
			availMoves = append(availMoves, BoardSquare{pos[0] - i, pos[1] + i})
			// towards bottom left
			availMoves = append(availMoves, BoardSquare{pos[0] - i, pos[1] - i})
			// towards bottom right
			availMoves = append(availMoves, BoardSquare{pos[0] + i, pos[1] - i})
		}
	case PieceKnight:
		// adding counter-clockwise from far right (q1 -> q4)
		availMoves = append(availMoves,
			BoardSquare{pos[0] + 2, pos[1] + 1},
			BoardSquare{pos[0] + 1, pos[1] + 2},
			BoardSquare{pos[0] - 1, pos[1] + 2},
			BoardSquare{pos[0] - 2, pos[1] + 1},
			BoardSquare{pos[0] - 2, pos[1] - 1},
			BoardSquare{pos[0] - 1, pos[1] - 2},
			BoardSquare{pos[0] + 1, pos[1] - 2},
			BoardSquare{pos[0] + 2, pos[1] - 1},
		)
	case PieceKing:
		// counter-clockwise from far right
		availMoves = append(availMoves,
			BoardSquare{pos[0] + 1, pos[1] + 0},
			BoardSquare{pos[0] + 1, pos[1] + 1},
			BoardSquare{pos[0] + 0, pos[1] + 1},
			BoardSquare{pos[0] - 1, pos[1] + 1},
			BoardSquare{pos[0] - 1, pos[1] + 0},
			BoardSquare{pos[0] - 1, pos[1] - 1},
			BoardSquare{pos[0] + 0, pos[1] - 1},
			BoardSquare{pos[0] + 1, pos[1] - 1},
		)
	case PieceQueen:
		for i := 0; i < 8; i++ {
			// move along row
			availMoves = append(availMoves, BoardSquare{i, pos[1]})
			// move along col
			availMoves = append(availMoves, BoardSquare{pos[0], i})

			// towards top right
			availMoves = append(availMoves, BoardSquare{pos[0] + i, pos[1] + i})
			// towards top left
			availMoves = append(availMoves, BoardSquare{pos[0] - i, pos[1] + i})
			// towards bottom left
			availMoves = append(availMoves, BoardSquare{pos[0] - i, pos[1] - i})
			// towards bottom right
			availMoves = append(availMoves, BoardSquare{pos[0] + i, pos[1] - i})
		}
	}

	// keep only valid moves (i.e. on board)
	validMoves := make([]BoardSquare, 0)
	for _, move := range availMoves {
		if !(move[0] < 0 || move[0] > 7 ||
			move[1] < 0 || move[1] > 7) {
			validMoves = append(validMoves, move)
		}
	}
	return validMoves
}

func (c *ComponentChessPiece) LockToGrid() error {
	worldlyComp, err := c.parentActor.GetComponent(ComponentTypeWorldly)
	if err != nil {
		return err
	}
	// see board.go for math
	worldlyComp.(*ComponentWorldly).SetPosition(
		BoardPieceOffsetX+BoardCellPaddingWidth+BoardCellWidth*float64(c.position[0]),
		BoardPieceOffsetY+BoardCellPaddingHeight+BoardCellHeight*float64(c.position[1]),
	)
	return nil
}

func (c *ComponentChessPiece) Update() error {
	if err := c.LockToGrid(); err != nil {
		return err
	}
	return nil
}

const ActorTypeChessPiece = "actor-chess-piece"

func NewActorChessPiece(parentScene SceneInterface, color BoardSide, pieceType ChessPiece,
	position BoardSquare, assetDir string) (ActorInterface, error) {

	actor := Actor{
		parentScene: parentScene,
		actorType:   ActorTypeChessPiece,
		id:          NewId("piece-" + string(color) + "-" + string(pieceType)),
		components:  make([]ComponentInterface, 0),
	}

	spriteFile := string(color) + "_" + string(pieceType) + ".png"
	sprite, err := NewBasicSpriteFromPath(assetDir + "/" + spriteFile)
	if err != nil {
		return nil, err
	}
	spriteComp, err := NewComponentDrawable(&actor, sprite, RenderLayerForegroundObject)
	if err != nil {
		return nil, err
	}
	actor.components = append(actor.components, spriteComp)

	worldly, err := NewComponentWorldly(&actor, 0, 0, PieceWidth, PieceHeight, 0)
	if err != nil {
		return nil, err
	}
	actor.components = append(actor.components, worldly)

	pieceComp, err := NewComponentChessPiece(&actor, color, pieceType, position)
	if err != nil {
		return nil, err
	}
	actor.components = append(actor.components, pieceComp)

	return &actor, nil
}
