package tailnet

type PreauthKey string

func (k PreauthKey) String() string {
	return string(k)
}

func PreauthKeyFromString(s string) (PreauthKey, error) {
	return PreauthKey(s), nil
}
