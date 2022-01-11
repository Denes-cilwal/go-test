package tutorial2

import "testing"

// test case  for emptyarguements ...
func TestHello(t *testing.T) {

	emptyHello := Hello("") // if empty, return david

	if emptyHello != "david!" {
		t.Errorf("Hello, failed expected %v but got %v", emptyHello, emptyHello)
	} else {
		t.Logf("Hello sucess expected %v also got %v", emptyHello, emptyHello)

	}

}

// test case for valid argument
func TestHelloValidArgs(t *testing.T) {

	result := Hello("david!")

	// test case for invalid argument
	if result != "Hello david!" {
		t.Errorf("Hello failed expected %v but got %v", result, result)

	} else {
		t.Logf("Hello sucess expected %v also got %v", result, result)

	}

}
