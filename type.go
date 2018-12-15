package gbfs

type Boolean bool

func (b Boolean) UnmarshalJSON(data []byte) error {
	s := string(data)

	if s == "1" || s == "true" {
		b = true
	}

	b = false

	return nil
}
