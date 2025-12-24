package main

import (
	"fmt"
	"math/rand"
	"time"
    // 引入你的 game 包
    // 如果你在本地开发，通常是 "项目名/game" 或者 "./game"
    // 假设你的 go.mod 名字是 "kaminotte"，那么这里是 "kaminotte/game"
    "kaminotte/game" 
)

// SimulateGame 现在变得很简单
// 注意：参数类型变成了 game.Board
func SimulateGame(board game.Board, currentPlayer int) int {
	simPlayer := currentPlayer
	for {
		// 调用方法：board.GetEmptyPoints()
		emptyPoints := board.GetEmptyPoints()
		if len(emptyPoints) == 0 {
			return 0
		}
		
		idx := rand.Intn(len(emptyPoints))
		move := emptyPoints[idx]
		
		// 这里的 board 是副本，直接改没事
		// 注意 Move.X (大写)
		board[move.X][move.Y] = simPlayer

		if board.CheckWin(move.X, move.Y, simPlayer) {
			return simPlayer
		}
		simPlayer = 3 - simPlayer
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 初始化一个空棋盘
	// board := game.Board{} 

	fmt.Println("正在进行 10000 次极速模拟测试...")
	start := time.Now()
	blackWins := 0
	
	for i := 0; i < 10000; i++ {
		// 传入空棋盘副本
		winner := SimulateGame(game.Board{}, 1)
		if winner == 1 {
			blackWins++
		}
	}
	
	fmt.Printf("耗时: %v\n", time.Since(start))
	fmt.Printf("黑棋胜场: %d\n", blackWins)
	
	// board.Show()
}