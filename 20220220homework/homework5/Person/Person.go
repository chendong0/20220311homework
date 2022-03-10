package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const Male string = "M"
const Female string = "F"

var sexArray = [2]string{Male, Female}

type Person struct {
	Name           string  `json:"name"`
	Sex            string  `json:"sex"`
	Tall           float64 `json:"tall"`
	Weight         float64 `json:"weight"`
	Age            int     `json:"age"`
	CurrentFatRate float64 `json:"current-fat-rate"`
	rank           int
}

func (p *Person) UpdateTallWeightAge(reqCtx context.Context, machine *MutexRank) {

	for {
		select {
		case <-reqCtx.Done():
			fmt.Println("[", p.Name, "]", "[UpdateTallWeightAge]", "Timeout! Exiting")
			return
		default:

			p.Tall = randFloat(1, 2)
			p.Weight = randFloat(40, 80)
			p.Age = randInt(15, 60)
			p.CurrentFatRate = calFatRate(p.Sex, p.Tall, p.Weight, p.Age)

			machine.updateRank()

			log.Printf("[Update Tall/Weight/Age] %+v", p)

		}

	}

}

func (p Person) Register(m *MutexRank) {
	m.register(p)
}

func calFatRate(sex string, tall float64, weight float64, age int) (currentFatRate float64) {
	var sexFactor float64

	bmi := weight / (tall * tall)

	if sex == Female {
		sexFactor = 0.0
	} else {
		sexFactor = 1.0
	}
	currentFatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(sexFactor)) / 100
	return
}

func (p Person) GetRank(reqCtx context.Context, manager *MutexRank) {
	for {
		select {
		case <-reqCtx.Done():
			fmt.Println("[", p.Name, "]", "[Query rank]", "Timeout! Exiting")
			return
		default:
			rank, ok := manager.getRank(p.Name)
			if ok {
				log.Println("[", p.Name, "]", "[Query rank]", rank, ",", p.CurrentFatRate, " ", time.Now())
			} else {
				log.Println("[", p.Name, "]", "[Query rank]", "  not found!")
			}

		}

	}
}

func randFloat(min, max float64) float64 {
	if min < 0 {
		min = 0.0
	}
	return min + rand.Float64()*(max-min)
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func NewPerson(name string) *Person {

	tall := randFloat(1, 2)
	weight := randFloat(40, 80)
	age := randInt(15, 60)
	sexIndex := randInt(0, 2)
	sex := sexArray[sexIndex]
	return &Person{
		Name:           name,
		Sex:            sex,
		Tall:           tall,
		Weight:         weight,
		Age:            age,
		CurrentFatRate: calFatRate(sex, tall, weight, age),
	}

}
