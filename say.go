//Package say documentation
package say

import "fmt"

//Hello returns a string saying hello to the name of the person has been provided
func Hello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
