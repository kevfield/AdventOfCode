package main

import (
	"fmt"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	input := `...........................A.............W........
.......x..........AZ......P.........m......k..W...
....v.....Z..K...V..........A.......R....f........
.............d..........V.............2...........
.........Z........d...V.....B....C.....e..........
............Z.vA...........m...................s.W
...x..o.....................ek....................
...O.....VK..................R......B.............
.........O..................my....kB..............
.v...................y..........C........B.....z..
xb.v..................................C..z........
.........................2.ey.....................
..K.......................y...................s...
..........................................z.......
.......................2.......R........z......F..
O.....d..................D..k.........F...........
..........O..........D.............s....E.........
..o......9....................D........C..s.......
.....o..r....................c...................m
.......P..............................e...........
..............1......................d............
...................o...............3..............
........................c.......24..........S.....
....x..................c.............F4.........S.
P............N..8.......W.D.......................
.....K............1.8.............................
.........P..............Q...M.....................
9...................R........4...8.....0..........
....n.........................F...................
........Y...n......1......................3..J....
.........................3...8....................
...n...M............................0.......ja...S
....f................................6.....S.....E
....................i.........M...J...............
...r..........q..........1......l..0.....L........
.........f7....i.Nc......................j.......l
..9....Y7......X.........Q...M5....J4...........3.
.Y........NX..I............Q........L.............
......Xw...nb.............0..............l6.......
......b.....f....5..q.....................a..6....
.......5..........iq.9.....p..........a...........
........X........I................p...6...........
..................N.............L.........j.......
...b.7.......................p....Q......E........
....Y....7......................p................j
.......r..........................................
.................i...................a............
..w.........5.....................................
......w........I..............J...................
.w............r.....................lL............`

	grid := strings.Split(input, "\n")
	width := len(grid[0])
	height := len(grid)

	antennas := parseGrid(grid)
	fmt.Println("Parsed Antennas:", antennas)

	antinodes := findAntinodes(antennas, width, height)
	fmt.Println("Unique antinodes:", len(antinodes))
}

func parseGrid(grid []string) map[rune][]Point {
	antennas := make(map[rune][]Point)
	for y, row := range grid {
		for x, char := range row {
			if char != '.' {
				antennas[char] = append(antennas[char], Point{x, y})
			}
		}
	}
	return antennas
}

func findAntinodes(antennas map[rune][]Point, width, height int) map[Point]struct{} {
	antinodes := make(map[Point]struct{})

	for char, points := range antennas {
		fmt.Printf("Processing frequency '%c' with %d antennas\n", char, len(points))
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				p1, p2 := points[i], points[j]
				fmt.Printf("Checking alignment: (%d, %d) and (%d, %d)\n", p1.x, p1.y, p2.x, p2.y)

				if isAligned(p1, p2) {
					antinode1, antinode2 := calculateAntinodes(p1, p2)
					fmt.Printf("Calculated antinodes: (%d, %d) and (%d, %d)\n", antinode1.x, antinode1.y, antinode2.x, antinode2.y)

					if antinode1.x != -1 && inBounds(antinode1, width, height) {
						antinodes[antinode1] = struct{}{}
					}
					if antinode2.x != -1 && inBounds(antinode2, width, height) {
						antinodes[antinode2] = struct{}{}
					}
				}
			}
		}
	}

	return antinodes
}

func isAligned(p1, p2 Point) bool {
	return p1.x == p2.x || p1.y == p2.y || abs(p1.x-p2.x) == abs(p1.y-p2.y)
}

func calculateAntinodes(p1, p2 Point) (Point, Point) {
	midX := (p1.x + p2.x) / 2
	midY := (p1.y + p2.y) / 2

	if (p1.x+p2.x)%2 != 0 || (p1.y+p2.y)%2 != 0 {
		return Point{-1, -1}, Point{-1, -1}
	}

	dx := p2.x - midX
	dy := p2.y - midY

	antinode1 := Point{midX - dx, midY - dy}
	antinode2 := Point{midX + dx, midY + dy}

	return antinode1, antinode2
}

func inBounds(p Point, width, height int) bool {
	return p.x >= 0 && p.x < width && p.y >= 0 && p.y < height
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
