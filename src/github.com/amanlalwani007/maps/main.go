package main

import "fmt"

func main() {
	// emailkeyvaluemapper := make(map[string]string)
	//assign kv
	// emailkeyvaluemapper["aman"] = "amanlalwani0807@gmail.com"
	// emailkeyvaluemapper["aman1"] = "lalwaniaman0807@gmail.com"
	// emailkeyvaluemapper["aman2"] = "lalwaniaman0807@gmail.com"
	//declare map and key value
	//HashMap<String,String> emailkeyvaluemapper=new HashMap<>;
	emailkeyvaluemapper := map[string]string{"aman": "amanlalwani0807@gmail.com", "aman1": "lalwaniaman0807@gmail.com", "aman2": "amanlalwani0807@gmail.com"}
	fmt.Println(emailkeyvaluemapper)
	fmt.Println(len(emailkeyvaluemapper))
	fmt.Println(emailkeyvaluemapper["aman"])
	//Delete from map
	delete(emailkeyvaluemapper, "aman1")
	fmt.Println(len(emailkeyvaluemapper))

}
