package types

func trimSpaces(part string) string {
	start := 0
	end := len(part) - 1
	for start <= end && (part[start] == ' ' || part[start] == '\t' || part[start] == '\n' || part[start] == '\r') {
		start++
	}
	for end >= start && (part[end] == ' ' || part[end] == '\t' || part[end] == '\n' || part[end] == '\r') {
		end--
	}
	if start > end {
		return ""
	}
	return part[start : end+1]
}
