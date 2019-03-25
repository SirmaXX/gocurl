package main  
  
import (  
    "fmt"  
    "os"
    "log"
   "net/http"
     "sync"
  "io/ioutil"

)  


func list(){

  var s, arg string  
    for i := 1; i < len(os.Args); i++ {  
        s += arg + os.Args[i]+" " 
      show(s)

    } 

}

func show(site string) *http.Response {
     var mutex = &sync.Mutex{}
    mutex.Lock()
    resp, err := http.Get(site)
    if err != nil {
        log.Fatal(err)
    }
 defer resp.Body.Close()
robots, err := ioutil.ReadAll(resp.Body)
 
    if err != nil {
        log.Fatal(err)
    }

 fmt.Printf("%s", robots)

  mutex.Unlock()
  return resp;
}



func main() {  

   list()
    
 }  
      
    


