package helpers

import (
	"fmt"
	"net"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func RegisterCustomValidator() {
	v := binding.Validator.Engine().(*validator.Validate)
	if err := v.RegisterValidation("list_uuid", ListUuidValidator); err != nil {
		panic(err)
	}
	if err := v.RegisterValidation("list_ip_address", ListIpAddressValidator); err != nil {
		panic(err)
	}
	if err := v.RegisterValidation("list_cidr", ListCidrValidator); err != nil {
		panic(err)
	}
	if err := v.RegisterValidation("list_mac", ListMacAddressValidator); err != nil {
		panic(err)
	}
}

func MacAddressValidator(mac string) (string, error) {
	if mac == "" {
		return "", nil
	}
	mac = strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(mac, "-", ""), ".", ""), ":", ""))
	if len(mac) != 12 {
		return "", fmt.Errorf("invalid mac address %s, length: %d", mac, len(mac))
	}
	re := regexp.MustCompile("^[0-9a-fA-F]{12}$")
	if !re.MatchString(mac) {
		return "", fmt.Errorf("invalid mac address: %s", mac)
	}
	// no for loop here because of performance
	return mac[:2] + ":" + mac[2:4] + ":" + mac[4:6] + ":" + mac[6:8] + ":" + mac[8:10] + ":" + mac[10:], nil
}

func ListMacAddressValidator(level validator.FieldLevel) bool {
	field := level.Field()
	for field.Kind() == reflect.Pointer {
		field = field.Elem()
	}

	if field.Kind() != reflect.Slice && field.Kind() != reflect.Array {
		return false
	}
	for i := 0; i < field.Len(); i++ {
		_field := field.Index(i).Interface()
		if str, ok := _field.(string); !ok {
			return false
		} else if _, err := MacAddressValidator(str); err != nil {
			return false
		}
	}
	return true
}

func ListUuidValidator(level validator.FieldLevel) bool {
	field := level.Field()
	for field.Kind() == reflect.Pointer {
		field = field.Elem()
	}

	if field.Kind() != reflect.Slice && field.Kind() != reflect.Array {
		return false
	}
	for i := 0; i < field.Len(); i++ {
		_field := field.Index(i).Interface()
		if str, ok := _field.(string); !ok {
			return false
		} else if _, err := uuid.Parse(str); err != nil {
			return false
		}
	}
	return true
}

func ListIpAddressValidator(level validator.FieldLevel) bool {
	field := level.Field()
	for field.Kind() == reflect.Pointer {
		field = field.Elem()
	}

	if field.Kind() != reflect.Slice && field.Kind() != reflect.Array {
		return false
	}
	for i := 0; i < field.Len(); i++ {
		_field := field.Index(i).Interface()
		if str, ok := _field.(string); !ok {
			return false
		} else if ip := net.ParseIP(str); ip == nil {
			return false
		}
	}
	return true
}

func ListCidrValidator(level validator.FieldLevel) bool {
	field := level.Field()
	for field.Kind() == reflect.Pointer {
		field = field.Elem()
	}

	if field.Kind() != reflect.Slice && field.Kind() != reflect.Array {
		return false
	}
	for i := 0; i < field.Len(); i++ {
		_field := field.Index(i).Interface()
		if str, ok := _field.(string); !ok {
			return false
		} else if _, _, err := net.ParseCIDR(str); err != nil {
			return false
		}
	}
	return true
}
