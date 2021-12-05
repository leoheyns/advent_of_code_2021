from itertools import*
from collections import*
import re
r=lambda a,b:cycle([a])if a==b else range(a,b+(c:=(a<b)*2-1),c)
print(sum(a>1 for a in Counter(chain(*[zip(r(c[0],c[2]),r(c[1],c[3]))for c in [list(map(int,re.split(",| -> ",l)))for l in open("i","r").readlines()]])).values()))
