package lesson

import "fmt"

// 参考 结英 分享的《Go 语言圣经》第六章《方法》方法声明/基于指针对象的方法
// 完成以下练习题，在你个人目录下创建 lesson9.go
// 1.定义了一个名为 Player 的结构体，其中包含两个字段 name 和 score，分别表示玩家的姓名和分数。
// 现在，你需要实现一个名为 UpdateScore() 的指针方法，它将更新玩家的分数，并返回修改后的玩家。
// 然后，你还需要实现一个名为 CompareScore() 的指针方法，它将对比两个玩家的分数，大于返回1，等于返回0，小于返回-1
// 注意：
// •UpdateScore 方法的调用不会影响到接收器本身的状态
// 提问：已知我们定义的 CompareScore 方法中接收器是指针，而参数是非指针，请对下面两个问题分别作答
// •CompareScore 方法 Player 值对象 可以作为接收器发起调用吗？
// •CompareScore 方法 Player 指针对象 可以作为入参吗？

type Player struct {
	Name  string
	Score int
}

func (p *Player) UpdateScore(s int) *Player {
	// 取入参的值，相当于浅克隆
	p1 := *p
	// 赋值
	p1.Score = s
	// 返回新对象的指针
	return &p1
}

func (p *Player) CompareScore(q Player) int {
	if p.Score > q.Score {
		return 1
	} else if p.Score < q.Score {
		return -1
	}
	return 0
}

// 提问：已知我们定义的 CompareScore 方法中接收器是指针，而参数是非指针，请对下面两个问题分别作答
// •CompareScore 方法 Player 值对象 可以作为接收器发起调用吗？
// 解答：
// 可以的，但是如果用值对象传递，那么方法内的修改，就仅在方法内生效，具体见 CompareScore1 和 TestCompareScore1
func (p Player) CompareScore1(q Player) int {
	p.Score = 100
	fmt.Println("程序中修改p：", p.Score)
	if p.Score > q.Score {
		return 1
	} else if p.Score < q.Score {
		return -1
	}
	return 0
}

// •CompareScore 方法 Player 指针对象 可以作为入参吗？
// 解答：
// 可以的，但是如果入参是指针，那么方法内的修改会在方法外也生效，具体见 CompareScore2 和 TestCompareScore2
func (p *Player) CompareScore2(q *Player) int {
	q.Score = 200
	fmt.Println("程序中修改q：", q.Score)
	if p.Score > q.Score {
		return 1
	} else if p.Score < q.Score {
		return -1
	}
	return 0
}
