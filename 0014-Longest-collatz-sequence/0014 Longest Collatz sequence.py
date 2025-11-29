"""
The following iterative sequence is defined for the set of positive integers:
n → n/2 (n is even)
n → 3n + 1 (n is odd)
Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
Which starting number, under one million, produces the longest chain?
NOTE: Once the chain starts the terms are allowed to go above one million.
Answer:  837799
Completed on Wed, 17 Aug 2022, 22:24
"""


def is_even(number):
    return True if number%2 == 0 else False


def length_collatz_sequence(number):
    sequence = [number]
    while sequence[-1] != 1:
        if is_even(number):
            number = int(number/2)
        else:
            number = 3 * number + 1
        sequence.append(number)

    return len(sequence)


list_one_million = [num for num in range(1, 1000000, 1)]
list_len_collatz_sequence = list(map(lambda x: length_collatz_sequence(x) , list_one_million))
idx = list_len_collatz_sequence.index(max(list_len_collatz_sequence))
print(list_one_million[idx])
