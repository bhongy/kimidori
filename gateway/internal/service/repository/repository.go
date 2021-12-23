// Package repository provides an interface to perform service discovery query
// for proxying requests to back-end services
package repository

import "errors"

var (
	ErrNotFound = errors.New("service not found")
)

type Backend struct {
	Origin string `json:"origin"`
}

type Repository interface {
	ByServiceName(serviceName string) (Backend, error)
}
