package main

import (
	"fmt"
	"strings"
)

// Point represents a position in the grid.
type Point struct {
	x, y int
}

// Directions for moving up, down, left, and right.
var directions = []Point{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

// BFS calculates the score for a trailhead by counting reachable '9' positions.
func BFS(start Point, heightMap [][]int) int {
	rows := len(heightMap)
	cols := len(heightMap[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	queue := []Point{start}
	visited[start.x][start.y] = true
	reachableNines := make(map[Point]bool)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		currentHeight := heightMap[current.x][current.y]

		// If we reach a '9', add it to the reachable set.
		if currentHeight == 9 {
			reachableNines[current] = true
			continue
		}

		// Explore all valid directions
		for _, dir := range directions {
			nextX := current.x + dir.x
			nextY := current.y + dir.y

			// Check boundaries and conditions
			if nextX >= 0 && nextX < rows && nextY >= 0 && nextY < cols {
				if !visited[nextX][nextY] && heightMap[nextX][nextY] == currentHeight+1 {
					visited[nextX][nextY] = true
					queue = append(queue, Point{nextX, nextY})
				}
			}
		}
	}

	// Return the count of unique reachable '9' positions
	return len(reachableNines)
}

// DFSWithBacktracking calculates all distinct trails starting from a given trailhead.
func DFSWithBacktracking(start Point, heightMap [][]int, visited [][]bool) int {
	rows := len(heightMap)
	cols := len(heightMap[0])
	var dfs func(Point) int

	// DFS function
	dfs = func(current Point) int {
		currentHeight := heightMap[current.x][current.y]

		// If we reach height 9, this is a distinct trail endpoint.
		if currentHeight == 9 {
			return 1
		}

		// Mark current position as visited
		visited[current.x][current.y] = true
		trails := 0

		// Explore all valid directions
		for _, dir := range directions {
			nextX := current.x + dir.x
			nextY := current.y + dir.y

			// Check boundaries
			if nextX < 0 || nextX >= rows || nextY < 0 || nextY >= cols {
				continue
			}

			// Check if not visited and the height increases by exactly 1
			if !visited[nextX][nextY] && heightMap[nextX][nextY] == currentHeight+1 {
				trails += dfs(Point{nextX, nextY})
			}
		}

		// Backtrack
		visited[current.x][current.y] = false
		return trails
	}

	return dfs(start)
}

// computeRatings calculates both the score and rating for all trailheads.
func computeRatings(heightMap [][]int) (int, int) {
	rows := len(heightMap)
	cols := len(heightMap[0])
	totalScore := 0
	totalRating := 0

	// Visited matrix for DFS
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// Iterate through the height map to find trailheads (height 0)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if heightMap[i][j] == 0 {
				trailhead := Point{i, j}

				// Part One: Count unique reachable '9's
				score := BFS(trailhead, heightMap)
				totalScore += score

				// Part Two: Count distinct hiking trails
				rating := DFSWithBacktracking(trailhead, heightMap, visited)
				totalRating += rating
			}
		}
	}

	return totalScore, totalRating
}

// parseInput converts the multi-line string input into a 2D slice of integers.
func parseInput(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	heightMap := make([][]int, len(lines))
	for i, line := range lines {
		heightMap[i] = make([]int, len(line))
		for j, char := range line {
			// Convert rune to integer (assuming '0' to '9')
			heightMap[i][j] = int(char - '0')
		}
	}
	return heightMap
}

func main() {
	// Input as a multi-line string
	input := `
987123434330121232101001234730123456781067632
876076576521010345692340349823212347892398701
945087689432105676787659856714503210987445610
332196576587654989801456787609654502376530923
211543210298923215432321098128778901430121894
300692340147210106523543210039569876589836765
456781678236103267015693016543410231276745650
576890549345234178106782187612320140345654321
985098432100125089235493498109876056034765012
834127102345456978340362569018765487123876678
123236221976347869651251078729034398101985589
014545340889298958707867897430120987012834432
105965456770107843216950956541231276543124501
896872378761016930345441019876501345678023670
787901069654325321210332398545432330589012981
107821543213034321089206787638901421432103210
215430694102123475670115896129876548901210349
126989780210014984308924925014578037654321458
037878921001235675217833210123669123109452367
549865438901045102346542106548754321278501476
678954987432696201256430087239689870347699985
230143006501787349961021298101236787656788014
123272112981010458872787034010345691875107623
054387623472129867763698125676210010961234510
565694502561036789854567012980387121250129878
676783411051045672343218763901296030343278569
989872123432345891050109654812345145467303450
012763094321056700891760345765432256958912341
103450185789763211709851236876301967843211032
814321276656854345612345654954101878701208983
923434434565956745678036783063210989870345674
874532345410345832989123192178981876781456564
265101654323234901808765013265432185692387565
103216765432101267814554323476501094501893474
232109856321011876923601098789678923432102985
343898707896540945498712367765672310567891078
456789010987231234321203456894581455454986569
556776125670102343100157654503490166303890432
543895434894321765212348983212321876212761201
432104898765010894301054581200110955211654300
301256567656987105498765690341034567300563212
434567430547896234787654785652123498456767843
321798121032345375696543098743096567877854952
210899021121036789781232143456787656928923761
326765430110145678710123232109876543210010890
`

	// Parse the input into a 2D height map
	heightMap := parseInput(input)

	// Compute the scores and ratings
	totalScore, totalRating := computeRatings(heightMap)

	// Output the results for both parts
	fmt.Println("Total Trailhead Score (Part One):", totalScore)
	fmt.Println("Total Trailhead Rating (Part Two):", totalRating)
}
