/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ChainPieChartChains struct {
	Key
	Attributes ChainPieChartChainsAttributes `json:"attributes"`
}
type ChainPieChartChainsResponse struct {
	Data     ChainPieChartChains `json:"data"`
	Included Included            `json:"included"`
}

type ChainPieChartChainsListResponse struct {
	Data     []ChainPieChartChains `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
}

// MustChainPieChartChains - returns ChainPieChartChains from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustChainPieChartChains(key Key) *ChainPieChartChains {
	var chainPieChartChains ChainPieChartChains
	if c.tryFindEntry(key, &chainPieChartChains) {
		return &chainPieChartChains
	}
	return nil
}
