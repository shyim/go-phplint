package phplint

import "testing"

func FuzzLintNeverPanics(f *testing.F) {
	f.Add([]byte("<?php echo 'ok';"), uint8(PHP86))
	f.Add([]byte("<?php function broken( {"), uint8(PHP74))
	f.Add([]byte{0xff, 0x00, '<', '?'}, uint8(PHP82))

	f.Fuzz(func(t *testing.T, source []byte, rawVersion uint8) {
		versions := SupportedVersions()
		version := versions[int(rawVersion)%len(versions)]
		_, _ = Lint("fuzz.php", source, Options{PHPVersion: version})
	})
}
