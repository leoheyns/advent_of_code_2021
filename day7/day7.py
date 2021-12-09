from statistics import*
# c=list(map(int,open('input').read().split(',')))
# print(sum(abs(x-median(c))for x in c))
#
c=list(map(int,open('input').read().split(',')))
print(min(sum((d:=abs(x-i))*(d+1)/2for x in c)for i in c))


# *c,=map(int,next(open('input')).split(','))
# print(min(sum((x-i)**2+abs(x-i)for x in c)/2for i in c))