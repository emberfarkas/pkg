// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: client/rabbitmq/conf.proto

package rabbitmq

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

// Validate checks the field values on RabbitConf with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RabbitConf) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RabbitConf with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RabbitConfMultiError, or
// nil if none found.
func (m *RabbitConf) ValidateAll() error {
	return m.validate(true)
}

func (m *RabbitConf) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Host

	// no validation rules for Port

	// no validation rules for VHost

	// no validation rules for Address

	if len(errors) > 0 {
		return RabbitConfMultiError(errors)
	}

	return nil
}

// RabbitConfMultiError is an error wrapping multiple validation errors
// returned by RabbitConf.ValidateAll() if the designated constraints aren't met.
type RabbitConfMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RabbitConfMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RabbitConfMultiError) AllErrors() []error { return m }

// RabbitConfValidationError is the validation error returned by
// RabbitConf.Validate if the designated constraints aren't met.
type RabbitConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RabbitConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RabbitConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RabbitConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RabbitConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RabbitConfValidationError) ErrorName() string { return "RabbitConfValidationError" }

// Error satisfies the builtin error interface
func (e RabbitConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRabbitConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RabbitConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RabbitConfValidationError{}

// Validate checks the field values on ConsumerConf with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ConsumerConf) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConsumerConf with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ConsumerConfMultiError, or
// nil if none found.
func (m *ConsumerConf) ValidateAll() error {
	return m.validate(true)
}

func (m *ConsumerConf) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Consumer

	if len(errors) > 0 {
		return ConsumerConfMultiError(errors)
	}

	return nil
}

// ConsumerConfMultiError is an error wrapping multiple validation errors
// returned by ConsumerConf.ValidateAll() if the designated constraints aren't met.
type ConsumerConfMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConsumerConfMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConsumerConfMultiError) AllErrors() []error { return m }

// ConsumerConfValidationError is the validation error returned by
// ConsumerConf.Validate if the designated constraints aren't met.
type ConsumerConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConsumerConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConsumerConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConsumerConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConsumerConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConsumerConfValidationError) ErrorName() string { return "ConsumerConfValidationError" }

// Error satisfies the builtin error interface
func (e ConsumerConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConsumerConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConsumerConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConsumerConfValidationError{}

// Validate checks the field values on ListenerConf with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListenerConf) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListenerConf with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListenerConfMultiError, or
// nil if none found.
func (m *ListenerConf) ValidateAll() error {
	return m.validate(true)
}

func (m *ListenerConf) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRabbit()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListenerConfValidationError{
					field:  "Rabbit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListenerConfValidationError{
					field:  "Rabbit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRabbit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerConfValidationError{
				field:  "Rabbit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetQueues() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListenerConfValidationError{
						field:  fmt.Sprintf("Queues[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListenerConfValidationError{
						field:  fmt.Sprintf("Queues[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerConfValidationError{
					field:  fmt.Sprintf("Queues[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListenerConfMultiError(errors)
	}

	return nil
}

// ListenerConfMultiError is an error wrapping multiple validation errors
// returned by ListenerConf.ValidateAll() if the designated constraints aren't met.
type ListenerConfMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListenerConfMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListenerConfMultiError) AllErrors() []error { return m }

// ListenerConfValidationError is the validation error returned by
// ListenerConf.Validate if the designated constraints aren't met.
type ListenerConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenerConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenerConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenerConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenerConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenerConfValidationError) ErrorName() string { return "ListenerConfValidationError" }

// Error satisfies the builtin error interface
func (e ListenerConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListenerConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenerConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenerConfValidationError{}

// Validate checks the field values on ProducerConf with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ProducerConf) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProducerConf with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProducerConfMultiError, or
// nil if none found.
func (m *ProducerConf) ValidateAll() error {
	return m.validate(true)
}

func (m *ProducerConf) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetRabbit()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProducerConfValidationError{
					field:  "Rabbit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProducerConfValidationError{
					field:  "Rabbit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRabbit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProducerConfValidationError{
				field:  "Rabbit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ContentType

	if len(errors) > 0 {
		return ProducerConfMultiError(errors)
	}

	return nil
}

// ProducerConfMultiError is an error wrapping multiple validation errors
// returned by ProducerConf.ValidateAll() if the designated constraints aren't met.
type ProducerConfMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProducerConfMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProducerConfMultiError) AllErrors() []error { return m }

// ProducerConfValidationError is the validation error returned by
// ProducerConf.Validate if the designated constraints aren't met.
type ProducerConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProducerConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProducerConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProducerConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProducerConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProducerConfValidationError) ErrorName() string { return "ProducerConfValidationError" }

// Error satisfies the builtin error interface
func (e ProducerConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProducerConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProducerConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProducerConfValidationError{}

// Validate checks the field values on ExchangeConf with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ExchangeConf) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExchangeConf with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ExchangeConfMultiError, or
// nil if none found.
func (m *ExchangeConf) ValidateAll() error {
	return m.validate(true)
}

func (m *ExchangeConf) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Exchange

	// no validation rules for Key

	if len(errors) > 0 {
		return ExchangeConfMultiError(errors)
	}

	return nil
}

// ExchangeConfMultiError is an error wrapping multiple validation errors
// returned by ExchangeConf.ValidateAll() if the designated constraints aren't met.
type ExchangeConfMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExchangeConfMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExchangeConfMultiError) AllErrors() []error { return m }

// ExchangeConfValidationError is the validation error returned by
// ExchangeConf.Validate if the designated constraints aren't met.
type ExchangeConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExchangeConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExchangeConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExchangeConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExchangeConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExchangeConfValidationError) ErrorName() string { return "ExchangeConfValidationError" }

// Error satisfies the builtin error interface
func (e ExchangeConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExchangeConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExchangeConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExchangeConfValidationError{}

// Validate checks the field values on QueueConf with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QueueConf) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueueConf with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QueueConfMultiError, or nil
// if none found.
func (m *QueueConf) ValidateAll() error {
	return m.validate(true)
}

func (m *QueueConf) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return QueueConfMultiError(errors)
	}

	return nil
}

// QueueConfMultiError is an error wrapping multiple validation errors returned
// by QueueConf.ValidateAll() if the designated constraints aren't met.
type QueueConfMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueueConfMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueueConfMultiError) AllErrors() []error { return m }

// QueueConfValidationError is the validation error returned by
// QueueConf.Validate if the designated constraints aren't met.
type QueueConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueueConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueueConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueueConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueueConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueueConfValidationError) ErrorName() string { return "QueueConfValidationError" }

// Error satisfies the builtin error interface
func (e QueueConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueueConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueueConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueueConfValidationError{}

// Validate checks the field values on AdminExchangeConf with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AdminExchangeConf) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AdminExchangeConf with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AdminExchangeConfMultiError, or nil if none found.
func (m *AdminExchangeConf) ValidateAll() error {
	return m.validate(true)
}

func (m *AdminExchangeConf) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Kind

	// no validation rules for Durable

	// no validation rules for AutoDelete

	// no validation rules for Internal

	// no validation rules for NoWait

	if len(errors) > 0 {
		return AdminExchangeConfMultiError(errors)
	}

	return nil
}

// AdminExchangeConfMultiError is an error wrapping multiple validation errors
// returned by AdminExchangeConf.ValidateAll() if the designated constraints
// aren't met.
type AdminExchangeConfMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AdminExchangeConfMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AdminExchangeConfMultiError) AllErrors() []error { return m }

// AdminExchangeConfValidationError is the validation error returned by
// AdminExchangeConf.Validate if the designated constraints aren't met.
type AdminExchangeConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdminExchangeConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdminExchangeConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdminExchangeConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdminExchangeConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdminExchangeConfValidationError) ErrorName() string {
	return "AdminExchangeConfValidationError"
}

// Error satisfies the builtin error interface
func (e AdminExchangeConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdminExchangeConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdminExchangeConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdminExchangeConfValidationError{}

// Validate checks the field values on AdminQueueConf with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AdminQueueConf) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AdminQueueConf with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AdminQueueConfMultiError,
// or nil if none found.
func (m *AdminQueueConf) ValidateAll() error {
	return m.validate(true)
}

func (m *AdminQueueConf) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Durable

	// no validation rules for AutoDelete

	// no validation rules for Exclusive

	// no validation rules for NoWait

	if len(errors) > 0 {
		return AdminQueueConfMultiError(errors)
	}

	return nil
}

// AdminQueueConfMultiError is an error wrapping multiple validation errors
// returned by AdminQueueConf.ValidateAll() if the designated constraints
// aren't met.
type AdminQueueConfMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AdminQueueConfMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AdminQueueConfMultiError) AllErrors() []error { return m }

// AdminQueueConfValidationError is the validation error returned by
// AdminQueueConf.Validate if the designated constraints aren't met.
type AdminQueueConfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdminQueueConfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdminQueueConfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdminQueueConfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdminQueueConfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdminQueueConfValidationError) ErrorName() string { return "AdminQueueConfValidationError" }

// Error satisfies the builtin error interface
func (e AdminQueueConfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdminQueueConf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdminQueueConfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdminQueueConfValidationError{}
