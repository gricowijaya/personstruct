// name the package of person struct
package personstruct

// import the fmt package
import(
    "fmt"
)

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
