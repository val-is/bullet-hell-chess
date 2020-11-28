package engine

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type ClickListener func() error

type MouseState int

const (
	MouseStateNone     MouseState = iota
	MouseStatePressed  MouseState = iota
	MouseStateReleased MouseState = iota
	MouseStateHeld     MouseState = iota
)

// component to deal with mouse input
// this sources a bounding box from worldly
const ComponentTypeClickable = "component-clickable"

type ComponentClickable struct {
	Component
	mouseState          MouseState
	mouseHover          bool
	mouseStateListeners map[MouseState][]ClickListener
	mouseHoverListeners []ClickListener
}

type ComponentClickableInterface interface {
	ComponentInterface

	AddStateListener(MouseState, ClickListener)
	AddHoverListener(ClickListener)

	MousePressed(x, y int) error
	MouseReleased(x, y int) error
	UpdateMousePos(x, y int) error

	CheckMouseHover(x, y int) (bool, error)

	GetMouseState() MouseState
	GetMouseHover() bool
}

func NewComponentClickable(parent ActorInterface) (ComponentClickableInterface, error) {
	return &ComponentClickable{
		Component:           Component{parent, ComponentTypeClickable},
		mouseState:          MouseStateNone,
		mouseHover:          false,
		mouseStateListeners: make(map[MouseState][]ClickListener),
		mouseHoverListeners: make([]ClickListener, 0),
	}, nil
}

func (c *ComponentClickable) AddStateListener(state MouseState, listener ClickListener) {
	c.mouseStateListeners[state] = append(c.mouseStateListeners[state], listener)
}

func (c *ComponentClickable) AddHoverListener(listener ClickListener) {
	c.mouseHoverListeners = append(c.mouseHoverListeners, listener)
}

func (c *ComponentClickable) MousePressed(x, y int) error {
	c.mouseState = MouseStatePressed
	for _, mouseListener := range c.mouseStateListeners[MouseStatePressed] {
		if err := mouseListener(); err != nil {
			return err
		}
	}
	return nil
}

func (c *ComponentClickable) MouseReleased(x, y int) error {
	c.mouseState = MouseStateReleased
	for _, mouseListener := range c.mouseStateListeners[MouseStateReleased] {
		if err := mouseListener(); err != nil {
			return err
		}
	}
	return nil
}

func (c *ComponentClickable) UpdateMousePos(x, y int) error {
	hover, err := c.CheckMouseHover(x, y)
	if err != nil {
		return err
	}
	c.mouseHover = hover
	return nil
}

func (c *ComponentClickable) CheckMouseHover(x, y int) (bool, error) {
	worldly, err := c.parentActor.GetComponent(ComponentTypeWorldly)
	if err != nil {
		return false, err
	}
	bbx, bby, bbw, bbh := worldly.(*ComponentWorldly).GetBoundingBox()
	return CheckBoundingBox(bbx, bby, bbw, bbh, float64(x), float64(y)), nil
}

func (c *ComponentClickable) GetMouseState() MouseState {
	return c.mouseState
}

func (c *ComponentClickable) GetMouseHover() bool {
	return c.mouseHover
}

func (c *ComponentClickable) Update() error {
	mx, my := ebiten.CursorPosition()
	if err := c.UpdateMousePos(mx, my); err != nil {
		return err
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && c.mouseHover {
		if err := c.MousePressed(mx, my); err != nil {
			return err
		}
	} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) ||
		c.mouseState == MouseStateHeld && !c.mouseHover {
		if err := c.MouseReleased(mx, my); err != nil {
			return err
		}
	} else if c.mouseState == MouseStatePressed {
		c.mouseState = MouseStateHeld
	} else if c.mouseState == MouseStateReleased {
		c.mouseState = MouseStateNone
	}

	if c.mouseHover {
		for _, hoverListener := range c.mouseHoverListeners {
			if err := hoverListener(); err != nil {
				return err
			}
		}
	}

	return nil
}
