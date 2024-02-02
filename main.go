package main

// main
func main() {
	service := NewServiceImpl()
	client := NewAPIClient("127.0.0.1:3000", service)
	client.Run()
}
