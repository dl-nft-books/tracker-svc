/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type NftListItem struct {
	Key
	Attributes NftListItemAttributes `json:"attributes"`
}
type NftListItemResponse struct {
	Data     NftListItem `json:"data"`
	Included Included    `json:"included"`
}

type NftListItemListResponse struct {
	Data     []NftListItem `json:"data"`
	Included Included      `json:"included"`
	Links    *Links        `json:"links"`
}

// MustNftListItem - returns NftListItem from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustNftListItem(key Key) *NftListItem {
	var nftListItem NftListItem
	if c.tryFindEntry(key, &nftListItem) {
		return &nftListItem
	}
	return nil
}
