package main

func main() {
	r := setupRouter()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
