package mcts

import (
	"math"
	"math/rand"
	"kaminotte/game"
)

// UCB1ExplorationConstant 是 UCB 公式里的 C 值
// 理论值是 1.414，你可以调整它来看看 AI 性格的变化
const UCB1ExplorationConstant = 1.414

// MCTSSearch 主函数：给我一个局面，我思考 iterations 次，告诉你怎么走
func MCTSSearch(rootBoard game.Board, nextPlayer int, iterations int) game.Point {
	// 1. 创建根节点 (root)
	// 根节点的 Parent 是 nil，Move 是无效值(-1, -1)
	root := NewNode(rootBoard, nil, game.Point{X: -1, Y: -1}, nextPlayer)

	// 2. 循环思考 iterations 次
	for i := 0; i < iterations; i++ {
		// --- MCTS 四部曲 ---

		// Step 1: Selection (选择)
		// 顺着树往下找，直到找到一个"还没被完全探索"的节点
		// (也就是这个节点还有空位没生成子节点，或者是游戏结束了)
		node := selectNode(root)

		// Step 2: Expansion (扩展)
		// 如果游戏没结束，在这个节点上长出一个新的树枝(子节点)
		// 也就是挑一个还没试过的空位走一步
		child := expand(node)

		// Step 3: Simulation (模拟)
		// 从这个新长出来的子节点开始，双方瞎下直到终局
		// 这个你最熟了，就是昨天的 PVP/PVC 逻辑
		winner := simulate(child)

		// Step 4: Backpropagation (回溯)
		// 把模拟的结果(赢/输)告诉这一路上的所有长辈节点
		backpropagate(child, winner)
	}

	// 3. 思考时间到，选择访问次数(Visits)最多的那个孩子作为下一步
	// 注意：这里通常不选胜率最高的，而是选Visits最多的，因为Visits多说明最靠谱
	bestChild := getBestChild(root)
	
	// 防御性编程：如果没有孩子（比如棋盘满了），返回一个特殊值
	if bestChild == nil {
		return game.Point{X: -1, Y: -1}
	}

	return bestChild.Move
}

// -----------------------------------------------------------------------
// 你的作业：请完成下面这 5 个核心函数
// -----------------------------------------------------------------------

// 作业 1: Selection
// 一直往下找 (node = bestChild)，直到 node 是叶子节点，或者 node 还有没尝试过的落子点
func selectNode(node *MCTSNode) *MCTSNode {
	// 提示：
	// 这是一个循环：for !isLeaf(node) && isFullyExpanded(node) { ... }
	// 在循环里，你需要找到 UCB 值最大的那个孩子，然后 node = bestChild
	// 只有当一个节点的所有可能落子都被长出子节点了，我们才用 UCB 去选深一层的
	// 如果这个节点还有空位没试过，那它就是我们要找的边缘，直接返回它，交给 expand 去扩展
	
	// TODO: 你的代码
	for !isLeaf(node) && isFullyExpanded(node) {
        
		bestChild := getBestChild(node)

        node = bestChild 
    }

    // 当跳出循环时，说明我们找到了一个“没满”或者“结束”的节点
    // 这正是 Step 2 (Expansion) 梦寐以求的输入对象
    return node
}

// 作业 2: Expansion
// 挑一个还没长出来的动作，生成一个新的子节点
func expand(node *MCTSNode) *MCTSNode {
	// 提示：
	// 1. 获取当前局面的所有空位 (board.GetEmptyPoints)
	// 2. 这里的难点是：要过滤掉那些已经在 node.Children 里的点
	//    (比如 (1,1) 已经在 Children 里了，就不能再 expand 它了)
	// 3. 从剩下的未尝试空位里，随机挑一个
	// 4. 执行落子 (PlaceStone)，创建新节点 (NewNode)
	// 5. 把新节点 append 到 node.Children 里
	// 6. 返回这个新节点
	
	// TODO: 你的代码
	emptyPoints := node.Board.GetEmptyPoints()

	var realPoints []game.Point
	for _, point := range emptyPoints {
		alreadyExpanded := false
		for _, child := range node.Children {
			if child.Move.X == point.X && child.Move.Y == point.Y {
				alreadyExpanded = true
				break
			}
		}
		if !alreadyExpanded {
			realPoints = append(realPoints, point)
		}
	}

	if len(realPoints) == 0 {
		return node
	}

	point := realPoints[rand.Intn(len(realPoints))]

	newBoard := node.Board.Clone()

	newBoard.PlaceStone(point.X, point.Y, node.NextPlayer)

	newNode := NewNode(newBoard, node, point, 3 - node.NextPlayer)

	node.Children = append(node.Children, newNode)

	return newNode
}

// 作业 3: Simulation
// 快速随机模拟直到游戏结束
func simulate(node *MCTSNode) int {
	// 提示：
	// 1. 记得复制一份棋盘！(currentBoard := node.Board)
	//    千万别在原来的节点上改，否则树就乱了
	// 2. 接下来就是你昨天写的"死循环随机落子"逻辑
	// 3. 返回赢家 (1 或 2，平局返回 0)

	// TODO: 你的代码
	currentBoard := node.Board.Clone()
    currentPlayer := node.NextPlayer

    for {
        emptyPoints := currentBoard.GetEmptyPoints()
        if len(emptyPoints) == 0 {
            return 0
        }
        
        idx := rand.Intn(len(emptyPoints))
        move := emptyPoints[idx]
        
        ok, _ := currentBoard.PlaceStone(move.X, move.Y, currentPlayer)
        if !ok { continue }

        if currentBoard.CheckWin(move.X, move.Y, currentPlayer) {
            return currentPlayer
        }
        currentPlayer = 3 - currentPlayer
    }

}

// 作业 4: Backpropagation
// 从当前节点(leaf)开始，一直往上找 Parent，更新数据
func backpropagate(node *MCTSNode, winner int) {
	for node != nil {
		node.Visits++
		// 如果赢家是该节点“产生时”的玩家，则计分

		if winner != 0 && winner != node.NextPlayer {
			node.Wins += 1.0
		} else if winner == 0 {
			node.Wins += 0.5 // 平局计 0.5
		}
		node = node.Parent // 向上爬
	}
}

// 作业 5: UCB 公式计算
// 计算一个节点的 UCB 值
func calculateUCB(node *MCTSNode) float64 {
	// 提示：
	// 公式：(Wins / Visits) + C * Sqrt( Log(Parent.Visits) / Visits )
	// 在 Go 里：
	// Log 是 math.Log()
	// Sqrt 是 math.Sqrt()
	// 注意防止除以 0 的情况
	if node.Visits == 0 {
		return math.MaxFloat64
	}
	UCB := node.Wins / float64(node.Visits) + UCB1ExplorationConstant * math.Sqrt(math.Log(float64(node.Parent.Visits))/float64(node.Visits))

	// TODO: 你的代码
	return UCB 
}

// -----------------------------------------------------------------------
// 辅助函数 (已经送给你了)
// -----------------------------------------------------------------------

// getBestChild: 找 Visits 最大的孩子 (最后一步决策用)
func getBestChild(root *MCTSNode) *MCTSNode {
	if len(root.Children) == 0 {
		return nil
	}
	maxVisits := -1
	var bestNode *MCTSNode
	for _, child := range root.Children {
		if child.Visits > maxVisits {
			maxVisits = child.Visits
			bestNode = child
		}
	}
	return bestNode
}

// isLeaf: 判断是不是终局 (或者刚开始的空树)
func isLeaf(node *MCTSNode) bool {
    // 只有当没有孩子，或者游戏本身已经分出胜负时，才是叶子
    // 简单版：没有孩子就是叶子
	return len(node.Children) == 0
}

// isFullyExpanded: 判断这个节点是不是所有孩子都生齐了
func isFullyExpanded(node *MCTSNode) bool {
    // 逻辑：如果孩子的数量 == 当前棋盘空位的数量，说明所有路都试过了
	emptyPoints := node.Board.GetEmptyPoints()
	return len(node.Children) == len(emptyPoints)
}