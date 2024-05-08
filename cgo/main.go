package main

/*
#cgo CFLAGS: -O0 -DUNW_LOCAL_ONLY

#include <libunwind.h>
#include <stdio.h>

void backtrace() {
	unw_cursor_t cursor;
	unw_context_t context;

	// Initialize cursor to current frame for local unwinding.
	unw_getcontext(&context);
	unw_init_local(&cursor, &context);

	// Unwind frames one by one, going up the frame stack.
	while (unw_step(&cursor) > 0) {
		unw_word_t offset, pc;
		unw_get_reg(&cursor, UNW_REG_IP, &pc);
		if (pc == 0) {
			break;
		}
	    printf("0x%llx:", pc);
		char sym[256];
		if (unw_get_proc_name(&cursor, sym, sizeof(sym), &offset) == 0) {
			printf(" (%s+0x%llx)\n", sym, offset);
		} else {
			printf(" -- error: unable to obtain symbol name for this frame\n");
		}
  	}
}

void two(int m) {
	int ret = 0;
	for(int i = 0; i < m; i++)
		ret += i;
	printf("two %d\n", ret);
	backtrace();
}

void one(int n) {
	printf("one %d\n", n);
	two(n * 2);
}
*/
import "C"

//go:noinline
func goone() {
	C.one(1000000000)
}

func main() {
	goone()
}
