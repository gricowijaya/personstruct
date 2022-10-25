// name the package of person struct
package personstruct

// import the fmt package
import(
    "fmt"
)

// create the person struct 
type person struct {
  // first name of the person
  first string 
  // last name of the person
  last string
  // age of the person
  age int
}

func main() {
    // example of person struct, we use firstPerson variable to create the person
    firstPerson := person {
      first: "James",
      last: "Bond",
    }
    
    // how to print the value first name of firstPerson
    fmt.Println("This is the First Name ", firstPerson.first)

    // how to print the value last name of firstPerson
    fmt.Println("This is the Last Name", firstPerson.last)
}
