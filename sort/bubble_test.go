package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 2, 5, 3, 1, 8, 3, 5, 0, 6, 2, 4, 6}
	fmt.Printf("bubble排序前：%v \n", arr)
	BubbleSort(arr)

	//  最好用bce 方式比较两个slice,以下反射不推荐
	okRes := []int{0, 1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6, 8}
	fmt.Printf("bubble排序后：%v \n", arr)
	ok := reflect.DeepEqual(arr, okRes)

	if ok {
		t.Log("测试通过")
	} else {
		t.Error("测试失败")
	}
}
