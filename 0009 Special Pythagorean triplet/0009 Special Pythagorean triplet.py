"""
A Pythagorean triplet is a set of three natural numbers, a<b<c, for which,
a²+b²=c²
There exists exactly one Pythagorean triplet for which a+b+c=1000.
Find the product abc.
Answer:  31875000
Completed on Mon, 15 Aug 2022, 17:54
"""

def is_pythagorean_triplet(a, b, c):
    if a**2 + b**2 == c**2:
        return True
    else:
        return False

def main():
    sum_a_b_c = 1000

    for a in range(1, sum_a_b_c):
        for b in range(a, sum_a_b_c):
            if a + b < sum_a_b_c:
                c = sum_a_b_c - a - b
                if is_pythagorean_triplet(a, b, c):
                    if is_pythagorean_triplet(a, b, c):
                        print('a:{} b:{} c:{} product:{}'.format(a, b, c, a*b*c))
                        break



if __name__ == '__main__':
    main()