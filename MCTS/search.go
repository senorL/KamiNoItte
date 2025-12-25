package mcts

import (
	"fmt"
	"kaminotte/game"
	"math"
	"math/rand"
)

// UCB1ExplorationConstant 是 UCB 公式里的 C 值，通常取 1.414 (根号2)
// 这个值决定了 AI 是更喜欢"稳扎稳打"(选胜率高的) 还是 "勇于探索"(选下得少的)
const UCB1ExplorationConstant = 1.414

// MCTSSearch 是主入口：给我一个局面，我算 iterations 次，告诉你下一步最好走哪
func MCTSSearch(rootBoard game.Board, nextPlayer int, iterations int) game.Point {
	// 1. 创建根节点
	// 注意：根节点的父节点是 nil，上一手棋我们不知道也可以填空
	root := NewNode(rootBoard, nil, game.Point{X: -1, Y: -1}, nextPlayer)

	// 2. 循环 iterations 次，不断让树变大
	for i := 0; i < iterations; i++ {
		// --- 你的任务：把这 4 个步骤串起来 ---
		
		// Step 1: Selection (选择)
		// 从根节点一直往下找，直到找到一个"边缘节点"(还没完全扩展的节点)
		node := selectNode(root)

		// Step 2: Expansion (扩展)
		// 如果这个节点还没结束游戏，给它长出一个新的子节点
		child := expand(node)

		// Step 3: Simulation (模拟)
		// 从这个新子节点开始，随机瞎下直到终局，看谁赢
		winner := simulate(child)

		// Step 4: Backpropagation (回溯)
		// 把模拟结果(赢/输)告诉这一路上的所有父节点
		backpropagate(child, winner)
	}

	// 3. 搜索结束，找出根节点下面访问次数(Visits)最多的那个孩子，就是我们的一手
	bestChild := getBestChild(root)
	return bestChild.Move
}

// -----------------------------------------------------------------------
// 下面是需要你补充完整的 4 个核心函数
// -----------------------------------------------------------------------

// Step 1: Selection
// 一直往下走，直到找到一个叶子节点
func selectNode(node *MCTSNode) *MCTSNode {
	// 提示：
	// 1. 如果 node 还有孩子没生出来 (比如棋盘有空位，但 Children 列表还没满)，
	//    那它就是我们要找的边缘节点，直接返回 node。
	// 2. 如果 node 的孩子全都生齐了，我们就得用 UCB 公式挑一个"最值得探索"的孩子，
	//    然后让那个孩子继续 selectNode (递归)。
	// 3. 如果 node 已经是终局了(没法再走了)，也直接返回 node。

	// TODO: 你的代码
	return node // 占位符
}

// Step 2: Expansion
// 挑选一个还没尝试过的动作，生成一个新的子节点
func expand(node *MCTSNode) *MCTSNode {
	// 提示：
	// 1. 获取所有合法的空位 (GetEmptyPoints)
	// 2. 检查哪些空位是 Children 里没有的 (也就是还没走过的路)
	// 3. 挑一个新路，创建 NewNode
	// 4. 把新节点加入到 node.Children
	// 5. 返回这个新节点

	// TODO: 你的代码
	return node // 占位符
}

// Step 3: Simulation
// 就是你之前写的 SimulateGame，逻辑基本一样
func simulate(node *MCTSNode) int {
	// 提示：
	// 1. 拿到 node.Board 的副本 (注意不要改坏了树上的节点)
	// 2. 拿到 node.NextPlayer
	// 3. 死循环：随机下子 -> 判赢 -> 换人
	// 4. 返回赢家 (1 或 2，平局 0)

	// TODO: 你的代码
	return 0 // 占位符
}

// Step 4: Backpropagation
// 从 leaf 开始，一直往上找 Parent，更新 Visits 和 Wins
func backpropagate(node *MCTSNode, winner int) {
	// 提示：
	// 1. 这是一个循环，直到 node == nil (根节点的父亲) 结束
	// 2. node.Visits 加 1
	// 3. 如果 winner 也就是 node.Parent.NextPlayer (谁下的这步棋导致了这个局面)，
	//    那么 node.Wins 加 1 (或者加 0.5 如果平局)
	//    注意：这里涉及到一个视角切换的问题，MCTS通常记录的是"这步棋对上一手下棋的人来说是否是好棋"

	// TODO: 你的代码
}

// -----------------------------------------------------------------------
// 辅助函数
// -----------------------------------------------------------------------

// 计算 UCB1 值：衡量一个节点及其父节点的潜力
// 公式：(Wins / Visits) + C * Sqrt(Log(Parent.Visits) / Visits)
func calculateUCB(node *MCTSNode) float64 {
	// TODO: 你的代码 (这是纯数学公式)
	UCB1 := (node.Wins / float64(node.Visits)) + UCB1ExplorationConstant * math.Sqrt(math.log(node.Parent.Visits) / float64(node.Visits))
	// 需要用到 math.Sqrt 和 math.Log
	return UCB1
}

// 在根节点的所有孩子里，找 Visits 最大的那个
func getBestChild(root *MCTSNode) *MCTSNode {
	// TODO: 你的代码 (简单的找最大值算法)
	max_visits := -1
	var max_children *MCTSNode
	for root.Children != nil {
		if root.Children.Visits > max_visits {
			max_visits = root.Children.Visits
			max_children = root.Children
		}
	}
	return max_children
}

// 辅助：判断这个节点是不是还能生孩子 (IsFullyExpanded)
func isFullyExpanded(node *MCTSNode) bool {
	// 检查 board.GetEmptyPoints 的数量 和 len(node.Children)
	// 如果相等，说明所有路都走过了
	
	return false 
}