[toc]

### 搜索插入位置

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 `O(log n)` 的算法。

 

**示例 1:**

```
输入: nums = [1,3,5,6], target = 5
输出: 2
```

**示例 2:**

```
输入: nums = [1,3,5,6], target = 2
输出: 1
```

**示例 3:**

```
输入: nums = [1,3,5,6], target = 7
输出: 4
```

 

**提示:**

- `1 <= nums.length <= 104`
- `-104 <= nums[i] <= 104`
- `nums` 为 **无重复元素** 的 **升序** 排列数组
- `-104 <= target <= 104`



#### 题解：

##### 方法一：二分查找法

在一个数组中查找一个指定元素，第一时间想到的就是二分查找法，使用这里使用二分查找法来解决

情况一：target存在nums中，直接使用二分查找可以快速找target。

情况二：如果target不做nums中，需要将target插入合适位置，并返回target的索引，这里可以发现一个规律，```nums[0] < target < nums[len()-1]```的情况,当target < nums[mid] 时，r = mid-1 即target应该插入nums[0] ~ nums[mid-1]之间，当target > nums[mid]时，l = mid + 1, 即target应该插入nums[mid+1] ~ nums[r]之间,但其实最后，插入的位置其实就是l； 这里有两种出界的情况：当nums[0] > target时，超出左边界，经过二分查找查找范围一直是往左缩减的，原l没有变化，所以直接返回l即可； 当nums[len(nums)-1] < target时，超出右界，经过二分查找查找范围一直是往又缩减的，原r没有变化，但是最终r == l 时的一次遍历，l = mid + 1, 就是超过r,所以直接返回l即可。

```go
func searchInsert(nums []int, target int) int {
    //二分查找法
    l := 0
    r := len(nums)-1
 
    for l <= r {
         mid := (r - l) >> 1 + l
        if nums[mid] == target {
            return mid
        }else if nums[mid] < target {
            l = mid+1
        }else {
            r = mid-1
        }
    }
    return l
}
```





### 最后一个单词的长度

给你一个字符串 `s`，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 **最后一个** 单词的长度。

**单词** 是指仅由字母组成、不包含任何空格字符的最大子字符串。

 

**示例 1：**

```
输入：s = "Hello World"
输出：5
解释：最后一个单词是“World”，长度为5。
```

**示例 2：**

```
输入：s = "   fly me   to   the moon  "
输出：4
解释：最后一个单词是“moon”，长度为4。
```

**示例 3：**

```
输入：s = "luffy is still joyboy"
输出：6
解释：最后一个单词是长度为6的“joyboy”。
```

 

**提示：**

- `1 <= s.length <= 104`
- `s` 仅有英文字母和空格 `' '` 组成
- `s` 中至少存在一个单词



#### 题解：

##### 方法一：暴力枚举

看到题目后第一个思路就是直接使用从字符串尾部暴力枚举，这里需要考虑的几种情况：首先我们需要将遍历的字符都存储进一个容器这里使用slice：

1. ```s = "Hello World"``` 

   这种情况是好解决的，直接从末尾对字符串遍历，直到找到```" "```。

2. ```s = "   fly me   to   the moon  "```

   这情况就比较复杂，最后一个单词前后都有```" "``` ， 这里可以声明一个bool变量，isEmpty = true， 只要是isEmpty = true的情况都可以说是从遍历开始存在```" "``` 当遍历到真正的字符时将isEmpty = true设置为false, 当下一次发现```" "```时，就说明我们找出了最后一个单词。

3. 时间上有较高效率，但是空间消耗是比较大的。

```go
func lengthOfLastWord(s string) int {
    EmptyLen := 0    //用于记录空格的长度
    isEmpty := true  //用于识别空格
    n := len(s)
    var result []string   //用户存储遍历的内容
    for i := n-1; i >= 0; i--{
        v := fmt.Sprintf("%c", s[i])  //转为字符
        if v != " " {
            result = append(result, v)
            isEmpty = false
        }else{
            if isEmpty {
                EmptyLen++
                result = append(result, v)
                continue
            }
            return len(result)-EmptyLen   //总遍历数-遍历空格数 = 最后一个单词的长度
        }
    }
    return len(result)-EmptyLen
}
```



