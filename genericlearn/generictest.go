package genericlearn

import "strconv"

type Number interface {
	int | int32 | int64 | uint32 | uint64 | float64
}

func Str2Number[N Number](strNumber string) (N, error) {

	var n N
	switch (interface{})(n).(type) {
	case int:
		cn, err := strconv.Atoi(strNumber)
		return N(cn), err
	case int32:
		cn, err := strconv.ParseInt(strNumber, 10, 32)
		return N(cn), err
	case int64:
		cn, err := strconv.ParseInt(strNumber, 10, 64)
		return N(cn), err
	case uint32:
		cn, err := strconv.ParseUint(strNumber, 10, 32)
		return N(cn), err
	case uint64:
		cn, err := strconv.ParseUint(strNumber, 10, 64)
		return N(cn), err
	case float64:
		cn, err := strconv.ParseFloat(strNumber, 64)
		return N(cn), err
	}
	return 0, nil
}
