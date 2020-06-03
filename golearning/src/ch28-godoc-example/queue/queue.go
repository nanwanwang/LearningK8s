package queue
//一个队列 可以存放任何类型
type Queue [] interface{}

//往队列中添加一个元素
func (queue *Queue) Push(n interface{} ){
	*queue=append(*queue,n)
}
//删除队列中的一个元素
func (queue *Queue) Pop() interface{} {
	head:=(*queue)[0]
	*queue=(*queue)[1:]
	return head
}
