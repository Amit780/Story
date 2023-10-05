/*

*/  

    .text

    .global _start

_start:		// label
    //setting up the stack
    //
//    la x2, my_stack_end
//    la x9, my_func
//    jalr x9

	li x1, 1000

//label:
//    addi x0,x0,1
//    j label
//
//

    .data
    .align 8
//
//my_var:
//    .dword
//
//my_stack:
//    .space 0x2000
//my_stack_end:

    .end
