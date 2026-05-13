package base

type Node struct {
	Key   string
	Value string
	Prev  *Node
	Next  *Node
}

type LruCache struct {
	count int
	size  int
	Head  *Node
	Tail  *Node
}

func NewLruCache(size int) *LruCache {
	return &LruCache{size: size}
}

func (l *LruCache) Put(key string, value string) {
	if l.size <= 0 {
		return
	}
	if current := l.Get(key); current != nil {
		*current = value
		return
	}
	newNode := &Node{Key: key, Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		l.count++
		return
	}

	newNode.Next = l.Head
	l.Head.Prev = newNode
	l.Head = newNode

	if l.count == l.size {
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
	} else {
		l.count++
	}
}

func (l *LruCache) Get(key string) *string {
	for current := l.Head; current != nil; current = current.Next {
		if current.Key != key {
			continue
		}

		if current == l.Head {
			return &current.Value
		}

		if current == l.Tail {
			l.Tail = current.Prev
			l.Tail.Next = nil
		} else {
			current.Prev.Next = current.Next
			current.Next.Prev = current.Prev
		}

		current.Prev = nil
		current.Next = l.Head
		l.Head.Prev = current
		l.Head = current

		return &current.Value
	}

	return nil
}
