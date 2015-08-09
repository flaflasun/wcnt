package wordcount

func (w WordCounters) Len() int {
	return len(w)
}

func (w WordCounters) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w WordCounters) Less(i, j int) bool {
	return w[i].Count < w[j].Count
}
