package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
)

type ClietnMessage struct {
	Id       int
	Name     string
	Messages string
	Gender   int
	Age      int
	Tall     float64
	Weight   float64
	Fatrate  float64
}

func ClientGender(id int, name string) *ClientMessage {
	var gender int
	term := rand.Float64()
	if term > 0.5 {
		gender = 1

	} else {
		gender = 0
	}
	return &ClientMessage{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    int(rand.Uint64() % 100),
		Tall:   1.0 + 1.0*rand.Float64(),
		Weight: 55.0 + 55.0*rand.Float64(),
	}

}
func (c *ClietnMessage) calFatRate() {
	bmi := c.Weight / (c.Tall * c.Tall)
	c.Fatrate = 1.2*bmi + 0.22*float64(c.Age) - 5.4 - 10.88*float64(c.Gender)
	return
}

func main() {
	NumberClients := 100
	server := NServer()

	wait := sync.WaitGroup{}
	wait.Add(NumberClients)
	for i := 0; i < NumberClients; i++ {
		id := i
		name := "user=" + strconv.Itoa(i)
		cc := ClientGender(id, name)
		go func(cc *ClientMessage) {
			defer wait.Done()
			for j := 0; j < 100; j++ {
				cc.calFatRate()
				cc.Message = fmt.Sprintf("我的信息是 %s:我的体脂率是: %.2f", cc.Name, cc.FatRate)
				err := server.Publish(*cc)
				if err != nil {
					log.Fatalln("上传失败:", err)
				}
			}

		}(cc)
	}

	wait.Wait()

	wait.Add(NumberClients)
	for i := 0; i < NumberClients; i++ {
		go func(i int) {
			defer wait.Done()
			for j := 0; j < 20; j++ {
				id := 5*i + int(rand.Uint64()%10)
				err := server.Delete(id)
				if err != nil {
					log.Fatalln("删除失败:", err)
				}
				res, err := server.Query()
				if err != nil {
					log.Fatalln("提交数据失败:", err)
				}
				term, _ := json.Marshal(res)
				fmt.Println(string(term))
			}
		}(i)
	}
	wait.Wait()

}
