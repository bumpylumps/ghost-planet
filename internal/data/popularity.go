package data

import (
	"fmt"
	"strconv"
)

type Popularity int32

func (p Popularity) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d stars", p)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
