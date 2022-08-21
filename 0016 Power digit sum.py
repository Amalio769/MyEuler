"""
2pow(15) = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2pow(1000)?
Answer:  1366
Completed on Sun, 21 Aug 2022, 18:43
"""

numero = pow(2,1000)
string = str(numero)
resultado = 0
for character in string:
    resultado += int(character)
print(resultado)
