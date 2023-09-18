/*
Бинарный поиск - это алгоритм; на входе он получает отсортированный список элементов.
Если элемент, который вы ищите, присутствует в списке, то бинарный поиск возвращает ту позицию, в которой был найден.
В противном случае бинарный поиск возвращает null.
*/
package introductiontoalgorithms

import "testing"

func Test_binarysearch(t *testing.T) {
	type args struct {
		list []int
		item int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "OK",
			args: args{
				list: []int{1, 3, 5, 7, 9},
				item: 3,
			},
			want: int(1),
		},
		{
			name: "OK",
			args: args{
				list: []int{1, 3, 5, 7, 9},
				item: 10,
			},
			want: int(-1),
		},
		{
			name: "OK",
			args: args{
				list: []int{1, 3, 5, 7, 9},
				item: 7,
			},
			want: int(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarysearch(tt.args.list, tt.args.item); got != tt.want {
				t.Errorf("binarysearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
