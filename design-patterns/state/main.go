package main

func main() {

	ctx := NewContext(&RequestState{})
	handler := &Handler{Name: "Aditya", Val: 120}
	ctx.Request(*handler)
	ctx.Start(*handler)
	ctx.Process(*handler)
	ctx.Close(*handler)
	ctx.Response(*handler)
}
