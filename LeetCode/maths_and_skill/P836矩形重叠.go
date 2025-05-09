package main

// 数学：几何
// 力扣：https://leetcode.cn/problems/rectangle-overlap/description/

// 矩形如果不重叠，从x轴和y轴上看两个矩形就变成了两条线段，这两条线段肯定是不相交的，也就是说左边的矩形的最右边小于右边矩形的最左边，也就是rec1[2] < rec2[0] || rec2[2] < rec1[0]；
// y轴同理，下面的矩形的最上边小于上面矩形的最下边，也就是rec1[3] < rec2[1] || rec2[3] < rec1[1]。因为题目要求重叠算相离，所以加上=，最后取反就行啦

// func isRectangleOverlap(rec1 []int, rec2 []int) bool {
// 	return !((rec1[2] <= rec2[0] || rec2[2] <= rec1[0]) || (rec1[3] <= rec2[1] || rec2[3] <= rec1[1]))
// }

// 或者：
// 图解：https://leetcode.cn/problems/rectangle-overlap/solutions/155825/tu-jie-jiang-ju-xing-zhong-die-wen-ti-zhuan-hua-we/
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 判断x轴是否相交
	xBool := !(rec1[2] <= rec2[0] || rec2[2] <= rec1[0]) // rec1[2] <= rec2[0] || rec2[2] <= rec1[0]如果为true，表示没有相交。然后取反
	// 判断y轴是否相交
	yBool := !(rec1[3] <= rec2[1] || rec2[3] <= rec1[1])

	return xBool && yBool
}
