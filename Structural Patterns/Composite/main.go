package main

import "fmt"

func main() {
	// Create files
	file1 := &File{name: "file1.txt", size: 100}
	file2 := &File{name: "file2.txt", size: 200}
	file3 := &File{name: "file3.txt", size: 150}

	// create directories
	root := &Directory{name: "root"}
	subDir := &Directory{name: "subDir"}

	// add files to subDir
	subDir.Add(file1)
	subDir.Add(file2)

	// add subdir to root
	root.Add(file3)
	root.Add(subDir)

	// print the structure
	fmt.Printf("Total size of %s: %d\n", root.Name(), root.Size())
	fmt.Printf("Total size of %s: %d\n", subDir.Name(), subDir.Size())
}
