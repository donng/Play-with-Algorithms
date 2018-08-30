package MaxHeap

type MaxHeap struct {
	data     []int
	count    int
	capacity int
}

// 构造一个空堆, 可容纳capacity个元素
func GetMaxHeap(capacity int) *MaxHeap {
	data := make([]int, capacity+1)
	count := 0

	return &MaxHeap{data, count, capacity}
}

// 返回堆中的元素个数
func (h *MaxHeap) Size() int {
	return h.count
}

// 返回一个布尔值, 表示堆中是否为空
func (h *MaxHeap) IsEmpty() bool {
	return h.count == 0
}

// 像最大堆中插入一个新的元素 item
func (h *MaxHeap) Insert(int int) {
	if h.count+1 > h.capacity {
		panic("out of range.")
	}
	h.data[h.count+1] = int
	h.count++
	h.shiftUp(h.count)
}

func (h *MaxHeap) ExtractMax() int {
	ret := h.data[1]

	swap(h.data, 1, h.count)
	h.count--
	h.shiftDown(1)

	return ret
}

//********************
//* 最大堆核心辅助函数
//********************
func (h *MaxHeap) shiftUp(k int) {
	for h.data[k] > h.data[k/2] && k > 1 {
		swap(h.data, k, k/2)
		k /= 2
	}
}

func (h *MaxHeap) shiftDown(k int) {
	for 2*k <= h.count {
		j := 2 * k
		if (j+1 <= h.count && h.data[j+1] > h.data[j]) {
			j++
		}
		if h.data[k] >= h.data[j] {
			break
		}
		swap(h.data, k, j)
		k = j
	}
}

// 交换堆中索引为i和j的两个元素
func swap(data []int, l int, r int) {
	data[l], data[r] = data[r], data[l]
}
