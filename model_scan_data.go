/*
 * ibkr
 *
 * ibkr API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ScanData struct for ScanData
type ScanData struct {
	ConId int64 `json:"conId,omitempty"`
	Currency string `json:"currency,omitempty"`
	Exchange string `json:"exchange,omitempty"`
	Ticker string `json:"ticker,omitempty"`
}