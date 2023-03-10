package main

type Array struct {
	Length uint
	Data   map[uint]interface{}
}

func (a Array) Get(position uint) interface{} {
	return a.Data[position]
}

func (a *Array) Append(value interface{}) {
	a.Data[a.Length] = value
	a.Length = a.Length + 1
}

func (a *Array) Map(function func(interface{}, uint) interface{}) Array {
	arr := Array{
		Length: 0,
		Data:   map[uint]interface{}{},
	}
	for i := 0; i < int(a.Length); i++ {
		arr.Append(function(a.Get(uint(i)), uint(i)))
	}
	return arr
}

func (a *Array) Unshift(value interface{}) {
	for i := 1; i < int(a.Length)+1; i++ {
		a.Data[uint(i)] = a.Data[uint(i-1)]
	}
	a.Data[0] = value
	a.Length += 1
}

func (a *Array) Pop() interface{} {
	tmp := a.Data[a.Length-1]
	a.Data[a.Length] = nil
	a.Length -= 1
	return tmp
}

func (a *Array) Concat(arr Array) {
	for i := 0; i < int(a.Length); i++ {
		a.Data[a.Length+uint(i)] = arr.Data[uint(i)]
	}
	a.Length += arr.Length
}

func (a *Array) Splice(position uint, num uint, replacement ...interface{}) {
	i := 0
	for j := 0; j < int(a.Length+uint(len(replacement))); j++ {
		if position <= uint(j) {
			if a.Data[uint(j)] != nil {
				if num == 0 {
					a.Data[uint(j+1)] = a.Data[uint(j)]
					a.Data[uint(j)] = replacement[i]
					i++
				} else {
					a.Data[uint(j)] = replacement[i]
				}
			} else {
				a.Data[uint(j)] = replacement[i]
				i++
			}
		}
	}
	a.Length += uint(len(replacement))
}
