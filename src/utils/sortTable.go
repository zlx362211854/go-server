package utils


import "serverGo/src/declare"

type TableList []declare.Table

func (this TableList) Len() int {
	return len(this)
}
func (this TableList) Less(i, j int) bool {
	return this[i].CreateTime > this[j].CreateTime
}
func (this TableList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}