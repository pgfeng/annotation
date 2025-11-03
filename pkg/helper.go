package pkg

import "strings"

// ParseKeyValues parses a key=value list into a map (keys lowercased).
// Values can be quoted with double quotes and may contain commas.
// Example input: `name="名称", required=true, default="默认值", summary="参数简介"`
func ParseKeyValues(s string) map[string]string {
	out := make(map[string]string)
	parts := SplitTopLevelComma(s)
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		idx := strings.Index(p, "=")
		if idx < 0 {
			continue
		}
		k := strings.ToLower(strings.TrimSpace(p[:idx]))
		val := strings.TrimSpace(p[idx+1:])

		// If wrapped with double quotes, remove them and unescape common sequences.
		if len(val) >= 2 && val[0] == '"' && val[len(val)-1] == '"' {
			inner := val[1 : len(val)-1]
			inner = strings.ReplaceAll(inner, `\"`, `"`)
			inner = strings.ReplaceAll(inner, `\\`, `\`)
			val = inner
		}
		out[k] = val
	}
	return out
}

// SplitTopLevelComma splits by commas but ignores commas inside double quotes.
func SplitTopLevelComma(s string) []string {
	var parts []string
	var sb strings.Builder
	inQuotes := false
	escaped := false
	for _, r := range s {
		if escaped {
			sb.WriteRune(r)
			escaped = false
			continue
		}
		if r == '\\' {
			escaped = true
			sb.WriteRune(r)
			continue
		}
		if r == '"' {
			inQuotes = !inQuotes
			sb.WriteRune(r)
			continue
		}
		if r == ',' && !inQuotes {
			parts = append(parts, sb.String())
			sb.Reset()
			continue
		}
		sb.WriteRune(r)
	}
	rest := strings.TrimSpace(sb.String())
	if rest != "" {
		parts = append(parts, rest)
	}
	return parts
}
