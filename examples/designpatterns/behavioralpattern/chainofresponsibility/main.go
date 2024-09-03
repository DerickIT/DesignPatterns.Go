package main

type Handler interface {
	HandleRequest(request string) string
	SetNextHandler(handler Handler)
}

type BaseHandler struct {
	nextHandler Handler
}

func (h *BaseHandler) SetNextHandler(handler Handler) {
	h.nextHandler = handler

}

type ConcreteHandler1 struct {
	BaseHandler
}

func (h *ConcreteHandler1) HandleRequest(request string) string {
	if request == "one" {
		return "Handled by ConcreteHandler1"
	}
	if h.nextHandler != nil {
		return h.nextHandler.HandleRequest(request)
	}
	return "Request not handled"
}

type ConcreteHandler2 struct {
	BaseHandler
}

func (h *ConcreteHandler2) HandleRequest(request string) string {
	if request == "two" {
		return "Handled by ConcreteHandler2"
	}
	if h.nextHandler != nil {
		return h.nextHandler.HandleRequest(request)
	}
	return "Request not handled"
}

func main() {
	handler1 := &ConcreteHandler1{}
	handler2 := &ConcreteHandler2{}

	handler1.SetNextHandler(handler2)

	println(handler1.HandleRequest("one"))
	println(handler1.HandleRequest("two"))
	println(handler1.HandleRequest("three"))
}
