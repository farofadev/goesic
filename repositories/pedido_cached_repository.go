package repositories

import (
	"time"

	"github.com/farofadev/goesic/models"
	"github.com/patrickmn/go-cache"
)

type PedidoCachedRepository struct {
	PedidoRepository
}

type PedidoCacheData struct {
	Pedido *models.Pedido
	Err    *error
}

var PedidoCacheStore = cache.New(10*time.Minute, 1*time.Minute)

func NewPedidoCachedRepository() *PedidoCachedRepository {
	return &PedidoCachedRepository{}
}

func NewPedidoCacheData() *PedidoCacheData {
	return &PedidoCacheData{}
}

func (repo *PedidoCachedRepository) FindById(id string) (*models.Pedido, error) {
	cacheKey := models.GetCacheKeyForPedidoId(id)

	cacheData := NewPedidoCacheData()

	if value, found := PedidoCacheStore.Get(cacheKey); found {
		cacheData := value.(*PedidoCacheData)

		return cacheData.Pedido, *cacheData.Err
	}

	pedido, err := repo.PedidoRepository.FindById(id)

	cacheData.Pedido = pedido
	cacheData.Err = &err

	PedidoCacheStore.Add(cacheKey, cacheData, 15*time.Second)

	return pedido, err
}
