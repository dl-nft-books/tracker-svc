/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type PricePieChartTokens struct {
	Key
	Attributes PricePieChartTokensAttributes `json:"attributes"`
}
type PricePieChartTokensResponse struct {
	Data     PricePieChartTokens `json:"data"`
	Included Included            `json:"included"`
}

type PricePieChartTokensListResponse struct {
	Data     []PricePieChartTokens `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
}

// MustPricePieChartTokens - returns PricePieChartTokens from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPricePieChartTokens(key Key) *PricePieChartTokens {
	var pricePieChartTokens PricePieChartTokens
	if c.tryFindEntry(key, &pricePieChartTokens) {
		return &pricePieChartTokens
	}
	return nil
}
