import euler

def check((n,d)):
    top, bottom = set(str(n)), set(str(d))
    remove = top.intersection(bottom)

    if len(remove) != 1: 
        return False

    common = remove.pop()
    if common == "0": 
        return False

    topL, bottomL = list(str(n)), list(str(d))
    topL.remove(common)
    bottomL.remove(common)
    newN, newD = int(topL.pop()), int(bottomL.pop())

    return n*newD == d*newN

fracs = ((n, d) for n in range(10,100) for d in range(n,100) if check((n,d)))
product = reduce(lambda (a,b), (c,d) : (a*c, b*d), fracs)
print product[1] / euler.GCD(product[0], product[1])
