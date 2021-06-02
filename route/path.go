package route

type Path struct {
	Root string
	Spu  string
}

func NewPath() *Path {
	return &Path{
		Spu:  "goods",
		Root: "/",
	}
}
