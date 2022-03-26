package apis


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

func (*Circle) TableName() string  {
	return "tesing db.circle"

}
