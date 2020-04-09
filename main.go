package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/home", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Handle("POST", "/create", PostRequest)
	server.Listen()
}
