import euler

top = 20

answer = 1

for i in xrange(1,top):
    answer = euler.LCM(i,answer)

print answer