package repository

type Base interface {
	Add() bool
}

func Add(base Base) bool {
	return base.Add()
}
