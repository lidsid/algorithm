package insertion

func Sort(item []int) {
	var n = len(item)
	for i:=1;i<n;i++ {
		j := i
		for j>0 {
			if item[j-1] > item[j] {
				item[j-1],item[j] = item[j],item[j-1]
			}
			j = j-1
		}
	}
}

