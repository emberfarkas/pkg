// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rest/conf.proto

package rest

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

// Validate checks the field values on Conf with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Conf) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Conf with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ConfMultiError, or nil if none found.
func (m *Conf) ValidateAll() error {
	return m.validate(true)
}

func (m *Conf) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Network

	// no validation rules for Address

	if all {
		switch v := interface{}(m.GetTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetMiddlewares() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ConfValidationError{
						field:  fmt.Sprintf("Middlewares[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ConfValidationError{
						field:  fmt.Sprintf("Middlewares[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfValidationError{
					field:  fmt.Sprintf("Middlewares[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ConfMultiError(errors)
	}

	return nil
}

// ConfMultiError is an error wrapping multiple validation errors returned by
// Conf.ValidateAll() if the designated constraints aren't met.
type ConfMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfMultiError) AllErrors() []error { return m }

// ConfValidationError is the validation error returned by Conf.Validate if the
// designated constraints aren't met.
type ConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfValidationError) ErrorName() string { return "ConfValidationError" }

// Error satisfies the builtin error interface
func (e ConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfValidationError{}
