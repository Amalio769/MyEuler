# If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these
# multiples is 23.
#
# Find the sum of all the multiples of 3 or 5 below 1000.
# Answer: 233168
# Completed on Thu, 16 Dec 2021, 08:08

list_3 = Enum.to_list(3..999//3)
list_5 = Enum.to_list(5..999//5)
list_total = list_3 ++ list_5
IO.puts(Enum.sum(Enum.uniq(list_total)))