package repositories

import (
	"log"
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

type PedidosCacheData struct {
	Pedidos *[]models.Pedido
	Err     *error
}

var PedidoCacheStore = cache.New(10*time.Minute, 1*time.Minute)

func NewPedidoCachedRepository() *PedidoCachedRepository {
	return &PedidoCachedRepository{}
}

func NewPedidoCacheData() *PedidoCacheData {
	return &PedidoCacheData{}
}

func NewPedidosCacheData() *PedidosCacheData {
	return &PedidosCacheData{}
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

func (repo *PedidoCachedRepository) FetchAll(a ...interface{}) (*[]models.Pedido, error) {

	cacheKey := "gshghsgashgashaghags"

	cacheData := NewPedidosCacheData()

	if value, found := PedidoCacheStore.Get(cacheKey); found {
		cacheData := value.(*PedidosCacheData)
		return cacheData.Pedidos, *cacheData.Err
	}

	pedidos, err := repo.PedidoRepository.FetchAll(a...)

	// page, error := utils.GetPaginatorParams(1, 25, []{pedidos})

	log.Printf("%T", pedidos)

	cacheData.Pedidos = pedidos
	cacheData.Err = &err

	PedidoCacheStore.Add(cacheKey, cacheData, 15*time.Second)

	return pedidos, err

}
