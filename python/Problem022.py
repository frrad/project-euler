def score(name):
    return sum([ord(letter)-64 for letter in name])

f = open("../problemdata/names.txt","r")
filedata = f.read()
f.close()

names = filedata.split(",")

cleaned =  [ bob.strip('"') for bob in names]
cleaned.sort()

total = 0
for i, name in enumerate(cleaned):
    total += score(name)*(i+1)

print total