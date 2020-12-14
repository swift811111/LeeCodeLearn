package main

func main() {
	
}

func rotate(matrix [][]int)  {
    x := int(math.Ceil(float64(len(matrix))/2))
    ix := 0
    iy := 0
    for i:=0; i<x; i++ {
        if len(matrix) % 2 > 0 && i == x-1 {
            return
        }
        for j:=0 ;j < len(matrix)-2*i-1; j++ {
            temp := matrix[ix+j][iy+len(matrix)-2*i-1]
            last := matrix[ix][iy+j]
            matrix[ix+j][iy+len(matrix)-2*i-1] = last
            last = temp
            
            temp = matrix[ix+len(matrix)-2*i-1][iy+len(matrix)-2*i-1-j]
            matrix[ix+len(matrix)-2*i-1][iy+len(matrix)-2*i-1-j] = last
            last = temp
            
            temp = matrix[ix+len(matrix)-2*i-1-j][iy]
            matrix[ix+len(matrix)-2*i-1-j][iy] = last
            last = temp
            
            temp = matrix[ix][iy+j]
            matrix[ix][iy+j] = last
            
        }
        ix++
        iy++
    }
}


// Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
// Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]