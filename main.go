package main

import "fmt"

// 定义常量：棋盘大小
const Size = 6

// 输入：棋盘，x坐标，y坐标，谁下的(1或2)
// 输出：是否成功(bool)，如果有错误返回error
func PlaceStone(board *[Size][Size]int, x int, y int, player int) (ok bool, message string) {
    // 1. 判断 x, y 是否越界 (小于0 或 大于等于Size)
    // if x < 0 || x >= Size ... { return false, "越界了" }
	if x < 0 || x >= Size {
		ok = false
		message = "越界"
		return ok, message
	}
    // 2. 判断这个位置是不是已经有子了 (board[x][y] != 0)
    if board[x][y] != 0 {
		ok = false
		message = "已有子!!!!"
		return ok, message
	}
    // 3. 如果没问题，修改 board[x][y] = player
    // 注意：这里传入的是 board 的指针，所以直接改不仅是改副本
    board[x][y] = player

    return true, ""
}


func PrintBoard(board *[Size][Size]int) {
		// 行号
		fmt.Print("   ")
		for i := 0; i < Size; i++ {
			if i < 10 {
				fmt.Printf("%d  ", i)
			} else {
				fmt.Printf("%d ", i)
			}
		}

		fmt.Println()

		for i := 0; i < Size; i++ {
			if i < 10 {
				fmt.Printf("%d  ", i)
			} else {
				fmt.Printf("%d ", i)
			}

			for j := 0; j < Size; j++ {
				switch board[i][j] {
				case 1:
					fmt.Print("●  ") // 黑子
				case 2:
					fmt.Print("○  ") // 白子
				default:
					fmt.Print(".  ") // 空位
				}
			}
			fmt.Println()
		}
		fmt.Println("棋盘初始化完成")
}

// CheckWin 检查 (x,y) 落子后是否获胜
// board: 棋盘
// x, y: 当前落子的坐标
// player: 当前玩家 (1或2)
func CheckWin(board *[Size][Size]int, x int, y int, player int) bool {
    // 1. 检查横向 (左右) ---------------------------------------
    count := 1 // 先把中间这一颗算上

    // 向左数 (y 减小)
    for i := 1; i < 5; i++ {
        // 如果越界 或者 颜色不对，就停
        if y-i < 0 || board[x][y-i] != player {
            break
        }
        count++
    }

    // 向右数 (y 增加)
    for i := 1; i < 5; i++ {
        // 如果越界 或者 颜色不对，就停
        if y+i >= Size || board[x][y+i] != player {
            break
        }
        count++
    }

    // 判断是否五连
    if count >= 5 {
        return true
    }

	count = 1
    // 2. 检查纵向 (上下) | ---------------------------------------
    // 提示：重置 count = 1
    // 向上是 x-i，向下是 x+i
	for i := 1; i < 5; i++ {
		if x-i < 0 || board[x-i][y] != player {
			break
		}
		count++
	}

	for i := 1; i < 5; i++ {
		if x+1 >= Size || board[x+i][y] != player {
			break
		}
		count++
	}

    // 判断是否五连
    if count >= 5 {
        return true
    }
	
	

	count = 1
    // 3. 检查左斜 (左上到右下) \ --------------------------------
    // 提示：左上是 (x-i, y-i)，右下是 (x+i, y+i)
    // 请你自己写...
	for i := 1; i < 5; i++ {
		if x-i < 0 || y-i < 0 || board[x-i][y-i] != player {
			break
		}
		count++
	}

	for i := 1; i < 5; i++ {
		if x+i >= Size || y+i >= Size || board[x+i][y+i] != player {
			break
		}
		count++
	}
	
	// 判断是否五连	
	if count >= 5 {
		return true
	}
	
	count = 1

    // 4. 检查右斜 (右上到左下) / --------------------------------
    // 提示：右上是 (x-i, y+i)，左下是 (x+i, y-i)
    // 请你自己写...
	for i := 1; i < 5; i++ {
		if x-i < 0 || y+i >= Size || board[x-i][y+i] != player {
			break
		}
		count++
	}

	for i := 1; i < 5; i++ {
		if x+i >= Size || y-i < 0 || board[x+i][y-i] != player {
			break
		}
		count++
	}			
	// 判断是否五连
	if count >= 5 {
		return true
	}	

    return false
}


func main() {
    // 任务：定义一个 15x15 的二维整型数组，叫 board
    // 0代表空，1代表黑子，2代表白子
    // 提示：Go的二维数组写法是 [Size][Size]int
	board := [Size][Size]int{}
    // var board [Size][Size]int  <-- 解除注释就是答案，但你试试用 := 怎么写？
	currentPlayer := 1
	
	for {
    	// 1. 把棋盘打印出来
		PrintBoard(&board)

		// 2. 提示用户输入坐标
        fmt.Printf("请玩家 %d 输入坐标 (x y): ", currentPlayer)
        var x, y int
        fmt.Scan(&x, &y) // & 取地址，把输入的值存进去
        
        // 3. 调用 PlaceStone 落子
        success, errMsg := PlaceStone(&board, x, y, currentPlayer)
        
        // 4. 如果失败，打印 errMsg，continue 继续下一次循环
		if success != true {
			fmt.Println(errMsg)
			fmt.Println()
			continue
		}

		if CheckWin(&board, x, y, currentPlayer) {
			PrintBoard(&board)
			fmt.Printf("玩家%d获胜\n", currentPlayer)
			break
		}

        // 5. 如果成功，切换玩家 (1变2，2变1)
        currentPlayer = 3 - currentPlayer 
	}

}