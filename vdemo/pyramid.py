l=list(map(int,input().split()))
d={}
c=0
for i in l:
    d[c]=i
    c=c+1

j=True
for i in range(1,999):
    if j:
        for char in str(i):
            if (str(i).count(char)>d[int(char)]):
                print(i)
                j=False
                break

    


