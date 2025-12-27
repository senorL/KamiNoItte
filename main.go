package main

import (
	"fmt"
	"math/rand"
	"time"

	mcts "kaminotte/MCTS"
	"kaminotte/game"
)

const MCTScount = 10000

func CVC() int {
	board := game.Board{}
	currentPlayer := 1

	fmt.Println("=== 机机对战开始 ===")
	for {
		// board.Show()
		// time.Sleep(time.Millisecond * 200) // 2. 稍微停顿，方便观察

		emptyPoints := board.GetEmptyPoints()
		if len(emptyPoints) == 0 {
			// fmt.Println("平局！棋盘已满。")
			return 0 // 3. 直接返回，不带数值
		}
		// 使用 MCTS 决策当前玩家的落子
		move := mcts.MCTSSearch(board, currentPlayer, MCTScount)
		// 防御：若返回无效点，则退回随机
		if move.X < 0 || move.Y < 0 {
			idx := rand.Intn(len(emptyPoints))
			move = emptyPoints[idx]
		}

		// 使用封装好的 PlaceStone 保证逻辑一致
		ok, _ := board.PlaceStone(move.X, move.Y, currentPlayer)
		if !ok {
			continue
		}

		if board.CheckWin(move.X, move.Y, currentPlayer) {
			// board.Show()
			// fmt.Printf("🎉 电脑 %d 获胜！\n", currentPlayer)
			return currentPlayer
		}
		currentPlayer = 3 - currentPlayer
	}
}

// func ComputerPlay(board game.Board) game.Point {
// 	emptyPoints := board.GetEmptyPoints()
// 	// 防御性编程：如果没有空位了（虽然理论上不会走到这）
// 	if len(emptyPoints) == 0 {
// 		return game.Point{X: -1, Y: -1}
// 	}
// 	move := emptyPoints[rand.Intn(len(emptyPoints))]
// 	return move
// }

func PVP() {
	board := game.Board{}
	currentPlayer := 1

	for {
		board.Show()
		fmt.Printf("轮到玩家 %d，请输入坐标 (x y): ", currentPlayer)

		var x, y int
		_, scanErr := fmt.Scan(&x, &y)
		if scanErr != nil {
			fmt.Println("输入错误，请输入两个整数！")
			// 4. 清理输入缓冲区，防止非法字符导致死循环
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		ok, msg := board.PlaceStone(x, y, currentPlayer)
		if !ok {
			fmt.Printf("无效落子: %s\n", msg)
			continue
		}

		if board.CheckWin(x, y, currentPlayer) {
			board.Show()
			fmt.Printf("恭喜！玩家 %d 获胜了！\n", currentPlayer)
			break
		}

		currentPlayer = 3 - currentPlayer
		fmt.Println("-----------------------")
	}
}

func PVC() {
	board := game.Board{}
	currentPlayer := 1

	fmt.Print("输入您为先手(1)还是后手(2): ")
	var player int
	fmt.Scan(&player)

	for {
		board.Show()
		var x, y int

		if currentPlayer == player {
			fmt.Printf("轮到玩家 %d，请输入坐标 (x y): ", currentPlayer)
			_, scanErr := fmt.Scan(&x, &y)
			if scanErr != nil {
				fmt.Println("输入错误，请输入两个整数！")
				var discard string
				fmt.Scanln(&discard)
				continue
			}
		} else {
			fmt.Println("电脑思考中...")
			time.Sleep(time.Second) // 增加代入感
			// move := ComputerPlay(board)
			move := mcts.MCTSSearch(board, currentPlayer, MCTScount)
			x, y = move.X, move.Y
		}

		ok, msg := board.PlaceStone(x, y, currentPlayer)
		if !ok {
			fmt.Printf("无效落子: %s\n", msg)
			continue
		}

		if board.CheckWin(x, y, currentPlayer) {
			board.Show()
			if currentPlayer == player {
				fmt.Printf("恭喜！您(玩家 %d) 获胜了！\n", currentPlayer)
			} else {
				fmt.Printf("😭 电脑获胜！再接再厉。\n")
			}
			break
		}

		currentPlayer = 3 - currentPlayer
		fmt.Println("-----------------------")
	}
}

func oldmain() {
	// 5. 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	fmt.Print(`
    请选择模式：
    1. 人人对战 (PVP)
    2. 人机对战 (PVC)
    3. 机机对战 (CVC)
    请输入数字: `)

	var playMode int
	fmt.Scan(&playMode)

	switch playMode {
	case 1:
		PVP()
	case 2:
		PVC()
	case 3:
		c1, c2, nowiner := 0, 0, 0
		for i := 0; i < 10; i++ {
			switch CVC() {
			case 1: c1++
			case 2: c2++
			default: nowiner++
			}
		}
		fmt.Printf(`
		computer1胜%d;
		computer2胜%d;
		平局%d;
		`, c1, c2, nowiner)
	default:
		fmt.Println("错误的输入，程序退出。")
	}
}
