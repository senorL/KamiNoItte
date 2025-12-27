package game

import (
	"fmt"
)

const Size = 15

type Board [Size][Size]int

type Point struct {
	X int
	Y int
}

// Clone 返回当前棋盘的拷贝（值拷贝）
// 注意：Board 是定长数组，值赋值本身就是深拷贝，这里显式拷贝更直观
func (b Board) Clone() Board {
	var nb Board
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			nb[i][j] = b[i][j]
		}
	}
	return nb
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
	dirs := [][2]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}}
	for _, d := range dirs {
		count := 1
		for i := 1; i < 5; i++ { // 正向
			nx, ny := x+d[0]*i, y+d[1]*i
			if nx < 0 || nx >= Size || ny < 0 || ny >= Size || b[nx][ny] != player {
				break
			}
			count++
		}
		for i := 1; i < 5; i++ { // 反向
			nx, ny := x-d[0]*i, y-d[1]*i
			if nx < 0 || nx >= Size || ny < 0 || ny >= Size || b[nx][ny] != player {
				break
			}
			count++
		}
		if count >= 5 {
			return true
		}
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
