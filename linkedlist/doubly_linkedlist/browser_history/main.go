package main

import "fmt"

type HistoryNode struct {
	URL  string
	Prev *HistoryNode
	Next *HistoryNode
}

type BrowserHistory struct {
	Head    *HistoryNode
	Current *HistoryNode
}

func NewBrowserHistory(url string) *BrowserHistory {
	initNode := &HistoryNode{
		URL: url,
	}

	return &BrowserHistory{
		Head:    initNode,
		Current: initNode,
	}
}

func (bh *BrowserHistory) Visit(url string) {
	newNode := &HistoryNode{
		URL:  url,
		Prev: bh.Current, // set previous node, if not cannot go back
	}

	// clear forward history
	if bh.Current.Next != nil {
		bh.Current.Next = nil
	}

	// connect current node to new node
	bh.Current.Next = newNode

	// update the current node to new node
	bh.Current = newNode
}

func (bh *BrowserHistory) Back(steps int) string {
	fmt.Printf("You clicked the back button %d times\n", steps)
	for steps > 0 && bh.Current.Prev != nil {
		bh.Current = bh.Current.Prev
		steps--
	}

	return bh.CurrentUrl()
}

func (bh *BrowserHistory) Forward(steps int) string {
	fmt.Printf("You clicked the forward button %d times\n", steps)
	for steps > 0 && bh.Current.Next != nil {
		bh.Current = bh.Current.Next
		steps--
	}

	return bh.CurrentUrl()
}

func (bh *BrowserHistory) CurrentUrl() string {
	return bh.Current.URL
}

func (bh *BrowserHistory) PrintHistory() {
	fmt.Println("Browsing History:")
	node := bh.Head
	for node != nil {
		prefix := ""
		if node == bh.Current {
			prefix = "* "
		}
		fmt.Printf("%s%s\n", prefix, node.URL)
		node = node.Next
	}
	fmt.Println()
}

func main() {
	bh := NewBrowserHistory("start.com")

	bh.Visit("google.com")
	bh.Visit("openai.com")
	bh.Visit("github.com")
	bh.Back(1)                    // 当前是 openai.com
	bh.Back(1)                    // 当前是 google.com
	bh.Forward(1)                 // 当前是 openai.com
	bh.Visit("stackoverflow.com") // 当前是 stackoverflow.com，github.com 被清除
	bh.PrintHistory()
	// 输出：
	// start.com
	// google.com
	// openai.com
	// * stackoverflow.com

}
