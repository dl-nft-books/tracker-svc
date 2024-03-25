/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AmountPieChart struct {
	Key
	Attributes AmountPieChartAttributes `json:"attributes"`
}
type AmountPieChartResponse struct {
	Data     AmountPieChart `json:"data"`
	Included Included       `json:"included"`
}

type AmountPieChartListResponse struct {
	Data     []AmountPieChart `json:"data"`
	Included Included         `json:"included"`
	Links    *Links           `json:"links"`
}

// MustAmountPieChart - returns AmountPieChart from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAmountPieChart(key Key) *AmountPieChart {
	var amountPieChart AmountPieChart
	if c.tryFindEntry(key, &amountPieChart) {
		return &amountPieChart
	}
	return nil
}
