package estructuras

import "fmt"

type ContenidoWeb struct {
	Multimedias []Multimedia
}

type Multimedia interface {
	Mostrar()
}

type Imagen struct {
	Titulo  string
	Formato string
	Canales string
}

func (i *Imagen) Mostrar() {
	fmt.Println(i.Titulo)
	fmt.Println(i.Formato)
	fmt.Println(i.Canales)
}

type Audio struct {
	Titulo   string
	Formato  string
	Duración int64
}

func (a *Audio) Mostrar() {
	fmt.Println(a.Titulo)
	fmt.Println(a.Formato)
	fmt.Println(a.Duración)
}

type Video struct {
	Titulo  string
	Formato string
	Frames  int64
}

func (v *Video) Mostrar() {
	fmt.Println(v.Titulo)
	fmt.Println(v.Formato)
	fmt.Println("frm:", v.Frames)
}
