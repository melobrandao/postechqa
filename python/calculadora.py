class Calculadora:
    """Calculadora de operações matemáticas básicas."""

    def somar(self, x: float, y: float) -> float:
        return x + y

    def subtrair(self, x: float, y: float) -> float:
        return x - y

    def multiplicar(self, x: float, y: float) -> float:
        return x * y

    def dividir(self, x: float, y: float) -> float:
        if y == 0:
            raise ZeroDivisionError("divisão por zero não é permitida")
        return x / y

    def fatorial(self, numero: int) -> int:
        if not isinstance(numero, int) or isinstance(numero, bool):
            raise TypeError("o fatorial aceita apenas números inteiros")
        if numero < 0:
            raise ValueError("não é possível calcular o fatorial de um número negativo")

        resultado = 1
        for fator in range(2, numero + 1):
            resultado *= fator
        return resultado

    def potencia(self, base: float, expoente: int) -> float:
        if base == 0 and expoente < 0:
            raise ZeroDivisionError("zero não pode ser elevado a um expoente negativo")
        return base**expoente
