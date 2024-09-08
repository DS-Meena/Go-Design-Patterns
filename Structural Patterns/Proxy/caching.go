package main

func main() {
	image1 := &ProxyImage{filename: "photo1.jpg"}
	image2 := &ProxyImage{filename: "photo2.jpg"}

	// Image will be loaded from disk
	image1.Display()

	// Image will not be loaded from disk, because it's already in cache
	image1.Display()

	// new image, so load from disk
	image2.Display()
}
