package listener

var mainTestTpl = `
package main_test

func TestMain(m *testing.M) {
    m.Run()
}

func TestListen(t *testing.T) {
}
`
