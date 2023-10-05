IMG_NAME=abc

set -o xtrace
riscv64-unknown-elf-g++ -x assembler-with-cpp -c -march=rv64imafdc -mcmodel=medany -mabi=lp64d  -ggdb -o abc.o abc.s
riscv64-unknown-elf-g++ -x assembler-with-cpp -c -march=rv64imafdc -mcmodel=medany -mabi=lp64d -ggdb -o sample.o sample.s
#riscv64-unknown-elf-g++                       -c -march=rv64imafdc -mcmodel=medany -ggdb -o  ./uart.o ../common_files/uart.cpp
riscv64-unknown-elf-g++ -o abc.elf -march=rv64imafdc -g -Xlinker -Map=output.map  -T linker_tut1.script -nostdlib abc.o sample.o
riscv64-unknown-elf-objdump -xsD abc.elf > abc.dis

set +o xtrace

########  Notes/Explanation   #########
###    Explanation  ####
#
# 1. The command on line 4 is used to build assembly files(denoted by .s extension) to object files(denoted by .o extension).
# 2. The command on Line 5 is used to build C++ files(denoted by .cpp extension) to object files.
# 3. Line 6 takes all object files thus produced and Linkes them together based on a Linker script provided(here linker_tut1.script) and produces a output map and an executable of .elf format.
# 4. The statement on Line 7 is to dump the .elf file to a disassembly file, which is not required by the simulator but is used for debugging purposes by the programmer.
#
###    Notes   ###
# 
# 1. A disassembly file shows the way the RISCV core executes all of the instructions written along with their corresponding code memory address.
#
######## end ##########
