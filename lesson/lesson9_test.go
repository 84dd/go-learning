package lesson

import (
	"fmt"
	"testing"
)

func TestUpdateScore(t *testing.T) {
	player := Player{"小明", 100}
	fmt.Println("调用前player()：", player.Score)
	newPlayer := player.UpdateScore(50)
	fmt.Println("调用后(player)：", player.Score)
	fmt.Println("调用后(newPlayer)：", newPlayer.Score)
}

func TestCompareScor1(t *testing.T) {
	p := Player{"小明", 150}
	q := Player{"小红", 50}
	fmt.Println(p.CompareScore(q))
	fmt.Println(p.CompareScore1(q))
}

func TestCompareScore1(t *testing.T) {
	p := Player{"小明", 150}
	q := Player{"小红", 50}
	fmt.Println("调用前p：", p.Score)
	fmt.Println(p.CompareScore1(q))
	fmt.Println("调用后p：", p.Score)
}

func TestCompareScore2(t *testing.T) {
	p := Player{"小明", 150}
	q := Player{"小红", 50}
	fmt.Println("调用前q：", q.Score)
	fmt.Println(p.CompareScore2(&q))
	fmt.Println("调用后q：", q.Score)
}
