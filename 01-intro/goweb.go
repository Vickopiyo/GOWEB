package main

import (
	"fmt"
	"log"
	"net/http"
)     

// Argument in URL --  https://localhost:3000/error?name=alex&age=26   

   
//  Handler  functions    

//  Handler 1--  homepage 
func Hola(rw http.ResponseWriter,   r*http.Request)  {    
    //  Fprintln formats to json kinda 
	fmt.Fprintln(rw,  "Hello Server!" )    
	// GET REQUEST---Method only...Each refresh makes get request           
	fmt.Println("The  request is " + r.Method)         
}            

// Handler 2 for not found!!    

func PageNotFound(rw http.ResponseWriter,    r*http.Request)  {     
//   replies to http with not found
   http.NotFound(rw, r)       
   
}       

// Handler 3 --  errors     

func Error(rw http.ResponseWriter, r*http.Request)  {         
    
	// http.status not found replaces  the hard coded intial--- 404    \

	   http.Error( rw,  "The page does not work!!", http.StatusNotFound)                      
}            

// Handler 4 --    a url which expects arguements ---localhost:3000/greeting?name=vick&age=26  thus outputs until you pass argument ?name=vick&age=26
// &{0xc0000aebe0 0xc000190200 {} 0x3c27e0 false false false false 0 {0 0} 0xc000054040 {0xc00004c1c0 map[] false false} map[] false 0 -1 0 false false [] 0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0] 0xc000052150 0} This page does not work! 404

func Greeting(rw http.ResponseWriter,  r *http.Request)  {    

    fmt.Println(r.URL)      
	    //outputs the query in url as raw as it is 
	fmt.Println(r.URL.RawQuery)      

	// maps the argument in the url( map[age:[26], name:[vick]])
	fmt.Println(r.URL.Query())  
       
	// hardcoding url arguments 
	name :=r.URL.Query().Get("name")
    age :=r.URL.Query().Get("age")          
	
  //FprintF--accepts three arguments--1.specified standard i or o 2.format string--contains strings ,   3.a--interface--specified constant variables used in the code 
//   it returns the number of bytes written and any write  error encountered 

	fmt.Fprintf(rw, "Hello %s, your age is %s ", name, age )   
	      
}         
          
func main() {           
	
	//MUX-- alias HTTP request multiplexer--circuit used to select  and route any one of the several input signals to a single output.    
    // Mux solves problem of reducancy of http.handlefunc instead it picks which route it has been passed to and serve response 
	mux :=http.NewServeMux()   
	 mux.HandleFunc("/", Hola)   
	 mux.HandleFunc("/page", PageNotFound)      
	 mux.HandleFunc("/error", Error)      
	 mux.HandleFunc("/greeting", Greeting)
          
	    //ROUTERS     
    //ROUTER 1 --- home route  and handler func (Hola)          
	//  http.HandleFunc("/", Hola)   
	// ROUTE 2 PAGE NOT FOUND      
    //   http.HandleFunc("/page", PageNotFound)       
	//   ROUTE 3 -- Error  
    //   http.HandleFunc("/error", Error)          
	//   ROUTE 4 --  Greeting URL  
	// http.HandleFunc("/greeting", Greeting)

    //creating  Server      
	fmt.Println("The server is running on port 3000")
	fmt.Println("Run Server: http://localhost:3000/")   
	// Before inputing mux as last arg--it returns page not found  

    log.Fatal(http.ListenAndServe("localhost:3000", mux))              

}                                                       
                                     