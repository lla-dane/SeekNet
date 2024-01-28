package main

import "github.com/fatih/color"

var (
	white = color.New(color.FgHiWhite).SprintfFunc()
	red = color.New(color.FgHiRed).SprintFunc()
	blue = color.New(color.FgBlue).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgHiYellow).SprintFunc()
	magenta = color.New(color.FgHiMagenta).SprintFunc()
)

var (
	bold = color.New(color.Bold).SprintFunc()
	underlined = color.New(color.Underline).SprintFunc()
)

var (
	colors = []*color.Color{
		color.New(color.FgRed),
		color.New(color.FgGreen),
		color.New(color.FgYellow),
		color.New(color.FgCyan),
		color.New(color.FgMagenta),
		color.New(color.FgWhite),
		color.New(color.FgHiRed),
		color.New(color.FgHiGreen),
		color.New(color.FgHiBlue),
	}
)