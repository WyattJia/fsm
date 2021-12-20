package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
* Finite-state machine
	* Acceptors
	* Classifiers
	* Transducers
*/

type State uint32

const (
	Locked State = iota
	Unlocked
)

const (
	CmdCoin = "coin"
	CmdPush = "push"
)

type Turnstile struct {
	State State
}

type CmdStateTupple struct {
	Cmd   string
	State State
}

type TransitionFunc func(state *State)

func (p *Turnstile) ExecuteCmd(cmd string) {
	tupple := CmdStateTupple{strings.TrimSpace(cmd), p.State}
	if f := StateTransitionTable[tupple]; f == nil {
		fmt.Println("unknown command, try again please")
	} else {
		f(&p.State)
	}
}

func prompt(s State) {
	m := map[State]string{
		Locked:   "Locked",
		Unlocked: "Unlocked",
	}
	fmt.Printf("当前的状态是[%s], 请输入命令：[coin|push]\n", m[s])
}

var StateTransitionTable = map[CmdStateTupple]TransitionFunc{
	{CmdCoin, Locked}: func(state *State) {
		fmt.Println("已解锁，请通行")
		*state = Unlocked
	},
	{CmdPush, Locked}: func(state *State) {
		fmt.Println("禁止通行，请先解锁")
	},
	{CmdCoin, Unlocked}: func(state *State) {
		fmt.Println("大兄弟，不要浪费钱了")
	},
	{CmdPush, Unlocked}: func(state *State) {
		fmt.Println("请尽快通行，然后将会锁定")
		*state = Locked
	},
}

func main() {
	// Init state
	state := Locked

	reader := bufio.NewReader(os.Stdin)
	prompt(state)

	for {
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		// 获取状态转换表中的值
		tupple := CommandStateTupple{strings.TrimSpace(cmd), state}

		if f := StateTransitionTable[tupple]; f == nil {
			fmt.Println("未知命令，请重新输入")
		} else {
			f(&state)
		}
	}
}
