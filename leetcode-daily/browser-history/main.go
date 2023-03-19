package main

type BrowserHistory struct {
	curr, head *Url
}

type Url struct {
	data string
	prev, next *Url
}


func Constructor(homepage string) BrowserHistory {
	hp := &Url{
		data: homepage,
	}
	b := BrowserHistory{
		head: hp,
	}
	b.curr = b.head
	return b
}


func (this *BrowserHistory) Visit(url string)  {
	newUrl := &Url{
		data: url,
	}

	this.curr.next = newUrl
	newUrl.prev = this.curr

	this.curr = newUrl
}


func (this *BrowserHistory) Back(steps int) string {
	for steps > 0 && this.curr.prev != nil {
		this.curr = this.curr.prev
		steps--
	}

	return this.curr.data
}


func (this *BrowserHistory) Forward(steps int) string {
	for this.curr.next != nil && steps > 0 {
		this.curr = this.curr.next
		steps--
	}

	return this.curr.data
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */