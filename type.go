package gbfs

type Boolean bool

func (b *Boolean) UnmarshalJSON(data []byte) (err error) {
	s := string(data)

	if s == "1" || s == "true" {
		*b = true

		return
	}

	*b = false

	return
}
