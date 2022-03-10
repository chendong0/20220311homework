package main

import (
	"math"
	"sort"
)

type RankProject struct {
	Name    string
	FatRate float64
	Sex     string
	Tall    int
	Weight  float64
	Age     int
}

type FatRateRank struct {
	items []RankProject
}

func (r *FatRateRank) inputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64
	for _, item := range fatRate {
		if minFatRate > item {
			minFatRate = item
		}
	}

	found := false
	for i, item := range r.items {
		if item.Name == name {
			if item.FatRate >= minFatRate {
				item.FatRate = minFatRate
			}
			r.items[i] = item
			found = true
			break
		}
	}
	if !found {
		r.items = append(r.items, RankProject{
			Name:    name,
			FatRate: minFatRate,
		})
	}
}

func (r *FatRateRank) getRank(name string) (rank int, fatRate float64) {
	sort.Slice(r.items, func(i, j int) bool {
		return r.items[i].FatRate < r.items[j].FatRate
	})
	frs := map[float64]struct{}{}
	for _, item := range r.items {
		frs[item.FatRate] = struct{}{}
		if item.Name == name {
			fatRate = item.FatRate
		}
	}
	rankArr := make([]float64, 0, len(frs))
	for k := range frs {
		rankArr = append(rankArr, k)
	}
	BubbleFatrate(&rankArr)
	//sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		if frItem == fatRate {
			rank = i + 1
			break
		}
	}

	return
}
