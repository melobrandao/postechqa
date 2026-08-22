package calculadora

import (
	"math"
	"testing"
)

type casoOperacao struct {
	nome     string
	x, y     float64
	esperado float64
}

func executarCasos(t *testing.T, casos []casoOperacao, fn func(float64, float64) float64) {
	t.Helper()
	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			if obtido := fn(c.x, c.y); !quaseIguais(obtido, c.esperado) {
				t.Errorf("resultado de (%v, %v) = %v; esperado %v", c.x, c.y, obtido, c.esperado)
			}
		})
	}
}

func quaseIguais(a, b float64) bool {
	const tolerancia = 1e-9
	return math.Abs(a-b) <= tolerancia
}

func deveCausarPanic(t *testing.T, mensagemEsperada string, fn func()) {
	t.Helper()
	defer func() {
		valor := recover()
		if valor == nil {
			t.Fatal("era esperado um panic, mas nenhum ocorreu")
		}
		if mensagem, ok := valor.(string); !ok || mensagem != mensagemEsperada {
			t.Errorf("panic = %q; esperado %q", mensagem, mensagemEsperada)
		}
	}()
	fn()
}

func TestSomar(t *testing.T) {
	executarCasos(t, []casoOperacao{
		{"positivos", 2, 3, 5},
		{"negativos", -2, -3, -5},
		{"positivo e negativo", 5, -3, 2},
		{"zero a direita", 8, 0, 8},
		{"zeros", 0, 0, 0},
		{"decimais", 0.1, 0.2, 0.3},
	}, Soma)
}

func TestSubtrair(t *testing.T) {
	executarCasos(t, []casoOperacao{
		{"positivos", 5, 3, 2},
		{"negativos", -5, -3, -2},
		{"resultado negativo", 3, 5, -2},
		{"subtrair negativo", 5, -3, 8},
		{"com zero", 8, 0, 8},
		{"decimais", 5.5, 2.5, 3},
	}, Subtracao)
}

func TestMultiplicar(t *testing.T) {
	executarCasos(t, []casoOperacao{
		{"positivos", 4, 2.5, 10},
		{"negativos", -4, -2, 8},
		{"positivo e negativo", 4, -2, -8},
		{"por zero", 8, 0, 0},
		{"decimais", 1.5, 2, 3},
	}, Multiplicacao)
}

func TestDividir(t *testing.T) {
	executarCasos(t, []casoOperacao{
		{"resultado inteiro", 10, 2, 5},
		{"resultado decimal", 1, 4, 0.25},
		{"dividendo negativo", -9, 2, -4.5},
		{"divisor negativo", 9, -2, -4.5},
		{"dois negativos", -9, -3, 3},
		{"zero dividido por numero", 0, 5, 0},
	}, Divisao)
}

func TestDividirPorZero(t *testing.T) {
	for _, divisor := range []float64{0, math.Copysign(0, -1)} {
		deveCausarPanic(t, "Resultado Indefinido, divisão por zero não é permitida", func() {
			Divisao(10, divisor)
		})
	}
}

func TestFatorial(t *testing.T) {
	casos := []struct {
		nome     string
		entrada  int
		esperado int
	}{
		{"limite zero", 0, 1},
		{"um", 1, 1},
		{"valor pequeno", 3, 6},
		{"valor maior", 5, 120},
		{"limite representativo", 10, 3628800},
	}

	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			if obtido := Fatorial(c.entrada); obtido != c.esperado {
				t.Errorf("Fatorial(%d) = %d; esperado %d", c.entrada, obtido, c.esperado)
			}
		})
	}
}

func TestFatorialNegativo(t *testing.T) {
	deveCausarPanic(t, "Não é possível calcular fatorial para números negativos", func() {
		Fatorial(-1)
	})
}

func TestPotencia(t *testing.T) {
	casos := []struct {
		nome     string
		base     float64
		expoente int
		esperado float64
	}{
		{"expoente positivo", 2, 3, 8},
		{"expoente negativo", 2, -3, 0.125},
		{"base negativa e expoente negativo", -2, -3, -0.125},
		{"expoente zero", 7, 0, 1},
		{"base zero e expoente zero", 0, 0, 1},
		{"base zero", 0, 3, 0},
		{"base negativa e expoente par", -2, 4, 16},
		{"base negativa e expoente impar", -2, 3, -8},
		{"base decimal", 1.5, 2, 2.25},
	}

	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			if obtido := Potencia(c.base, c.expoente); !quaseIguais(obtido, c.esperado) {
				t.Errorf("Potencia(%v, %d) = %v; esperado %v", c.base, c.expoente, obtido, c.esperado)
			}
		})
	}
}
