#include "stdio.h"

void testScanfs() {
    char c;

    printf("输入任意字符:");

    scanf("%c", &c);

    printf("%c ascii is : %d\n",c,c);
}