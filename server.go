package main

import (
	"encoding/json"
	mcts "kaminotte/MCTS"
	"kaminotte/game"
	"math/rand"
	"net/http"
	"time"
)

// 全局棋盘状态（简单起见，仅支持单人单局）
var currentBoard = game.Board{}

type MoveRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type MoveResponse struct {
	Board  game.Board `json:"board"`
	Winner int        `json:"winner"`
	Error  string     `json:"error,omitempty"`
}

func handleMove(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	var req MoveRequest
	json.NewDecoder(r.Body).Decode(&req)

	// 1. 玩家落子 (玩家1)
	ok, msg := currentBoard.PlaceStone(req.X, req.Y, 1)
	if !ok {
		json.NewEncoder(w).Encode(MoveResponse{Error: msg})
		return
	}

	// 检查玩家是否获胜
	if currentBoard.CheckWin(req.X, req.Y, 1) {
		json.NewEncoder(w).Encode(MoveResponse{Board: currentBoard, Winner: 1})
		return
	}

	// 2. AI 落子 (玩家2)
	// 使用你已经实现的 MCTSSearch，设定迭代次数为 5000 或更高以增加难度
	aiMove := mcts.MCTSSearch(currentBoard, 2, 10000)
	if aiMove.X != -1 {
		currentBoard.PlaceStone(aiMove.X, aiMove.Y, 2)
		// 检查 AI 是否获胜
		if currentBoard.CheckWin(aiMove.X, aiMove.Y, 2) {
			json.NewEncoder(w).Encode(MoveResponse{Board: currentBoard, Winner: 2})
			return
		}
	}

	// 3. 返回最新局面
	json.NewEncoder(w).Encode(MoveResponse{Board: currentBoard, Winner: 0})
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 托管静态文件 (HTML)
	http.Handle("/", http.FileServer(http.Dir("./web")))

	// API 接口
	http.HandleFunc("/move", handleMove)
	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		currentBoard = game.Board{}
	})

	println("服务器已启动，请访问: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
