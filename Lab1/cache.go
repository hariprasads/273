package main

// DO NOT CHANGE THIS CACHE SIZE VALUE
const CACHE_SIZE int = 3
var memory []map[int]int
var make_map int=0

func Set(key int, value int) {
	// TODO: add your code here!
	if make_map == 0 {
		memory = make([]map[int]int,CACHE_SIZE)
		make_map = 1
	}
	current:=0
	element := map[int]int{key:value}
	for i:=0; i < CACHE_SIZE; i++ {
		temp := memory[i]
		if(temp[key]!=0) {
			current=1
			for j:=i; j>0; j-- {
				memory[j] = memory[j-1]
			}
			memory[0]=element
			break
		}
	}

	if(current == 0){
		for i:=CACHE_SIZE-1; i>0; i-- {
			memory[i]=memory[i-1]
		}
		memory[0]=element
	}

}

func Get(key int) int {
		// TODO: add your code here!
	current:=0
	value:= 0
for i:=0; i<CACHE_SIZE; i++ {
	element := memory[i]
	if(element[key]!=0){
		value = element[key]
		current=1
		for j:=i;j>0;j-- {
			memory[j]=memory[j-1]
		}
		memory[0]=element
		break
	}
}

if(current==0) {
	return -1
}else {
	return value
}

}
