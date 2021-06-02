package common

type DisPatch struct {
	Name string
}

func (dis *DisPatch) GetName() string {
	return dis.Name
}
