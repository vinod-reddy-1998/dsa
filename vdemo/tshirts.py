l=list(map(int,input().split()))
c=0
for item in l:
    # print("item",item)
    for i in range(1,item+1):
        k=int(str(i)[::-1])
        # print("k",k)
        if i+k==item and len(str(i))==len(str(k)):
            # print("val",i,k)
            c+=1
            break
print(c)