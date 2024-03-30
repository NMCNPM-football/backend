// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: club.proto

package gen

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ClubProfileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ClubProfileResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClubProfileResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ClubProfileResponseMultiError, or nil if none found.
func (m *ClubProfileResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ClubProfileResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ClubProfileResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ClubProfileResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClubProfileResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ClubProfileResponseMultiError(errors)
	}

	return nil
}

// ClubProfileResponseMultiError is an error wrapping multiple validation
// errors returned by ClubProfileResponse.ValidateAll() if the designated
// constraints aren't met.
type ClubProfileResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClubProfileResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClubProfileResponseMultiError) AllErrors() []error { return m }

// ClubProfileResponseValidationError is the validation error returned by
// ClubProfileResponse.Validate if the designated constraints aren't met.
type ClubProfileResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClubProfileResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClubProfileResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClubProfileResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClubProfileResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClubProfileResponseValidationError) ErrorName() string {
	return "ClubProfileResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ClubProfileResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClubProfileResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClubProfileResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClubProfileResponseValidationError{}

// Validate checks the field values on ClubProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ClubProfileRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClubProfileRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ClubProfileRequestMultiError, or nil if none found.
func (m *ClubProfileRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ClubProfileRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NameClub

	// no validation rules for NameAward

	// no validation rules for NameStadium

	// no validation rules for SeaSon

	// no validation rules for Achievement

	// no validation rules for OwnerBy

	// no validation rules for UpdateBy

	// no validation rules for UpdateAt

	if len(errors) > 0 {
		return ClubProfileRequestMultiError(errors)
	}

	return nil
}

// ClubProfileRequestMultiError is an error wrapping multiple validation errors
// returned by ClubProfileRequest.ValidateAll() if the designated constraints
// aren't met.
type ClubProfileRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClubProfileRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClubProfileRequestMultiError) AllErrors() []error { return m }

// ClubProfileRequestValidationError is the validation error returned by
// ClubProfileRequest.Validate if the designated constraints aren't met.
type ClubProfileRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClubProfileRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClubProfileRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClubProfileRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClubProfileRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClubProfileRequestValidationError) ErrorName() string {
	return "ClubProfileRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ClubProfileRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClubProfileRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClubProfileRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClubProfileRequestValidationError{}

// Validate checks the field values on ClubProfileResponse_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ClubProfileResponse_Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClubProfileResponse_Data with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ClubProfileResponse_DataMultiError, or nil if none found.
func (m *ClubProfileResponse_Data) ValidateAll() error {
	return m.validate(true)
}

func (m *ClubProfileResponse_Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NameClub

	// no validation rules for NameAward

	// no validation rules for Shorthand

	// no validation rules for NameStadium

	// no validation rules for Achievement

	// no validation rules for UpdateBy

	// no validation rules for OwnerBy

	// no validation rules for SeaSon

	// no validation rules for UpdateAt

	if len(errors) > 0 {
		return ClubProfileResponse_DataMultiError(errors)
	}

	return nil
}

// ClubProfileResponse_DataMultiError is an error wrapping multiple validation
// errors returned by ClubProfileResponse_Data.ValidateAll() if the designated
// constraints aren't met.
type ClubProfileResponse_DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClubProfileResponse_DataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClubProfileResponse_DataMultiError) AllErrors() []error { return m }

// ClubProfileResponse_DataValidationError is the validation error returned by
// ClubProfileResponse_Data.Validate if the designated constraints aren't met.
type ClubProfileResponse_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClubProfileResponse_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClubProfileResponse_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClubProfileResponse_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClubProfileResponse_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClubProfileResponse_DataValidationError) ErrorName() string {
	return "ClubProfileResponse_DataValidationError"
}

// Error satisfies the builtin error interface
func (e ClubProfileResponse_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClubProfileResponse_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClubProfileResponse_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClubProfileResponse_DataValidationError{}
