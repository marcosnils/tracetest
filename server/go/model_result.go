/*
 * Project X
 *
 * OpenAPI definition for project X endpoint and resources
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

type Result struct {

	// ID
	Id string `json:"id,omitempty"`

	// trace ID
	Traceid string `json:"traceid,omitempty"`

	// parent span ID
	Spanid string `json:"spanid,omitempty"`

	Successful Assertion `json:"successful,omitempty"`

	Failed Assertion `json:"failed,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	CompletedAt time.Time `json:"completedAt,omitempty"`
}

// AssertResultRequired checks if the required fields are not zero-ed
func AssertResultRequired(obj Result) error {
	if err := AssertAssertionRequired(obj.Successful); err != nil {
		return err
	}
	if err := AssertAssertionRequired(obj.Failed); err != nil {
		return err
	}
	return nil
}

// AssertRecurseResultRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Result (e.g. [][]Result), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseResultRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aResult, ok := obj.(Result)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertResultRequired(aResult)
	})
}
