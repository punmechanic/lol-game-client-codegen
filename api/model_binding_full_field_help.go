/*
 * LoLClient
 *
 * League of Legends Game Client
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// BindingFullFieldHelp Describes a member of a struct.
type BindingFullFieldHelp struct {
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	Offset int32 `json:"offset,omitempty"`
	Optional bool `json:"optional,omitempty"`
	Type BindingFullTypeIdentifier `json:"type,omitempty"`
}
