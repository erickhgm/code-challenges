package anycallback

type Callback func(n []int) bool

func AnyCallback(n []int, c Callback) bool {
	return c(n)
}
