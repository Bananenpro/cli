package cli

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

var (
	// ErrCanceled indicates that the user has interrupted the programm, e.g. by using Ctrl + C.
	ErrCanceled = errors.New("Canceled.")
)

// Input ask the user to input a line of text.
func Input(prompt string, validators ...Validator) (string, error) {
	opts := make([]survey.AskOpt, 0, len(validators)+1)
	opts = append(opts, survey.WithValidator(survey.Required))
	for _, v := range validators {
		opts = append(opts, survey.WithValidator(survey.Validator(v)))
	}

	var result string
	err := survey.AskOne(&survey.Input{
		Message: prompt,
	}, &result, opts...)
	if err == terminal.InterruptErr {
		err = ErrCanceled
	}
	return result, err
}

// YesNo asks the user a yes/no question.
func YesNo(question string, defaultValue bool) (yes bool, err error) {
	err = survey.AskOne(&survey.Confirm{
		Message: question,
		Default: defaultValue,
	}, &yes, survey.WithValidator(survey.Required))
	if err == terminal.InterruptErr {
		err = ErrCanceled
	}
	return yes, err
}

// Select asks the user to select on of many options. It returns the index of the chosen option.
func Select(msg string, options []string) (int, error) {
	var index int
	err := survey.AskOne(&survey.Select{
		Message: msg,
		Options: options,
	}, &index, survey.WithValidator(survey.Required))
	if err == terminal.InterruptErr {
		err = ErrCanceled
	}
	return index, err
}

// SelectString asks the user to select on of many options. It returns the entry in options with the chosen index.
// It panics if the length of displayOptions differs from the length of options.
func SelectString(msg string, displayOptions []string, options []string) (string, error) {
	if len(displayOptions) != len(options) {
		panic("Lengths of displayOptions and options don't match")
	}
	var index int
	err := survey.AskOne(&survey.Select{
		Message: msg,
		Options: options,
	}, &index, survey.WithValidator(survey.Required))
	if err == terminal.InterruptErr {
		err = ErrCanceled
	}
	return options[index], err
}
