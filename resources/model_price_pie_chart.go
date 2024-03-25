/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type PricePieChart struct {
	Key
	Attributes PricePieChartAttributes `json:"attributes"`
}
type PricePieChartResponse struct {
	Data     PricePieChart `json:"data"`
	Included Included      `json:"included"`
}

type PricePieChartListResponse struct {
	Data     []PricePieChart `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustPricePieChart - returns PricePieChart from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPricePieChart(key Key) *PricePieChart {
	var pricePieChart PricePieChart
	if c.tryFindEntry(key, &pricePieChart) {
		return &pricePieChart
	}
	return nil
}
