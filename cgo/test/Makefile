DEBUG_OPT :=
OPT_LV := -O3
CFLAGS=-std=c99 -g ${OPT_LV} -Wall -Werror
OBJS=AuxIndexStructures_c.o Clustering_c.o error_c.o faiss_c.o Index_c.o index_factory_c.o

all: libfaiss.a

libfaiss.a: $(OBJS)
	ar -rcs libfaiss.a *.o

.PHONY: debug
debug: override OPT_LV := -O0
debug: override DEBUG_OPT := debug
debug: all

.PHONY: clean
clean:
	rm -f *.o *.a
