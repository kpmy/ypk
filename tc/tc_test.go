package tc

import (
	"errors"
	"testing"
)

var e0 = errors.New("some shit")
var e1 = errors.New("other shit")

func TestTry(t *testing.T) {
	Try(func(...interface{}) interface{} {
		t.Log("try")
		panic(e0)
		t.Log("never happens")
		return nil
	}).Catch(e0, func(error) {
		t.Log(e0)
	}).Catch(e1, func(error) {
		t.Log(e1)
	}).Catch(nil, func(e error) {
		t.Log(e)
	}).Finally(func() {
		t.Log("finally")
	}).Do()

	Try(func(...interface{}) interface{} {
		t.Log("try")
		panic(e1)
		t.Log("never happens")
		return nil
	}).Catch(e0, func(error) {
		t.Log(e0)
	}).Catch(e1, func(error) {
		t.Log(e1)
	}).Catch(nil, func(e error) {
		t.Log(e)
	}).Finally(func() {
		t.Log("finally")
	}).Do()

	Try(func(...interface{}) interface{} {
		t.Log("try")
		panic(errors.New("runtime shit"))
		t.Log("never happens")
		return nil
	}).Catch(e0, func(error) {
		t.Log(e0)
	}).Catch(e1, func(error) {
		t.Log(e1)
	}).Catch(nil, func(e error) {
		t.Log(e)
	}).Finally(func() {
		t.Log("finally")
	}).Do()

	t.Log(Try(func(x ...interface{}) interface{} {
		return x[0]
	}).Do("ok"))

	Do(func() {
		Throw("not", " ", "ok")
	}).Catch(nil, func(e error) {
		t.Log(e)
	}).Do()

	Do(func() {
		x := false
		Assert(x, 100, "sooo bad")
	}).Catch(nil, func(e error) {
		t.Log(e)
	}).Do()

	Do(func() {
		Halt(100, "sooooooo bad")
	}).Catch(nil, func(e error) {
		t.Log(e)
	}).Do()
}
