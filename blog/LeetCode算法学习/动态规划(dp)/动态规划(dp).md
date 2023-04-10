[toc]

## 动态规划

动态规划(dp)

* 递归+记忆法 ==> 递推
* 定义模型状态
* 构造动态方程
* 拆分子问题

###  难度：简单

### 70、爬楼梯

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



### 题解

#### 方法一：回溯(超时)

```go
func climbStairs(n int) int {
	tag := make([]int, n+1)
	return stairsdp(n, tag)
}

func stairsdp(n int, tag []int) int {
	if n == 2 {
		return 2
	}
	if n == 1 {
		return 1
	}

	if tag[n] != 0 {
		tag[n] = stairsdp(n-1, tag) + stairsdp(n-2, tag)
		return tag[n]
	}
	return stairsdp(n-1, tag) + stairsdp(n-2, tag)
}
```



#### 方法二：动态规划

```go
func climbStairs(n int) int {
	//核心是统计爬楼梯方式的数量
	//用dp来记录到第i阶楼梯的方式数，使用dp记录
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
```

压缩内存

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



### 难度：中等

### 120、三角形最小路径和

给定一个三角形 `triangle` ，找出自顶向下的最小路径和。

每一步只能移动到下一行中相邻的结点上。**相邻的结点** 在这里指的是 **下标** 与 **上一层结点下标** 相同或者等于 **上一层结点下标 + 1** 的两个结点。也就是说，如果正位于当前行的下标 `i` ，那么下一步可以移动到下一行的下标 `i` 或 `i + 1` 。

 

**示例 1：**

```
输入：triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
输出：11
解释：如下面简图所示：
   2
  3 4
 6 5 7
4 1 8 3
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
```

**示例 2：**

```
输入：triangle = [[-10]]
输出：-10
```

 

**提示：**

- `1 <= triangle.length <= 200`
- `triangle[0].length == 1`
- `triangle[i].length == triangle[i - 1].length + 1`
- `-104 <= triangle[i][j] <= 104`

 

**进阶：**

- 你可以只使用 `O(n)` 的额外空间（`n` 为三角形的总行数）来解决这个问题吗？



### 题解

#### 方法一：贪心法

使用贪心法，每次求最优，从局部最优，到全局最优。提供一种方法，但是不是本题解题方案。

```go
func minimumTotal(triangle [][]int) int {
	signIndex := 0
	deth := 0
	//遍历二维数组
	for i := 0; i < len(triangle); i++ {
		arr := triangle[i]
		if len(arr) == 1 {
			deth += arr[signIndex]
			continue
		}
		if arr[signIndex] <= arr[signIndex+1] {
			deth += arr[signIndex]
		} else {
			deth += arr[signIndex+1]
			signIndex = signIndex + 1
		}
	}
	return deth
}
```



#### 方法二：回溯法(超时)

使用回溯法自底向上进行回溯，判断最小值。回溯法能解决本题，但会超时。

```go
func minimumTotal1(triangle [][]int) int {
	//回溯法(超时）无法剪枝
	return dfs(triangle, 0, 0)
}

func dfs(triangle [][]int, row int, signIndex int) int {
	//递归到底
	if row == len(triangle) {
		return 0
	}
	//下放-回溯
	ldeth := dfs(triangle, row+1, signIndex)
	rdeth := dfs(triangle, row+1, signIndex+1)

	if ldeth <= rdeth {
		return ldeth + triangle[row][signIndex]
	} else {
		return rdeth + triangle[row][signIndex]
	}
}
```



#### 方法三：动态规划(dp)

自底向上进行递推，```dp[i][j] = mini(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]```状态方程。

```go
func minimumTotal2(triangle [][]int) int {
	//动态规划，自底向上，dp状态，dp方程
	h := len(triangle)
	dp := make([][]int, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}

	//自底向上遍历二维数组
	for i := h - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			//到达二维数组底层，将底层所有元素给到dp
			if i == h-1 {
				//dp状态
				dp[i][j] = triangle[i][j]
			} else {
				//dp状态方程
				//回到上一层时，需要判断的当前层的子层的最小值
				dp[i][j] = mini(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
			}
		}
	}
	return dp[0][0]
}

func mini(x, y int) int {
	if x > y {
		return y
	}
	return x
}

```

