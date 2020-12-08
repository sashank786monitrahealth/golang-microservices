package main;

import(
    "fmt"
)

type person struct{
  first string
}

// any TYPE that has the methods specified by the interface 
// is also the interface type
// an interface says 
// "Hey baby, if you have these methods, then you are my type."

type human interface{
    speak()
}


func (p person) speak(){

   fmt.Println(p.first);
}


func main(){

   var p person = person{
                          first: "James",
                        };
                  
   // in go a VALUE can be of more than one TYPE
   // in this example, p1 is both TYPE person and TYPE human      
   var h human;
   
   h = p;
   h.speak();
   
   fmt.Printf("%T\n",h)
                        

}
