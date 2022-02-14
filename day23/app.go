package main

import (
	"advent/utils"
	"bufio"
	"fmt"
	"os"
	"time"
)

//  #############
//  #...........#
//  ###D#C#A#B###
//    #D#C#B#A#
//    #D#B#A#C#
//    #D#C#B#A#
//    #########

//  #############
//  #...........#
//  ###B#C#B#D###
//    #D#C#B#A#
//    #D#B#A#C#
//    #A#D#C#A#
//    #########

type Action struct {
	Kind   int //0 - room to hall, 1 - hall to room, 2 - room to room
	Hall   int
	Room   int
	Place  int
	Room2  int
	Place2 int
}

var Rooms = [4][4]int{{1000, 1000, 1000, 1000}, {100, 100, 10, 100}, {1, 10, 1, 10}, {10, 1, 100, 1}} // E=47232

// var Rooms = [4][4]int{{10, 1000, 1000, 1}, {100, 100, 10, 1000}, {10, 10, 1, 100}, {1000, 1, 100, 1}} //test 1 E=44169
// var Rooms = [4][4]int{{1, 1, 1, 1}, {10, 10, 10, 10}, {1000, 100, 100, 100}, {100, 1000, 1000, 1000}} //test 2 E=4600

var Hall = [11]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

var MaxLevel = 50

var reader = bufio.NewReader(os.Stdin)

func main() {
	start := time.Now()

	// print(Rooms, Hall)
	bestEnergy := -1
	deepMove(Rooms, Hall, 0, &bestEnergy, 0)

	fmt.Println("bestEnergy:", bestEnergy)

	elapsed := time.Since(start)
	fmt.Println("done", elapsed)
}

func deepMove(rooms [4][4]int, hall [11]int, energy int, bestEnergy *int, level int) {

	// print(rooms, hall, energy)
	// reader.ReadString('\n')

	actions := possibleActions(rooms, hall)

	// fmt.Println(level, len(actions), energy)

	for i := 0; i < len(actions); i++ {

		if level == 0 {
			fmt.Printf("%d - %d / %d\n", level, i+1, len(actions))
		}

		newRooms, newHall, spentEnergy := applyAction(actions[i], rooms, hall)

		newEnergy := energy + spentEnergy

		if isDone(newRooms) {
			if *bestEnergy == -1 || newEnergy < *bestEnergy {
				*bestEnergy = energy + spentEnergy
				fmt.Println("Solution:", *bestEnergy, "L:", level)
			}
		} else {
			var cheatEnergy int
			if *bestEnergy > -1 {
				cheatEnergy = calcCheatEnergy(rooms, hall, false)
			}

			// if level < MaxLevel && (*bestEnergy == -1 || newEnergy < *bestEnergy) {
			// 	if *bestEnergy > -1 && energy+cheatEnergy > *bestEnergy {
			// 		print(rooms, hall, energy)
			// 		fmt.Println("Cheat stop:", *bestEnergy, energy+cheatEnergy)
			// 		reader.ReadString('\n')
			// 	}
			// }

			if level < MaxLevel && (*bestEnergy == -1 || newEnergy < *bestEnergy) && (*bestEnergy == -1 || energy+cheatEnergy < *bestEnergy) {
				deepMove(newRooms, newHall, newEnergy, bestEnergy, level+1)
			}
		}
	}
}

func calcCheatEnergy(rooms [4][4]int, hall [11]int, v bool) int {
	result := 0
	inPlace := [4]int{0, 0, 0, 0}

	for j := 3; j >= 0; j-- {
		for r := 0; r < 4; r++ {
			if rooms[r][j] > 0 {
				ameba := rooms[r][j]
				ownRoomIndex := ownRoom(ameba)
				c1 := ownRoomIndex == r
				c2 := (j == 3 || j == 2 && rooms[r][3] == ameba || j == 1 && rooms[r][3] == ameba && rooms[r][2] == ameba || j == 0 && rooms[r][3] == ameba && rooms[r][2] == ameba && rooms[r][1] == ameba)
				if c1 && c2 {
					inPlace[ownRoomIndex]++
				}
			}
		}
	}

	for r := 0; r < 4; r++ {
		for j := 3; j >= 0; j-- {
			if rooms[r][j] > 0 {
				ameba := rooms[r][j]
				ownRoomIndex := ownRoom(ameba)

				c1 := j != 3
				c2 := j == 2 && rooms[r][3] != ameba
				c3 := j == 1 && (rooms[r][3] != ameba || rooms[r][2] != ameba)
				c4 := j == 0 && (rooms[r][3] != ameba || rooms[r][2] != ameba || rooms[r][1] != ameba)

				if ownRoomIndex == r && c1 && (c2 || c3 || c4) {
					result = result + ((4-inPlace[r])+j+3)*ameba
					inPlace[ownRoomIndex]++
				} else if ownRoomIndex != r {
					result = result + ((4-inPlace[ownRoomIndex])+j+utils.AbsInt(r-ownRoomIndex)*2+1)*ameba
					inPlace[ownRoomIndex]++
				}
			}
		}
	}

	if v {
		fmt.Println("R=", result)
	}

	for i := 0; i < 11; i++ {
		if hall[i] > 0 {
			ameba := hall[i]
			ownRoomIndex := ownRoom(ameba)

			h := ((4 - inPlace[ownRoomIndex]) + utils.AbsInt(i-2*(ownRoomIndex+1))) * ameba
			result = result + h
			if v {
				fmt.Printf("H = (4 - %d + abs(%d - 2 * (%d + 1) )) * %d = %d\n", inPlace[ownRoomIndex], i, ownRoomIndex, ameba, h)
			}

			inPlace[ownRoomIndex]++
		}
	}

	return result
}

func possibleActions(rooms [4][4]int, hall [11]int) []Action {
	result := make([]Action, 0)

	freePlaces := [4]int{-1, -1, -1, -1}
	freePlaces[0] = freeAndCorrect(rooms[0], 1)
	freePlaces[1] = freeAndCorrect(rooms[1], 10)
	freePlaces[2] = freeAndCorrect(rooms[2], 100)
	freePlaces[3] = freeAndCorrect(rooms[3], 1000)

	wanderers := [4]int{-1, -1, -1, -1}
	wanderers[0] = getWanderer(rooms[0], 1)
	wanderers[1] = getWanderer(rooms[1], 10)
	wanderers[2] = getWanderer(rooms[2], 100)
	wanderers[3] = getWanderer(rooms[3], 1000)

	//room to room
	for r := 0; r < 4; r++ {
		if wanderers[r] > -1 {
			ameba := rooms[r][wanderers[r]]
			if notInOwnRoom(ameba, r) {
				roomIndex := ownRoom(ameba)
				roomEntrance := 2 + r*2
				place2Index := freeAndCorrect(rooms[roomIndex], ameba)
				if place2Index > -1 && freeWayFromHall(hall, roomEntrance, roomIndex) {
					temp := Action{Kind: 2, Hall: -1, Room: r, Place: wanderers[r], Room2: roomIndex, Place2: place2Index}
					result = append(result, temp)

					// fmt.Printf("R-R: %+v\n", temp)
					// fmt.Println(hall, roomEntrance, roomIndex)
				}
			}
		}
	}

	//hall to room
	for i := 0; i < 11; i++ {
		if hall[i] > 0 {
			ameba := hall[i]
			for r := 0; r < 4; r++ {
				if ownRoom(ameba) == r && freePlaces[r] > -1 && freeWayFromHall(hall, i, r) {
					result = append(result, Action{Kind: 1, Hall: i, Room: r, Place: freePlaces[r], Room2: -1, Place2: -1})
				}
			}
		}
	}

	//room to hall
	for r := 0; r < 4; r++ {
		if wanderers[r] > -1 {
			for h := 0; h < 11; h++ {
				if h != 2 && h != 4 && h != 6 && h != 8 && freeWayFromRoom(hall, h, r) {
					result = append(result, Action{Kind: 0, Hall: h, Room: r, Place: wanderers[r], Room2: -1, Place2: -1})
				}
			}
		}
	}

	//sort by
	//  moving in correct room
	//  spent energy asc

	return result
}

func ownRoom(ameba int) int {
	if ameba == 1 {
		return 0
	}
	if ameba == 10 {
		return 1
	}
	if ameba == 100 {
		return 2
	}
	if ameba == 1000 {
		return 3
	}
	return -1
}

func notInOwnRoom(ameba int, roomIndex int) bool {
	return !(roomIndex == 0 && ameba == 1) && !(roomIndex == 1 && ameba == 10) && !(roomIndex == 2 && ameba == 100) && !(roomIndex == 3 && ameba == 1000)
}

func freeWayFromRoom(hall [11]int, hallIndex, roomIndex int) bool {
	enterIndex := 2 + roomIndex*2

	if enterIndex > hallIndex {
		for i := hallIndex; i <= enterIndex; i++ {
			if hall[i] > 0 {
				return false
			}
		}
	} else {
		for i := enterIndex; i <= hallIndex; i++ {
			if hall[i] > 0 {
				return false
			}
		}
	}

	return true
}

func getWanderer(room [4]int, correct int) int {
	if room[0] > 0 && (room[0] != correct || room[0] == correct && (room[1] != correct || room[2] != correct || room[3] != correct)) {
		return 0
	}
	if room[1] > 0 && room[0] == 0 && (room[1] != correct || room[1] == correct && (room[2] != correct || room[3] != correct)) {
		return 1
	}
	if room[2] > 0 && room[0] == 0 && room[1] == 0 && (room[2] != correct || room[2] == correct && room[3] != correct) {
		return 2
	}
	if room[3] > 0 && room[3] != correct && room[0] == 0 && room[1] == 0 && room[2] == 0 {
		return 3
	}
	return -1
}

func freeWayFromHall(hall [11]int, hallIndex, roomIndex int) bool {
	enterIndex := 2 + roomIndex*2

	if enterIndex > hallIndex {
		for i := hallIndex + 1; i <= enterIndex; i++ {
			if hall[i] > 0 {
				return false
			}
		}
	} else {
		for i := enterIndex; i < hallIndex; i++ {
			if hall[i] > 0 {
				return false
			}
		}
	}

	return true
}

func freeAndCorrect(room [4]int, correct int) int {
	if room[3] == 0 {
		return 3
	}
	if room[2] == 0 && room[3] == correct {
		return 2
	}
	if room[1] == 0 && room[3] == correct && room[2] == correct {
		return 1
	}
	if room[0] == 0 && room[3] == correct && room[2] == correct && room[1] == correct {
		return 0
	}
	return -1
}

func applyAction(action Action, rooms [4][4]int, hall [11]int) ([4][4]int, [11]int, int) {
	var base int
	var path int

	if action.Kind == 0 {
		base = rooms[action.Room][action.Place]
		rooms[action.Room][action.Place] = 0
		hall[action.Hall] = base
		path = action.Place + utils.AbsInt(action.Hall-2*(action.Room+1)) + 1
	} else if action.Kind == 1 {
		base = hall[action.Hall]
		hall[action.Hall] = 0
		rooms[action.Room][action.Place] = base
		path = action.Place + utils.AbsInt(action.Hall-2*(action.Room+1)) + 1
	} else if action.Kind == 2 {
		base = rooms[action.Room][action.Place]
		rooms[action.Room][action.Place] = 0
		rooms[action.Room2][action.Place2] = base
		path = action.Place + action.Place2 + 1 + utils.AbsInt(action.Room-action.Room2)*2 + 1
	}

	return rooms, hall, base * path
}

func isDone(rooms [4][4]int) bool {
	return rooms[0][0] == 1 && rooms[0][1] == 1 && rooms[0][2] == 1 && rooms[0][3] == 1 &&
		rooms[1][0] == 10 && rooms[1][1] == 10 && rooms[1][2] == 10 && rooms[1][3] == 10 &&
		rooms[2][0] == 100 && rooms[2][1] == 100 && rooms[2][2] == 100 && rooms[2][3] == 100 &&
		rooms[3][0] == 1000 && rooms[3][1] == 1000 && rooms[3][2] == 1000 && rooms[3][3] == 1000
}

var EtoS = map[int]string{0: ".", 1: "A", 10: "B", 100: "C", 1000: "D"}

func print(rooms [4][4]int, hall [11]int, energy int) {
	fmt.Println("--==", energy, "==--")
	fmt.Println("--==", calcCheatEnergy(rooms, hall, true), "==--")

	for i := 0; i < 11; i++ {
		fmt.Print(EtoS[hall[i]])
	}
	fmt.Println()
	for i := 0; i < 4; i++ {
		fmt.Print(" ")
		for j := 0; j < 4; j++ {
			fmt.Print(" ")
			fmt.Print(EtoS[rooms[j][i]])
		}
		fmt.Println()
	}
	fmt.Println()
}
