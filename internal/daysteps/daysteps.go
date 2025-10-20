package daysteps

import "test-repo/spentcalories"

// DaySteps хранит данные о шагах за день.
type DaySteps struct {
	Steps int
}

// Add добавляет шаги из тренировки в DaySteps.
func (ds *DaySteps) Add(input string) error {
	steps, _, _, err := spentcalories.ParseTraining(input)
	if err != nil {
		return err
	}
	ds.Steps += steps
	return nil
}

// Steps возвращает количество шагов.
func (ds *DaySteps) Steps() int {
	return ds.Steps
}
