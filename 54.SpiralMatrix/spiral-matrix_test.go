package _4_SpiralMatrix

import (
	"reflect"
	"testing"
)

//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
func TestSpiralOrder(t *testing.T) {
	martix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	re := spiralOrder(martix)
	excepted := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	if !reflect.DeepEqual(re, excepted) {
		t.Errorf("failed! %v", re)
	}
}

//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
func TestSpiralOrder2(t *testing.T) {
	martix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	re := spiralOrder(martix)
	excepted := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
	if !reflect.DeepEqual(re, excepted) {
		t.Errorf("failed! %v", re)
	}
}
