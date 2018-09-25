/*
* Shi Ruitao 2017-12-7
*/

package ago
import "fmt"

type Stringer interface {
	String() string
}

type test struct {}

func (this *test) String() {
	fmt.Println("String()")
}
func main() {

	var v interface {}
	var ok bool
	v = &test{}

	if sv, ok := v.(Stringer); ok {
		sv.String()
		fmt.Printf("v implements String(): %s\n", sv) // note: sv, not v
	}
	fmt.Println(v, ok)
}