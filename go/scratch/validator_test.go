package main

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

type P1 struct {
	F int `vd:"eq=10"`
}

type P2 struct {
	F int `vd:"omitempty,eq=10"`
}

func TestValidator(t *testing.T) {
	var err error
	validate := validator.New()
	validate.SetTagName("vd")

	t.Run("dive", func(t *testing.T) {
		t.Run("without dive", func(t *testing.T) {
			type B struct {
				F1 int `vd:"eq=10"`
			}
			type A struct {
				Elems []B `vd:""`
			}
			err = validate.Struct(A{Elems: []B{{10}, {-10}}})
			fmt.Println(err)
		})

		t.Run("with dive", func(t *testing.T) {
			type B struct {
				F1 int `vd:"eq=10"`
			}
			type A struct {
				Elems []B `vd:"dive"`
			}
			err = validate.Struct(A{Elems: []B{{10}, {-10}}})
			fmt.Println(err)
		})
	})

	t.Run("omitempty", func(t *testing.T) {
		type B struct {
			F1 int
		}
		type A struct {
			Elems []B `vd:"omitempty,eq=2"`
		}
		err = validate.Struct(A{Elems: []B{{10}, {-10}}})
		fmt.Println(err)
		err = validate.Struct(A{Elems: []B{{10}}})
		fmt.Println(err)
		err = validate.Struct(A{Elems: []B{}})
		fmt.Println(err)
		err = validate.Struct(A{Elems: nil})
		fmt.Println(err)
	})
}
