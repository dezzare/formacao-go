package main

import "fmt"

func main() {

    const ebulicaoKelvin = 373.0

    ebulicaoCelsius := ebulicaoKelvin - 273.0
    ebulicaoFahrenheit := (ebulicaoCelsius * 9/5) + 32
    
	fmt.Printf("A temperatura de ebulição da água em Kelvin = %f\n", ebulicaoKelvin)

    fmt.Printf("A temperatura de ebulição da água em graus Celsius = %f\n", ebulicaoCelsius)


    fmt.Printf("A temperatura de ebulição da água em Fahrenheit = %f\n", ebulicaoFahrenheit)


}
