package helper

type Speeder interface {
	maxSpeed() int
}

func NewCar(speeder Speeder) *Car{
	return &Car{
		Speeder: speeder,
	}
}

type Car struct {
	Speeder Speeder
}

func (c Car) Speed() int {

	defaultSpeed := 80

	if c.Speeder.maxSpeed() < 10 {
		return 20
	}

	if defaultSpeed > c.Speeder.maxSpeed(){
		return c.Speeder.maxSpeed()
	}
	return defaultSpeed
}

type DefaultEngine struct {

}

func (e DefaultEngine) maxSpeed() int{
	return 50
}

type TurboEngine struct {

}

func (e TurboEngine) maxSpeed() int{
	return 100
}