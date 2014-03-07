def spell(number):
    if number < 20:
        small = {0:'', 1:'one', 2:"two", 3:"three", 4:"four", 5:"five",6:"six",7: "seven", 8:"eight",  9: "nine",  10: "ten",  11 : "eleven",  12 : "twelve",  13:  "thirteen" , 14 : "fourteen", 15 : "fifteen" ,  16:  "sixteen" ,  17 : "seventeen",  18 : "eighteen",  19 : "nineteen" }
        return small[number]

    if number > 19 and number < 100:
        tens = number / 10
        ones = number % 10

        big = {2:"twenty ", 3:"thirty ", 4:"forty ", 5:"fifty ", 6:"sixty ", 7:"seventy ", 8:"eighty ", 9:"ninety "}
        return big[tens]+spell(ones)

    if number > 99 and number < 1000:
        ending = number % 100
        hundreds = number / 100

        if ending == 0:
            return spell(hundreds) + "hundred"
        else:
            return spell(hundreds) + " hundred and " + spell(ending)

    if number == 1000:
        return "one thousand"


def count(word):
    if len(word) == 0:
        return 0
    if word[0] == " ":
        return count(word[1:])
    return 1 + count(word[1:])


total = 0
for x in xrange(1,1001):
    # print spell(x)
    total += count(spell(x))

print total