package mcts // 包名是 mcts

import (
	"kaminotte/game" // 引入游戏规则包
)

// MCTSNode 代表搜索树上的一个节点
type MCTSNode struct {
	// 1. 基础关系信息 (树的骨架)
	Parent   *MCTSNode   // 我的父亲是谁？(根节点的父亲是 nil)
	Children []*MCTSNode // 我有哪些孩子？(下一步的所有可能走法)

	// 2. 决策数据 (用来判断这步棋好不好)
	Visits int     // 这个节点被访问了多少次？(N)
	Wins   float64 // 经过这个节点赢了多少次？(Q)
	// 注意：Wins 用 float64 是为了以后能处理平局(0.5分)或者更复杂的评分

	// 3. 游戏状态信息 (当前局面)
	Board      game.Board // 当前棋盘长什么样
	Move       game.Point // 是哪一步棋导致变成了这个局面？(记录一下 x,y)
	NextPlayer int        // 接下来轮到谁下？
}

// NewNode 创建一个新节点 (构造函数)
func NewNode(board game.Board, parent *MCTSNode, move game.Point, nextPlayer int) *MCTSNode {
	return &MCTSNode{
		Board:      board,
		Parent:     parent,
		Move:       move,
		NextPlayer: nextPlayer,
		Children:   make([]*MCTSNode, 0), // 初始化一个空的孩子列表
	}
}