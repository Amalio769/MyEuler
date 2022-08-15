"""
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
Answer:  142913828922
Completed on Mon, 15 Aug 2022, 18:40
"""


def is_prime(number, prime_list):
    for prime in prime_list:
        if number % prime == 0:
            return False
    return True


def primes_below_number(num: int):
    prime_list = [2,]
    for x in range(3, num):
        if is_prime(x, prime_list):
            prime_list.append(x)
    return prime_list


def main():
    print(sum(primes_below_number(2000000)))


if __name__ == '__main__':
    main()
