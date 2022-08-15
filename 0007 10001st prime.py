"""
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
Answer:  104743
Completed on Mon, 15 Aug 2022, 15:54
"""

def is_prime_factor(number):
    for i in range(2, number):
        if number % i == 0:
            return False
    return True


def main():
    number = 2
    list_prime = []
    while True:
        if is_prime_factor(number):
            list_prime.append(number)
        number += 1
        if len(list_prime) == 10001:
            break
    print(list_prime[-1])

if __name__ == '__main__':
    main()
