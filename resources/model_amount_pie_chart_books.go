/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AmountPieChartBooks struct {
	Key
	Attributes AmountPieChartBooksAttributes `json:"attributes"`
}
type AmountPieChartBooksResponse struct {
	Data     AmountPieChartBooks `json:"data"`
	Included Included            `json:"included"`
}

type AmountPieChartBooksListResponse struct {
	Data     []AmountPieChartBooks `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
}

// MustAmountPieChartBooks - returns AmountPieChartBooks from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAmountPieChartBooks(key Key) *AmountPieChartBooks {
	var amountPieChartBooks AmountPieChartBooks
	if c.tryFindEntry(key, &amountPieChartBooks) {
		return &amountPieChartBooks
	}
	return nil
}
