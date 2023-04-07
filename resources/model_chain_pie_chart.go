/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ChainPieChart struct {
	Key
	Attributes ChainPieChartAttributes `json:"attributes"`
}
type ChainPieChartResponse struct {
	Data     ChainPieChart `json:"data"`
	Included Included      `json:"included"`
}

type ChainPieChartListResponse struct {
	Data     []ChainPieChart `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustChainPieChart - returns ChainPieChart from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustChainPieChart(key Key) *ChainPieChart {
	var chainPieChart ChainPieChart
	if c.tryFindEntry(key, &chainPieChart) {
		return &chainPieChart
	}
	return nil
}
