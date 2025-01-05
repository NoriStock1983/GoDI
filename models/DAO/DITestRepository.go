package dao

type DITestImpl struct{}

func (d DITestImpl) GetAllDiTest() (string, error) {
	return "test", nil
}
