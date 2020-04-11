package main

import (
	"github.com/manuelmunoz00/gomodules/pkginterfaz"
	"github.com/manuelmunoz00/myinterface"
)

func main() {
	//invocando desde package local
	p := pkginterfaz.Perro{}
	pkginterfaz.MoverAnimal(p)

	//invocando desde module go en repo github
	p1 := myinterface.Perro{}
	myinterface.MoverAnimal(p1)
	p2 := myinterface.Pescado{}
	myinterface.MoverAnimal(p2)

}
