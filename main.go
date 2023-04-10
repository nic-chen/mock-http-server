package main

func main() {
	r := setupRouter()
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
