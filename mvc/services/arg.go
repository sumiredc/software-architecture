package services

func GetArgOrDefault[V any](v *V, def V) V {
	if v == nil {
		return def
	}
	return *v
}
