package main

func main()  {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil{
		fmt.Printf("Request failed %s\n",err)
	}else{
		data, _:= ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
 
}