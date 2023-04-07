/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type StatisticsByBook struct {
	Key
	Attributes StatisticsByBookAttributes `json:"attributes"`
}
type StatisticsByBookResponse struct {
	Data     StatisticsByBook `json:"data"`
	Included Included         `json:"included"`
}

type StatisticsByBookListResponse struct {
	Data     []StatisticsByBook `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
}

// MustStatisticsByBook - returns StatisticsByBook from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustStatisticsByBook(key Key) *StatisticsByBook {
	var statisticsByBook StatisticsByBook
	if c.tryFindEntry(key, &statisticsByBook) {
		return &statisticsByBook
	}
	return nil
}
