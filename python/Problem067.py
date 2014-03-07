data = open('../problemdata/triangle.txt','r').read().strip()
triangle = [map(int,line.split(" ")) for line in data.split('\n')]

for lindex in range(len(triangle)-1, 0,-1):
    for j in range(lindex):
        triangle[lindex-1][j] += max(triangle[lindex][j], triangle[lindex][j+1]) 

print triangle[0][0]