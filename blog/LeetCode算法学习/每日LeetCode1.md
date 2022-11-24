[toc]



### 两数之和

给定一个整数数组 `nums` 和一个整数目标值 `target`，请你在该数组中找出 **和为目标值** *`target`* 的那 **两个** 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

 

**示例 1：**

```
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
```

**示例 2：**

```
输入：nums = [3,2,4], target = 6
输出：[1,2]
```

**示例 3：**

```
输入：nums = [3,3], target = 6
输出：[0,1]
```

 

**提示：**

- `2 <= nums.length <= 104`
- `-109 <= nums[i] <= 109`
- `-109 <= target <= 109`
- **只会存在一个有效答案**



#### 题解：

##### 方法一：暴力枚举

对数组每一个元素都进行遍历，尝试所有组合

```go
func twoSum(nums []int, target int) []int {
    for j := 0; j < len(nums); j++ {
        for i := j + 1; i < len(nums); i++ {
            if target == nums[j] + nums[i] {
                return []int{j, i}
            }
        }
    }
    return []int{}
}
```



##### 方法二：哈希map

创建一个map,我们以nums[i]作为map的key, i作为map的value,行程映射，然后对数组遍历，每遍历一个数a时，将target减去a,得到b, 然后就以b为keyz直接对map进行查询，如果map中存在b这个key,就返回对应的value; 如果不存在就以b为map的key, i 为map的value放入map中，进入下一次遍历。

```go
func twoSum(nums []int, target int) []int {
    //使用map
    m := make(map[int]int)
    for i := 0; i < len(nums); i++ {
        b := target - nums[i] 
        if v, ok := m[b]; ok {
            return []int{i, v}
        }else {
            m[nums[i]] = i
        }
    }
    return []int{}
}
```



### 回文数

给你一个整数 `x` ，如果 `x` 是一个回文整数，返回 `true` ；否则，返回 `false` 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

- 例如，`121` 是回文，而 `123` 不是。

 

**示例 1：**

```
输入：x = 121
输出：true
```

**示例 2：**

```
输入：x = -121
输出：false
解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
```

**示例 3：**

```
输入：x = 10
输出：false
解释：从右向左读, 为 01 。因此它不是一个回文数。
```

 

**提示：**

- `-231 <= x <= 231 - 1`

 

**进阶：**你能不将整数转为字符串来解决这个问题吗？



#### 题解：

##### 方法一：转字符串

将数字转为字符串，然后将字符串反转即可

```go
func isPalindrome(x int) bool {
    strx := fmt.Sprintf("%d", x)
    var s []byte
    for i :=len(strx); i > 0; i-- {
        s = append(s, strx[i-1])
    }
    if strx == string(s) {
        return true
    }
    return false
}
```



##### 方法二：数学方法

```go
func isPalindrome(x int) bool {
    //将负数和大于零且个位数为0数排除
    if x < 0 || (x != 0 && x % 10 == 0 ){
        return false
    }

    revertedNumber := 0
    for x > revertedNumber {
        revertedNumber = revertedNumber * 10 + x % 10
        x /= 10
    }

    return revertedNumber == x || revertedNumber / 10 == x

}
```





### 罗马数字转整数

罗马数字包含以下七种字符: `I`， `V`， `X`， `L`，`C`，`D` 和 `M`。

```
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
```

例如， 罗马数字 `2` 写做 `II` ，即为两个并列的 1 。`12` 写做 `XII` ，即为 `X` + `II` 。 `27` 写做 `XXVII`, 即为 `XX` + `V` + `II` 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 `IIII`，而是 `IV`。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 `IX`。这个特殊的规则只适用于以下六种情况：

- `I` 可以放在 `V` (5) 和 `X` (10) 的左边，来表示 4 和 9。
- `X` 可以放在 `L` (50) 和 `C` (100) 的左边，来表示 40 和 90。 
- `C` 可以放在 `D` (500) 和 `M` (1000) 的左边，来表示 400 和 900。

给定一个罗马数字，将其转换成整数。

 

**示例 1:**

```
输入: s = "III"
输出: 3
```

**示例 2:**

```
输入: s = "IV"
输出: 4
```

**示例 3:**

```
输入: s = "IX"
输出: 9
```

**示例 4:**

```
输入: s = "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
```

**示例 5:**

```
输入: s = "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.
```

 

**提示：**

- `1 <= s.length <= 15`
- `s` 仅含字符 `('I', 'V', 'X', 'L', 'C', 'D', 'M')`
- 题目数据保证 `s` 是一个有效的罗马数字，且表示整数在范围 `[1, 3999]` 内
- 题目所给测试用例皆符合罗马数字书写规则，不会出现跨位等情况。
- IL 和 IM 这样的例子并不符合题目要求，49 应该写作 XLIX，999 应该写作 CMXCIX 。
- 关于罗马数字的详尽书写规则，可以参考 [罗马数字 - Mathematics ](https://b2b.partcommunity.com/community/knowledge/zh_CN/detail/10753/罗马数字#knowledge_article)。



#### 题解：

##### 方法一：

* 情况一：正常情况下，罗马数字的左边大于右边，我们把每一个字符看作是一个数字，使用map来做映射，然后使用key对应的value进行叠加即可，例如：```XXVII``` 就是：X + X + V + I + I = 10+10+5+1+1 =  27。

* 情况二：特殊情况，罗马数字中小的数字在大的数字的左边，我们需要在遍历是需要判断相邻s[i]和s[i+1]对应的值是否满足：V1 < V2, 如果满足我们需要将小的数减去，在下一次遍历时，再加上对应值即可， 例如：

  ```XIV```: X - I + V = 10 -  1 + 5 = 14

```go
func romanToInt(s string) int {
	m := map[string]int{
		"I":1,
		"V":5,
		"X":10,
		"L":50,
		"C":100,
		"D":500,
		"M":1000,
	}
		total := 0
    for i := 0; i < len(s); i++ {
        //len(s)-1特殊情况最多倒数第二个字符并且防止i+1越界
        if i < len(s)-1 && m[string(s[i])] < m[string(s[i+1])] {
            total -= m[string(s[i])]
        }else {
            total += m[string(s[i])]
        }
    }
    return total
}
```



### 最长公共前缀

编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 `""`。

 

**示例 1：**

```
输入：strs = ["flower","flow","flight"]
输出："fl"
```

**示例 2：**

```
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
```

 

**提示：**

- `1 <= strs.length <= 200`
- `0 <= strs[i].length <= 200`
- `strs[i]` 仅由小写英文字母组成



#### 题解：

##### 方法一：纵向法

这里我们以[]string中任意一个字符串为参考k，然后用k中的第一个元素依次去和[]string中的所有字符串的第一个字符比较，如果相等依次类推的进入下一次比较，例如：```strs = ["flower","flow","flight"]```

* 以s1 = ```flower```为参考，将s1[0]去和s2[0]比较->相同; 然后用s1[0]和s3[0]比较->相同。
* 将s1[1]去和s2[1]比较->相同; 然后用s1[1]和s3[1]比较->相同。
* 将s1[2]去和s2[2]比较->相同; 然后用s1[2]和s3[2]比较->不相同，直接返回```s[:2] = fl```

这里需要注意的是：我们以k为参考，但是如果出现len(k) < len(s2)或者len(s3)的情况，也需要将s[:n]直接返回，原因是，最长公共前缀最长为[]string中的最短字符串。

```go
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    for i := 0; i < len(strs[0]); i++ {
        for j := 1; j < len(strs); j++ {
            if  i == len(strs[j]) ||strs[j][i] != strs[0][i] {   //i == len(strs[j])找出最短字符串直接返回
                return strs[0][:i]
            }
        }
    }
    return strs[0]
}
```





