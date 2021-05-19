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

Kết quả số cặp (bên phải) dựa trên số phần tử (bên trái)
2 -> 1
3 -> 3
4 -> 6
5 -> 10
6 -> 14

Ta sẽ thấy quy luật là
2 + 1 = 3 
3 + 3 = 6
4 + 6 = 10
...
Cộng 2 số trước sẽ ra kết quả của lần tính sau