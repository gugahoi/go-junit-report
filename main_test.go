package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestXmlToTemplate(t *testing.T) {
	testCases := []struct {
		desc     string
		expected string
	}{
		{
			desc:     "fixtures/with_failure.xml",
			expected: withFailureExpectedOutput,
		},
		{
			desc:     "fixtures/no_failure.xml",
			expected: "",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// arrange
			writer := bytes.NewBufferString("")
			f, _ := ioutil.ReadFile(tC.desc)

			// assign
			xmlToTemplate(writer, defaultTemplate, f)

			// assert
			if writer.String() != tC.expected {
				t.Fatalf("expected\n====\n%s\ngot:\n====\n%s\n", tC.expected, writer.String())
			}
		})
	}
}

var withFailureExpectedOutput = `<details>
	<summary>
		<code>
		View: Banner Renders toprsuccess in View: Banner Renders toprsuccess
		</code>
	</summary>
	<code>
		<pre>
		Error: Warning: Failed prop type: Invalid prop "format" of value "topr" supplied to "Banner", expected one of ["top","inline"].
		in Banner
		at BufferedConsole.Object.<anonymous>.console.error (/code/tools/jest-setup.js:74:9)
		at printWarning (/code/node_modules/prop-types/node_modules/fbjs/lib/warning.js:33:15)
		at warning (/code/node_modules/prop-types/node_modules/fbjs/lib/warning.js:57:20)
		at checkPropTypes (/code/node_modules/prop-types/checkPropTypes.js:52:11)
		at validatePropTypes (/code/node_modules/react/cjs/react.development.js:1247:5)
		at Object.createElement (/code/node_modules/react/cjs/react.development.js:1297:5)
		at Object.<anonymous> (/code/src/components/banner/banner.spec.js:11:47)
		at Object.asyncFn (/code/node_modules/jest-jasmine2/build/jasmine_async.js:124:345)
		at resolve (/code/node_modules/jest-jasmine2/build/queue_runner.js:46:12)
		at new Promise (<anonymous>)
		at mapper (/code/node_modules/jest-jasmine2/build/queue_runner.js:34:499)
		at promise.then (/code/node_modules/jest-jasmine2/build/queue_runner.js:74:39)
		at <anonymous>
			
		</pre>
	</code>
</details>
`
