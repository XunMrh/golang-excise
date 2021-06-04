package main

type IntSet struct {
	words []uint64
}

func (s *IntSet) len() int {
	var count int
	for _, word := range s.words {
		for i := 0; i < 64; i++ {
			if (word & (1 << i)) != 0 {
				count++
			}
		}
	}
	return count
}

func (s *IntSet) remove(x int) bool {
	word, index := x/64, x%64
	if word >= len(s.words) {
		index = 0
		return false
	}
	s.words[word] &= (^(1 << index))
	return true
}

func (s *IntSet) clear() {
	for i, _ := range s.words {
		s.words[i] &= 0
	}
}

func (s *IntSet) copy() *IntSet {
	var temp IntSet
	for _, word := range s.words {
		temp.words = append(temp.words, word)
	}
	return &temp
}

func (s *IntSet) addAll(vars ...int) {
	for _, elem := range vars {
		word, index := elem/64, elem%64
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= (1 << index)
	}
}

func (s *IntSet) JiaoJi(t *IntSet) {
	minLen := len(s.words)
	if minLen > len(t.words) {
		minLen = len(t.words)
	}
	for i := 0; i < minLen; i++ {
		s.words[i] &= t.words[i]
	}
	s.words = s.words[:minLen]
}

func (s *IntSet) ChaJi(t *IntSet) {
	for i, _ := range s.words {
		s.words[i] &= ^t.words[i]
	}
	for len(s.words) != 0 && s.words[len(s.words)-1] == 0 {
		s.words = s.words[:len(s.words)-1]
	}
}

func (s *IntSet) BingChaJi(t *IntSet) {
	for i, _ := range t.words {
		if i < len(s.words) {
			s.words[i] ^= t.words[i]
		} else {
			s.words = append(s.words, t.words[i])
		}
	}
	for len(s.words) != 0 && s.words[len(s.words)-1] == 0 {
		s.words = s.words[:len(s.words)-1]
	}
}

func (s *IntSet) Elem() []int {
	var res []int
	for i, _ := range s.words {
		for mask := 0; mask < 64; mask++ {
			if s.words[i]&(1<<mask) != 0 {
				res = append(res, i*64+mask)
			}
		}
	}
	return res
}

const unionsize = 32 << (^uint(0) >> 63)

type mIntSet struct {
	words []uint
}

func (s *mIntSet) add(x int) {
	word, index := x/unionsize, x%unionsize
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= (1 << index)
}
