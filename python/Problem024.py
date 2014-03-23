import euler

def permute(digits, index):
    if len(digits)==0: return []
    base = euler.factorial( len(digits) -1)
    here = digits.pop(index/base)
    rest = permute(digits, index % base)
    answer = [here]
    answer.extend(rest)
    return answer

def permutation(n, of):
    return "".join(permute(map(str,range(of)), n-1))
 
print permutation(10**6, 10)