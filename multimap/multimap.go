package multimap

type MultiMap map[string][]string

func (m MultiMap) Get(k string) string {
	if _, ok := m[k]; ok {
		return m[k][0]
	}
	return ""
}

func (m MultiMap) Add(k, v string) {
	m[k] = append(m[k], v)
}
