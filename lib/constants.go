package lib

import (
	"fmt"
	"mq/lib/errors"
	"mq/lib/utils"
)

// Exchanges

const (
	ProjectExchange = "ProjectExchange"
)

// Exchange Types

const (
	ExchangeTypeTopic = "topic"
)

const (
	RoutingKeyBook = "Book"
	RoutingKeyShop = "Shop"
)

type WorkerType string

const (
	None WorkerType = ""
	All  WorkerType = "all"
	Book WorkerType = "book"
	Shop WorkerType = "shop"
)

func GetGenericName(routingKey string) string {
	return fmt.Sprintf("%s_%s", routingKey, utils.GenerateUUID())
}

func GetWorkerFromString(value string) (WorkerType, error) {
	switch value {
	case "all":
		return All, nil
	case "book":
		return Book, nil
	case "shop":
		return Shop, nil
	default:
		return None, errors.ErrInvalidWorkerType
	}
}
