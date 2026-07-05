package main

type IContext interface {
	Request(handler Handler)
	Start(handler Handler)
	Process(handler Handler)
	Close(handler Handler)
	Response(handler Handler)
	SetState(state IState)
}

type Context struct {
	state IState
}

// SetState implements [IContext].
func (c *Context) SetState(state IState) {
	c.state = state
}

// Close implements [IContext].
func (c *Context) Close(handler Handler) {
	c.state.Close(c, handler)
}

// Process implements [IContext].
func (c *Context) Process(handler Handler) {
	c.state.Process(c, handler)
}

// Request implements [IContext].
func (c *Context) Request(handler Handler) {
	c.state.Request(c, handler)
}

// Response implements [IContext].
func (c *Context) Response(handler Handler) {
	c.state.Response(c, handler)
}

// Start implements [IContext].
func (c *Context) Start(handler Handler) {
	c.state.Start(c, handler)
}

func NewContext(state IState) IContext {
	return &Context{state: state}
}
