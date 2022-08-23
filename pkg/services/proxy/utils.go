package proxy

// CompareRequestStructs compare internal fields of ClientRequest structs
func CompareRequestStructs(a *ClientRequest, b *ClientRequest) bool {
	if (a.Method != b.Method) || (a.Url != b.Url) || (len(a.Headers) != len(b.Headers)) {
		return false
	}
	for key, element := range a.Headers {
		bValue, ok := b.Headers[key]
		if !ok {
			return false
		}
		if element != bValue {
			return false
		}
	}
	return true
}

// CompareResponseStructs compare internal fields of ClientResponse structs
func CompareResponseStructs(a *ClientResponse, b *ClientResponse) bool {
	if (a.ID != b.ID) || (a.Status != b.Status) || (a.Length != b.Length) || (len(a.Headers) != len(b.Headers)) {
		return false
	}
	for key, element := range a.Headers {
		bValue, ok := b.Headers[key]
		if !ok {
			return false
		}
		if element != bValue {
			return false
		}
	}
	return true
}
