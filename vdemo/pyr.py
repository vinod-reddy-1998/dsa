r = int(input("Enter number of rows: "))

k = 0
count=0
count1=0

for i in range(1, r+1):
    for j in range(1, i+1):
       print((r-i)*" ",j,end=",")
    print()
