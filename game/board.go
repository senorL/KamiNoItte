package game

import "fmt"

const Size = 6

type Board [Size][Size]int

type Point struct {
	X int
	Y int
}

func (b *Board) PlaceStone(x, y int, player int) (result bool, err string) {
	if x < 0 || x >= Size || y < 0 || y >= Size {
		return false, "越界"
	}
	if b[x][y] != 0 { // 直接用 b[x][y]
		return false, "已有子"
	}
	b[x][y] = player
	return true, ""
}

func (b *Board) CheckWin(x, y int, player int) bool {
	// 1. 检查横向 (左右) ---------------------------------------
    count := 1 // 先把中间这一颗算上

    // 向左数 (y 减小)
    for i := 1; i < 5; i++ {
        // 如果越界 或者 颜色不对，就停
        if y-i < 0 || b[x][y-i] != player {
            break
        }
        count++
    }

    // 向右数 (y 增加)
    for i := 1; i < 5; i++ {
        // 如果越界 或者 颜色不对，就停
        if y+i >= Size || b[x][y+i] != player {
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
		if x-i < 0 || b[x-i][y] != player {
			break
		}
		count++
	}

	for i := 1; i < 5; i++ {
		if x+i >= Size || b[x+i][y] != player {
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
		if x-i < 0 || y-i < 0 || b[x-i][y-i] != player {
			break
		}
		count++
	}

	for i := 1; i < 5; i++ {
		if x+i >= Size || y+i >= Size || b[x+i][y+i] != player {
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
		if x-i < 0 || y+i >= Size || b[x-i][y+i] != player {
			break
		}
		count++
	}

	for i := 1; i < 5; i++ {
		if x+i >= Size || y-i < 0 || b[x+i][y-i] != player {
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

func (b *Board) Show() {
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
				switch b[i][j] {
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

func (b *Board) GetEmptyPoints() []Point {
	var emptyPoints []Point
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if b[i][j] == 0 {
				emptyPoints = append(emptyPoints, Point{X: i, Y: j}) // 注意 X, Y 大写
			}
		}
	}
	return emptyPoints
}
