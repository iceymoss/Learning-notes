[toc]



###  广度优先算法

迷宫的广度优先搜索是基于广度优先算法来实现的，在爬虫领域也会经常使用广度优先算法，首先来了解一下广度优先算法。
我在之前对文章中具体介绍了广度优先算法，可参见[图论系列之[「广度优先遍历及无权图的最短路径(ShortPath)」](https://learnku.com/articles/57541),这里我们需要使用一个辅助队列，用来存放访问过的节点，下图是我们的迷宫实例(6行5列）：

![「Golang成长之路」迷宫的广度优先搜索](https://cdn.learnku.com/uploads/images/202110/15/69310/iWEaRmcxvh.png!large)
从左上角为入口，右下角为出口，数字1表示墙(走不通)，数字0表示通道，这里我们需要规定节点的遍历方向，即：上左下右(逆时针)，下面来具体介绍广度优先算法在迷宫中的逻辑：
1. 当前节点为0(入口)，将该0放入辅助队列中，然后拿出节点0并标记，再将该0节点相邻的节点以上左下右(逆时针)的顺序放入队列中，此时的队列中，***队列：0、1***

2. 接着从队列中拿出0，在将该0相邻的节点(0、0）放入队列中，此时的***队列：1、0、0***

3. 接着再拿出队列中的第一个节点"1",但节点1是墙，将其舍弃此时的***队列：0、0***
  4.再拿出队列中的第一个节点"0",然后将0节点相邻的节点1、1放入队列中，此时的***队列：0、1、1***
  5接着再拿出队列中第一个节点“0”，而该节点相邻的节点有：1、0、1、0
  此时只有该0节点的右邻节点未被访问，即：0，放入队列，此时的***队列：0、1、1、1***
  ……
  ……
  以此类推，就可找到出口

  
### 文件读入

我们需要将：
```go
6 5
0 1 0 0 0
0 0 0 1 0
0 1 0 1 0
1 1 1 0 0
0 1 0 0 1
0 1 0 0 0
```
读入，并创建一个6行5列的二维slice
程序如下：
```go
//读入文件
func ReadMaze(filename string) [][]int{
   file, err := os.Open(filename)
   if err != nil{
      panic(err)
   }
   var row, col int
  fmt.Fscanf(file, "%d %d", &row, &col)   //fmt.Fscanf()?
  maze := make([][]int, row)
   for i := range maze{
      maze[i]= make([]int, col)
      for j := range maze[i]{
         fmt.Fscanf(file, "%d", &maze[i][j])
      }

   }
   return maze
}
```


### 广度优先算法的实现

我们需要将数字抽象为节点：
```go
type point struct {
	i, j int
}
```
还需要对每一个节点的相邻节点进行访问：
```go
var dirs = [4]point{
	{-1,0}, {0,-1}, {1,0}, {0,1}}
```
这是广度优先算法的核心：
```go
func walk(maze [][]int, start, end point) [][]int{
   //创建迷宫正确路线
  steps := make([][]int, len(maze))
   for i := range steps{
      steps[i] = make([]int, len(maze[i]))

      }
        //创建队列Q
  Q := []point{start}

      for len(Q) > 0 {
         cur := Q[0]
         Q = Q[1: ]

         if cur == end{
            break
  }

         for _, dir := range dirs{
            next := cur.add(dir)

            val, ok := next.at(maze)
            if !ok || val == 1{    //1为墙，ok == fals为越界
  continue
  }
            val, ok = next.at(steps) //判断是否被访问，steps中的值为零表示为墙
  if !ok || val != 0 {
               continue
  }
            if next == start{
               continue
  }

            curSteps, _ := cur.at(steps)
            steps[next.i][next.j] = curSteps + 1

  Q = append(Q, next)

            //maze at next is 0
 //and steps at next is 0 指到过的点 //and next == 0  }
      }
      return steps
   }
```
我们还需要对坐标进行变更新：
```go
//坐标变换
func (p point)add(r point) point{
	return point{p.i + r.i, p.j + r.j}
}
```
在遍历过程中需要检查是否越界：
```go
//判断越界
func (p point)at(grid [][]int)(int, bool){
   if p.i < 0 || p.i >= len(grid) {
      return 0, false
  }
    if p.j < 0 || p.j >= len(grid[p.i]){
       return 0, false
  }
   return grid[p.i][p.j],true
}
```


### 调用

```go
func main() {
   maze := ReadMaze("maze/maze.in")

   steps := walk(maze, point{0,0}, point{len(maze)-1, len(maze[0])-1})

   for i := range steps{
      for j := range steps[i]{
         fmt.Printf("%3d " ,steps[i][j])
      }
      fmt.Println()
   }
 }
```

打印结果：
```go
  0   0   4   5   6 
  1   2   3   0   7 
  2   0   4   0   8 
  0   0   0  10   9 
  0   0  12  11   0 
  0   0  13  12  13 
```