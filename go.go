package main
import (
  "fmt"
)
var jumlah int32
var angka int32
var hasil int32
var x int32

func main() {


 
    fmt.Scanf("%d", &jumlah)

  
    if jumlah < 1 || jumlah > 100{
      return
    
    }else{
   
      
    num(jumlah)

      
    return 
}  
   
}





func sum(x int32, total int32) int32 {
  
 
  
  if x==0{
    return total
         
  }else{
 

  
fmt.Scanf("%d", &angka)
  
if angka < -100 || angka > 100 {
return 0
}else if angka < 0 {
                return sum(x - 1, total)

                           }else{
                            total += angka * angka
                             return sum(x - 1, total)
                              }
  }

}


func num(jumlah int32) int32 {
 
  

if jumlah == 0 {
return 0
}else{
fmt.Scanf("%d", &x)
if x < 1 || x > 100 {
return 0
}else{
       hasil = sum(x, 0)
       fmt.Println(hasil)
       num(jumlah - 1)
       return hasil
  }    



}
}





