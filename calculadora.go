package calculadora

func Soma(x, y float64) float64 {
	return x + y
}

func Subtracao(x, y float64) float64 {
	return x - y
}

func Multiplicacao(x, y float64) float64 {
	return x * y
}

func Divisao(x, y float64) float64 {
	if y == 0 {
		panic("Resultado Indefinido, divisão por zero não é permitida")
	}
	return x / y
}

func Fatorial(x int) int {
	if x < 0 {
		panic("Não é possível calcular fatorial para números negativos")
	}
	if x == 0 || x == 1 {
		return 1
	}
	return x * Fatorial(x-1)
}

func Potencia(base float64, expoente int) float64 {
	if expoente == 0 {
		return 1
	}
	if expoente < 0 {
		return 1 / Potencia(base, -expoente)
	}
	result := base
	for i := 1; i < expoente; i++ {
		result *= base
	}
	return result
}
