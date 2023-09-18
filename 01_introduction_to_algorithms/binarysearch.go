/*
Бинарный поиск - это алгоритм; на входе он получает отсортированный список элементов.
Если элемент, который вы ищите, присутствует в списке, то бинарный поиск возвращает ту позицию, в которой был найден.
В противном случае бинарный поиск возвращает null.
*/
package introductiontoalgorithms

func binarysearch(list []int, item int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		middle := (low + high) / 2
		if list[middle] == item { //Найдено
			return middle
		} else if list[middle] > item { // Много
			high = middle - 1
		} else { // Мало
			low = middle + 1
		}
	}
	return -1
}
