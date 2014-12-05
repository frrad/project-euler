import copy

top = 40

class partition:
    def __init__(self):
        self.data = dict()

    def __hash__(self):
        return hash(str(self))

    def __eq__(self, other):
        return str(self) == str(other)

    def __str__(self):
        self.clean()
        return str([self.data.get(i,0) for i in xrange(top)])
        
    # Returns number of segments in partition
    def count(self):
        self.clean()
        return sum((self.data[key] for key in self.data))

    # Number of blocks in entire confiugration
    def size(self):
        self.clean()
        return sum((self.data[key]*key for key in self.data))    

    def clean(self):
        cleansed = dict()
        for key in self.data:
            if key > 0 and self.data[key] != 0:
                cleansed[key] = self.data[key]
        self.data = cleansed

class max_state:
    def __init__(self):
        self.data = dict()

    def __str__(self):
        return str(self.data)

    # Enforces min of n by promoting things which are too small
    def promote(self, n):
        if n not in self.data:
            self.data[n] = 0

        for key in self.data:
            if key < n:
                self.data[n] += self.data[key]
                self.data[key] = 0

    # Returns number of segments in partition
    def count(self):
        return sum((self.data[key] for key in self.data))

    # Number of blocks in entire confiugration
    def size(self):
        return sum((self.data[key]*key for key in self.data))    

    def average(self):
        return float(self.size()) / float(self.count())

    def __add__(self, other):
        kees = set(self.data.keys())
        kees =         kees.union(set(other.data.keys())  )

        ans = max_state()
        for key in kees:
            ans.data[key] = other.data.get(key, 0) + self.data.get(key,0)

        return ans
            

    def __mul__(self, other):
        ans = copy.deepcopy(self)
        for key in ans.data:        
            ans.data[key] = other * ans.data[key]
        return ans

memo = dict()
def underlying(part):
    if part in memo:
        return copy.deepcopy(memo[part])

    if part.size() == 1:
        ans = max_state()
        ans.data[1] = 1

        memo[part] = ans
        return underlying(part)


    accumulate = max_state()
    count = part.count()
    for case in part.data:
        downsize = copy.deepcopy(part)
        downsize.data[case] -= 1

        for i in xrange(case):
            doomsize = copy.deepcopy(downsize)
            doomsize.data[i] = doomsize.data.get(i, 0) + 1
            doomsize.data[case -i-1] = doomsize.data.get(case-i-1, 0) + 1

            twinkle = underlying(doomsize)

            if case == 1:
                twinkle.promote(count)

            accumulate += twinkle * part.data[case]


    memo[part] = accumulate
    return underlying(part)

for i in xrange(4, 41):
    tester = partition()
    tester.data[i] = 1 

    print i,  underlying(tester),  underlying(tester).average()

