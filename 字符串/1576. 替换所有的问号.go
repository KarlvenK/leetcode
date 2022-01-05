package main

func modifyString(s string) string {
	arr := []byte(s)
	n := len(s)

	get_char := func(i int) byte {
		return byte(int('a') + i)
	}

	var dfs func(int)
	dfs = func(cur int) {
		if cur != n-1 {
			if arr[cur+1] == '?' {
				dfs(cur + 1)
			}
			if cur == 0 {
				for i := 0; i < 26; i++ {
					if get_char(i) != arr[cur+1] {
						arr[cur] = get_char(i)
						return
					}
				}
			} else {
				for i := 0; i < 26; i++ {
					c := get_char(i)
					if c != arr[cur-1] && c != arr[cur+1] {
						arr[cur] = c
						return
					}
				}
			}

		} else {
			if n == 1 {
				arr[cur] = 'a'
				return
			} else {
				for i := 0; i < 26; i++ {
					if get_char(i) != arr[cur-1] {
						arr[cur] = get_char(i)
						return
					}
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if arr[i] == '?' {
			dfs(i)
		}
	}
	return string(arr)
}
