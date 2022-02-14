package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const FileName = "input.txt"

const KindInp = 0
const KindAdd = 1
const KindMul = 2
const KindDiv = 3
const KindMod = 4
const KindEql = 5

const RegW = 0
const RegX = 1
const RegY = 2
const RegZ = 3

type Command struct {
	Kind   int
	Reg1   int
	Reg2   int
	Val2   int
	UseVal bool
}

type Program []Command

type State struct {
	W, X, Y, Z int
}

type Uniq map[int][9]int

type CorrectZ map[int]map[int]bool

var programs = LoadInput()

func main() {
	start := time.Now()

	correctZ := CorrectZ{0: {}, 1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}, 10: {}, 11: {}, 12: {}, 13: {}, 14: {0: true}}
	deepRunCheck([]int{0}, 0, correctZ)

	fmt.Println("K=", 0, "L=", len(correctZ[0]), correctZ[0])
	for k := 1; k <= 13; k++ {
		fmt.Println("K=", k, "L=", len(correctZ[k]))
	}

	deepNumbers(0, 0, correctZ, [14]int{})

	elapsed := time.Since(start)
	fmt.Println("done", elapsed)
}

func deepNumbers(Z int, pos int, correctZ CorrectZ, nums [14]int) {

	// fmt.Println("DN=", Z, pos, nums)

	zzz := correctZ[pos+1]

	for input := 9; input >= 1; input-- {
		newState := applyProgram(Z, programs[pos], input)

		if zzz[newState.Z] && pos < 13 {
			nums[pos] = input
			deepNumbers(newState.Z, pos+1, correctZ, nums)
		} else if pos == 13 && newState.Z == 0 {
			nums[pos] = input
			fmt.Println("Solution:", nums)
		}
	}

}

func deepRunCheck(variants []int, pos int, correctZ CorrectZ) {

	newVariants := []int{}
	uniq := map[int]bool{}

	for i := 0; i < len(variants); i++ {
		for input := 9; input >= 1; input-- {
			newState := applyProgram(variants[i], programs[pos], input)

			if pos == 13 && correctZ[14][newState.Z] {
				// fmt.Println("Solution:", newState, "N=", input, "input Z = ", variants[i])
				correctZ[13][variants[i]] = true
			}

			if !uniq[newState.Z] {
				newVariants = append(newVariants, newState.Z)
				uniq[newState.Z] = true
			}
		}
	}

	if pos < 13 {
		deepRunCheck(newVariants, pos+1, correctZ)

		zzz := correctZ[pos+1]

		for i := 0; i < len(variants); i++ {
			for input := 9; input >= 1; input-- {
				newState := applyProgram(variants[i], programs[pos], input)
				if zzz[newState.Z] {
					correctZ[pos][variants[i]] = true
				}
			}
		}
	}
	fmt.Println("finished:", pos)
}

// func runCheck(programs []Program) {
// 	variants := []int{0}

// 	// uniqs := make(Uniq, 14)

// 	for pos := 0; pos <= 13; pos++ {

// 		newVariants := []int{}
// 		uniq := map[int]bool{}

// 		for i := 0; i < len(variants); i++ {
// 			zzz, exPos := correctZ[pos]
// 			if !exPos || exPos && zzz[variants[i]] {

// 				for input := 9; input >= 1; input-- {
// 					newState := applyProgram(variants[i], programs[pos], input)

// 					if pos == 11 && correctZ[12][newState.Z] {
// 						fmt.Println("Solution 11:", newState, "N=", input, "input Z = ", variants[i])
// 					}

// 					// if pos == 12 && correctZ[13][newState.Z] {
// 					// 	fmt.Println("Solution 12:", newState, "N=", input, "input Z = ", variants[i])
// 					// }

// 					// if pos == 13 && newState.Z == 0 {
// 					// 	fmt.Println("Solution:", newState, "N=", input, "input Z = ", variants[i])
// 					// }

// 					if !uniq[newState.Z] {
// 						newVariants = append(newVariants, newState.Z)
// 						uniq[newState.Z] = true
// 					}
// 				}

// 			}
// 		}
// 		variants = newVariants
// 		fmt.Println("Pos:", pos, "  V=", len(newVariants))
// 	}

// }

func applyProgram(Z int, pr Program, input int) State {
	state := State{Z: Z}
	for i := 0; i < len(pr); i++ {
		state = applyCommand(state, pr[i], input)
		// fmt.Printf("%+v  %s\n", state, cmdAsText(pr[i]))
	}
	return state
}

// inp a - Read an input value and write it to variable a.
// add a b - Add the value of a to the value of b, then store the result in variable a.
// mul a b - Multiply the value of a by the value of b, then store the result in variable a.
// div a b - Divide the value of a by the value of b, truncate the result to an integer, then store the result in variable a. (Here, "truncate" means to round the value toward zero.)
// mod a b - Divide the value of a by the value of b, then store the remainder in variable a. (This is also called the modulo operation.)
// eql a b - If the value of a and b are equal, then store the value 1 in variable a. Otherwise, store the value 0 in variable a.

func applyCommand(state State, command Command, input int) State {
	if command.Kind == KindInp {
		setReg(&state, command.Reg1, input)
		return state
	}

	v1 := getReg(state, command.Reg1)
	var v2 int
	if command.UseVal {
		v2 = command.Val2
	} else {
		v2 = getReg(state, command.Reg2)
	}

	if command.Kind == KindDiv {
		setReg(&state, command.Reg1, v1/v2)
	} else if command.Kind == KindEql {
		if v1 == v2 {
			setReg(&state, command.Reg1, 1)
		} else {
			setReg(&state, command.Reg1, 0)
		}
	} else if command.Kind == KindAdd {
		setReg(&state, command.Reg1, v1+v2)
	} else if command.Kind == KindMod {
		setReg(&state, command.Reg1, v1%v2)
	} else if command.Kind == KindMul {
		setReg(&state, command.Reg1, v1*v2)
	} else {
		panic("Command ???")
	}

	return state
}

func setReg(state *State, reg int, val int) {
	if reg == RegW {
		state.W = val
	}
	if reg == RegX {
		state.X = val
	}
	if reg == RegY {
		state.Y = val
	}
	if reg == RegZ {
		state.Z = val
	}
}

func getReg(state State, reg int) int {
	if reg == RegW {
		return state.W
	}
	if reg == RegX {
		return state.X
	}
	if reg == RegY {
		return state.Y
	}
	if reg == RegZ {
		return state.Z
	}
	panic("Read wrong reg")
}

func print(programs []Program) {
	for i, p := range programs {
		fmt.Println("Seq", i)
		for _, c := range p {
			fmt.Println("    ", cmdAsText(c))
		}
	}
}

func cmdAsText(c Command) string {
	if c.Kind == KindInp {
		return fmt.Sprintf("inp %s", regAsText(c.Reg1))
	}

	var part2 string
	if c.UseVal {
		part2 = strconv.Itoa(c.Val2)
	} else {
		part2 = regAsText(c.Reg2)
	}

	if c.Kind == KindAdd {
		return fmt.Sprintf("add %s %s", regAsText(c.Reg1), part2)
	}
	if c.Kind == KindDiv {
		return fmt.Sprintf("div %s %s", regAsText(c.Reg1), part2)
	}
	if c.Kind == KindEql {
		return fmt.Sprintf("eql %s %s", regAsText(c.Reg1), part2)
	}
	if c.Kind == KindMod {
		return fmt.Sprintf("mod %s %s", regAsText(c.Reg1), part2)
	}
	if c.Kind == KindMul {
		return fmt.Sprintf("mul %s %s", regAsText(c.Reg1), part2)
	}

	panic(fmt.Sprintf("Wrong command: %d", c.Kind))
}

func regAsText(reg int) string {
	if reg == RegW {
		return "w"
	}
	if reg == RegX {
		return "x"
	}
	if reg == RegY {
		return "y"
	}
	if reg == RegZ {
		return "z"
	}
	panic(fmt.Sprintf("Wrong reg: %d", reg))
}

func LoadInput() []Program {
	result := make([]Program, 0)

	file, _ := os.Open(FileName)
	scanner := bufio.NewScanner(file)

	var temp Program

	for scanner.Scan() {
		text := scanner.Text()
		command := parseCommand(text)

		if command.Kind == KindInp {
			if len(temp) > 0 {
				result = append(result, temp)
				temp = Program{}
			}
		}
		temp = append(temp, command)
	}

	if len(temp) > 0 {
		result = append(result, temp)
	}

	return result
}

func parseCommand(text string) Command {
	parts := strings.Split(text, " ")

	reg1 := registrByName(parts[1])
	var reg2 int
	var val2 int
	isVal := false

	if len(parts) == 3 {
		temp, err := strconv.ParseInt(parts[2], 10, 64)
		if err == nil {
			isVal = true
			val2 = int(temp)
		} else {
			reg2 = registrByName(parts[2])
		}
	}

	if parts[0] == "inp" {
		return Command{Kind: KindInp, Reg1: reg1, Reg2: reg2, Val2: val2, UseVal: isVal}
	}
	if parts[0] == "add" {
		return Command{Kind: KindAdd, Reg1: reg1, Reg2: reg2, Val2: val2, UseVal: isVal}
	}
	if parts[0] == "mul" {
		return Command{Kind: KindMul, Reg1: reg1, Reg2: reg2, Val2: val2, UseVal: isVal}
	}
	if parts[0] == "div" {
		return Command{Kind: KindDiv, Reg1: reg1, Reg2: reg2, Val2: val2, UseVal: isVal}
	}
	if parts[0] == "mod" {
		return Command{Kind: KindMod, Reg1: reg1, Reg2: reg2, Val2: val2, UseVal: isVal}
	}
	if parts[0] == "eql" {
		return Command{Kind: KindEql, Reg1: reg1, Reg2: reg2, Val2: val2, UseVal: isVal}
	}

	panic(fmt.Sprintf("Wrong text: %s", text))
}

// w, x, y, z
func registrByName(name string) int {
	if name == "w" {
		return RegW
	}
	if name == "x" {
		return RegX
	}
	if name == "y" {
		return RegY
	}
	if name == "z" {
		return RegZ
	}
	panic(fmt.Sprintf("Wrong reg: %s", name))
}
