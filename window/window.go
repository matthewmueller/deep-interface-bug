package window

// EventTarget interface
type EventTarget interface {
	DispatchEvent() Event
}

// Node interface
type Node interface {
	EventTarget

	ParentElement() HTMLElement
}

// Element interface
type Element interface {
	Node
}

// HTMLElement interface
type HTMLElement interface {
	Element
}

// Event interface
type Event interface {
	SrcElement() Element
}
