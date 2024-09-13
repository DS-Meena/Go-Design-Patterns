package main

// Component interface
type FileSystemComponent interface {
	Name() string
	Size() int
}

// Leaf
type File struct {
	name string
	size int
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Size() int {
	return f.size
}

// Composite
type Directory struct {
	name        string
	componennts []FileSystemComponent
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Size() int {
	total := 0
	for _, component := range d.componennts {
		total += component.Size()
	}

	return total
}

func (d *Directory) Add(component FileSystemComponent) {
	d.componennts = append(d.componennts, component)
}
