package utils
import (
	"log"
)

func ErrCheck(err error, params interface{}) (interface{}, error) {
	if err != nil {
		log.Fatalln(err.Error())
		return params, err
	}
	return params, nil
}