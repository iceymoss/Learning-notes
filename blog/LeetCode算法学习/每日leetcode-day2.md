[toc]



### 有效括号

给定一个只包括 `'('`，`')'`，`'{'`，`'}'`，`'['`，`']'` 的字符串 `s` ，判断字符串是否有效。

有效字符串需满足：

1. 左括号必须用相同类型的右括号闭合。
2. 左括号必须以正确的顺序闭合。
3. 每个右括号都有一个对应的相同类型的左括号。

 

**示例 1：**

```
输入：s = "()"
输出：true
```

**示例 2：**

```
输入：s = "()[]{}"
输出：true
```

**示例 3：**

```
输入：s = "(]"
输出：false
```

 

**提示：**

- `1 <= s.length <= 104`

- `s` 仅由括号 `'()[]{}'` 组成

  

#### 题解

##### 方法一：使用辅助-栈

第一次做这道题时，想了10来分钟，没有任何思路，不知道怎么实现，后面看了看题解 ，可以用栈来完成这道题，所以在此之前必须要知道栈是什么：一种基础数据结构，最大的特点就是后进先出

###### 思路

这里首先要知道的是，应对括号必须是一个一对一的映射关系，所以这里需要使用map，将```map[')'] = ( ```然后我们使用栈，从第一个字符开始遍历，只要是左括号，都将压入栈中，当遍历到右括号时，就将栈顶元素拿出和此时遍历到的字符对比，如果相同，则说明目前遍历到的括号是有效的，然后进入下一次遍历，按照此逻辑以此类推； 如果不相同，就可以说明，当前遍历到的括号一定不是有效括号，直接返回false。

```go
func isValid(s string) bool {
    //字符串为奇数
    if len(s) % 2 == 1 {
        return false
    }

    bt := []byte(s)
    //栈
    stack := []byte{}
    m := map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}
    for i := 0; i < len(s); i++ {
        //左括号入栈
        if _, ok := m[bt[i]]; !ok {
            stack = append(stack, bt[i])
        }else{
            if len(stack) > 0 {
                stack_len := len(stack)
                target := stack[stack_len-1]
                if m[bt[i]] != target || len(stack) == 0 {
                    return false
                }

                stack = stack[:len(stack)-1] 
            }else{
                return false
            }
            
        }
    }

    return len(stack) == 0

}
```





### 合并链表

将两个升序链表合并为一个新的 **升序** 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

 

**示例 1：**

![img](https://assets.leetcode.com/uploads/2020/10/03/merge_ex1.jpg)

```
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
```

**示例 2：**

```
输入：l1 = [], l2 = []
输出：[]
```

**示例 3：**

```
输入：l1 = [], l2 = [0]
输出：[0]
```

 

**提示：**

- 两个链表的节点数目范围是 `[0, 50]`
- `-100 <= Node.val <= 100`
- `l1` 和 `l2` 均按 **非递减顺序** 排列



#### 题解

##### 方法一：迭代发

使用迭代法，当两个链表都不为nil 时，通过迭代的方法对两个链表往后遍历， 将两个链表的数据拿出依次对比；这里其实就是跟使用链表的操作差不多； 如果两个链表中为nil 时，那么直接将另一个不为nil的链表直接追加到结果链表后(因为这里是有序链表)。

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    //结果链表
    result := &ListNode{}
    // result := new(ListNode)
    cur_result := result
    
    //都不为空的情况
    for list1 != nil && list2 != nil {
        if list1.Val <=  list2.Val {
            cur_result.Next = list1
            list1 = list1.Next
            cur_result = cur_result.Next
        }else {
            cur_result.Next = list2
             list2 = list2.Next
            cur_result = cur_result.Next
        }
    }

    	//list1为空
	if list1 == nil {
		cur_result.Next = list2
	}

	//list2为空
	if list2 == nil {
		cur_result.Next = list1
	}

    return result.Next
}
```









