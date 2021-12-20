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

func prompt(s State) {
	m := map[State]string{
		Locked:   "Locked",
		Unlocked: "Unlocked",
	}
	fmt.Printf("当前的状态是[%s], 请输入命令：[coin|push]\n", m[s])
}

func step(state State, cmd string) State {
	if cmd != CmdCoin && cmd != CmdPush {
		fmt.Println("未知命令，请重新输入")
		return state
	}

	switch state {
	case Locked:
		if cmd == CmdCoin {
			fmt.Println("已解锁，请通行")
			state = Unlocked
		} else {
			fmt.Println("禁止通行，请先解锁")
		}
	case Unlocked:
		if cmd == CmdCoin {
			fmt.Println("大兄弟，别浪费钱了，现在已经解锁了")
		} else {
			fmt.Println("请通行，通行之后将会关闭")
			state = Locked
		}
	}
	return state
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
		state = step(state, strings.TrimSpace(cmd))
	}
}
