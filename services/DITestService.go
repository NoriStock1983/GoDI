package services

import (
	dao "GODI/models/DAO"
	"fmt"
)

// interfaces
type IDiTest interface {
	GetAllDiTest() (string, error)
}

func NewDiTest(test dao.DITestImpl) {

	t, err := test.GetAllDiTest()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(t)

}
