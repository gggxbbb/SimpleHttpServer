package main

func raise404(rep *response) {
	rep.code = 404
	rep.status = "ERROR"
	rep.body = []byte("Error: 404 Not Found")
	rep.Send()
}
