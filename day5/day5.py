from itertools import*
from collections import*
import re
r=lambda a,b:range(a,b+(c:=(a<b)*2-1),c)if a-b else[a]*9**4
print(sum(a>1for a in Counter(chain(*[zip(r(a,c),r(b,d))for a,b,c,d in[tuple(map(int,re.split(",| -> ",l)))for l in open("i")]])).values()))

from itertools import*
import re
r=lambda a,b:range(a,b+(c:=(a<b)*2-1),c)if a-b else[a]*9**4
c=list(chain(*[zip(r(a,c),r(b,d))for a,b,c,d in[tuple(map(int,re.split(",| -> ",l)))for l in open("i")]]))
print(sum(c.count(v)>1for v in set(c)))