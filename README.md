### TDD or test driven development in Golang

### TDD- Benefits 
 - Small regression suite
     - You have test case around all code and if you change any part of 
       code and if you chnage any part of code for a new feature you will 
       get advanced feeback incase something in code code.

 - Reduction in Bugs
 - Cleaner and Simple Code
 - TDD follows a methodology i.e ( RED, GREEN , REFRACTOR)     
     <details>
     <summary>TDD methodology</summary>

    | methodology          | Desc                                                       |
    | -------------------- | ---------------------------------------------------------- |
    | `RED`                | Write a test case that gives error.                        |
    | `GREEN`              | Do changes to code to make the test case pass.             |
    | `REFRACTOR`          | If need some Refractoring do it here.                      |
     </details>
 - Avoids duplication of Code
 - Refractoring code improves the code
 - TDD drive the code design and approach
 - Productivity increase


### TDD Test Steps: 
Following steps define how to perform TDD test,

 - Add a test.
 - Run all tests and see if any new test fails.
 - Write some code.
 - Run tests and Refactor code.
 - Repeat.

### Structure of a test case:
 - SetUp: gets the UUT(unit under test ready : dependencies that have to be injected or input paramters , initiate this here ),unit can be particular module , functions, struct and others..
 - Execution: trigger or run the UUT and capture all output like return value and output parameters
 - Validation: Check if the output matches the expected output.
 - Cleanup: restore the UUT on the actual or old state , letting the other test to start from its execution.


### Unit test 
      - A unit test is a program that tests a unit component by all possible means and compares the result to the expected output.
      - Unit Components : functions, struct, methods,

```unit test sample

import "testing"
func TestAbc(t *testing.T) {
    t.Error() // to indicate test failed
}

```      
