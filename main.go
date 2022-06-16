package main

import "fmt"

// 定义节点结构
type LinkNode struct {
	data int
	next *LinkNode
	prev *LinkNode
}
// 定义链表结构
type LinkTable struct {
	len int
	head *LinkNode
	tail *LinkNode
}

// 链表初始化
func InitLinkTable() *LinkTable {
	Table := new(LinkTable)
	Table.len = 0
	Table.head = nil
	Table.tail = nil
	return Table
}

// 节点初始化
func NewLinkNode(d int) *LinkNode{
	return &LinkNode{data : d}
}

// 在链表头部插入节点（头插法）
func (Table *LinkTable) PushFront(node *LinkNode) int{
	if Table == nil || node == nil{
		return -1
	}
	if Table.head == nil{
		Table.head = node
		Table.tail = node
		node.next = nil
		node.prev = nil
	}else{
		Table.head.prev = node
		node.next = Table.head
		Table.head = node
	}
	Table.len++
	return 0
}

// 在链表尾部插入节点（尾插法）
func (Table *LinkTable) PushBack(node *LinkNode) int{
	if Table == nil || node == nil{
		return -1
	}
	if Table.head == nil{
		Table.head = node
		Table.tail = node
		node.next = nil
		node.prev = nil
	}else{
		Table.tail.next = node
		node.prev = Table.tail
		Table.tail = node
	}
	Table.len++
	return 0
}

// 在链表中删除某指定节点
func (Table *LinkTable) DelNode(node *LinkNode) int{
	if Table == nil || node == nil || Table.head == nil{
		return -1
	}
	if Table.head == node{
		Table.head = Table.head.next
		Table.len--
		if Table.head == nil{
			Table.tail = nil
		}
		return 0
	}
	q := Table.head
	for q != nil{
		if q.next == node{
			q.next = q.next.next
			q.next.prev = q
			Table.len--
			if Table.len == 0{
				Table.tail = nil
			}
			return 0
		}
		q = q.next
	}
	return -1
}

// 在链表中查找指定节点
func (Table *LinkTable)FindNode(node *LinkNode) bool{
	if Table == nil || node == nil{
		return false
	}
	q := Table.head
	for q != nil{
		if q == node{
			return true
		}
		q = q.next
	}
	return false
}

// 获取链表头节点
func (Table *LinkTable)GetHead() *LinkNode{
	if Table == nil || Table.head == nil{
		return nil
	}
	return Table.head
}

// 获取链表尾节点
func (Table *LinkTable) GetTail() *LinkNode{
	if Table == nil || Table.tail == nil{
		return nil
	}
	return Table.tail
}

// 获取链表长度
func (Table *LinkTable) GetLen() int{
	if Table == nil{
		return 0
	}
	return Table.len
}

// 打印链表
func (Table *LinkTable) PrintTable(){
	if Table == nil || Table.head == nil{
		fmt.Println("为空链")
	}else {
		q := Table.head
		for q != nil{
			fmt.Print(q.data)
			fmt.Print(" -> ")
			q = q.next
		}
	}
	fmt.Println("nil")
}

func main(){
	//初始化链表
	Table := InitLinkTable()
	//添加链表节点 分别使用尾插 头插
	Table.PushBack(NewLinkNode(1))
	p := NewLinkNode(2)
	Table.PushBack(p)
	Table.PushBack(NewLinkNode(3))
	Table.PushFront(NewLinkNode(4))
	//打印链表
	fmt.Println("打印链表:")
	Table.PrintTable()
	//获取链表长度
	fmt.Println("链表长度为：", Table.GetLen())
	//删除指定节点
	fmt.Println("删除数据部分为2的指定节点后打印链表：")
	Table.DelNode(p)
	Table.PrintTable()
	//获取删除节点后链表长度
	fmt.Println("删除节点后链表长度为：", Table.GetLen())
}