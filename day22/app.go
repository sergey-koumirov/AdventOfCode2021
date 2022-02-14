package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Cube struct {
	X1, X2 int
	Y1, Y2 int
	Z1, Z2 int
	Parent int
	On     bool
}

type Interval struct {
	V1, V2 int
}

func main() {
	start := time.Now()

	cubes := LoadInput()
	// fmt.Println(data)

	fmt.Println("pre1")
	pre1 := splitCubesByCubes(cubes)
	fmt.Println("pre2")
	pre2 := splitCubesByCubes(pre1)
	fmt.Println("finalShards")
	finalShards := splitCubesByCubes(pre2)

	fmt.Println("cubes: ", len(cubes))
	// fmt.Println("points: ", len(points))
	fmt.Println("-------- shards: ", len(finalShards))

	// printShards(finalShards)

	for _, cube := range cubes {
		fmt.Printf("%+v %d\n", cube, vol(cube))
		for i := 0; i < len(finalShards); i++ {
			if shardInside(finalShards[i], cube) {
				finalShards[i].On = cube.On
			}
		}
	}

	// printShards(finalShards)

	total := int64(0)
	for i := 0; i < len(finalShards); i++ {
		if finalShards[i].On {
			total = total + vol(finalShards[i])
		}
	}
	fmt.Println("total: ", total)

	elapsed := time.Since(start)
	fmt.Println("done", elapsed)
}

func splitCubesByCubes(cubes []Cube) []Cube {
	result := make([]Cube, 0)

	for i := 0; i < len(cubes); i++ {

		stone := cubes[i]

		shards := make([]Cube, 0)
		shards = append(shards, Cube{X1: stone.X1, X2: stone.X2, Y1: stone.Y1, Y2: stone.Y2, Z1: stone.Z1, Z2: stone.Z2})

		for j := 0; j < len(cubes); j++ {

			if i != j {
				hammer := cubes[j]
				if xInt(stone, hammer) && yInt(stone, hammer) && zInt(stone, hammer) {

					newShards := make([]Cube, 0, len(shards)*2)

					for k := 0; k < len(shards); k++ {
						shard := shards[k]
						if xInt(shard, hammer) && yInt(shard, hammer) && zInt(shard, hammer) && vol(shard) > 1 {
							shardsK := breakCubeByCube(hammer, shards[k])
							for n := 0; n < len(shardsK); n++ {
								if uniqShard(&newShards, shardsK[n]) {
									shardsK[n].Parent = i
									newShards = append(newShards, shardsK[n])
								}
							}
						} else {
							if uniqShard(&newShards, shard) {
								shard.Parent = i
								newShards = append(newShards, shard)
							}
						}
					}

					shards = newShards
				}
			}

			// fmt.Println("i,j=", i, j)
			// printShards(shards)
		}

		for n := 0; n < len(shards); n++ {
			if uniqShard(&result, shards[n]) {
				result = append(result, shards[n])
			}
		}

		// fmt.Println("----------------------------------------------------------------------------")
	}

	checkIntersections(result)

	return result
}

func checkIntersections(result []Cube) {
	cnt := 0

	for x := 0; x < len(result); x++ {
		for y := 0; y < len(result); y++ {
			if x != y && xInt(result[x], result[y]) && yInt(result[x], result[y]) && zInt(result[x], result[y]) {
				//fmt.Println("!!!xy=", x, y, result[x].Parent, result[y].Parent)
				cnt++
			}
		}
	}

	if cnt > 0 {
		fmt.Println("Warning intersections:", cnt)
	}
}

func printShards(result []Cube) {
	fmt.Println("--------")
	for _, s := range result {
		fmt.Printf("x %d..%d  y %d..%d  z %d..%d %t +%d\n", s.X1, s.X2, s.Y1, s.Y2, s.Z1, s.Z2, s.On, vol(s))
	}
	fmt.Println("--------")
}

func eqCubes(c, s Cube) bool {
	return c.X1 == s.X1 && c.X2 == s.X2 && c.Y1 == s.Y1 && c.Y2 == s.Y2 && c.Z1 == s.Z1 && c.Z2 == s.Z2
}

func breakCubeByCube(h, s Cube) []Cube {
	tempX := make([]Cube, 0)

	if xBtw(s, h) && yInt(h, s) && zInt(h, s) {
		vals := splitIntoIntervals(s.X1, s.X2, h.X1, h.X2)
		// fmt.Println("i-X:", s.X1, s.X2, h.X1, h.X2, vals)
		for _, val := range vals {
			tempX = append(tempX, Cube{X1: val.V1, X2: val.V2, Y1: s.Y1, Y2: s.Y2, Z1: s.Z1, Z2: s.Z2})
		}
	} else {
		tempX = append(tempX, Cube{X1: s.X1, X2: s.X2, Y1: s.Y1, Y2: s.Y2, Z1: s.Z1, Z2: s.Z2})
	}

	tempY := make([]Cube, 0)
	if yBtw(s, h) && xInt(h, s) && zInt(h, s) {
		vals := splitIntoIntervals(s.Y1, s.Y2, h.Y1, h.Y2)

		for i := 0; i < len(tempX); i++ {
			tx := tempX[i]

			if xInt(h, tx) && yInt(h, tx) && zInt(h, tx) {
				for _, val := range vals {
					tempY = append(tempY, Cube{X1: tx.X1, X2: tx.X2, Y1: val.V1, Y2: val.V2, Z1: tx.Z1, Z2: tx.Z2})
				}
			} else {
				tempY = append(tempY, tx)
			}
		}
	} else {
		tempY = tempX
	}

	tempZ := make([]Cube, 0)
	if zBtw(s, h) && xInt(h, s) && yInt(h, s) {
		vals := splitIntoIntervals(s.Z1, s.Z2, h.Z1, h.Z2)
		for i := 0; i < len(tempY); i++ {
			ty := tempY[i]
			for _, val := range vals {
				tempZ = append(tempZ, Cube{X1: ty.X1, X2: ty.X2, Y1: ty.Y1, Y2: ty.Y2, Z1: val.V1, Z2: val.V2})
			}
		}
	} else {
		tempZ = tempY
	}

	return tempZ
}

func xBtw(stone, hammer Cube) bool {
	return stone.X1 <= hammer.X1 && hammer.X1 <= stone.X2 || stone.X1 <= hammer.X2 && hammer.X2 <= stone.X2
}

func yBtw(stone, hammer Cube) bool {
	return stone.Y1 <= hammer.Y1 && hammer.Y1 <= stone.Y2 || stone.Y1 <= hammer.Y2 && hammer.Y2 <= stone.Y2
}

func zBtw(stone, hammer Cube) bool {
	return stone.Z1 <= hammer.Z1 && hammer.Z1 <= stone.Z2 || stone.Z1 <= hammer.Z2 && hammer.Z2 <= stone.Z2
}

func splitIntoIntervals(stoneX1, stoneX2, x1, x2 int) []Interval {
	result := make([]Interval, 0)

	var leftX int
	var rightX int

	if stoneX1 < x1 {
		result = append(result, Interval{V1: stoneX1, V2: x1 - 1})
		leftX = x1
	} else {
		leftX = stoneX1
	}

	if x2 < stoneX2 {
		result = append(result, Interval{V1: x2 + 1, V2: stoneX2})
		rightX = x2
	} else {
		rightX = stoneX2
	}

	result = append(result, Interval{V1: leftX, V2: rightX})

	return result
}

func yInt(c1, c2 Cube) bool {
	return c1.Y2 >= c2.Y1 && c2.Y2 >= c1.Y1
}

func xInt(c1, c2 Cube) bool {
	return c1.X2 >= c2.X1 && c2.X2 >= c1.X1
}

func zInt(c1, c2 Cube) bool {
	return c1.Z2 >= c2.Z1 && c2.Z2 >= c1.Z1
}

func vol(shard Cube) int64 {
	return int64(shard.X2-shard.X1+1) * int64(shard.Y2-shard.Y1+1) * int64(shard.Z2-shard.Z1+1)
}

func shardInside(shard Cube, cube Cube) bool {

	xx := cube.X1 <= shard.X1 && shard.X2 <= cube.X2
	yy := cube.Y1 <= shard.Y1 && shard.Y2 <= cube.Y2
	zz := cube.Z1 <= shard.Z1 && shard.Z2 <= cube.Z2

	return xx && yy && zz
}

func addShards(cubes *[]Cube, shards []Cube, shardEx map[string]bool) {

	for i := 0; i < len(shards); i++ {

		k := strconv.FormatInt(int64(shards[i].X1), 10) + "-" + strconv.FormatInt(int64(shards[i].X2), 10) + "-" +
			strconv.FormatInt(int64(shards[i].Y1), 10) + "-" + strconv.FormatInt(int64(shards[i].Y2), 10) + "-" +
			strconv.FormatInt(int64(shards[i].Z1), 10) + "-" + strconv.FormatInt(int64(shards[i].Z2), 10)

		if !(shardEx)[k] {
			(shardEx)[k] = true
			*cubes = append(*cubes, shards[i])
		}

		// if uniqShard(cubes, shards[i]) {
		// 	*cubes = append(*cubes, shards[i])
		// }
	}
}

func uniqShard(cubes *[]Cube, s Cube) bool {

	for i := 0; i < len(*cubes); i++ {
		c := (*cubes)[i]
		if c.X1 == s.X1 && c.X2 == s.X2 && c.Y1 == s.Y1 && c.Y2 == s.Y2 && c.Z1 == s.Z1 && c.Z2 == s.Z2 {
			return false
		}
	}
	return true
}

func LoadInput() []Cube {
	result := make([]Cube, 0)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		parts := strings.Split(text, " ")

		nums := strings.Split(parts[1], ",")

		x1, x2 := extractTwo(nums[0])
		y1, y2 := extractTwo(nums[1])
		z1, z2 := extractTwo(nums[2])

		result = append(result, Cube{X1: x1, X2: x2, Y1: y1, Y2: y2, Z1: z1, Z2: z2, On: parts[0] == "on"})
	}

	return result
}

func extractTwo(text string) (int, int) {
	parts := strings.Split(text[2:], "..")

	temp1, _ := strconv.ParseInt(parts[0], 10, 64)
	temp2, _ := strconv.ParseInt(parts[1], 10, 64)

	return int(temp1), int(temp2)
}
