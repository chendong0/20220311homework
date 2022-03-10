package main

import (
	"fmt"
	"log"
)

func main() {
	input := &inputSQL{}
	calc := &Calc{}
	rank := &FatRateRank{}
	records := NewRecord("/31565/chendong/test.txt")

	for {
		ran := input.GetInput()
		if err := records.savePersonalInformation(ran); err != nil {
			log.Fatal("fail to save file：", err)
		}
		fr, err := calc.FatRate(ran)
		if err != nil {
			log.Fatal("calc fail：", err)
		}
		rank.inputRecord(ran.Name, fr)

		rankResult, _ := rank.getRank(ran.Name)
		fmt.Println(" rank  output ：", rankResult)
	}
}
