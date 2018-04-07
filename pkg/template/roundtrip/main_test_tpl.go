package roundtrip

var mainTestTpl = `
package main_test

func TestMain(m *testing.M) {
    m.Run()
}

func TestRoundtrip(t *testing.T) {
}
`
