package area

import "math"

const Pi = 3.1416

//Circ é responsavel pelo calculo da area da circuferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect é respo pelo calculo da area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é visivel! A fun Calcula a area do triangulo equilatero
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
