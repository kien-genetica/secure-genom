package adapter

import "github.com/kien6034/secure-genom/internal/domain/entity"

type DataIDGetter struct{}

func NewDataIDGetter() *DataIDGetter {
	return &DataIDGetter{}
}

func (g *DataIDGetter) GetID(data *entity.Data) string {
	return data.ID
}
