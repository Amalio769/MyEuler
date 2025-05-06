"""
Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names,
begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value
by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the
938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
Answer:  871198282
Completed on Thu, 25 Aug 2022, 06:41
"""

def string_value(text):
    alphabet = "abcdefghijklmnopqrstuvwxyz"
    value = 0

    for letter in text:
        value += alphabet.find(letter.lower())+1

    return value

my_file = open("0022_names.txt", "r")
content = my_file.read()
content = content.replace('"','')
content_list = content.split(",")
my_file.close()

content_list.sort()

value = 0
for name in content_list:
    value += (content_list.index(name) + 1) * string_value(name)

print(value)
