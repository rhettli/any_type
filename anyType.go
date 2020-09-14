package any_type

import "strconv"

type SafeList struct {
	list  []*AnyType
	index int
}

func (ls *SafeList) Get(key int) *AnyType {
	return ls.list[key]
}
func (ls *SafeList) Next() *AnyType {
	if ls.index == len(ls.list) {
		return &AnyType{V: ""}
	}
	o := ls.list[ls.index]
	ls.index++
	return o
}

func (ls *SafeList) First() *AnyType {
	ls.index = 1
	return ls.list[0]
}

type AnyType struct {
	V string
}

func (receiver *AnyType) Int() int {
	i, _ := strconv.Atoi(receiver.V)
	return i
}
func (receiver *AnyType) Float64() float64 {
	i, _ := strconv.ParseFloat(receiver.V, 10)
	return i
}
func (receiver *AnyType) Uint64() uint64 {
	i, _ := strconv.ParseUint(receiver.V, 10, 32)
	return i
}

func (receiver *AnyType) Uint8() uint8 {
	i, _ := strconv.ParseUint(receiver.V, 10, 32)
	return uint8(i)
}

func (receiver *AnyType) Float32() float32 {
	i, _ := strconv.ParseFloat(receiver.V, 10)
	return float32(i)
}

func (receiver *AnyType) Bool() bool {
	return receiver.V == "true"
}

func (receiver *AnyType) String() string {
	return receiver.V
}
