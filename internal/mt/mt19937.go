package mt

const (
	n = 624
	m = 397

	notSeeded = n + 1

	hiMask uint32 = 0x80000000 // 1 << 31
	loMask uint32 = 0x7FFFFFFF // ~hiMask

	matrix_a uint32 = 0x9908B0DF
)

type MT19937 struct {
	state []int32
	index int
}

func New(seed int32) *MT19937 {
	res := &MT19937{
		state: make([]int32, notSeeded),
		index: notSeeded,
	}
	res.Seed(seed)
	return res
}

func (mt *MT19937) Seed(seed int32) {
	x := mt.state
	x[0] = seed
	for i := int32(1); i < n; i++ {
		step1 := x[i-1]
		step2 := x[i-1] >> 30
		step3 := step1 ^ step2
		step4 := 0x6C078965 * step3
		step5 := step4 + i

		x[i] = step5
	}
	mt.index = n
}

func (mt *MT19937) Int31() int32 {
	if mt.index >= n {
		if mt.index == notSeeded {
			mt.Seed(4357)
		}
		mt.twist()
	}

	x := mt.state[mt.index]
	x ^= (x >> 11)
	x ^= int32(uint32(x<<7) & 0x9D2C5680)
	x ^= int32(uint32(x<<15) & 0xEFC60000)
	x ^= (x >> 18)

	mt.index++
	return x
}

func (mt *MT19937) twist() {
	_mt := mt.state

	for i := 0; i < n-m; i++ {
		x := twistA(_mt[i], _mt[i+1])
		_mt[i] = twistB(x, _mt[i+m])
	}
	for i := n - m; i < n-1; i++ {
		x := twistA(_mt[i], _mt[i+1])
		_mt[i] = twistB(x, _mt[i+(m-n)])
	}

	x := twistA(_mt[n-1], _mt[0])
	_mt[n-1] = twistB(x, _mt[m-1])

	mt.index = 0
}

func twistA(a, b int32) int32 {
	return int32(uint32(a)&hiMask | uint32(b)&loMask)
}

func twistB(x, a int32) int32 {
	return int32(uint32(a^(x>>1)) ^ (uint32(x&1) * matrix_a))
}
