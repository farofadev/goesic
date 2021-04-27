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

// Tempo de cache de cada chave
var PedidoCacheDuration = 15 * time.Second

// Depois desse tempo o cache será apagado por completo (todas as chaves)
var PedidoCacheStoreExpiration = 10 * time.Minute

// De quanto em quanto tempo o gerenciador do cache vai limpar as chaves que já expiraram
var PedidoCacheStoreCleanupInterval = 1 * time.Minute

// Armazenamento do Cache
var PedidoCacheStore = cache.New(PedidoCacheStoreExpiration, PedidoCacheStoreCleanupInterval)

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

	PedidoCacheStore.Add(cacheKey, cacheData, PedidoCacheDuration)

	return pedido, err
}