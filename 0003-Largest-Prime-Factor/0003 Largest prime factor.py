"""
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
Answer:  6857
Completed on Sun, 14 Aug 2022, 22:37
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
    prime_factor_list = [1, ]
    while True:
        x = first_prime_factor(number)
        if x:
            prime_factor_list.append(x)
            number = number / x
        else:
            break

    return prime_factor_list


def main():
    print(prime_factors(600851475143)[-1])


if __name__ == '__main__':
    main()
