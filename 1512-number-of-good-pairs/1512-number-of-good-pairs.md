# 1512-number-of-good-pairs 

## Cách 2

Thử tính số cặp đối với trường hợp mảng có các item giống hệt nhau, để tìm ra công thức chung

[1,1]

0,1

[1,1,1]

0,1
0,2
1,2

[1,1,1,1]

0,1
0,2
0,3
1,2
1,3
2,3

[1,1,1,1,1]

0,1
0,2
0,3
0,4
1,2
1,3
1,4
2,3
2,4
3,4

[1,1,1,1,1,1]

0,1
0,2
0,3
0,4
0,5
1,2
1,3
1,4
1,5
2,3
2,4
2,5
3,4
3,5

Thử tính toán để tìm quy luật số cặp
số phần tử -> số cặp
        1 -> 0
        2 -> 1
        3 -> 3
        4 -> 6
        5 -> 10
        6 -> 15

Ta sẽ thấy quy luật là
  Nếu ta cộng số phần tử xuất hiện + số cặp của lần trước == số cặp của 
  số phần tử và lần xuất hiện sau
    1 + 0 = 2 (TH có 2 phần tử)
    2 + 1 = 3 (TH có 3 phần tử)
    3 + 3 = 6 ...
    4 + 6 = 10
    ...
  Hoặc cách tính khác là giả sử số lần xuất hiện là A sẽ bằng tổng của từ 1 -> A - 1
    VD: ta có 5 số 1 thì số lần xuất hiện của good pair của số 1 sẽ là
      1 + 2 + 3 = 6



------------------------------------------
Tính thử
1 2 2 1 1 1 1 2 2
            i       

1 2
---
4 2

result: 0 + 1 (2) = 1 + 1 (1) = 2 + 2 (1) = 4 + 3 (1) = 7 + 4 (1) = 11

------------------------------------------
Tính thử
1 1 1 1 1 1
      i    

1
---
3

result: 0 + 0 (s1) = 0 + 1 (s2) = 1 + 2 (s3) = 3 + 3 (s4) = 6
sX: "step X"