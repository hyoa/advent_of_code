package day

import (
	_file "advent_of_code/helper/file"
	"sort"
	"strconv"
	"strings"
)

type Day9 struct {
	tiles        []tile
	tilesHashmap map[string]int
	rowCount     int
	columnCount  int
}

type tile struct {
	number int
	x      int
	y      int
}

func CreateDay9(path string) Day9 {
	r := _file.ReadTextFile(path)

	tiles := make([]tile, 0)
	tilesHashmap := make(map[string]int)
	row := 0
	columnCount := len(r[0])
	for _, s := range r {
		numbers := strings.Split(s, "")
		for k, n := range numbers {
			v, _ := strconv.Atoi(n)
			tiles = append(tiles, tile{
				number: v,
				x:      k,
				y:      row,
			})

			tilesHashmap[strconv.Itoa(k)+"|"+strconv.Itoa(row)] = v
		}
		row++
	}

	return Day9{
		tiles:        tiles,
		tilesHashmap: tilesHashmap,
		rowCount:     row,
		columnCount:  columnCount,
	}
}

func (d Day9) GetStep1Result() int {

	lowestTiles := getLowestPoint(d)
	sum := 0
	for _, v := range lowestTiles {
		sum += v + 1
	}

	return sum
}

func getLowestPoint(d Day9) map[string]int {
	lowestTiles := make(map[string]int)

	for i := 0; i < d.rowCount; i++ {
		for j := 0; j < d.columnCount; j++ {
			pos := strconv.Itoa(j) + "|" + strconv.Itoa(i)
			posL := strconv.Itoa(j-1) + "|" + strconv.Itoa(i)
			posR := strconv.Itoa(j+1) + "|" + strconv.Itoa(i)
			posT := strconv.Itoa(j) + "|" + strconv.Itoa(i-1)
			posB := strconv.Itoa(j) + "|" + strconv.Itoa(i+1)

			// fmt.Println("-----")
			vPosL := d.tilesHashmap[posL]
			vPosR := d.tilesHashmap[posR]
			vPosT := d.tilesHashmap[posT]
			vPosB := d.tilesHashmap[posB]
			if i == 0 {
				vPosT = 9

			}
			if i == d.rowCount-1 {
				vPosB = 9
			}
			if j == 0 {
				vPosL = 9
			}
			if j == d.columnCount-1 {
				vPosR = 9
			}

			// fmt.Printf("%d [%s] l: %d, r: %d, t:%d , b:%d\n", d.tilesHashmap[pos], pos, vPosL, vPosR, vPosT, vPosB)
			if d.tilesHashmap[pos] < vPosL && d.tilesHashmap[pos] < vPosR && d.tilesHashmap[pos] < vPosB && d.tilesHashmap[pos] < vPosT {
				lowestTiles[pos] = d.tilesHashmap[pos]
			}
		}
	}

	return lowestTiles
}

func getBasinsSizeForLowestPoint(startX, startY int, tiles map[string]int, rowCount int, columnCount int) int {
	all := getHighestClosestTiles(startX, startY, tiles, rowCount, columnCount)

	return len(all) + 1
}

func getHighestClosestTiles(x, y int, tiles map[string]int, rowCount int, columnCount int) []string {
	neighbors := make([]string, 0)

	pos := strconv.Itoa(x) + "|" + strconv.Itoa(y)
	posL := strconv.Itoa(x-1) + "|" + strconv.Itoa(y)
	posR := strconv.Itoa(x+1) + "|" + strconv.Itoa(y)
	posT := strconv.Itoa(x) + "|" + strconv.Itoa(y-1)
	posB := strconv.Itoa(x) + "|" + strconv.Itoa(y+1)

	incremented := tiles[pos] + 1

	if incremented == 9 {
		return neighbors
	}

	if tiles[posL] < 9 && tiles[posL] > tiles[pos] {
		neighbors = append(neighbors, posL)
	}
	if tiles[posR] < 9 && tiles[posR] > tiles[pos] {
		neighbors = append(neighbors, posR)
	}
	if tiles[posT] < 9 && tiles[posT] > tiles[pos] {
		neighbors = append(neighbors, posT)
	}
	if tiles[posB] < 9 && tiles[posB] > tiles[pos] {
		neighbors = append(neighbors, posB)
	}

	for _, n := range neighbors {
		s := strings.Split(n, "|")

		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])

		neighbors = append(neighbors, getHighestClosestTiles(x, y, tiles, rowCount, columnCount)...)
	}

	return uniqueString(neighbors)
}

func (d Day9) GetStep2Result() int {
	basinsSize := make([]int, 0)
	lowestPoints := getLowestPoint(d)
	for k := range lowestPoints {
		pos := strings.Split(k, "|")
		x, _ := strconv.Atoi(pos[0])
		y, _ := strconv.Atoi(pos[1])
		s := getBasinsSizeForLowestPoint(x, y, d.tilesHashmap, d.rowCount, d.columnCount)
		basinsSize = append(basinsSize, s)
	}

	sort.Ints(basinsSize)

	multiply := 1
	for _, v := range basinsSize[len(basinsSize)-3:] {
		multiply *= v
	}

	return multiply
}

func uniqueString(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
