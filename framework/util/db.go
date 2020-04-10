package util

// AddSlashes 转义sql语句
func AddSlashes(v string) string {
	pos := 0
	buf := make([]byte, len(v)*2)
	for i := 0; i < len(v); i++ {
		c := v[i]
		if c == '\'' || c == '"' || c == '\\' {
			buf[pos] = '\\'
			buf[pos+1] = c
			pos += 2
		} else {
			buf[pos] = c
			pos++
		}
	}
	return string(buf[:pos])
}
