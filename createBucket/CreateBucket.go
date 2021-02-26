package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {

	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI4G1ymHvHH8pUe1HmHFSL", "SZYRvjdKzjJIfcLlDbFMAJTGYDz38P")
	//
	//if err != nil{
	//	fmt.Println("Error:",err)
	//	os.Exit(-1)
	//}
	//
	//err = client.CreateBucket("mybucket5231",oss.ACL(oss.ACLPublicRead))
	//
	//if err != nil{
	//	fmt.Println("Error: ",err)
	//	os.Exit(-1)
	//}

	//***************************************************************************************//

	lsRes, err := client.ListBuckets()
	if err != nil {
		// handel error
	}
	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}
