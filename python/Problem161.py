import copy

mitts = [[(0, 0), (1,0), (0,1)],
         [(0, 1), (1,1), (0,0)],
         [(1, 0), (0,0), (1,1)],
         [(1, 0), (0,0), (1,-1)],

         [(0, 0), (1,0), (2,0)],
         [(0, 0), (0,1), (0,2)],
]

memo = dict()

class board():
    def __init__(self, h, w):
        self.state = [[True for x in xrange(w)] for y in xrange(h)]

    def __str__(self):
        ans = ""
        for line in self.state:
            for place in line:
                ans += " " if not place else "X"
            ans += "\n"
        return ans[:-1]

    # Rotate / flip to canonical orientation
    def canonize(self):
        self.trim()

        measure = str(self)
        best_state = copy.deepcopy(self.state)

        for i in xrange(3):
            self.rotate()
            if str(self) < measure:
                measure = str(self)
                best_state = copy.deepcopy(self.state)
        
        self.flip()

        for i in xrange(4):
            self.rotate()
            if str(self) < measure:
                measure = str(self)
                best_state = copy.deepcopy(self.state)

        self.state = best_state

    def trim(self):
        if self.width() == 0 or self.width() == 0: 
            self.state = [[]]
            return

        start_size = self.size() + 1

        while self.size() != start_size:
            start_size = self.size()
            for i in range(4):
                self.rotate()
                if len(self.state) < 1:
                    continue
                if len(self.state[0]) < 1:
                    continue
                if len(self.state[0]) == 1:
                    if not self.state[0][0]:
                        self.state.pop(0)
                        continue
                    else:
                        continue
                if not reduce(lambda a,b: a or b, self.state[0]) :
                    self.state.pop(0)


    def size(self):
        return self.length() + self.width()

    def length(self):
        return len(self.state)
    
    def width(self):
        if self.length() >=1:
            return len(self.state[0])
        else:
            return 0

    # Rotate 90 degrees clockwise.
    def rotate(self):
        if self.tiles() ==0 : return
        self.flip()
        new_state = [[self.state[i][j] for i in xrange(len(self.state))] for j in xrange(len(self.state[0]))]
        self.state = new_state

    def flip(self):
        for line in self.state: line.reverse()

    # Number of tiles
    def tiles(self):
        if self.length() == 0 or self.width() == 0: return 0
        return sum([sum([1 if i else 0 for i in line]) for line in self.state])


    def tilings(self):
        self.canonize()

        if self.tiles() == 0:
            return 1

        if self.tiles()%3 != 0:
            return 0
            print "invalid"

        if hash(str(self)) in memo:
            return memo[hash(str(self))]

        ans = 0

        x, y = 0, 0

        while not self.state[0][y]:
            y += 1

        for mitt in mitts:
            skip = False
            for delta in mitt:
                if x+delta[0] < 0 or x + delta[0] >= len(self.state) or y + delta[1] < 0 or y + delta[1] >= len(self.state[0]) or not self.state[x+delta[0]][y+delta[1]]:
                    skip = True
                    break
            if skip: continue

            new_state = copy.deepcopy(self.state)         
            for delta in mitt:
                new_state[x+delta[0]][y+delta[1]] = False
            downstairs = board(len(new_state), len(new_state[0]))
            downstairs.state = new_state
            ans += downstairs.tilings()

        memo[hash(str(self))] = ans
        return ans


j = 9
for i in xrange(2,13):
    test = board(j,i)
    print i, "x", j
    print "tilings:", test.tilings()
    print "============"
