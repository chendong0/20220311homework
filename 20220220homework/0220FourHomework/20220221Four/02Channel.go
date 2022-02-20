package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type Persons struct {
	Name        string
	BaseFatRate float64
	NewFatRate  float64
	Ranking     int
}

func ReMutexRank(maxSignUp int) *MutexRank {
	return &MutexRank{
		Ranking: make([]Persons, 0, maxSignUp),
		SignUp:  map[string]Persons{}, //漏了花括号
	}

}

type MutexRank struct {
	Ranking []Persons
	SignUp  map[string]Persons
	locks   sync.RWMutex //替换成sync.Mutex更简单
}

func randNum(minor, maximum float64) float64 {
	if minor < 0 {
		minor = 0.00
	}
	return minor + rand.Float64()*(maximum-minor)
}

func NewRand(Name string, Ranking chan []Persons) {
	slicePersons := <-Ranking
	for _, Persons := range slicePersons {
		if Persons.Name == Name {
			fmt.Printf("%s的体脂最新排名为%d \n", Name, Persons.Ranking)
		}
	}
}

func main() {
	for {
		ChannelPerson := make(chan Persons, 1000)
		ChRanking := make(chan []Persons, 1000)
		slicePersons := []Persons{}
		AllSignUp := 1000

		//Admin := ReMutexRank(AllSignUp)
		rand.Seed(time.Now().Unix())

		wg := sync.WaitGroup{}
		wg.Add(AllSignUp)
		rand.Seed(time.Now().Unix())

		for i := 0; i < AllSignUp; i++ {
			go func(i int, wg *sync.WaitGroup) {
				defer wg.Done()
				var Persons = Persons{
					Name:        fmt.Sprintf("姓名为%d", i),
					BaseFatRate: randNum(0, 0.4),
				}

				ChannelPerson <- Persons

				NewRand(Persons.Name, ChRanking)
			}(i, &wg)
		}

		finishedFileCount := 0
		for Persons := range ChannelPerson {
			finishedFileCount++
			slicePersons = append(slicePersons, Persons)
			if finishedFileCount == AllSignUp {
				close(ChannelPerson)
			}
		}

		sort.Slice(slicePersons, func(i, j int) bool {
			return slicePersons[i].BaseFatRate < slicePersons[j].BaseFatRate
		})

		for i, _ := range slicePersons {
			slicePersons[i].Ranking = i + 1
		}

		for i := 0; i < AllSignUp; i++ {
			ChRanking <- slicePersons
		}

		wg.Wait()
	}

}
