/**
트리의 지름 = 임의의 두 점 사이의 거리 중 가장 긴 것을 의미 한다.
단순히 트리의 지름을 묻는 문제이므로 -> 각 점 마다 모든 점에 대한 가장 짧은 거리 중 긴 길이를 알 필요는 없다.
한 점(A)을 기준으로 가장 먼 점(B)을 찾고 -> 해당 점(B)을 기준으로 다시 가장 먼 점(C)을 찾는다.
C <-> B의 거리가 트리의 지름이 된다.
*/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	MAX_INT_VALUE = math.MaxInt64
	MIN_INT_VALUE = -math.MaxInt64 - 1
)

type Edge struct {
	connected int
	distance  int
}

type Node struct {
	connected []Edge
}

var nodeCounts int
var nodes []Node
var visited []bool

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &nodeCounts)

	prepareNodes()

	for i := 0; i < nodeCounts; i++ {
		var nodeId, from, to int
		fmt.Fscan(reader, &nodeId)

		for ; ; {
			fmt.Fscan(reader, &from)
			if from == -1 {
				break
			}

			fmt.Fscan(reader, &to)
			createEdge(nodeId, from, to)
		}
	}

	answer := getAnswer()
	fmt.Printf("%d\n", answer)
}

func prepareNodes() {
	nodes = make([]Node, nodeCounts+1)
}

func createEdge(nodeId int, to int, distance int) {
	nodes[nodeId].connected = append(nodes[nodeId].connected, Edge{connected: to, distance: distance})
}

func getAnswer() int {
	farthestNodeIdFromFirstNode, _ := getFarthestNodeAndDistance(1)
	_, diameter := getFarthestNodeAndDistance(farthestNodeIdFromFirstNode)
	return diameter
}

var farthestNodeId, farthestDistance int

func getFarthestNodeAndDistance(nodeId int) (int, int) {
	initializeBeforeDfs()
	dfs(nodeId, 0)
	return farthestNodeId, farthestDistance
}

func initializeBeforeDfs() {
	visited = make([]bool, nodeCounts+1)
	farthestNodeId, farthestDistance = MAX_INT_VALUE, MIN_INT_VALUE
}

func dfs(nodeId int, distance int) {
	if visited[nodeId] {
		return
	}
	visited[nodeId] = true

	if farthestDistance < distance {
		farthestDistance = distance
		farthestNodeId = nodeId
	}

	for _, edge := range nodes[nodeId].connected {
		dfs(edge.connected, distance+edge.distance)
	}
}
