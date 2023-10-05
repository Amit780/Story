IMG_NAME=ABC

set -o xtrace
set -e

riscv64-unknown-elf-g++ -x assembler-with-cpp -c -march=rv64imafd -mabi=lp64d -mcmodel=medany -ggdb -o abc.o abc.s
#riscv64-unknown-elf-g++ -x assembler-with-cpp -c -march=rv64imafdc -mabi=lp64d -mcmodel=medany -ggdb -o sample.o sample.s
#riscv64-unknown-elf-g++ -c -march=rv64imafdc -mcmodel=medany -ggdb -o build/main.o src/main.cpp
riscv64-unknown-elf-g++ -o ${IMG_NAME}.elf -g -Xlinker -Map=output.map  -T linker.script -march=rv64imafd  -nostdlib abc.o #build/main.o
#-mabi=ilp32f
riscv64-unknown-elf-objdump -xsD ${IMG_NAME}.elf > ${IMG_NAME}.dis

set +o xtrace
