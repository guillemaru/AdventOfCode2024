package main

var updates = [][]int{
	{79, 33, 22, 64, 73, 93, 58},
	{27, 93, 55, 97, 43, 67, 84, 25, 34},
	{28, 55, 24, 77, 17, 14, 44},
	{79, 67, 97, 29, 47, 96, 45, 25, 44, 55, 43, 76, 77},
	{87, 84, 17, 25, 45, 44, 88},
	{88, 93, 72, 58, 73, 12, 57, 49, 51, 33, 64},
	{11, 28, 55, 24, 29, 56, 75},
	{73, 58, 65, 11, 93, 63, 81, 98, 33, 22, 79, 72, 64},
	{17, 44, 25, 43, 28, 77, 47},
	{65, 27, 78, 49, 79, 64, 52, 98, 12, 51, 38, 58, 73, 82, 81, 72, 14, 63, 33, 45, 88},
	{29, 43, 47, 87, 75, 25, 34, 26, 32, 79, 82},
	{47, 67, 84, 75, 25, 34, 26, 32, 14, 44, 79, 58, 82, 72, 94, 64, 98},
	{38, 24, 31, 11, 82, 81, 72, 73, 63, 78, 65, 52, 98},
	{75, 25, 76, 67, 43, 56, 87, 45, 44, 96, 17, 34, 97, 58, 14, 79, 32, 29, 47, 24, 84, 77, 26},
	{72, 73, 98, 81, 78, 22, 27, 11, 57, 28, 55, 24, 16},
	{52, 22, 57, 55, 24, 16, 76, 77, 43, 47, 67, 87, 25},
	{77, 56, 47, 87, 84, 17, 25, 34, 32, 14, 45, 44, 58, 88, 82, 72, 94},
	{84, 87, 29, 97, 55, 75, 32, 57, 31, 17, 76, 56, 16, 93, 25, 67, 26},
	{97, 31, 52, 25, 16, 11, 17},
	{38, 33, 65, 12, 78, 63, 52, 22, 27, 11, 31, 93, 57, 28, 55, 24, 16, 97, 76, 77, 56, 43, 47},
	{43, 47, 67, 87, 75, 96, 14, 45, 79, 58, 51, 72, 94, 73, 64},
	{87, 96, 97, 28, 11, 43, 55, 31, 47, 76, 24, 25, 93, 57, 26, 34, 16},
	{63, 22, 11, 31, 93, 57, 24, 16, 97, 76, 77, 56, 47, 67, 87, 17, 75},
	{58, 82, 12, 78, 32, 94, 79, 72, 14, 26, 34, 81, 98, 45, 51, 63, 33},
	{32, 14, 45, 44, 58, 51, 88, 72, 73, 98, 81, 49, 38, 65, 12, 78, 63, 52, 22},
	{64, 49, 38, 52, 27, 31, 16, 97, 76},
	{93, 47, 28, 27, 33, 55, 67},
	{44, 79, 58, 88, 82, 72, 94, 73, 81, 49, 33, 65, 12, 63, 22, 11, 31},
	{88, 82, 72, 94, 73, 64, 98, 81, 49, 38, 33, 65, 12, 78, 63, 52, 22, 27, 11, 31, 93, 57, 55},
	{84, 43, 11, 29, 17, 22, 55, 25, 75, 28, 97, 31, 56, 76, 27, 47, 77, 67, 24, 57, 52},
	{73, 64, 98, 81, 49, 38, 33, 65, 12, 78, 63, 52, 22, 27, 11, 31, 93, 57, 24, 97, 29},
	{93, 57, 28, 16, 97, 76, 43, 47, 67, 87, 75, 25, 96, 34, 26},
	{65, 78, 52, 27, 11, 31, 93, 57, 55, 97, 29, 76, 77, 56, 47, 67, 87},
	{97, 29, 76, 56, 43, 67, 84, 75, 25, 96, 34, 26, 32, 45, 44, 79, 58},
	{98, 82, 28, 49, 63, 27, 11, 24, 64, 72, 12, 55, 22, 94, 65, 52, 78, 38, 57},
	{29, 28, 55, 34, 87, 76, 96, 84, 45, 25, 77, 56, 32, 75, 16, 14, 57, 43, 26, 24, 67, 17, 47},
	{11, 65, 43, 29, 63, 55, 56, 47, 57, 31, 78, 28, 77, 33, 38},
	{52, 22, 11, 28, 55, 24, 16, 97, 29, 76, 47, 17, 25},
	{24, 16, 97, 29, 77, 56, 47, 75, 25, 96, 34, 45, 44, 79, 58},
	{75, 45, 73, 51, 79, 96, 81, 49, 17, 88, 98, 26, 84, 25, 14, 87, 72, 44, 94, 34, 58},
	{87, 84, 17, 25, 96, 34, 32, 14, 45, 79, 58, 51, 88, 82, 72, 94, 73, 64, 98, 81, 49},
	{65, 12, 78, 52, 22, 27, 11, 31, 93, 57, 28, 55, 24, 16, 97, 29, 76, 77, 56, 43, 47, 67, 87},
	{33, 78, 22, 27, 31, 28, 76, 56, 67},
	{25, 14, 32, 34, 87, 84, 28, 55, 77, 56, 47, 75, 24, 96, 16, 93, 97, 26, 17, 29, 76},
	{84, 82, 43, 51, 94, 67, 25, 45, 64, 17, 75, 88, 79, 58, 73},
	{16, 84, 87, 77, 25, 96, 24, 34, 45, 47, 44, 56, 79, 29, 43, 14, 17, 75, 97, 32, 67, 58, 76},
	{43, 67, 84, 96, 34, 32, 58, 88, 72, 73, 64},
	{12, 22, 16, 28, 84, 24, 27},
	{94, 17, 82, 64, 79, 87, 67, 45, 26, 72, 75, 32, 88, 14, 73, 51, 96, 98, 58, 81, 44, 84, 25},
	{67, 84, 75, 25, 96, 34, 26, 14, 45, 44, 79, 58, 51, 88, 82, 72, 94, 73, 64},
	{77, 28, 55, 22, 31, 87, 24, 65, 78},
	{94, 11, 63, 28, 55, 78, 73, 22, 27, 81, 24, 72, 31, 65, 12},
	{96, 25, 73, 43, 64, 17, 87},
	{29, 76, 77, 56, 43, 47, 67, 87, 84, 17, 75, 25, 96, 34, 26, 32, 45, 44, 79, 58, 51, 88, 82},
	{81, 49, 38, 33, 65, 78, 22, 27, 11, 31, 93, 57, 16, 97, 76, 77, 56},
	{16, 97, 29, 76, 77, 56, 43, 47, 67, 87, 84, 17, 75, 25, 96, 34, 26, 14, 45, 44, 79, 58, 51},
	{67, 87, 84, 17, 75, 25, 96, 34, 32, 14, 45, 44, 79, 58, 51, 88, 82, 72, 94, 73, 64, 98, 81},
	{79, 44, 94, 43, 82, 51, 88, 56, 14, 73, 34},
	{84, 67, 29, 34, 25, 76, 77},
	{58, 81, 64, 84, 25, 72, 88, 51, 82},
	{11, 17, 87, 97, 22, 76, 84, 24, 16, 47, 29, 27, 67, 57, 56, 31, 75, 63, 77, 43, 28},
	{64, 98, 81, 65, 78, 63, 31, 57, 28, 55, 76},
	{72, 81, 49, 65, 52, 22, 11, 31, 28},
	{31, 93, 28, 16, 97, 76, 77, 56, 47, 87, 84, 17, 25, 96, 34, 26, 32},
	{88, 72, 65, 12, 63, 22, 11},
	{45, 38, 65, 49, 26, 72, 58, 33, 32, 81, 94, 98, 34, 14, 79, 51, 64, 96, 78},
	{51, 49, 12, 22, 27},
	{73, 64, 98, 81, 49, 38, 33, 78, 63, 52, 11, 31, 93, 57, 55, 24, 16, 97, 29},
	{64, 34, 84, 81, 73, 26, 17, 14, 96, 67, 98, 82, 79, 72, 45},
	{98, 81, 38, 65, 52, 31, 57, 28, 16, 29, 77},
	{64, 98, 49, 78, 22, 27, 11, 28, 24, 29, 76},
	{75, 25, 26, 79, 58, 51, 82, 64, 98, 33, 65},
	{76, 77, 67, 87, 84, 75, 25, 96, 34, 26, 45, 44, 79, 58, 51, 88, 72},
	{26, 76, 75, 32, 34, 24, 45, 28, 84, 96, 17, 97, 43, 44, 25},
	{24, 29, 67, 87, 25, 34, 32, 44, 58},
	{49, 38, 33, 78, 27, 11, 93, 57, 55, 97, 29, 76, 77, 56, 43},
	{29, 87, 76, 14, 45, 96, 57},
	{38, 82, 65, 98, 63, 51, 31, 57, 33, 22, 58, 93, 78, 73, 81, 94, 12, 27, 11, 52, 49},
	{26, 45, 79, 88, 82, 72, 94, 98, 81, 49, 38, 63, 52},
	{14, 45, 44, 51, 88, 64, 98, 38, 65, 12, 78, 63, 27},
	{87, 75, 26, 32, 14, 44, 72},
	{73, 22, 12, 14, 63, 52, 82, 58, 78},
	{31, 28, 33, 67, 47, 22, 43, 65, 57, 76, 11},
	{94, 73, 64, 98, 49, 38, 65, 78, 52, 11, 93, 57, 16},
	{77, 67, 84, 17, 75, 34, 32, 44, 79, 58, 51, 72, 94},
	{38, 33, 65, 12, 63, 52, 27, 31, 93, 57, 16, 97, 76},
	{25, 11, 56, 77, 22, 57, 67, 16, 52, 29, 97, 75, 17, 28, 55},
	{49, 81, 33, 22, 65, 27, 51, 57, 82, 78, 11, 38, 93, 31, 63, 88, 98, 72, 73, 52, 28, 12, 64},
	{34, 26, 32, 14, 45, 44, 79, 58, 51, 88, 82, 72, 94, 73, 64, 98, 81, 49, 38, 65, 12, 78, 63},
	{64, 58, 88, 22, 44, 65, 51, 63, 31},
	{12, 49, 31, 72, 73, 79, 52, 11, 78, 64, 33, 93, 94},
	{31, 93, 57, 55, 16, 97, 29, 76, 43, 87, 84, 17, 25, 96, 34, 26, 32},
	{96, 34, 26, 32, 14, 44, 79, 58, 51, 88, 82, 72, 94, 73, 64, 98, 81, 49, 38, 33, 65, 12, 78},
	{56, 43, 47, 87, 84, 17, 25, 96, 26, 14, 45, 44, 79, 58, 88, 82, 72, 94, 73},
	{56, 47, 87, 84, 25, 96, 34, 45, 79, 58, 51, 88, 72},
	{63, 72, 88, 64, 79, 44, 58, 33, 26, 32, 49, 98, 45, 94, 12, 34, 65, 78, 81, 38, 82},
	{64, 33, 78, 63, 27, 93, 28, 55, 76},
	{45, 44, 58, 94, 12, 78, 11},
	{38, 33, 65, 12, 78, 63, 52, 22, 27, 11, 93, 28, 55, 24, 16, 97, 29, 76, 77, 43, 47},
	{27, 17, 78, 31, 16, 57, 63, 87, 56, 97, 29, 67, 24, 55, 22, 93, 84, 11, 77},
	{67, 55, 25, 57, 43, 84, 31},
	{55, 16, 76, 77, 96, 26, 79},
	{87, 84, 75, 25, 96, 34, 26, 32, 14, 44, 79, 58, 51, 88, 82, 72, 94, 73, 64, 98, 81},
	{63, 52, 22, 27, 11, 31, 93, 57, 28, 55, 24, 16, 97, 29, 76, 77, 56, 43, 47, 67, 84, 17, 75},
	{73, 98, 81, 49, 38, 33, 65, 12, 63, 52, 22, 27, 11, 93, 57, 28, 55, 16, 29},
	{82, 73, 49, 33, 52},
	{31, 51, 38, 81, 98, 11, 64, 27, 79, 72, 44, 88, 63},
	{88, 82, 73, 64, 98, 38, 33, 65, 12, 78, 63, 22, 27, 11, 31, 93, 57, 28, 55},
	{44, 79, 58, 51, 88, 82, 72, 94, 73, 98, 81, 49, 38, 33, 65, 78, 63, 52, 22, 11, 31},
	{96, 93, 29, 16, 17, 57, 47, 97, 24, 87, 25, 32, 67, 43, 76, 84, 14, 77, 28, 26, 56},
	{25, 96, 26, 45, 58, 51, 72, 94, 64, 98, 81, 65, 12},
	{98, 38, 17, 26, 75, 34, 73, 51, 84},
	{52, 22, 27, 31, 28, 55, 24, 16, 97, 77, 56, 43, 47, 67, 87, 17, 25},
	{64, 25, 87, 75, 34, 67, 72, 17, 47, 94, 43},
	{98, 38, 65, 12, 78, 63, 52, 22, 11, 93, 57, 28, 55, 29, 76},
	{93, 75, 57, 22, 63, 97, 67},
	{14, 84, 96, 72, 43, 88, 58, 79, 34},
	{78, 57, 72, 12, 55, 73, 49, 31, 65, 11, 22, 27, 33, 24, 16},
	{33, 78, 98, 51, 63, 81, 31, 52, 82, 22, 65, 73, 94, 79, 58, 12, 27},
	{17, 75, 88, 82, 38, 81, 96, 64, 34, 45, 72, 26, 33},
	{22, 58, 44, 45, 12, 65, 64, 79, 49, 72, 14, 81, 82, 94, 52, 51, 33, 88, 78, 32, 73, 98, 63},
	{49, 38, 33, 12, 63, 52, 22, 27, 11, 31, 93, 57, 28, 55, 24, 16, 97, 29, 76, 77, 56},
	{47, 87, 84, 75, 25, 96, 34, 26, 45, 44, 58, 51, 88, 72, 94, 73, 98},
	{38, 57, 52, 56, 11, 29, 22, 12, 24, 81, 76},
	{82, 79, 98, 27, 51, 63, 94, 65, 52, 44, 78, 11, 88, 58, 22, 33, 12, 81, 45, 64, 38, 73, 72},
	{79, 43, 56, 67, 26, 84, 58, 25, 72, 44, 87, 17, 96, 82, 51},
	{33, 52, 38, 63, 94, 31, 82, 51, 64, 11, 58, 22, 81, 73, 98, 78, 93, 27, 88, 57, 12},
	{97, 27, 56, 38, 65, 78, 22, 28, 12, 33, 93},
	{16, 97, 22, 87, 29, 43, 31, 77, 76},
	{28, 16, 29, 67, 87, 14, 44},
	{44, 79, 51, 88, 82, 72, 94, 98, 49, 38, 33, 12, 78, 63, 52, 22, 31},
	{76, 77, 56, 43, 67, 87, 84, 17, 75, 25, 96, 34, 26, 32, 14, 45, 44, 79, 58, 51, 88, 82, 72},
	{28, 16, 94, 81, 73, 38, 11, 31, 12, 97, 65, 93, 64, 24, 55},
	{65, 12, 78, 52, 22, 27, 11, 31, 93, 57, 28, 55, 24, 16, 97, 29, 76, 77, 56, 43, 47, 67, 87},
	{45, 94, 64, 33, 52},
	{81, 28, 98, 93, 88, 33, 38, 64, 52, 22, 72, 31, 49, 65, 82, 73, 94},
	{79, 58, 72, 94, 73, 64, 81, 38, 33, 65, 12, 52, 22, 27, 11, 31, 93},
	{26, 32, 45, 44, 58, 51, 82, 72, 94, 73, 64, 81, 49, 33, 78, 63, 52},
	{47, 93, 76, 22, 67, 52, 55, 29, 97, 11, 12, 16, 78, 24, 57, 84, 63},
	{79, 58, 51, 88, 94, 64, 98, 49, 38, 33, 65, 12, 78, 63, 11, 31, 93},
	{25, 34, 26, 32, 44, 58, 94, 64, 98, 81, 38, 65, 12},
	{47, 93, 22, 29, 56, 57, 43, 11, 55, 65, 76, 97, 12, 16, 31, 87, 78},
	{38, 98, 78, 11, 77, 16, 93},
	{34, 17, 96, 26, 77, 75, 82, 72, 14, 43, 44, 32, 84, 94, 79},
	{75, 79, 51, 88, 82, 64, 38, 33, 65},
	{94, 64, 34, 12, 51, 72, 81, 49, 44, 78, 33, 98, 82, 65, 32, 79, 26, 63, 45},
	{22, 11, 31, 97, 29, 76, 56, 67, 96},
	{88, 76, 17, 32, 29, 97, 84, 45, 77, 34, 79, 47, 44, 67, 26, 25, 56},
	{55, 24, 97, 29, 76, 77, 56, 43, 67, 87, 84, 17, 75, 25, 96, 26, 14, 45, 79},
	{31, 93, 57, 28, 97, 29, 77, 43, 87, 17, 32},
	{27, 11, 31, 93, 28, 55, 24, 16, 97, 29, 76, 77, 47, 67, 87, 84, 96},
	{29, 76, 77, 56, 43, 47, 87, 84, 75, 25, 96, 34, 26, 32, 14, 44, 79, 88, 82},
	{94, 64, 38, 65, 63, 52, 27, 31, 28, 24, 97},
	{81, 98, 26, 38, 72, 82, 84, 49, 44},
	{51, 82, 72, 94, 64, 98, 49, 38, 12, 78, 63, 52, 22, 27, 31},
	{14, 75, 32, 88, 51, 43, 45},
	{34, 26, 82, 73, 64, 33, 65, 12, 63},
	{32, 17, 75, 25, 94, 81, 79, 51, 98, 26, 82},
	{34, 65, 98, 96, 49, 79, 72, 51, 78, 64, 81, 44, 73, 12, 94},
	{24, 47, 75, 34, 32},
	{38, 33, 65, 12, 78, 63, 52, 22, 11, 31, 57, 28, 55, 16, 97, 29, 76, 77, 56, 43, 47},
	{51, 94, 73, 64, 98, 81, 33, 12, 78, 22, 27, 11, 31, 93, 28},
	{33, 63, 52, 22, 27, 11, 31, 93, 28, 55, 97, 29, 67},
	{65, 12, 78, 63, 52, 22, 27, 31, 93, 28, 55, 24, 97, 29, 76, 77, 56, 43, 47, 67, 87},
	{34, 84, 87, 43, 29, 96, 26, 32, 56, 45, 44, 47, 28, 97, 75},
	{96, 58, 26, 51, 47, 16, 43, 87, 45, 67, 76},
	{27, 97, 76, 77, 67, 87, 84, 75, 34},
	{93, 77, 65, 29, 33, 27, 16, 38, 28, 81, 97, 98, 63, 76, 11, 55, 57, 31, 52},
	{11, 55, 24, 97, 43, 47, 84, 17, 26},
	{12, 73, 98, 81, 44, 45, 26, 88, 14, 82, 79, 64, 25, 94, 32, 72, 58, 33, 65, 38, 51, 49, 96},
	{64, 44, 25, 58, 17, 88, 43, 79, 14, 75, 87, 45, 72, 47, 82, 51, 94, 67, 26},
	{94, 12, 38, 65, 44, 73, 64, 51, 52, 79, 88, 27, 81, 49, 11},
	{28, 14, 47, 77, 55, 43, 32, 75, 96, 29, 16, 57, 67, 34, 97, 26, 25, 87, 76, 17, 56, 24, 45},
	{63, 52, 22, 11, 57, 28, 16, 29, 43, 47, 87, 84, 75},
	{98, 81, 33, 65, 12, 78, 52, 22, 11, 31, 93, 57, 55, 24, 16, 97, 29, 76, 77},
	{47, 67, 87, 84, 75, 34, 26, 14, 58, 82, 98},
	{64, 73, 79, 45, 65, 98, 44, 81, 78, 14, 82},
	{93, 56, 67, 25, 96, 32, 14},
	{45, 79, 82, 72, 94, 73, 64, 98, 81, 65, 12, 27, 11},
	{87, 17, 25, 34, 26, 45, 44, 79, 88, 82, 73, 64, 98, 81, 49},
	{12, 22, 63, 65, 38, 28, 94, 55, 49, 64, 24},
	{28, 55, 24, 16, 97, 29, 76, 56, 43, 47, 67, 87, 84, 17, 75, 96, 34, 26, 32, 45, 44},
	{88, 79, 84, 67, 26, 51, 17},
	{96, 88, 56, 47, 67, 58, 79, 34, 51, 76, 72, 43, 14, 45, 82, 87, 44, 26, 84, 75, 77},
	{38, 33, 65, 12, 78, 63, 52, 22, 27, 11, 31, 93, 28, 55, 24, 16, 97, 29, 76, 77, 56, 43, 47},
}