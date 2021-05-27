/*
1. ../../bin/go build -buildmode=c-archive hello.go
2. gcc main.c hello.a -lpthread
3. ./a.out http://baidu.com
*/
#include "hello.h"
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
int main(int argc, char*argv[]) {
  if (argc != 2) {
    printf("Usage %s $url\n", argv[0]);
    return -1;
  }
  printf("Browsering %s\n", argv[1]);
  GoString url = {
    argv[1],
    strlen(argv[1]),
  };
  char* body = Browser(url);
  puts(body);
  free(body);
  return 0;
}
