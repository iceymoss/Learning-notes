package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Player 表示玩家的数据结构
type Player struct {
	ID       int
	Username string
	Level    int
	HP       int
	Conn     *websocket.Conn
}

// Game 表示游戏的数据结构
type Game struct {
	Player1 *Player
	Player2 *Player
}

// Match 表示匹配功能的数据结构
type Match struct {
	mutex sync.Mutex
	queue []*Player
}

// NewMatch 创建一个新的匹配实例
func NewMatch() *Match {
	return &Match{
		queue: make([]*Player, 0),
	}
}

// AddPlayer 将玩家添加到匹配队列
func (m *Match) AddPlayer(player *Player) {
	m.mutex.Lock()
	m.queue = append(m.queue, player)
	m.mutex.Unlock()
	fmt.Printf("玩家 %s 加入匹配队列\n", player.Username)
	fmt.Println("匹配队列长度：", len(m.queue))
	if len(m.queue) >= 2 {
		m.startGame()
	}
}

// startGame 开始游戏
func (m *Match) startGame() {
	if len(m.queue) >= 2 {
		m.mutex.Lock()
		player1 := m.queue[0]
		player2 := m.findClosestLevelPlayer(player1)
		m.mutex.Unlock()
		if player2 != nil {
			fmt.Printf("玩家 %s 与玩家 %s 的对决开始！\n", player1.Username, player2.Username)

			game := &Game{
				Player1: player1,
				Player2: player2,
			}
			game.startGame()
		}

		// 从匹配队列中移除已匹配的玩家
		m.mutex.Lock()
		m.queue = m.queue[2:]
		m.mutex.Unlock()
	}
}

// findClosestLevelPlayer 找到与给定玩家最接近等级的玩家
func (m *Match) findClosestLevelPlayer(player *Player) *Player {
	closestLevel := player.Level
	var closestPlayer *Player

	for _, p := range m.queue {
		if p != player {
			levelDiff := abs(player.Level - p.Level)
			if closestPlayer == nil || levelDiff < closestLevel {
				closestLevel = levelDiff
				closestPlayer = p
				fmt.Println("加入的玩家：", closestPlayer)
			}
		}
	}

	return closestPlayer
}

// abs 返回x的绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// startGame 开始游戏逻辑
func (g *Game) startGame() {
	// 在这里实现游戏的逻辑
	fmt.Printf("玩家 %s 与玩家 %s 的对决中...\n", g.Player1.Username, g.Player2.Username)
	// ...
	Msg1 := make(chan string)
	Msg2 := make(chan string)
	go func() {
		for {
			_, msg1, err := g.Player1.Conn.ReadMessage()
			if err != nil {
				fmt.Printf("接收信息失败：%s", err)
			}
			fmt.Printf("接受到%s消息：%s\n", g.Player1.Username, string(msg1))
			Msg1 <- string(msg1)
			_, msg2, err := g.Player2.Conn.ReadMessage()
			if err != nil {
				fmt.Printf("接收信息失败：%s", err)
			}
			fmt.Printf("接受到%s消息：%s\n", g.Player1.Username, string(msg2))
			Msg2 <- string(msg2)
		}
	}()

	go func() {
		for {
			msg1, _ := json.Marshal(<-Msg1)
			g.Player2.Conn.WriteMessage(1, msg1)
			msg2, _ := json.Marshal(<-Msg2)
			g.Player1.Conn.WriteMessage(1, msg2)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 10)
			msg1, _ := json.Marshal("系统公告：我是小黑子")
			g.Player2.Conn.WriteMessage(1, msg1)
			msg2, _ := json.Marshal("系统公告：我是小黑子")
			g.Player1.Conn.WriteMessage(1, msg2)
		}
	}()
	// 模拟游戏时间
	time.Sleep(5 * time.Second)

	// 游戏结束
	fmt.Printf("玩家 %s 与玩家 %s 的对决结束！\n", g.Player1.Username, g.Player2.Username)
}

func main() {
	match := NewMatch()

	http.HandleFunc("/match", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
		if err != nil {
			fmt.Println("websocket upgrade error:", err)
			return
		}

		rand.Seed(time.Now().Unix())
		// 生成一个随机整数
		randomInt := rand.Intn(100) // 生成0到99之间的随机整数
		fmt.Println("随机整数:", randomInt)
		// 假设连接时传入玩家的ID、用户名和等级
		player := &Player{
			ID:       randomInt,
			Username: fmt.Sprintf("玩家%d", randomInt),
			Level:    10,
			HP:       100,
			Conn:     conn,
		}

		match.AddPlayer(player)
	})
	http.ListenAndServe(":8080", nil)
}
