/*
* Shi Ruitao  2017-12-12
*/

package ago
import "fmt"
func main() {
	var i int = 1
	for {
		for {
			if i == 1 {
				fmt.Println("1.1")
				break
				fmt.Println("1.2")
			}
		}
		if i += 1; i == 2 {
			fmt.Println("2.1")
			break
			fmt.Println("2.2")
		}
	}
}