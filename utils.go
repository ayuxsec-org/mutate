package mutate

func prefixString(prefix string, slice []string) []string {
	out := make([]string, len(slice))
	for i, v := range slice {
		out[i] = prefix + v
	}
	return out
}

func dedupeSlice(slice []string) (out []string) {
	seen := make(map[string]struct{})

	for _, v := range slice {
		if _, exists := seen[v]; exists {
			continue
		}
		seen[v] = struct{}{}

		out = append(out, v)
	}

	return
}
