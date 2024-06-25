package average

type AverageNumberStore map[int]int

type AverageNumberCounter struct {
	Count       int
	Average     float64
	LatestValue int
	Store       AverageNumberStore
}

func NewAverageNumberCounter() AverageNumberCounter {
	var avg = AverageNumberCounter{}
	avg.Reset()
	return avg
}

func (avg *AverageNumberCounter) Reset() {
	avg.Count = 0
	avg.Average = 0
	avg.LatestValue = 0
	avg.Store = map[int]int{}
}

func (avg *AverageNumberCounter) incrementStore(num int) {
	avg.Store[num]++
}

func (avg *AverageNumberCounter) AddNewNumber(next int) {
	avg.LatestValue = next
	avg.Average = (avg.Average*float64(avg.Count) + float64(next)) / (float64(avg.Count) + 1)
	avg.Count++
	avg.incrementStore(next)
}

func (avg *AverageNumberCounter) AddNewNumbers(next []int) {
	for _, v := range next {
		avg.AddNewNumber(v)
	}
}
