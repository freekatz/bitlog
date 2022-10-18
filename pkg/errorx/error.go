package errorx

import "errors"

var (
	/*
		ErrorX cli errors
	*/
	ErrPrintAndExit    = errors.New("print and exit")
	ErrNoInputConfig   = errors.New("have no input config")
	ErrNoInputFilepath = errors.New("have no input filepath")

	/*
		ErrorX runtime errors
	*/
	ErrPackageLogsFailed = errors.New("package logs failed")

	/*
		ErrorX internal errors
	*/
	ErrConfigInvalid         = errors.New("config invalid")
	ErrEnvKeyInvalid         = errors.New("env key invalid")
	ErrEnvLookupFailed       = errors.New("env lookup failed")
	ErrLogTreeRootIncomplete = errors.New("log tree root incomplete")
	ErrLogTreeNotFoundChild  = errors.New("log tree not found child")

	/*
		ErrorX file errors
	*/
	ErrFileNotExisted   = errors.New("file not existed")
	ErrFileHasExisted   = errors.New("file has existed")
	ErrFileTypeNotMatch = errors.New("file type not match")
)
