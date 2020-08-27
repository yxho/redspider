package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"redspider/engine"
	"redspider/persist"
)

type ItemService struct {
	Client *elastic.Client
}

func (s *ItemService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, item)
	if err == nil {
		*result = "ok"
	}
	return err
}
