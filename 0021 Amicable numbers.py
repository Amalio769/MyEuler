"""
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The
proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
Answer:  31626
Completed on Wed, 24 Aug 2022, 15:19
"""


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


def d(n):
    list = factors_of_number(n)
    return sum(list)-n

def is_amicable(a):
    if d(d(a)) == a and a != d(a):
        return True
    else:
        return False

amicable_list = []
for i in range(2, 10000, 1):
    print(i)
    if is_amicable(i):
        amicable_list.append(i)

print(sum(amicable_list))
