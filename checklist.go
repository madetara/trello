// Copyright © 2016 Aaron Longwell
//
// Use of this source code is governed by an MIT licese.
// Details in the LICENSE file.

package trello

import (
	"errors"
	"fmt"
)

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

var allowedFilters = map[string]struct{}{
	"":     struct{}{},
	"all":  struct{}{},
	"none": struct{}{},
}

var allowedFields = map[string]struct{}{
	"":         struct{}{},
	"all":      struct{}{},
	"name":     struct{}{},
	"nameData": struct{}{},
	"pos":      struct{}{},
	"state":    struct{}{},
	"type":     struct{}{},
}

// CheckItemState represents a CheckItem when it appears in CheckItemStates on a Card.
type CheckItemState struct {
	IDCheckItem string `json:"idCheckItem"`
	State       string `json:"state"`
}

// GetChecklistitems gets items from specified checklist
// Allowed filters: all, none
// Allowed fields: all, name, nameData, pos, state, type
func (c *Client) GetChecklistitems(id, filter, fields string) (result []CheckItem, err error) {
	if _, ok := allowedFields[fields]; !ok {
		err = errors.New("not allowed field used")
		return
	}
	if _, ok := allowedFilters[filter]; !ok {
		err = errors.New("not allowed filter used")
		return
	}

	path := fmt.Sprintf("checklists/%s/checkItems", id)
	args := Arguments{
		"filter": filter,
		"fields": fields,
	}
	err = c.Get(path, args, result)
	return
}

// func (c *Client) GetChecklistitem(checklistID, checklistItemID, fields string) (result *CheckItem, err error) {

// }

// CreateChecklistItem creates specified checklist item
//
// Takes item's "name", "pos" ("top" or "bottom"), and "checked" (true of false) as arguments
func (c *Client) CreateChecklistItem(item *CheckItem, args Arguments) (result *CheckItem, err error) {
	path := fmt.Sprintf("checklists/%s/checkItems", item.IDChecklist)
	err = c.Post(path, args, result)
	return
}
