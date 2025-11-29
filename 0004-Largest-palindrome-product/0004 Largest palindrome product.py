"""
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is
9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
Answer:  906609
Completed on Sun, 14 Aug 2022, 23:32
"""

def is_palindrome(number):
    str_number = str(number)
    len_str_number = len(str_number)
    for i in range(int(len_str_number//2)):
        if str_number[i] != str_number[-1*(i+1)]:
            return False
    return True

def main():
    palindrome_list = []
    for x in range(999, 99, -1):
        for y in range(999, 99, -1):
            xy = x * y
            # print('x:{} y:{} xy:{}'.format(x, y, xy))
            if is_palindrome(xy):
                palindrome_list.append(xy)
    print(max(palindrome_list))



if __name__ == '__main__':
    main()