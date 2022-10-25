package personstruct

import("fmt")

func main() {
    // example of creation
    firstPerson := person {
      first: "James",
      last: "Bond",
    }
    
    fmt.Println("This is the First Person", firstPerson.first)
}
