"""
The sum of the squares of the first ten natural numbers is, (1)2+(2)2+...+(10)2 = 385

The square of the sum of the first ten natural numbers is, (1+2+...+10)2 = (55)2 = 3025

Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is .
3025 -385 = 2640
Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
Answer:  25164150
Completed on Mon, 15 Aug 2022, 15:42
"""

def sum_square(number):
    result = 0
    for num in range(1, number+1, 1):
        result += num**2
    return result


def square_sum(number):
    result = 0
    for num in range(1, number+1, 1):
        result += num
    return result**2

def main():
    print(square_sum(100) - sum_square(100))

if __name__ == '__main__':
    main()