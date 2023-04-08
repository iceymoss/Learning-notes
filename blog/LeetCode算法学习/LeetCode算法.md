[toc]



## 112、路径总和

给你二叉树的根节点 `root` 和一个表示目标和的整数 `targetSum` 。判断该树中是否存在 **根节点到叶子节点** 的路径，这条路径上所有节点值相加等于目标和 `targetSum` 。如果存在，返回 `true` ；否则，返回 `false` 。

**叶子节点** 是指没有子节点的节点。

 

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/01/18/pathsum1.jpg)

```
输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出：true
解释：等于目标和的根节点到叶节点路径如上图所示。
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg)

```
输入：root = [1,2,3], targetSum = 5
输出：false
解释：树中存在两条根节点到叶子节点的路径：
(1 --> 2): 和为 3
(1 --> 3): 和为 4
不存在 sum = 5 的根节点到叶子节点的路径。
```

**示例 3：**

```
输入：root = [], targetSum = 0
输出：false
解释：由于树是空的，所以不存在根节点到叶子节点的路径。
```

 

**提示：**

- 树中节点的数目在范围 `[0, 5000]` 内
- `-1000 <= Node.val <= 1000`
- `-1000 <= targetSum <= 1000`



### 题解

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    //从root节点开始
    return dfs(root, targetSum, 0)

}


func dfs(node *TreeNode, tag int, item int) bool {
    //进入当前层，将当前节点只相加
    item += node.Val
    //递归终止条件: 必须为叶子节点
    if node.Left == nil && node.Right == nil && item == tag {
        return true
    }
    
    //进入下一层时，需要判断当前层的子节点不为空
    var isHavL, isHavR bool
    if node.Left != nil {
        isHavL = dfs(node.Left, tag, item)
    }
    
    if node.Right != nil {
         isHavR = dfs(node.Right, tag, item)   
    }
    
    //如果找到结果，就返回true
    if isHavL {
        return true
    }

    if isHavR {
        return true
    }
    //没有找到返回false
    return false    
}
```





## 125、验证回文串

如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 **回文串** 。

字母和数字都属于字母数字字符。

给你一个字符串 `s`，如果它是 **回文串** ，返回 `true` ；否则，返回 `false` 。

 

**示例 1：**

```
输入: s = "A man, a plan, a canal: Panama"
输出：true
解释："amanaplanacanalpanama" 是回文串。
```

**示例 2：**

```
输入：s = "race a car"
输出：false
解释："raceacar" 不是回文串。
```

**示例 3：**

```
输入：s = " "
输出：true
解释：在移除非字母数字字符之后，s 是一个空字符串 "" 。
由于空字符串正着反着读都一样，所以是回文串。
```

 

**提示：**

- `1 <= s.length <= 2 * 105`
- `s` 仅由可打印的 ASCII 字符组成

### 题解

```go
func isPalindrome(s string) bool {
    //难点，如何将字符串脱非数字字母字符
  	//将字符串转为小写
    str := strings.ToLower(s)
  	//双指针法
    l, r := 0, len(str)-1
    for l <= r {
        if !isVail(str[l]) {
            l++
            continue
        }
        if !isVail(str[r]) {
            r--
            continue
        }
        if str[l] != str[r] {
            return false
        }
        l++
        r--
    }
    return true
}

//判断字符是否有效
func isVail(b byte) bool {
    if ('a' <= b && b <= 'z') || ('0' <= b && '9' >= b) {
        return true
    }
    return false
}
```





## 136、只出现一次的数字

给你一个 **非空** 整数数组 `nums` ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。

 

**示例 1 ：**

```
输入：nums = [2,2,1]
输出：1
```

**示例 2 ：**

```
输入：nums = [4,1,2,1,2]
输出：4
```

**示例 3 ：**

```
输入：nums = [1]
输出：1
```

 

**提示：**

- `1 <= nums.length <= 3 * 104`
- `-3 * 104 <= nums[i] <= 3 * 104`
- 除了某个元素只出现一次以外，其余每个元素均出现两次。





### 题解

```go
func singleNumber(nums []int) int {
    //使用hashe表
    m := make(map[int]int)
    var res int
  	//统计次数
    for i := 0; i < len(nums); i++ {
        if _, ok := m[nums[i]]; !ok {
            m[nums[i]] = 1
        }else {
            m[nums[i]] += 1
        }
    }
    
  	//判断结果
    for i := 0; i < len(nums); i++{
        if m[nums[i]] == 1 {
            res = nums[i]
            break
        }
    }
    return res
}
```

