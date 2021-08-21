package main

import "fmt"

type TongsOfPatty interface { //외부 공개 메서드 = 관계
	String() string // PutOfPatty는 String()이라는 외부 공개 메서드
}

type Patty interface { //오로지 관계만 선언해본다
	GetOneTongs() TongsOfPatty // Jam은 GetOnsSpoon()이라는 외부 공개 메서드
}

type Bigburger struct {
	val string
}

func (b *Bigburger) PutOfPatty(Patty Patty) { //오랜지잼이든 스트로베리잼이든 상관없다.
	Tongs := Patty.GetOneTongs()
	b.val += Tongs.String()
}

func (b *Bigburger) String() string { //Bread 메서드2 String
	return "Bigburger " + b.val + " + coke"
}

type PorkPatty struct {
}

func (j *PorkPatty) GetOneTongs() TongsOfPatty {
	return &TongsOfPorkPatty{}
}

type CheesePatty struct { ///5555
}

func (j *CheesePatty) GetOneTongs() TongsOfPatty { ///우럭을 접시위로 올린다
	return &TongsOfSlicePatty{}
}

type DubblePatty struct {
}

func (j *DubblePatty) GetOneTongs() TongsOfPatty {
	return &TongsOfDubblePatty{}
}

type TongsOfPorkPatty struct {
}

func (s *TongsOfPorkPatty) String() string {
	return "+ PorkPatty"
}

type TongsOfSlicePatty struct { ///5555
}

func (s *TongsOfSlicePatty) String() string { ///접시를 불러온다
	return "+ CheesePatty"
}

type TongsOfDubblePatty struct {
}

func (s *TongsOfDubblePatty) String() string {
	return "+ DPatty"
}

func main() {
	Bigburger := &Bigburger{}
	// jam := &StrawberryJam{}
	// jam := &OrangeJam{}
	Patty := &CheesePatty{}
	Bigburger.PutOfPatty(Patty)

	fmt.Println(Bigburger)
}
