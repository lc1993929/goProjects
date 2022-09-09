package main

func main() {
	for i := 0; i < 5; i++ {
		i := i
		defer func() {
			println(i)
		}()
	}
}
