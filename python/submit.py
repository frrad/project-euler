#!/usr/bin/env python2

#Submitter tool. For now only offline. Checks against database of correct answers.


import sys, subprocess, time

answerPath = '''/home/frederick/.euler-tools/known.txt'''
infinity = 9999999999

def command(problem):
    return ['python2', 'Problem%03d.py'%problem]

def validate(lookup, q, a):
    if not q in lookup:
        print "Problem %s not in table" % q
        return False, False
    if lookup[q]==a:
        return True, True
    else:
        return True, False

def progress(start,end,solved):
    tosolve = end-start + 1
    percent = float(solved)/tosolve

    draw = int(tosolve * percent)
    undraw = tosolve - draw

    fancy = "["+ draw*'-' + undraw*" "+"]"

    return '{} {}/{}'.format(fancy ,solved,tosolve)

def solve(problem):
    start = time.time()
    print "Solving Problem #%d" % problem
    cmd = command(problem)

    p = subprocess.Popen(cmd, stdout=subprocess.PIPE,stderr=subprocess.PIPE)
    out, err = p.communicate()
    if err != "":
        print "There was an error executing %s" % cmd
        return False, None, infinity

    answer = out.split("\n")[-2]
    elapsed = time.time()-start

    print "Problem Solved: %s %.2fs" % (answer.ljust(15),elapsed)

    return True, answer, elapsed


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
        if len(endpts[0]) == 0:
            start = 1
        else:
            start =int(endpts[0])

        end =  int(endpts[1])
        solved = 0
        times = dict()
        for problem in xrange(start,end+1):
            works, answer, times[problem] = solve(problem)
 
            message = "Incorrect."
            table, correct = validate(lookup,problem,answer)
            if works and correct:
                message = "Correct!"
                solved += 1

                 
            print '{} {}'.format(message,progress(start,end,solved))

        speed = sorted(range(start,end +1), key=lambda x:times[x])
        show = "\n"
        speed.reverse()
        for problem in speed:
            show += "%d (%.2fs), " % (problem, times[problem])
        print show.rstrip(' ,')

    else: #just one problem is specified
        problem = int(sys.argv[1]) 
        works, answer, _ = solve(problem)
        inTable, correct = validate(lookup,problem,answer)
        if not inTable: sys.exit(1) #exit with error
        if works and correct:
            print "Correct!"
        else:
            print "Incorrect."


else: 
    print "Wrong number of arguments"
