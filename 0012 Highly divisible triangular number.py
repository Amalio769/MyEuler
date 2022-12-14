"""
The sequence of triangle numbers is generated by adding the natural numbers. So the 7th triangle number would be
1 + 2 + 3 + 4 + 5 + 6 + 7 = 28. The first ten terms would be:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

Let us list the factors of the first seven triangle numbers:

 1: 1
 3: 1,3
 6: 1,2,3,6
10: 1,2,5,10
15: 1,3,5,15
21: 1,3,7,21
28: 1,2,4,7,14,28
We can see that 28 is the first triangle number to have over five divisors.

What is the value of the first triangle number to have over five hundred divisors?
Answer:  76576500
Completed on Wed, 17 Aug 2022, 07:43
"""


def triangle_number_generator(number):
    return sum([num for num in range(1, number+1, 1)])


def is_prime_factor(number):
    for i in range(2, number):
        if number % i == 0:
            return False
    return True


def first_prime_factor(number):
    for i in range(2, int(number+1)):
        if is_prime_factor(i):
            if number % i == 0:
                return i
    return False


def prime_factors(number):
    prime_factor_list = {1}
    while True:
        x = first_prime_factor(number)
        if x:
            prime_factor_list.add(x)
            number = number / x
        else:
            break

    return prime_factor_list


def factors_of_number(number):
    prime_factor_list = prime_factors(number)
    number_factors_before = len(prime_factor_list)
    number_factors_after = 0
    while number_factors_before != number_factors_after:
        number_factors_before = len(prime_factor_list)
        for num1 in prime_factor_list.copy():
            for num2 in prime_factor_list.copy():
                num = num1 * num2
                if number % num == 0:
                    prime_factor_list.add(num)
                    prime_factor_list.add(int(number / num))
        number_factors_after = len(prime_factor_list)
    return prime_factor_list


def main():
    idx = 10000
    while True:
        print(idx)
        if len(factors_of_number(triangle_number_generator(idx))) > 500:
            break
        else:
            idx += 1

    print('{} -> {}: {}'.format(idx, triangle_number_generator(idx), len(factors_of_number(triangle_number_generator(idx)))))


if __name__ == '__main__':
    main()
