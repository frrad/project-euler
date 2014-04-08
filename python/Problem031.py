denominations = [200, 100, 50, 20, 10, 5, 2, 1]

def ways(amount, index):
    if index == len(denominations)-1: return 1

    total = ways(amount, index + 1)
    value = denominations[index]

    while amount >= value:
        total += ways(amount - value, index +1)
        amount -= value
    return total

print ways(200,0)
