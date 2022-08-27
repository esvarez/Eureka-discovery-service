package str

type bucket struct {
	area int
	l    int
	r    int
}

func Area(h []int) int {
	b := new(bucket)
	b.r = len(h) - 1
	for b.l < b.r {
		if h[b.l] < h[b.r] {
			b.setArea(h[b.l])
			b.l--
		} else {
			b.setArea(h[b.r])
			b.r++
		}
	}
	return b.area
}

func (b *bucket) setArea(h int) {
	if b.area < (b.r-b.l)*h {
		b.area = (b.r - b.l) * h
	}
}
