package daysteps

type DaySteps struct {
    steps int
}

func (ds *DaySteps) Add(input string) error {
    steps, _, _, err := parseTraining(input)
    if err != nil {
        return err
    }
    ds.steps += steps
    return nil
}

func (ds *DaySteps) Steps() int {
    return ds.steps
}
