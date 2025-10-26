package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetParams(r *http.Request) (int, int, error) {
	q := r.URL.Query()
	a := q.Get("a")
	intA, err := strconv.Atoi(a)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid parameter 'a'")
	}
	b := q.Get("b")
	intB, err := strconv.Atoi(b)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid parameter 'b'")
	}
	return intA, intB, nil
}
