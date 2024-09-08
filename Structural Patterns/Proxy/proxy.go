package main

// ProxyImage is the surrogate
type ProxyImage struct {
	realImage *RealImage
	filename  string
}

func (p *ProxyImage) Display() {
	// caching
	if p.realImage == nil {
		p.realImage = &RealImage{filename: p.filename}
		p.realImage.LoadFromDisk()
	}
	p.realImage.Display()
}
