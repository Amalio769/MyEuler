"""
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
Answer:  232792560
Completed on Mon, 15 Aug 2022, 07:37
"""
import math


def is_prime_number(number):
    for i in range(2, number):
        if number % i == 0:
            return False
    return True


def prime_numbers(number):
    prime_number_list = [1, ]
    for i in range(2, number+1):
        if is_prime_number(i):
            prime_number_list.append(i)

    return prime_number_list


def higest_power_less_than_number(number, less_than_number):
    i=1
    list_power = []
    if number == 1:
        return 1
    while True:
        if pow(number, i) < less_than_number:
            list_power.append(i)
            i += 1
        else:
            break

    return pow(number, max(list_power))


def smallest_multiple(number):
    list_prime_numbers = prime_numbers(number)
    list_prime_power_numbers = []
    for num in list_prime_numbers:
        list_prime_power_numbers.append(higest_power_less_than_number(num, number))

    return math.prod(list_prime_power_numbers)


def main():
    print(smallest_multiple(20))


if __name__ == '__main__':
    main()
