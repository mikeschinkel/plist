package plist_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/micromdm/plist"
)

// Value is key to my use-case
type Value struct {
	Value interface{}
}

func TestDecodeWithDict(t *testing.T) {
	r := bytes.NewBuffer([]byte(dockYAML))
	decoder := plist.NewXMLDecoder(r)
	pm := ProfileManifest{}
	err := decoder.Decode(&pm)
	if err != nil {
		t.Errorf("Failed to decode profile manifest: %s", err.Error())
	}
	fmt.Println(pm)
}

//// UnmarshalPlist only captures the value and assigns to Value.value which is of
//func (v *Value) UnmarshalPlist(f func(interface{}) error) (err error) {
//	var value any
//	err = f(&value)
//	v.Value = value
//	return nil
//}

type ManifestKey struct {
	Name    string        `plist:"pfm_name"`
	Default interface{}   `plist:"pfm_default,omitempty"`
	Subkeys []ManifestKey `plist:"pfm_subkeys,omitempty"`
}
type ProfileManifest struct {
	Domain  string        `plist:"pfm_domain"`
	Subkeys []ManifestKey `plist:"pfm_subkeys"`
}

// dockYAML is a subset of https://github.com/ProfileManifests/ProfileManifests/tree/master/Manifests/ManifestsApple/com.apple.dock.plist
const dockYAML = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
  <dict>
    <key>pfm_domain</key>
    <string>com.apple.dock</string>
    <key>pfm_subkeys</key>
    <array>
      <dict>
        <key>pfm_name</key>
        <string>static-apps</string>
        <key>pfm_subkeys</key>
        <array>
          <dict>
            <key>pfm_name</key>
            <string>StaticItem</string>
            <key>pfm_subkeys</key>
            <array>
              <dict>
                <key>pfm_default</key>
                <string>file-tile</string>
                <key>pfm_name</key>
                <string>tile-type</string>
              </dict>
              <dict>
                <key>pfm_name</key>
                <string>tile-data</string>
                <key>pfm_subkeys</key>
                <array>
                  <dict>
                    <key>pfm_name</key>
                    <string>label</string>
                  </dict>
                  <dict>
										<key>pfm_default</key>
										<integer>100</integer>
										<key>pfm_name</key>
										<string>Amount</string>
                  </dict>
                  <dict>
										<key>pfm_default</key>
										<real>1.234</real>
										<key>pfm_name</key>
										<string>Fraction</string>
                  </dict>
                  <dict>
                    <key>pfm_name</key>
                    <string>file-label</string>
                  </dict>
                  <dict>
                    <key>pfm_name</key>
                    <string>file-type</string>
                  </dict>
                  <dict>
                    <key>pfm_name</key>
                    <string>file-data</string>
                    <key>pfm_subkeys</key>
                    <array>
                      <dict>
												<key>pfm_default</key>
												<string>example.com</string>
                        <key>pfm_name</key>
                        <string>_CFURLString</string>
                      </dict>
                      <dict>
												<key>pfm_default</key>
												<string>whatever</string>
                        <key>pfm_name</key>
                        <string>_CFURLStringType</string>
                      </dict>
                    </array>
                  </dict>
                  <dict>
                    <key>pfm_name</key>
                    <string>url</string>
                  </dict>
                </array>
              </dict>
            </array>
          </dict>
        </array>
      </dict>
    </array>
  </dict>
</plist>`
