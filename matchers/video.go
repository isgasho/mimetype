package matchers

import (
	"bytes"
)

// WebM matches a WebM file.
func WebM(in []byte) bool {
	return isMatroskaFileTypeMatched(in, "webm")
}

// Mkv matches a mkv file.
func Mkv(in []byte) bool {
	return isMatroskaFileTypeMatched(in, "matroska")
}

// isMatroskaFileTypeMatched is used for webm and mkv file matching.
// It checks for .Eß£ sequence. If the sequence is found,
// then it means it is Matroska media container, including WebM.
// Then it verifies which of the file type it is representing by matching the
// file specific string.
func isMatroskaFileTypeMatched(in []byte, flType string) bool {
	if bytes.HasPrefix(in, []byte("\x1A\x45\xDF\xA3")) {
		return isFileTypeNamePresent(in, flType)
	}
	return false
}

// isFileTypeNamePresent accepts the matroska input data stream and searches
// for the given file type in the stream. Return whether a match is found.
// The logic of search is: find first instance of \x42\x82 and then
// search for given string after one byte of above instance.
func isFileTypeNamePresent(in []byte, flType string) bool {
	var ind int
	if len(in) >= 4096 { // restricting length to 4096
		ind = bytes.Index(in[0:4096], []byte("\x42\x82"))
	} else {
		ind = bytes.Index(in, []byte("\x42\x82"))
	}
	if ind > 0 {
		// filetype name will be present exactly
		// one byte after the match of the two bytes "\x42\x82"
		return bytes.HasPrefix(in[ind+3:], []byte(flType))
	}
	return false
}

// Flv matches a Flash video file.
func Flv(in []byte) bool {
	return bytes.HasPrefix(in, []byte("\x46\x4C\x56\x01"))
}

// Mpeg matches a Moving Picture Experts Group file.
func Mpeg(in []byte) bool {
	return len(in) > 3 && bytes.Equal(in[:3], []byte("\x00\x00\x01")) &&
		in[3] >= 0xB0 && in[3] <= 0xBF
}

// QuickTime matches a QuickTime File Format file.
func QuickTime(in []byte) bool {
	return len(in) > 12 &&
		(bytes.Equal(in[4:12], []byte("ftypqt  ")) ||
			bytes.Equal(in[4:8], []byte("moov")))
}

// Avi matches an Audio Video Interleaved file.
func Avi(in []byte) bool {
	return len(in) > 16 &&
		bytes.Equal(in[:4], []byte("RIFF")) &&
		bytes.Equal(in[8:16], []byte("AVI LIST"))
}
