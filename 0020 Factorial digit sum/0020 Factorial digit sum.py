"""
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
Answer:  648
Completed on Wed, 24 Aug 2022, 06:37
"""

def factorial(num):
    if num == 1:
        return 1
    else:
        return num * factorial(num-1)

num_string = str(factorial(100))
sum_num = 0
for i in range(len(num_string)):
    sum_num += int(num_string[i])
print(sum_num)
