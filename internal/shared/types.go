package shared

type (
    Element interface {int32 | int64 | float32 | float64}
    Operation[E Element] func(left, right, result []E) int
)
