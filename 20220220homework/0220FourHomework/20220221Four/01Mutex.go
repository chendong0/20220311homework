package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Persons struct {
	//个人体脂排行榜基础信息
	Name        string
	BaseFatRate float64
	NewFatRate  float64
}

type MutexRank struct {
	Ranking []Persons
	SignUp  map[string]Persons
	locks   sync.RWMutex //替换成sync.Mutex更简单
}

//定义一个函数
func ReMutexRank(maxSignUp int) *MutexRank {
	return &MutexRank{
		Ranking: make([]Persons, 0, maxSignUp),
		SignUp:  map[string]Persons{}, //漏了花括号
	}

}

//创建另外一个函数
func Add(ad []Persons, p Persons, i int) []Persons {
	return append(ad[:i], append([]Persons{p}, ad[i:]...)...)
}

//函数方法实例化
func (m *MutexRank) sign(p Persons) {
	m.locks.Lock()
	defer m.locks.Unlock()

	_, ok := m.SignUp[p.Name] //注册新用户
	if ok == false {
		for i, v := range m.Ranking {
			if p.NewFatRate <= v.NewFatRate {
				m.Ranking = Add(m.Ranking, p, i)
				m.SignUp[p.Name] = p
				return
			}

		}
		m.Ranking = append(m.Ranking, p)
		m.SignUp[p.Name] = p

	}

}

func (m *MutexRank) ExportRanking() {
	for i, personnel := range m.Ranking {
		fmt.Println(i+1, ":", personnel.Name, "  ", personnel.NewFatRate)

	}

}

func (m *MutexRank) topRanking(p Persons) (int, bool) {
	m.locks.RLock()
	defer m.locks.RUnlock()

	for i, persons := range m.Ranking {
		if persons.Name == p.Name {
			return i + 1, true
		}
	}
	return 0, false

}

func Delete(d []Persons, i int) []Persons {
	return append(d[:i], d[i+1:]...)

}

func (m *MutexRank) PrintExportRankingLocd() {
	m.locks.RLock()
	defer m.locks.RLock()
	m.ExportRanking()

}

func (m *MutexRank) AutoUpdateFatRate(p Persons) (int, bool) {
	m.locks.Lock()
	defer m.locks.Unlock()

	_, ok := m.SignUp[p.Name]

	if ok == true {
		for i, v := range m.Ranking { //主意是用for循环,用错if
			if p.Name == v.Name {
				m.Ranking = Delete(m.Ranking, i)
				break
			}
		}
		for i1, v1 := range m.Ranking {
			if p.BaseFatRate <= v1.BaseFatRate {
				m.Ranking = Add(m.Ranking, p, i1)
				return i1 + 1, true
			}

		}
		m.Ranking = append(m.Ranking, p)
		return len(m.Ranking), true
	} else {
		return 0, false
	}

}

func randNum(minor, maximum float64) float64 {
	if minor < 0 {
		minor = 0.00
	}
	return minor + rand.Float64()*(maximum-minor)

}

func main() {

	AllSignUp := 1000
	Finish := make(chan os.Signal, 1)
	signal.Notify(Finish, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	Admin := ReMutexRank(AllSignUp)
	rand.Seed(time.Now().Unix())

	var ww sync.WaitGroup
	ww.Add(AllSignUp)

	for i := 0; i < AllSignUp; i++ {
		go func(i int) {
			defer ww.Done()
			name := fmt.Sprintf("姓名%d", i)
			base := randNum(0, 0.4)
			Admin.sign(Persons{Name: name, BaseFatRate: base, NewFatRate: base})
		}(i)

	}
	ww.Wait()
	Admin.ExportRanking()
	reCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		fmt.Println("主程序:正在取消上下文语境")
		cancel()
	}()

	for _, persons := range Admin.SignUp {
		ww.Add(2)
		go func(persons Persons) {
			defer ww.Done()

			for {
				select {
				case <-reCtx.Done():
					fmt.Println("<", persons.Name, ">", "最新体脂率排行榜", "退出超时")
					return
				default:
					minnorFatRate := persons.BaseFatRate - 0.2
					persons.NewFatRate = randNum(minnorFatRate, persons.BaseFatRate+0.20)
					rank, ok := Admin.topRanking(persons)
					if ok {
						fmt.Println("<", persons.Name, ">", "最新体脂率排行榜", "~", rank, "~", persons.NewFatRate, "//", time.Now())
					} else {
						fmt.Println("<", persons.Name, ">", "最新体脂率排行榜", "无效")
					}
				}

			}
		}(persons)

		go func(persons Persons) {
			defer ww.Done()

			for {
				select {
				case <-reCtx.Done():
					fmt.Println("<", persons.Name, ">", "确认榜单", "+", persons.NewFatRate, " ", time.Now())
					return
				default:
					rank, ok := Admin.topRanking(persons)
					if ok {
						fmt.Println("<", persons.Name, ">", "最新体脂率排行榜", "~", rank, "~", persons.NewFatRate, "//", time.Now())
					} else {
						fmt.Println("<", persons.Name, ">", "确认榜单", "无效")
					}
				}
			}

		}(persons)
	}
	select {
	case <-Finish:
		fmt.Println("主程序终止")
		cancel()
		ww.Wait()
	case <-reCtx.Done():
		fmt.Println("主程序超时退出")
		ww.Wait()
	}
	fmt.Println("终于完成,体脂率排行榜")
	Admin.ExportRanking()

	fmt.Println("程序成功运行.")
}
