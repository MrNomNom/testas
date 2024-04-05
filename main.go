package main

func main() {
	child := newChildService()
	main := newMainService(child)
	main.Main(false)
}
