[toc]



### 加一

给定一个由 **整数** 组成的 **非空** 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储**单个**数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

 

**示例 1：**

```
输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
```

**示例 2：**

```
输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
```

**示例 3：**

```
输入：digits = [0]
输出：[1]
```

 

**提示：**

- `1 <= digits.length <= 100`
- `0 <= digits[i] <= 9`



#### 题解：

##### 方法一：转数字

就是将slice中的元素拼接成对应的数，然后加一，当这个数比较小还好，当他的值达到一定程度就会发生溢出，使用这个方法是不可行的

下面提供这个方法的代码：

```go
func plusOne(digits []int) []int {
    var num int
    for i := range digits {
        num = (num + digits[i])*10
    }
    
    if num == 0 {
        return []int{1}
    }

    var res []int
    str := strconv.Itoa((num/10)+1)
    for _, v := range str {
        
        res = append(res, int(v)-48)
    }
    return res
}
```



正确的代码：

思路：从后往前遍历，遍历到9时则说明需要在9的前一个数+1， 然后这个9变为0， 所以当遇到9时，先不做任何处理，找到不为9的值后将其+1， 然后在从此时遍历到的位置向后依次将9改为0; 当数组中前是9时，直接申请一个比之前空间多1的slice，然后将slice[0] == 1, 其他值默认为0。

```go
func plusOne(digits []int) []int {
   //直接找到不为9的数，然后直接+1
   for i := len(digits)-1; i >= 0; i-- {
       if digits[i] != 9 {
           digits[i]++
           
           //对于找到9有边是数后，其遍历到的9都应该为0
           for j := i + 1; j < len(digits); j++ {
               digits[j] = 0
           }
           return digits
       }
   }

   //处理全为9的情况
   res := make([]int, len(digits)+1)
   res[0] = 1
   return res
}
```



### 二进制求和

给你两个二进制字符串 `a` 和 `b` ，以二进制字符串的形式返回它们的和。

 

**示例 1：**

```
输入:a = "11", b = "1"
输出："100"
```

**示例 2：**

```
输入：a = "1010", b = "1011"
输出："10101"
```

 

**提示：**

- `1 <= a.length, b.length <= 104`
- `a` 和 `b` 仅由字符 `'0'` 或 `'1'` 组成
- 字符串如果不是 `"0"` ，就不含前导零

#### 题解：

##### 方法一：转为10进制

使用这里我们可以自己来实现一个进制转换函数,然后将其相加，再转为二进制

代码一：

对于不大的数据是可以的，但是对于很大的数据就会溢出，不可行。

```go
func addBinary(a string, b string) string {
    dnum_a := Str2DEC(a)
    dnum_b := Str2DEC(b)

    res := fmt.Sprintf("%b", dnum_a + dnum_b)
    return res
}

func Str2DEC(s string) (num int) {
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		num += (int(s[l-i-1]) - 48) << uint8(i)
	}
	return
}
```



###### 使用math包

```go
func addBinary(a string, b string) string {
    ai, _ := new(big.Int).SetString(a, 2)
	bi, _ := new(big.Int).SetString(b, 2)

	ai.Add(ai, bi)
	return ai.Text(2)
}
```

