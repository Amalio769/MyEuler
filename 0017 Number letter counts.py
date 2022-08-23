"""
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19
letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one
hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
Answer:  21124
Completed on Tue, 23 Aug 2022, 19:28
"""

units = ['one',
         'two',
         'three',
         'four',
         'five',
         'six',
         'seven',
         'eight',
         'nine']
teens = ['ten',
         'eleven',
         'twelve',
         'thirteen',
         'fourteen',
         'fifteen',
         'sixteen',
         'seventeen',
         'eighteen',
         'nineteen']
tens = ['twenty',
        'thirty',
        'forty',
        'fifty',
        'sixty',
        'seventy',
        'eighty',
        'ninety']
lista = []
for unit in units:
    lista.append(unit)
for teen in teens:
    lista.append(teen)
for ten in tens:
    lista.append(ten)
    for unit in units:
        lista.append('{} {}'.format(ten, unit))
for unit_h in units:
    lista.append('{} hundred'.format(unit_h))
    for unit in units:
        lista.append('{} hundred and {}'.format(unit_h, unit))
    for teen in teens:
        lista.append('{} hundred and {}'.format(unit_h, teen))
    for ten in tens:
        lista.append('{} hundred and {}'.format(unit_h, ten))
        for unit in units:
            lista.append('{} hundred and {} {}'.format(unit_h, ten, unit))
lista.append('one thousand')
lista_without_spaces = list(map(lambda x: x.replace(" ", "") , lista))
lista_len = list(map(lambda x: len(x), lista_without_spaces))


print(lista)
print(lista_without_spaces)
print(lista_len)
print(sum(lista_len))
