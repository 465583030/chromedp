package domstorage

// AUTOGENERATED. DO NOT EDIT.

import (
	cdp "github.com/knq/chromedp/cdp"
)

// EventDomStorageItemsCleared [no description].
type EventDomStorageItemsCleared struct {
	StorageID *StorageID `json:"storageId"`
}

// EventDomStorageItemRemoved [no description].
type EventDomStorageItemRemoved struct {
	StorageID *StorageID `json:"storageId"`
	Key       string     `json:"key"`
}

// EventDomStorageItemAdded [no description].
type EventDomStorageItemAdded struct {
	StorageID *StorageID `json:"storageId"`
	Key       string     `json:"key"`
	NewValue  string     `json:"newValue"`
}

// EventDomStorageItemUpdated [no description].
type EventDomStorageItemUpdated struct {
	StorageID *StorageID `json:"storageId"`
	Key       string     `json:"key"`
	OldValue  string     `json:"oldValue"`
	NewValue  string     `json:"newValue"`
}

// EventTypes all event types in the domain.
var EventTypes = []cdp.MethodType{
	cdp.EventDOMStorageDomStorageItemsCleared,
	cdp.EventDOMStorageDomStorageItemRemoved,
	cdp.EventDOMStorageDomStorageItemAdded,
	cdp.EventDOMStorageDomStorageItemUpdated,
}
