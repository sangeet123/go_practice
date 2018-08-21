package main

import (
	"bytes"
	"fmt"
)

type BitSet struct {
	words []int64
}

func (b *BitSet) Add(x int) {
	p, o := x/64, uint(x%64)

	for p >= len(b.words) {
		b.words = append(b.words, 0)
	}

	b.words[p] |= (0x1 << o)
}

func (b *BitSet) IsSet(x int) bool {
	p, o := x/64, uint(x%64)
	return len(b.words) > p && (b.words[p]&(1<<o) != 0)
}

func (b *BitSet) Union(u BitSet) {

	for i, v := range u.words {
		if i < len(b.words) {
			b.words[i] |= u.words[i]
		} else {
			b.words = append(b.words, v)
		}
	}
}

func (b *BitSet) Len() int {
	count := 0
	for _, v := range b.words {
		count += CountBitSet(v)
	}

	return count
}

func (b *BitSet) Remove(x int) {
	p, o := x/64, uint(x%64)

	if p >= len(b.words) {
		return
	}

	b.words[p] &^= (1 << o)
}

func (b *BitSet) Clear() {
	b.words = []int64(nil)
}

func (b *BitSet) Copy() *BitSet {
	r := new(BitSet)

	for _, v := range b.words {
		r.words = append(r.words, v)
	}

	return r
}

func CountBitSet(n int64) int {
	if n == 0 {
		return 0
	}
	return 1 + CountBitSet(n&(n-1))
}

func (b *BitSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	for i, v := range b.words {
		if v == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if v&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}

				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x BitSet

	x.Add(1)
	x.Add(144)
	x.Add(9)

	fmt.Println(x.String())
	fmt.Println(x.Copy())
	x.Remove(144)
	fmt.Println(x.String())
	fmt.Println(x.Len())
}
