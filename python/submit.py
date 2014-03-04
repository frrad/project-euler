#!/usr/bin/env python

#Submitter tool. For now only offline. Checks against database of correct answers.


import sys, subprocess

answerPath = '''../eulerdata/known.txt'''

def command(problem):
    return ['python', 'Problem%03d.py'%problem]

def validate(lookup, q, a):
    if not q in lookup:
        print "Problem %s not in table" % q
        sys.exit(1) #exit with error
    if lookup[q]==a:
        print "Correct!"
        return True
    else:
        print "Incorrect."
        return False

def solve(problem):
    print "Solving Problem #%d" % problem
    cmd = command(problem)

    p = subprocess.Popen(cmd, stdout=subprocess.PIPE,stderr=subprocess.PIPE)
    out, err = p.communicate()
    if err != "":
        print "There was an error executing %s" % cmd
        sys.exit(1)

    answer = out.split("\n")[-2]
    print "Problem Solved: %s" % answer
    return answer


f = open(answerPath, "r")
data =  f.read().split("\n")
f.close()

lookup = dict()

for line in data:
    pair = line.split(":")
    if len(pair)==2:
        lookup[int(pair[0])] = pair[1]
    else:
        pass#maybe we shouldn't fail silently?

if len(sys.argv) == 3:
    #two arguments: problem and solution
    problem, solution = int(sys.argv[1]), sys.argv[2]
    validate(lookup,problem,solution)


elif len(sys.argv) == 2:
    #one argument: problem number (or range). run and check answer
    if "-" in sys.argv[1]: #we're dealing with a range
        endpts = sys.argv[1].split("-")
        start, end = int(endpts[0]), int(endpts[1])
        for problem in xrange(start,end+1):
            answer = solve(problem)
            validate(lookup,problem,answer)

    else: #just one problem is specified
        problem = int(sys.argv[1]) 
        answer = solve(problem)
        validate(lookup,problem,answer)

else: 
    print "Wrong number of arguments"
