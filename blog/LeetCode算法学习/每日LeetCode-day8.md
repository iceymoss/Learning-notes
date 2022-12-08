[toc]



### 相同树

给你两棵二叉树的根节点 `p` 和 `q` ，编写一个函数来检验这两棵树是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

 

**示例 1：**

![img](https://assets.leetcode.com/uploads/2020/12/20/ex1.jpg)

```
输入：p = [1,2,3], q = [1,2,3]
输出：true
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2020/12/20/ex2.jpg)

```
输入：p = [1,2], q = [1,null,2]
输出：false
```

**示例 3：**

![img](https://assets.leetcode.com/uploads/2020/12/20/ex3.jpg)

```
输入：p = [1,2,1], q = [1,1,2]
输出：false
```

 

**提示：**

- 两棵树上的节点数目都在范围 `[0, 100]` 内
- `-104 <= Node.val <= 104`



#### 题解：

##### 方法一：前序遍历

###### 思路

仔细看这道题，可以发现对树的遍历是一个前序遍历的过程，所以这里将比较两棵树就变成了树的前序遍历，可以使用递归或者迭代，这里我先使用递归实现：

```go
func Tarvse(root *TreeNode) {
	if root == nil {
		return
	}
  
  //对节点操作
  fmt.Println(root.Val)
	Tarvse(root.Left)
	Tarvse(root.Right)
}
```

我的第一个思路就是将两颗树，分别进行遍历，然后将遍历到的结果放入Slice中，然后对slice进行比较即可。

###### 实现：

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	//思路：转换为二叉树的前序遍历： 1.递归； 2.迭代
	arrq := make([]int, 0)
	arrp := make([]int, 0)

	Tarvse(p, &arrp)
	Tarvse(q, &arrq)

	if len(arrp) != len(arrq) {
		return false
	}

	for i := 0; i < len(arrp); i++ {
		if arrp[i] != arrq[i] {
			return false
		}
	}
	return true
}

func Tarvse(root *TreeNode, arr *[]int) {
    //递归
	if root == nil {
		return
	}
    
    //左右孩子为空的情况
	if root.Right == nil || root.Right == nil {
		*arr = append(*arr, 0)
	}

	*arr = append(*arr, root.Val)
	Tarvse(root.Left, arr)
	Tarvse(root.Right, arr)
}
```

这个思路简单，但是在时间复杂度和空间复杂度的消耗还是挺大的。

迭代：

```go
func preorder(root *TreeNode) []int { //非递归前序遍历
	res := []int{}
	if root == nil {
		return res
	}
	
	stack := []*TreeNode{} //定义一个栈储存节点信息
	for root != nil || len(stack) != 0 {
		if root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}
```



##### 方法二： 深度优先遍历

###### 思路

同时对两颗树的节点进行遍历：我们需要先捋清楚，会出现那几种情况： 1. 当两树都为空，则一定不相等；2. 有且仅有一棵树为空，则一定不相等；3. 当两颗树都不为空时，需要向判断两树中的值是否相等，然后再判断其结构。

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
    //当两个树为空时，一个不相同
    if p == nil && q == nil {
        return true
    }

    //有且只有一棵树为空时
    if p == nil || q == nil {
        return false
    }

    //当两树都不为空时， 先把判断值是否相等
    if p.Val != q.Val {
        return false
    }
  
 		//使用递归进行深度遍历，在遍历的同时会判断每一个节点的情况
    return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```

这样时间复杂度和空间复杂度相比于方法一有了质的提升

我在LeetCode的题解上看到这个：直呼lNB

```go
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil || q == nil {
        return p == q
    }
    return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
```





### 对称二叉树

给你一个二叉树的根节点 `root` ， 检查它是否轴对称。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/02/19/symtree1.jpg)

```
输入：root = [1,2,2,3,4,4,3]
输出：true
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2021/02/19/symtree2.jpg)

```
输入：root = [1,2,2,null,3,null,3]
输出：false
```

 

**提示：**

- 树中节点数目在范围 `[1, 1000]` 内
- `-100 <= Node.val <= 100`

 

#### 题解

##### 方法一：递归

###### 思路

判断树是否为轴对称， 其实就是判断树的左右子树是否互为镜像，方法是将树的左右子树分为两颗树，判断两棵树是否为镜像树。

要判断树为镜像，需要满足如下情况:

>1. 树根节点相等
>
>2. 每个树的右子树都与另一个树的左子树镜像对称

这里用递归方法实现，通过「同步移动」两个指针的方法来遍历这棵树，指针p和q ，刚开始p和q都指向根节点，随后当q右移时，p左移； 当q左移时，p右移。然后检查q和p指向的节点的值是否相等，如果相等接着递归左右子树是否对称。

###### 实现

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
   return check(root, root)
}

func check(p, q *TreeNode) bool {
    //树为空,刚开始指向同一棵树根节点
    if p == nil && q == nil {
        return true
    }

    //两树结构不对称
    if p == nil || q == nil {
        return false 
    }

    //当前节点值是否相等
    if p.Val != q.Val {
        return false
    }
		
    return check(p.Right, q.Left) && check(p.Left, q.Right)
}

```



##### 方法二：迭代

###### 思路：树的广度优先遍历

使用迭代方法对树进行遍历，使用辅助队列，将子树根节点分为u和v，每次放入对称位置元素队列中（每两个连续的结点应该是相等的，而且它们的子树互为镜像），然后将两个结点的左右子结点按相反的顺序插入队列中。当队列为空时，或者我们检测到树不对称（即从队列中取出两个不相等的连续结点）时，算法结束

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
    u, v := root, root

    //使用辅助队列
    q := []*TreeNode{}
    q = append(q, u)
    q = append(q, v)

    //
    for len(q) > 0 {
        u, v = q[0], q[1]
        q = q[2:]

        //空节点
        if u == nil && v == nil {
            continue
        }

        //结构不对称
        if u == nil || v == nil {
            return false
        }

        //对称位置的值
        if u.Val != v.Val {
            return false
        }

        //连续的添加对称节点
        q = append(q, u.Left)
        q = append(q, v.Right)

        q = append(q, u.Right)
        q = append(q, v.Left) 
    }

    return true 
}
```









