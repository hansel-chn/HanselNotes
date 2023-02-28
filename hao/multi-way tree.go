package main

import (
	"fmt"
	"strings"
)

//type tree struct {
//	root treeNode
//}

// type treeMap map[int]treeNode
var threshold int

type treeNode struct {
	val string
	//containedUrl []string
	childSet  map[string]*treeNode
	childNums int
}

type routeTmpl struct {
	dynamicRouteTmpl []string
	staticRouteTmpl  []string
}

func splitUrl(url string) (urlSegment []string) {
	return strings.Split(url, "/")
}

func insertTreeNode(node *treeNode, urlSegment []string) {
	for _, seg := range urlSegment {
		if _, ok := node.childSet[seg]; !ok {
			node.childNums++
			node.childSet[seg] = &treeNode{
				val: seg,
			}
		}
		node = node.childSet[seg]
	}
}

func combineNode(node1 *treeNode, node2 *treeNode) {
	if nil == node2.childSet {
		return
	}
	for key, node2Child := range node2.childSet {
		if _, ok := node1.childSet[key]; !ok {
			node1.childSet[key] = node2Child
			node1.childNums++
		} else {
			combineNode(node1.childSet[key], node2Child)
		}

	}
}

// combine childSet
func combineChildSet(nodeSet map[string]*treeNode) *treeNode {
	var newNode *treeNode
	for _, node := range nodeSet {
		if nil == newNode {
			newNode = node
			continue
		}
		combineNode(newNode, node)
	}
	return newNode
}

func aggregateTree(node *treeNode) {
	for nil != node.childSet {
		if node.childNums > threshold {
			//	combine
			node.childSet = map[string]*treeNode{"*": combineChildSet(node.childSet)}
			node.childNums = 1
		} else {
			// based on feature
		}
		for _, childNode := range node.childSet {
			aggregateTree(childNode)
		}
	}
}

func extractRoute(treeSet map[int]*treeNode) map[int]*routeTmpl {
	routeMap := make(map[int]*routeTmpl)

	for urlLen, rootNode := range treeSet {
		routeMap[urlLen] = &routeTmpl{
			dynamicRouteTmpl: []string{},
			staticRouteTmpl:  []string{},
		}

		temp := make([]string, 0)
		flag := 0
		var backtrace = func(node *treeNode) {}
		backtrace = func(node *treeNode) {
			if nil == node.childSet {
				if flag > 0 {
					routeMap[urlLen].dynamicRouteTmpl = append(routeMap[urlLen].dynamicRouteTmpl, strings.Join(temp, "/"))
				} else {
					routeMap[urlLen].staticRouteTmpl = append(routeMap[urlLen].staticRouteTmpl, strings.Join(temp, "/"))
				}
				return
			}
			for key, childNode := range node.childSet {
				if "*" == key {
					flag++
				}
				temp = append(temp, key)
				backtrace(childNode)

				if "*" == key {
					flag--
				}
				temp = temp[:len(temp)-1]
			}

		}
		backtrace(rootNode)
	}
	return routeMap
}

func main() {
	urls := make([]string, 0)
	// 不同长度路由树的集合
	treeSet := make(map[int]*treeNode)
	// 不同长度路由tmpl的集合

	// create url tree
	for _, url := range urls {
		urlSegment := splitUrl(url)

		if _, ok := treeSet[len(urlSegment)]; !ok {
			treeSet[len(urlSegment)] = &treeNode{}
		}
		rootNode := treeSet[len(urlSegment)]
		insertTreeNode(rootNode, urlSegment)
	}

	for _, rootNode := range treeSet {
		aggregateTree(rootNode)
	}

	tmpl := extractRoute(treeSet)
	fmt.Println(tmpl)

}
