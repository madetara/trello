package trello

func testCheckistItem() ChecklistItem {
	c := testClient()
	c.BaseURL = mockResponse("/checklists/5914b/checkItems", "checklistItem_example.json").URL
	checklistItem, err := c.Get
}
