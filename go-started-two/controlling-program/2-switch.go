package switcexp

type HTTPRequest struct {
	Method string
}

func switcexp() {
	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		println("Get Request")
	case "DELETE":
		println("Delete Request")
	case "POST":
		println("Post Request")
	case "PUT":
		println("Put Request")
	default:
		println("unhandled method")
	}
}
