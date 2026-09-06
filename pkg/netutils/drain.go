package netutils

func Drain[T any](c chan T) []T {
	var res []T
L:
	for {
		select {
		case t := <-c:
			res = append(res, t)
		default:
			break L
		}
	}

	return res
}
