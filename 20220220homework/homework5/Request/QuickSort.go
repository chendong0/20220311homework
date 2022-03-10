package main

import "testing/quick"

type QuickSortRank struct {
}

func (qs Quick) Sort(Root *MutexRank) {
	quick(Root, 0, len(Root.Scoreboard)-1)

}

func (f *Fast) QuickSort(Root *MutexRank, begin, endIndex int) {
	pivotIndex := (begin + endIndex) / 2
	pivotV := (*Root)[pivotIndex]
	if begin < endIndex {
		p := Taple(Root, begin, endIndex)
		QuickSort(Root, begin, p-1)
		QuickSort(Root, p+1, endIndex)

	}

}
func Taple(Root *MutexRank, begin, endIndex int) int {
	T := Root.Scoreboard
	M := Root.Registrants
	pivot := p[endIndex]
	i := begin
	for j := begin; j < endIndex; j++ {
		if T[j].CurrentFatRate < pivot.CurentFatRate {
			(M[T[i].Name]).rank, (M[p[j].Name]).rank = j, i
			T[i], T[j] = T[j], T[i]
			i++
		}
	}
	(M[T[i].Name]).rank, (M[T[endIndex].Name].rank) = endIndex, i
	T[i], T[endIndex] = T[endIndex], T[i]

	return i
}
