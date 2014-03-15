months = dict({1:31,2:28,3:31,4:30,5:31,6:30,7:31,8:31,9:30,10:31,11:30,12:31})

def isLeap(year):
    if year % 4 == 0:
        if year %100 == 0:
            if year % 400 == 0:
                return True
            return False
        return True
    return False


def mutate(month, year, dow):
    dow += months[month]
    if month == 2 and isLeap(year):
        dow += 1
    dow %= 7

    month = (month + 1) % 12
    if month == 0: month = 12

    if month == 1: year += 1

    return month, year, dow


#1 January 1900 Monday
(month, year, dow) = (1 , 1900, 1)
 
count = 0

while year<=2000:
    month, year, dow = mutate(month, year, dow)
    if dow == 0 and year > 1900:
        count += 1

print count