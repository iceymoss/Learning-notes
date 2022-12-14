### 算术平方根

给你一个非负整数 `x` ，计算并返回 `x` 的 **算术平方根** 。

由于返回类型是整数，结果只保留 **整数部分** ，小数部分将被 **舍去 。**

**注意：**不允许使用任何内置指数函数和算符，例如 `pow(x, 0.5)` 或者 `x ** 0.5` 。

 

**示例 1：**

```
输入：x = 4
输出：2
```

**示例 2：**

```
输入：x = 8
输出：2
解释：8 的算术平方根是 2.82842..., 由于返回类型是整数，小数部分将被舍去。
```

 

**提示：**

- `0 <= x <= 231 - 1`

#### 题解：

##### 方法一：取对数

题目中说了不能直接使用内建方法和`x ** 0.5`， 但是没有说不能使用库函数要知道go有强大的math包的，那么如何求x的算术平方根呢？用用数学方法嘛，这貌似太难了，注意题目只是不允许使用内建方法和`x ** 0.5`， 那么我们将`x ** 0.5`用数学方法转换成另一种表达形式不就行了吗，对x取对数：`x ** 0.5` 就变成了
$$
ans = x^{1/2} = (e^{ln^x})^{1/2} = e^{1/2lnx}
$$
这样就可以避免`x ** 0.5`了， 但是这里需要注意：例如当 x=2147395600时，$$ e^{1/2lnx} $$
  的计算结果与正确值 463404634046340 相差 10−1110^{-11}10 −11 ，这样在对结果取整数部分时，会得到 463394633946339 这个错误的结果。

```go
func mySqrt(x int) int {
    if x == 0 {
        return 0
    }

    ans := int(math.Exp(0.5 * math.Log(float64(x))))
    if (ans+1) * (ans+1) <= x {
        return ans + 1
    }
    return ans
}
```



##### 方法二：二分查找法

由于 ```x``` 平方根的整数部分 ans 是满足 ```k^2 ≤ x  ```的最大 k值，因此我们可以对 k 进行二分查找，从而得到答案。

二分查找的下界为 0，上界可以粗略地设定为 x。在二分查找的每一步中，我们只需要比较中间元素mid的平方和x的大小关系，并通过比较的结果调整上下界的范围。由于我们所有的运算都是整数运算，不会存在误差，因此在得到最终的答案 。

```go
func mySqrt(x int) int {
   //二分查找法
   if x == 0 {
       return 0
   }
   ans := -1
   l, r := 0, x
   for l <= r {
       mid :=l + (r - l)/2
       if mid * mid <= x {
           ans = mid
           l = mid + 1
       }else {
           r = mid - 1
       }
   }
   return ans
}
```





### 爬楼梯

假设你正在爬楼梯。需要 `n` 阶你才能到达楼顶。

每次你可以爬 `1` 或 `2` 个台阶。你有多少种不同的方法可以爬到楼顶呢？

 

**示例 1：**

```
输入：n = 2
输出：2
解释：有两种方法可以爬到楼顶。
1. 1 阶 + 1 阶
2. 2 阶
```

**示例 2：**

```
输入：n = 3
输出：3
解释：有三种方法可以爬到楼顶。
1. 1 阶 + 1 阶 + 1 阶
2. 1 阶 + 2 阶
3. 2 阶 + 1 阶
```

 

**提示：**

- `1 <= n <= 45`



#### 题解：

##### 方法一：动态规划

想一下，当走第n个台阶时，上一步是不是有两种可能：

1. 上一步走了一个台阶： n-1
2. 上一步走了两个台阶：n-2

那么走到第n个台阶时出现的可能就是f(n) = f(n-1) + f(n-2)

接着这个问题是不是可以将f(n-1)和f(n-2)按照同样的思路拆分呢？答案：可以

这个过程就涉及到了递归思想，我们回到刚开始走台阶的时候：

当走第一个台阶时：1种可能: 即 走一个台阶； 当走到第二个台阶时：两种可能， 即：一台阶 x 2或者直接走两个台阶。其实这两种情况就是```base case``` 即：递归终止条件。

递归：

```go
func climbStairs(n int) int {
    if n == 1 {
        return 1;
    }else if n == 2 {
        return 2
    }
    return climbStairs(n-1) + climbStairs(n-2)
}
```

但是这里消耗的时间和时间是比较大的，所以更好的方法是使用迭代：

迭代：

```go
func climbStairs(n int) int {
    if n == 1 {
        return 1;
    }
    dp := make([]int, n+1)  //n+1保证n如何变化，不会越界
    dp[1], dp[2] = 1, 2
    for i := 3; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    return dp[n]
}
```



当然这里可以进一步优化：不需要使用slice

```go
func climbStairs(n int) int {
   per := 1
   cur := 1
   for i := 2; i < n+1; i++ {
       tmp := cur
       cur = cur + per
       per = tmp
   }
   return cur
}
```

总结：

>##### 可以用动态规划的问题都能用递归
>* 从子问题入手，解决原问题，分两种做法：自顶向下和自底向上
>* 前者对应递归，借助函数调用自己，是程序解决问题的方式，它不会记忆解
>* 后者对应动态规划，利用迭代将子问题的解存在数组里，从数组0位开始顺序往后计算
>* 递归的缺点在于包含重复的子问题（没有加记忆化的情况下），动态规划的效率更高



详细题解可参考：[手画图解」详解爬楼梯问题 | 从递归，讲到动态规划](https://leetcode.cn/problems/climbing-stairs/solutions/270926/cong-zhi-jue-si-wei-fen-xi-dong-tai-gui-hua-si-lu-/?languageTags=golang)





