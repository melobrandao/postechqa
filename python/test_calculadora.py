import pytest

from calculadora import Calculadora


@pytest.fixture
def calculadora() -> Calculadora:
    return Calculadora()


@pytest.mark.parametrize(
    ("x", "y", "esperado"),
    [(2, 3, 5), (-2, -3, -5), (5, -3, 2), (0, 0, 0), (0.1, 0.2, 0.3)],
)
def test_somar(calculadora, x, y, esperado):
    assert calculadora.somar(x, y) == pytest.approx(esperado)


@pytest.mark.parametrize(
    ("x", "y", "esperado"),
    [(5, 3, 2), (-5, -3, -2), (3, 5, -2), (5, -3, 8), (5.5, 2.5, 3)],
)
def test_subtrair(calculadora, x, y, esperado):
    assert calculadora.subtrair(x, y) == pytest.approx(esperado)


@pytest.mark.parametrize(
    ("x", "y", "esperado"),
    [(4, 2.5, 10), (-4, -2, 8), (4, -2, -8), (8, 0, 0), (1.5, 2, 3)],
)
def test_multiplicar(calculadora, x, y, esperado):
    assert calculadora.multiplicar(x, y) == pytest.approx(esperado)


@pytest.mark.parametrize(
    ("x", "y", "esperado"),
    [(10, 2, 5), (1, 4, 0.25), (-9, 2, -4.5), (9, -2, -4.5), (-9, -3, 3), (0, 5, 0)],
)
def test_dividir(calculadora, x, y, esperado):
    assert calculadora.dividir(x, y) == pytest.approx(esperado)


def test_dividir_por_zero(calculadora):
    with pytest.raises(ZeroDivisionError, match="divisão por zero"):
        calculadora.dividir(10, 0)


@pytest.mark.parametrize(("numero", "esperado"), [(0, 1), (1, 1), (3, 6), (5, 120), (10, 3_628_800)])
def test_fatorial(calculadora, numero, esperado):
    assert calculadora.fatorial(numero) == esperado


def test_fatorial_negativo(calculadora):
    with pytest.raises(ValueError, match="número negativo"):
        calculadora.fatorial(-1)


@pytest.mark.parametrize("entrada", [1.5, "5", True])
def test_fatorial_rejeita_valores_nao_inteiros(calculadora, entrada):
    with pytest.raises(TypeError, match="números inteiros"):
        calculadora.fatorial(entrada)


@pytest.mark.parametrize(
    ("base", "expoente", "esperado"),
    [(2, 3, 8), (2, -3, 0.125), (-2, -3, -0.125), (7, 0, 1), (0, 0, 1), (0, 3, 0), (-2, 4, 16), (-2, 3, -8), (1.5, 2, 2.25)],
)
def test_potencia(calculadora, base, expoente, esperado):
    assert calculadora.potencia(base, expoente) == pytest.approx(esperado)


def test_zero_com_expoente_negativo(calculadora):
    with pytest.raises(ZeroDivisionError, match="expoente negativo"):
        calculadora.potencia(0, -1)
