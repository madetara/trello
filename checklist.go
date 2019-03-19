// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import "fmt"

// Checklist represents Trello card's checklists.
// A card can have one zero or more checklists.
// https://developers.trello.com/reference/#checklist-object
type Checklist struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	IDBoard    string      `json:"idBoard,omitempty"`
	IDCard     string      `json:"idCard,omitempty"`
	Pos        float64     `json:"pos,omitempty"`
	CheckItems []CheckItem `json:"checkItems,omitempty"`
}

// CheckItem is a nested resource representing an item in Checklist.
type CheckItem struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	State       string  `json:"state"`
	IDChecklist string  `json:"idChecklist,omitempty"`
	Pos         float64 `json:"pos,omitempty"`
}

// CheckItemState represents a CheckItem when it appears in CheckItemStates on a Card.
type CheckItemState struct {
	IDCheckItem string `json:"idCheckItem"`
	State       string `json:"state"`
}

// CreateChecklistItem creates specified checklist item
func (c *Client) CreateChecklistItem(item *CheckItem, args Arguments) error {
	path := fmt.Sprintf("checklists/%s/checkItems", item.IDChecklist)
	return c.Post(path, args, CheckItem{})
}
