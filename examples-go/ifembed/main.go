package main

import "fmt"

type ContentSpecs interface {
	Specs() string
}

type ContentRating interface {
	Rating() string
}

type Content struct {
	ContentName string
}

func (c *Content) Rating() string {
	return "content-good"
}

type AudioContent struct {
	*Content
}

func (ac *AudioContent) Specs() string {
	return "audio;16-bits;2-channels"
}

func (ac *AudioContent) Rating() string {
	return "audio-excellent"
}

func (c *Content) DisplayContentSpecsInner() {
	impl := interface{}(c).(ContentSpecs)
	fmt.Println(impl.Specs())
}

func (c *AudioContent) DisplayContentSpecs() {
	impl := interface{}(c).(ContentSpecs)
	fmt.Println(impl.Specs())
}

// any object that supports this interface ContentSpecs can be passed
// as the argument cs1. "duck-typing"
func PrintSpecs(cs1 ContentSpecs) {
	fmt.Println(cs1.Specs())
}

// any object that supports this interface ContentRating can be passed
// as the argument cr1. "duck-typing"
func PrintRating(cr1 ContentRating) {
	fmt.Println(cr1.Rating())
}

func main() {
	var c Content
	c.ContentName = "test"
	fmt.Println(c.Rating())

	var ac AudioContent
	fmt.Println(ac.Rating())
	fmt.Println(ac.Specs())

	fmt.Println("Testing interfaces and duck-typing:")
	PrintRating(&c)
	PrintRating(&ac)
	PrintSpecs(&ac)

	fmt.Println("Invoking interface of the outer struct... This is successful.")
	ac.DisplayContentSpecs()

	fmt.Println("Invoking interface of the inner struct... This will fail at runtime.")
	ac.DisplayContentSpecsInner()
}
