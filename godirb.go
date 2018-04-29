package main


import (
   "net/http"
   // "net/url"
   "fmt"
   // "os"
   // "strconv"
)

//TODO
// maybe use slices here?
var result_slice []string


func bruteforce_url(url_c chan string,  {

   for{
      url := <- url_c
      resp, resperr := http.Get(url)

      if resperr != nil {
         fmt.Println("\033[31m%s \033[0m\n", resperr)
      } else {
         status_code := resp.StatusCode

         // 401 unauthorized
         if status_code == 401{
            result_slice = append(result_slicem, url)
         }

         // 200 Found
         } else if status_code == 200{
            result_slice = append(result_slicem, url)


         // 403 Forbidden
         } else if status_code == 403{

            result_slice = append(result_slicem, url)
         }
   }
}



// gets robots.txt entires, returns string array with identified ressources
func getRobots(url string) []string {

      robots_url := url
   if last_char := url[len(url)-1:]; last_char == "/"{
      robots_url = robots_url + "robots.txt"
   } else{
      robots_url = robots_url + "/robots.txt"
   }

   resp, resperr := http.Get(robots_url)

   if resperr != nil {
      fmt.Println("\033[31m%s \033[0m\n", resperr)
   } else {
      status_code := resp.StatusCode

      // 401 unauthorized
      if status_code == 401{

      // 200 Found
      } else if status_code == 200{

      // 403 Forbidden
      } else if status_code == 403{

      }



   }


}

func generate_target_urls(target string){

}
