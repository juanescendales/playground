package collection

type Node struct {
	key   string
	value []byte
	next  *Node
	prev  *Node
}

func NewNode(key string, value []byte) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

func (n *Node) SetKey(key string) {
	n.key = key
}

func (n *Node) Key() string {
	return n.key
}

func (n *Node) SetValue(value []byte) {
	n.value = value
}

func (n *Node) Value() []byte {
	return n.value
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func New() *LinkedList {
	return &LinkedList{
		head: nil,
		tail: nil,
	}
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) OrderedKeys() []string {
	keys := make([]string, l.size)
	i := 0
	for node := l.head; node != nil; node = node.next {
		keys[i] = node.Key()
		i++
	}
	return keys
}

func (l *LinkedList) Push(node *Node) {
	if l.head == nil {
		l.head = node
		l.tail = node
		l.size++
		return
	}

	l.head.prev = node
	node.next = l.head
	l.head = node
	l.size++
}

func (l *LinkedList) ToHead(node *Node) {
	if node == l.head {
		return
	}

	l.removeNode(node)
	l.Push(node)
}

func (l *LinkedList) Pop() *Node {
	if l.tail == nil {
		return nil
	}
	node := l.tail
	l.removeNode(l.tail)
	return node
}

func (l *LinkedList) removeNode(node *Node) {
	if node == l.head {
		l.head = node.next
		if l.head != nil {
			l.head.prev = nil
		}
		l.size--
		return
	}

	if node == l.tail {
		l.tail = node.prev
		if l.tail != nil {
			l.tail.next = nil
		}
		l.size--
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
}
