package main

import (
	"fmt"

	"./estructuras"
)

func crearImagen() estructuras.Imagen {
	var aux0, aux1, aux2 string
	fmt.Println("Introduce el titulo:")
	fmt.Scan(&aux0)
	fmt.Println("Introduce el formato")
	fmt.Scan(&aux1)
	fmt.Println("Introduce el canales")
	fmt.Scan(&aux2)
	imagen := estructuras.Imagen{Titulo: aux0, Formato: aux1, Canales: aux2}
	return imagen
}

func crearAudio() estructuras.Audio {
	var aux0, aux1 string
	var aux2 int64
	fmt.Println("Introduce el titulo:")
	fmt.Scan(&aux0)
	fmt.Println("Introduce el formato")
	fmt.Scan(&aux1)
	fmt.Println("Introduce la duración")
	fmt.Scan(&aux2)
	audio := estructuras.Audio{Titulo: aux0, Formato: aux1, Duración: aux2}
	return audio
}

func crearVideo() estructuras.Video {
	var aux0, aux1 string
	var aux2 int64
	fmt.Println("Introduce el titulo:")
	fmt.Scan(&aux0)
	fmt.Println("Introduce el formato")
	fmt.Scan(&aux1)
	fmt.Println("Introduce los Frames")
	fmt.Scan(&aux2)
	video := estructuras.Video{Titulo: aux0, Formato: aux1, Frames: aux2}
	return video
}

func imprimirSlice(multi []estructuras.Multimedia) {
	for i, m := range multi {
		fmt.Println("::::::Elemento: ", i, "::::::::::")
		m.Mostrar()

	}
}
func main() {
	opc := 2

	CW := estructuras.ContenidoWeb{Multimedias: []estructuras.Multimedia{}}
	menu := "Menu \n1: crear imagen \n2: crear video \n3: crear Audio \n4: mostar contenido web \n0: salir \nopc: "

	for ok := true; ok; ok = (opc != 0) {

		fmt.Println(menu)
		fmt.Scan(&opc)
		if opc == 1 {
			imag := crearImagen()
			CW.Multimedias = append(CW.Multimedias, &imag)
		}
		if opc == 2 {
			imag := crearVideo()
			CW.Multimedias = append(CW.Multimedias, &imag)
		}
		if opc == 3 {
			imag := crearAudio()
			CW.Multimedias = append(CW.Multimedias, &imag)
		}
		if opc == 4 {
			imprimirSlice(CW.Multimedias)
		}
	}
}
