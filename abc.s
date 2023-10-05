.text

.global _start
_start:
	li x6, 0x2
	li x7, 0x5
	add x6, x7, x6

	la x7, my_label 
	lw x8, 0(x7)

	lw x9, 4(x7)
	add x9, x8, x9

.data
.balign 0

my_label:
	.word 0x8
	.word 0x5


