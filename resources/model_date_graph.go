/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type DateGraph struct {
	Key
	Attributes DateGraphAttributes `json:"attributes"`
}
type DateGraphResponse struct {
	Data     DateGraph `json:"data"`
	Included Included  `json:"included"`
}

type DateGraphListResponse struct {
	Data     []DateGraph `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustDateGraph - returns DateGraph from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustDateGraph(key Key) *DateGraph {
	var dateGraph DateGraph
	if c.tryFindEntry(key, &dateGraph) {
		return &dateGraph
	}
	return nil
}
