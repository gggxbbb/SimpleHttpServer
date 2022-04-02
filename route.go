package main

type onReq func(*request, *response) []byte

var routes map[string]onReq

func requestHandler(req *request, rep *response) {
	fun, ok := routes[req.path]
	if ok {
		body := fun(req, rep)
		rep.WriteBody(body)
		rep.Send()
	} else {
		raise404(rep)
	}
}

func newRoute(path string, fun onReq) {
	routes[path] = fun
}

func initRoutes() {
	newRoute("/", index)
}
