package storage

import "strconv"

func (store *Storage) GetString(index int) (string, error) {
	val, err := store.RedisClient.Get(store.ctx, strconv.Itoa(index)).Result()
	return val, err
}

func (store *Storage) LoadPage(pageIndex int, pageSize int) []string {
	var result []string
	start := pageIndex * pageSize
	for i := start; i < start+pageSize; i++ {
		row, err := store.GetString(i)
		if err != nil {
			result = append(result, row)
		}
	}
	store.SetPulled(true)
	return result
}
