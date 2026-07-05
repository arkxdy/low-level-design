package main

import "fmt"

// IState interface definition
type IState interface {
	Request(ctx *Context, handler Handler)
	Start(ctx *Context, handler Handler)
	Process(ctx *Context, handler Handler)
	Close(ctx *Context, handler Handler)
	Response(ctx *Context, handler Handler)
}

// --- StartState ---
type StartState struct{}

func (s *StartState) Request(ctx *Context, handler Handler) {
	fmt.Println("Already Started: Request ignored.")
}
func (s *StartState) Start(ctx *Context, handler Handler) { fmt.Println("System is already starting.") }
func (s *StartState) Process(ctx *Context, handler Handler) {
	fmt.Println("Processing started...")
	ctx.SetState(&ProcessState{})
}
func (s *StartState) Close(ctx *Context, handler Handler) {
	fmt.Println("Closing from Start.")
	ctx.SetState(&CloseState{})
}
func (s *StartState) Response(ctx *Context, handler Handler) {
	fmt.Println("Cannot respond before process.")
}

// --- ProcessState ---
type ProcessState struct{}

func (p *ProcessState) Request(ctx *Context, handler Handler) {
	fmt.Println("Processing... Request ignored.")
}
func (p *ProcessState) Start(ctx *Context, handler Handler)   { fmt.Println("Already Processing.") }
func (p *ProcessState) Process(ctx *Context, handler Handler) { fmt.Println("Still processing...") }
func (p *ProcessState) Close(ctx *Context, handler Handler) {
	fmt.Println("Closing from Process.")
	ctx.SetState(&CloseState{})
}
func (p *ProcessState) Response(ctx *Context, handler Handler) {
	fmt.Println("Generating Response...")
	ctx.SetState(&ResponseState{})
}

// --- CloseState ---
type CloseState struct{}

func (c *CloseState) Request(ctx *Context, handler Handler) {
	fmt.Println("System Closed: Request ignored.")
}
func (c *CloseState) Start(ctx *Context, handler Handler) {
	fmt.Println("System Closed: Cannot start.")
}
func (c *CloseState) Process(ctx *Context, handler Handler) {
	fmt.Println("System Closed: Cannot process.")
}
func (c *CloseState) Close(ctx *Context, handler Handler) { fmt.Println("System already closed.") }
func (c *CloseState) Response(ctx *Context, handler Handler) {
	fmt.Println("System Closed: No response.")
}

// --- ResponseState ---
type ResponseState struct{}

func (r *ResponseState) Request(ctx *Context, handler Handler) {
	fmt.Println("Response provided. Resetting/Closing.")
}
func (r *ResponseState) Start(ctx *Context, handler Handler) {
	fmt.Println("Response provided. Resetting/Closing.")
}
func (r *ResponseState) Process(ctx *Context, handler Handler) {
	fmt.Println("Response provided. Resetting/Closing.")
}
func (r *ResponseState) Close(ctx *Context, handler Handler) {
	fmt.Println("Closing after response.")
	ctx.SetState(&CloseState{})
}
func (r *ResponseState) Response(ctx *Context, handler Handler) {
	fmt.Println("Response already sent.")
}

// --- RequestState ---
type RequestState struct{}

func (r *RequestState) Request(ctx *Context, handler Handler) {
	fmt.Println("Request Received.")
	ctx.SetState(&StartState{})
}
func (r *RequestState) Start(ctx *Context, handler Handler)    { fmt.Println("Must Request first.") }
func (r *RequestState) Process(ctx *Context, handler Handler)  { fmt.Println("Must Request first.") }
func (r *RequestState) Close(ctx *Context, handler Handler)    { fmt.Println("System Idle: Closing.") }
func (r *RequestState) Response(ctx *Context, handler Handler) { fmt.Println("Must Request first.") }

type Handler struct {
	Name string
	Val  int
}
