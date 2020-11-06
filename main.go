package area

import "math"

//Pi é uma proporção numéria definida po uma relação entre o perímetro e o diametro
//PI é pública porque começa com letra maiúscuila
const Pi = 3.1416

//Circ é responsável por calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi

}

//Rect cácula a áreta de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura

}

//não é visível porque tem o unerline na frente
func _TrianguloEquilatero(base, altura float64) float64 {
	return (base * altura) / 2

}
