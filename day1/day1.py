import functools

f = open("input", "r")
depths = [int(line) for line in f.read().split("\n")]
counts = 0
counts_dec = 0
for i in range(len(depths)- 1) :
  if  depths[i + 1] > depths[i]:
    counts += 1
    print(f"{i + 1}  {depths[i + 1]}  {'(increased)' if depths[i + 1] > depths[i] else '(decreased)'}")

  else:
    counts_dec += 1

print(counts, counts_dec)