package plist

import (
	"errors"
)

var (
	ErrUnexpectedElement             = errors.New("plist: unexpected element")
	ErrUnexpectedNonEmptyXMLCharData = errors.New("plist: unexpected non-empty xml.CharData")
	ErrFailedToParseXMLStartElement  = errors.New("plist: failed to parse XML start element")
)
