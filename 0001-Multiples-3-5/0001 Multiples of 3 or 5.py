"""
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these
multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
Answer: 233168
Completed on Thu, 16 Dec 2021, 08:08
"""


def main():
    list_numbers = []
    for i in range(1, 1000):
        if i % 3 == 0 or i % 5 == 0:
            list_numbers.append(i)
    print(sum(list_numbers))


if __name__ == '__main__':
    main()
