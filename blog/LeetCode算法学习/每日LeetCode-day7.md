### 删除排序链表中的重复元素

给定一个已排序的链表的头 `head` ， *删除所有重复的元素，使每个元素只出现一次* 。返回 *已排序的链表* 。

 

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/01/04/list1.jpg)

```
输入：head = [1,1,2]
输出：[1,2]
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2021/01/04/list2.jpg)

```
输入：head = [1,1,2,3,3]
输出：[1,2,3]
```

 

**提示：**

- 链表中节点数目在范围 `[0, 300]` 内
- `-100 <= Node.val <= 100`
- 题目数据保证链表已经按升序 **排列**



#### 题解

##### 方法一：一次循环

读懂题目后，第一个思路就是直接使用for对其遍历，使用cur := head ， 当遍历的当前节点的下一个节点不为空时cur.Next != nil 然后比较cur.Val和cur.Next.Val是否相等，如果不相等，直接遍历下一个节点：cur = cur.Next ；如果cur.Val和cur.Next.Val相等时，就将cur.Next = cur.Next.Next这样就将重复用cur.Next.Next覆盖了。当cur.Next.Next == nil 时，那么自然cur也就是链表的最后一个元素了。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
    cur := head
    if head == nil {
        return head
    }
     for cur.Next != nil {
         if cur.Val != cur.Next.Val {
             cur = cur.Next
         }else {
             cur.Next = cur.Next.Next
         }
     }
    return head
}
```





### 合并两个有序数组

给你两个按 **非递减顺序** 排列的整数数组 `nums1` 和 `nums2`，另有两个整数 `m` 和 `n` ，分别表示 `nums1` 和 `nums2` 中的元素数目。

请你 **合并** `nums2` 到 `nums1` 中，使合并后的数组同样按 **非递减顺序** 排列。

**注意：**最终，合并后数组不应由函数返回，而是存储在数组 `nums1` 中。为了应对这种情况，`nums1` 的初始长度为 `m + n`，其中前 `m` 个元素表示应合并的元素，后 `n` 个元素为 `0` ，应忽略。`nums2` 的长度为 `n` 。

 

**示例 1：**

```
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
```

**示例 2：**

```
输入：nums1 = [1], m = 1, nums2 = [], n = 0
输出：[1]
解释：需要合并 [1] 和 [] 。
合并结果是 [1] 。
```

**示例 3：**

```
输入：nums1 = [0], m = 0, nums2 = [1], n = 1
输出：[1]
解释：需要合并的数组是 [] 和 [1] 。
合并结果是 [1] 。
注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。
```

 

**提示：**

- `nums1.length == m + n`
- `nums2.length == n`
- `0 <= m, n <= 200`
- `1 <= m + n <= 200`
- `-109 <= nums1[i], nums2[j] <= 109`

 #### 题解

##### 方法一：先追加，后排序

最简单的方法是直接使用将nums2追加到nums1中，然后对nums1排序

```go
func merge(nums1 []int, m int, nums2 []int, n int)  {
    nums1 = nums1[:m]
    nums1 = append(nums1, nums2...)

  	//选择排序
    for i := 0; i < len(nums1);i++ {
        for j := i+1; j < len(nums1); j++ {
            if nums1[i] >=  nums1[j] {
                nums1[i], nums1[j] = nums1[j], nums1[i]
            }
        }
    }
}
```

当然也可以直接使用内置方法copy()：这样时间上得到了优化

```go
func merge(nums1 []int, m int, nums2 []int, n int)  {
    copy(nums1[m:], nums2)
  
   	//这里的排序算法可以直接用：sort.Ints(nums1)
    for i := 0; i < len(nums1);i++ {
        for j := i+1; j < len(nums1); j++ {
            if nums1[i] >=  nums1[j] {
                nums1[i], nums1[j] = nums1[j], nums1[i]
            }
        }
    }
}
```



##### 方法二：双指针法

通过指针p1和p2,来控制slice出现越界，使用双指针方法。这一方法将两个数组看作队列，每次从两个数组头部取出比较小的数字放到结果中。

```go
func merge(nums1 []int, m int, nums2 []int, n int)  {
    //双指针法
    p1, p2 := 0, 0
    result := make([]int, 0, m+n)
    for{
        if p1 == m {
            result = append(result, nums2[p2:]...)
            break
        }
        
        if p2 == n {
            result = append(result, nums1[p1:]...)
            break
        }

        if nums1[p1] <= nums2[p2] {
           result = append(result, nums1[p1])
            p1++
        }else {
           result = append(result, nums2[p2])
            p2++
        }
    }
    copy(nums1, result)
}
```

