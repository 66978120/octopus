// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: server/platform-server/internal/conf/conf.proto

package conf

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on Bootstrap with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Bootstrap) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetApp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "App",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetServer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Server",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				field:  "Service",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// BootstrapValidationError is the validation error returned by
// Bootstrap.Validate if the designated constraints aren't met.
type BootstrapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BootstrapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BootstrapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BootstrapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BootstrapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BootstrapValidationError) ErrorName() string { return "BootstrapValidationError" }

// Error satisfies the builtin error interface
func (e BootstrapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBootstrap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BootstrapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BootstrapValidationError{}

// Validate checks the field values on App with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *App) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Version

	// no validation rules for IsDev

	// no validation rules for LogLevel

	return nil
}

// AppValidationError is the validation error returned by App.Validate if the
// designated constraints aren't met.
type AppValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppValidationError) ErrorName() string { return "AppValidationError" }

// Error satisfies the builtin error interface
func (e AppValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppValidationError{}

// Validate checks the field values on Server with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Server) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetHttp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ServerValidationError{
				field:  "Http",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ServerValidationError is the validation error returned by Server.Validate if
// the designated constraints aren't met.
type ServerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServerValidationError) ErrorName() string { return "ServerValidationError" }

// Error satisfies the builtin error interface
func (e ServerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServerValidationError{}

// Validate checks the field values on Data with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Data) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for BaseServerAddr

	if v, ok := interface{}(m.GetRedis()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DataValidationError{
				field:  "Redis",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for BaseServerRequestTimeout

	return nil
}

// DataValidationError is the validation error returned by Data.Validate if the
// designated constraints aren't met.
type DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataValidationError) ErrorName() string { return "DataValidationError" }

// Error satisfies the builtin error interface
func (e DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataValidationError{}

// Validate checks the field values on Service with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Service) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TokenExpirationSec

	return nil
}

// ServiceValidationError is the validation error returned by Service.Validate
// if the designated constraints aren't met.
type ServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ServiceValidationError) ErrorName() string { return "ServiceValidationError" }

// Error satisfies the builtin error interface
func (e ServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ServiceValidationError{}

// Validate checks the field values on Server_HTTP with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Server_HTTP) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Network

	// no validation rules for Addr

	if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Server_HTTPValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for JwtSecrect

	return nil
}

// Server_HTTPValidationError is the validation error returned by
// Server_HTTP.Validate if the designated constraints aren't met.
type Server_HTTPValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Server_HTTPValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Server_HTTPValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Server_HTTPValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Server_HTTPValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Server_HTTPValidationError) ErrorName() string { return "Server_HTTPValidationError" }

// Error satisfies the builtin error interface
func (e Server_HTTPValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sServer_HTTP.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Server_HTTPValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Server_HTTPValidationError{}

// Validate checks the field values on Data_Redis with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Data_Redis) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Addr

	// no validation rules for Username

	// no validation rules for Password

	return nil
}

// Data_RedisValidationError is the validation error returned by
// Data_Redis.Validate if the designated constraints aren't met.
type Data_RedisValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Data_RedisValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Data_RedisValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Data_RedisValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Data_RedisValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Data_RedisValidationError) ErrorName() string { return "Data_RedisValidationError" }

// Error satisfies the builtin error interface
func (e Data_RedisValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData_Redis.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Data_RedisValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Data_RedisValidationError{}
