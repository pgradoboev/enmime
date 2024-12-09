package enmime

type Encoder struct {
	base64Percent int
}

type EncoderOption interface {
	apply(p *Encoder)
}

func Base64PercentValue(n int) EncoderOption {
	return base64PercentValue(n)
}

type base64PercentValue int

func (o base64PercentValue) apply(p *Encoder) {
	p.base64Percent = int(o)
}

func NewEncoder(ops ...EncoderOption) *Encoder {
	e := Encoder{
		base64Percent: 20,
	}

	for _, o := range ops {
		o.apply(&e)
	}

	return &e
}
