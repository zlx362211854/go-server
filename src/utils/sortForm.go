package utils

import "serverGo/src/declare"

type FormList []declare.Form

func (this FormList) Len() int {
	return len(this)
}
func (this FormList) Less(i, j int) bool {
	return this[i].CreateTime > this[j].CreateTime
}
func (this FormList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}