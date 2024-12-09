package enmime

// Option to configure MIME part encoding
type EncoderOption interface {
	apply(p *Encoder)
}

type Encoder struct {
	// Enforce Quoted encoding when selecting Content Transfer Encoding
	enforceQuotedCTE bool
}

// EnforceCteQuoted sets enforceCteQuotedOption option.
func EnforceCteQuoted(b bool) EncoderOption {
	return enforceCteQuotedOption(b)
}

type enforceCteQuotedOption bool

func (o enforceCteQuotedOption) apply(p *Encoder) {
	p.enforceQuotedCTE = bool(o)
}

func NewEncoder(ops ...EncoderOption) *Encoder {
	e := Encoder{
		enforceQuotedCTE: false,
	}

	for _, o := range ops {
		o.apply(&e)
	}

	return &e
}
