package main 
import(
"encoding/json"
"strings"
"fmt"
"io/ioutil"
)


func loadData(file string)[]byte {
    data , _ := ioutil.ReadFile(file)
        return data
    }

type Busstop struct{
    Township string `json:"township"` 
    Lines []string
    
    }

func in(x string , list []string)bool{
    for _,k := range list{
        if k ==x { 
        return true
        }
    }
   return false
}


func main(){

    bl := loadData("data/bus-stop-ids-of-lines.js")
    var buslines map[string][]string
    json.Unmarshal([]byte(bl),&buslines)

    b := loadData("data/bus-stop-data-by-id.js")
    var bustops_list map[string]Busstop
    json.Unmarshal([]byte(b),&bustops_list)
    var tship_names []string

    for _,j := range bustops_list {
        if !in(j.Township,tship_names){
           tship_names =append(tship_names,j.Township) 
            }
        }

var township map[string][]string
township = make (map[string][]string)


for _,i :=range tship_names{
    township[i]= getTotalBustop(i,bustops_list)
    }

var township_with_busline map[string][]string
township_with_busline = make(map[string][]string)

for k,v := range township{
    var data []string 
    for _,i :=range  v {
        for  b_line,b_ids :=range buslines {
            if in(i , b_ids){
                if !in(b_line,data){
                data=append(data,b_line)}
                }
            }

    }
    township_with_busline[k]=data
}


for k,v :=range township_with_busline{
   fmt.Printf(k) 
   fmt.Printf(strings.Join(v," , "))
   fmt.Printf("\n")


    }







}


func getTotalBustop(t string , list map[string]Busstop)[]string{
    var bus_nos []string
    for k,v := range list {
        if t == v.Township{
            bus_nos = append(bus_nos,k)
            }
        }
    return bus_nos
    }
