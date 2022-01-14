package cache

type TemporaryCache struct {
	wrapper map[string]interface{}
}

func (cache *TemporaryCache) Put(key string, value interface{}) {
	if cache.wrapper == nil {
		cache.wrapper = make(map[string]interface{})
	}
	cache.wrapper[key] = value
}

func (cache *TemporaryCache) Get(key string) interface{} {
	var temporary = cache.wrapper[key]
	if temporary != nil {
		delete(cache.wrapper, key)
		return temporary
	}
	return nil
}
