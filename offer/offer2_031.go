package offer

type LRUCache struct {
	m        map[int]*Node22
	head     *Node22
	tail     *Node22
	len      int
	capacity int
}
type Node22 struct {
	key   int
	value int
	pre   *Node22
	next  *Node22
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, len: 0, m: make(map[int]*Node22)}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.m[key]
	if !ok {
		return -1
	}
	if node == this.tail {
		return node.value
	} else if node == this.head { //如果该结点为头结点
		this.head.next.pre = nil
		this.head = this.head.next
	} else if node != this.tail { //如果该结点为中间结点
		node.pre.next = node.next
		node.next.pre = node.pre
	} //从当中删去
	//添加到尾结点
	node.pre = this.tail
	this.tail.next = node

	this.tail = node
	return this.tail.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.m[key]
	if ok { //是否已经存在
		node.value = value
		if node == this.tail {
			return
		} else if node == this.head {
			this.head.next.pre = nil
			this.head = this.head.next
		} else if node != this.tail {
			node.next.pre = node.pre
			node.pre.next = node.next
		}
		node.pre = this.tail
		this.tail.next = node

		this.tail = node
		return
	} else { //不存在

		this.m[key] = &Node22{key: key, value: value}
		node = this.m[key]
		if this.len < this.capacity { //是否已经满
			this.len++
			if this.tail != nil {
				this.tail.next = node
				node.pre = this.tail

			} else {
				this.tail = node
				this.head = node
			}
			this.tail = node

		} else {
			delete(this.m, this.head.key)
			this.head = this.head.next
			this.tail.next = node
			node.pre = this.tail

			this.tail = node
		}
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
