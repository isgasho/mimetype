package matchers

import "bytes"

// Zip matches a zip archive.
func Zip(in []byte) bool {
	return len(in) > 3 &&
		in[0] == 0x50 && in[1] == 0x4B &&
		(in[2] == 0x3 || in[2] == 0x5 || in[2] == 0x7) &&
		(in[3] == 0x4 || in[3] == 0x6 || in[3] == 0x8)
}

// SevenZ matches a 7z archive.
func SevenZ(in []byte) bool {
	return len(in) > 6 &&
		bytes.Equal(in[:6], []byte{0x37, 0x7A, 0xBC, 0xAF, 0x27, 0x1C})
}

// Epub matches an EPUB file.
func Epub(in []byte) bool {
	if len(in) < 58 {
		return false
	}

	return bytes.Equal(in[30:58], []byte("mimetypeapplication/epub+zip"))
}

// Jar matches a Java archive file.
func Jar(in []byte) bool {
	return bytes.Contains(in, []byte("META-INF/MANIFEST.MF"))
}

// Gzip matched gzip files based on http://www.zlib.org/rfc-gzip.html#header-trailer.
func Gzip(in []byte) bool {
	return len(in) > 2 && bytes.Equal(in[:2], []byte{0x1f, 0x8b})
}

// Crx matches a Chrome extension file: a zip archive prepended by "Cr24".
func Crx(in []byte) bool {
	return len(in) > 4 && bytes.Equal(in[:4], []byte("Cr24"))
}

// Tar matches a (t)ape (ar)chive file.
func Tar(in []byte) bool {
	return len(in) > 262 &&
		bytes.Equal(in[257:262], []byte("ustar"))
}
