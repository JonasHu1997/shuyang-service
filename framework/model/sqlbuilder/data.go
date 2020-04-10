package sqlbuilder

import (
	"gin-service/framework/util"
)

func (im MapD) T() *MapDTransfer {
	i := 0
	m := new(MapDTransfer)
	data := make([]string, 0, 5)
	for k, v := range im {
		if k != "" {
			m.field = append(m.field, "`"+k+"`")
			data = append(data, "'"+util.AddSlashes(v)+"'")
			i++
		}
	}
	m.value[0] = data
	return m
}

func (arr MapDArr) T() *MapDTransfer {
	i := 0
	m := new(MapDTransfer)
	fieldPos := make(map[string]int, 5)
	for c, im := range arr {
		data := make([]string, i)
		for k, v := range im {
			if k != "" {
				if c == 0 {
					m.field = append(m.field, "`"+k+"`")
					fieldPos[k] = i
					data = append(data, "'"+util.AddSlashes(v)+"'")
					i++
				} else {
					data[fieldPos[k]] = "'" + util.AddSlashes(v) + "'"
				}
			}
		}
		m.value = append(m.value, data)
	}
	return m
}

func (m *MapDTransfer) Field() []string {
	return m.field
}

func (m *MapDTransfer) Value() [][]string {
	return m.value
}
