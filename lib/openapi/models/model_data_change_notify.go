/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type DataChangeNotify struct {
	OriginalCallbackReference []string     `json:"originalCallbackReference,omitempty" bson:"originalCallbackReference"`
	UeId                      string       `json:"ueId,omitempty" bson:"ueId"`
	NotifyItems               []NotifyItem `json:"notifyItems,omitempty" bson:"notifyItems"`
}
