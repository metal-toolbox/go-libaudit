package aucoalesce

// CompoundEventError is returned by CoalesceMessages when a compound
// audit event cannot be parsed.
type CompoundEventError struct {
	message string
}

func (o *CompoundEventError) Error() string {
	return o.message
}
