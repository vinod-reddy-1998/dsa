l=["(","{","["]
m=input()
n=[]
t=False
for item in m:
    if item in l:
        n.append(item)
        print(n)
    else:
        if item=="}":
            if n[-1]=="{":
                n.pop()
                print(n)
                t=True
            else:
                t=False
        elif item==")":
            if n[-1]=="(":
                n.pop()
                print(n)
                t=True
            else:
                t=False
        elif item=="]":
            if n[-1]=="[":
                n.pop()
                print(n)
                t=True
            else:
                t=False

if t:
    print("they are equal")
else:
    print("they are not equal..")