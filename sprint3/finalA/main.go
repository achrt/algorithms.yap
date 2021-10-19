package main

/*
-- ПРИНЦИП РАБОТЫ --
По условию задачи мы имеем "сломанный массив", в котором сбит порядок элементов.
Из этого следует, что массив упорядоченный, за некоторым исключением, поэтому
можно применить бинарный поиск. Реализация соответсвует алгоритму бинарного поиска, но во время выбора
в какой части продолжать поиск, осуществляется проверка, является ли подмассив сломанным. Если нет
то сверяем искомый элемент К по принципу left < K < right. Если подмассив сломан, то проверяем
правую границу K < right.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Корректность обеспечивается условием задачи - массив был отсортирован по возрастанию.
Формула сравнения left < K < right в нашем случае будет работать.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Алгоритм на каждой итерации делит входные данные на 2 части и осуществляет поиск только в одной из них.
В худшем случае К будет найдено за O(log n)
.
-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Пространственная сложность O(n), т.к. используется структура со ссылкой на исходный массив
и 2 константы для указателей.
*/

// https://contest.yandex.ru/contest/23815/run-report/55015069/

func brokenSearch(arr []int, k int) int {
	d := &Data{
		digits: &arr,
		lenght: len(arr),
	}
	d.SearchInLoop(0, d.Len()-1, k)
	return d.position
}

type Data struct {
	digits   *[]int
	lenght   int
	position int
}

func (d *Data) Len() int {
	return d.lenght
}

func (d *Data) SearchInLoop(left, right, k int) {
	for {
		middle := (right + left) / 2
		if middle == left {
			for i := left; i <= right; i++ {
				if (*d.digits)[i] == k {
					d.position = i
					return
				}
			}
			d.position = -1
			break
		}

		// проверка правых краев разбитого массива
		if (*d.digits)[right] == k {
			d.position = right
			break
		}
		if (*d.digits)[middle] == k {
			d.position = middle
			break
		}

		// порядок не сломан, число К в диапазоне, поиск в левой части
		if (*d.digits)[left] < (*d.digits)[middle] && (*d.digits)[left] <= k && k <= (*d.digits)[middle] {
			right = middle
			continue
		}

		// порядок сломан, K <= middle, поиск в левой части
		if (*d.digits)[left] > (*d.digits)[middle] && k <= (*d.digits)[middle] {
			right = middle
			continue
		}

		// порядок не сломан, число К в диапазоне, поиск в правой части
		if (*d.digits)[middle] < (*d.digits)[right] && (*d.digits)[middle] <= k && k <= (*d.digits)[right] {
			left = middle
			continue
		}

		// порядок сломан, последний кейс, поиск в правой части
		left = middle
	}
}
