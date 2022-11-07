package main

import "fmt"

/**
 * Oct 25, 2022
 * ptp-glng033
 * GLNG033ONL1006
 * iota
 * simplifies constants predefined value
 * iota value: index of constant
 * ref: https://go.dev/ref/spec#Iota
 */

//-------------------------------------------------------------
/* [START] BASIC CONST() USAGE */
// in const(), every variables which is not inline-assigned
// will be assigned with same value from previous expression:
// const (
/* VAR 		EXPR	RESULT*/
// 	MYCONST1 = 1 //1
// 	MYCONST2     //1
// 	MYCONST3     //1
// 	MYCONST4     //1
// 	MYCONST5     //1
// 	MYCONST6     //1
// )

// const (
/* VAR 		EXPR	RESULT*/
// 	MYCONST1 = 0     //0
// 	MYCONST2         //0
// 	MYCONST3 = 5     //5
// 	MYCONST4         //5
// 	MYCONST5 = 1 + 1 //2
// 	MYCONST6         //2
// )

/* [END] BASIC CONST() USAGE */

//-------------------------------------------------------------
/* [START] IOTA USAGE */
// with basic pricipal of const() assignation process,
// iota could be used to make constant values unique.
// you can name it as enumeration as well (?)
// const (
/* VAR 		EXPR	RESULT*/
// 	MYCONST1 = iota //0
// 	MYCONST2     //1
// 	MYCONST3     //2
// 	MYCONST4     //3
// 	MYCONST5     //4
// 	MYCONST6     //5
// )

const (
	/* VAR 		EXPR	RESULT*/
	MYCONST1 = 0        //0
	MYCONST2            //0
	MYCONST3 = iota     //2
	MYCONST4            //3
	MYCONST5 = iota * 2 //8
	MYCONST6            //10
)

/* [END] IOTA USAGE */

func main() {
	fmt.Printf("MYCONST1: [%d]\nMYCONST2: [%d]\nMYCONST3: [%d]\nMYCONST4: [%d]\nMYCONST5: [%d]\nMYCONST6: [%d]\n",
		MYCONST1, MYCONST2, MYCONST3, MYCONST4, MYCONST5, MYCONST6)
}
