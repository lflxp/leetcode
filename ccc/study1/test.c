#include <stdio.h>
#include <limits.h>
#include <float.h>

int x=1;
int y=2;
int addtwonum() {
    return x+y;
};
int main() 
{
    /*first prom */
    printf("Hello world\n");
    printf("Storage size for int: %lu\n", sizeof(int));
    printf("float 最大值: %d\n", sizeof(float));
    printf("float %E\n",FLT_MAX);
    printf("精确度: %d\n",FLT_DIG);

    int a,b;
    int c;

    a = 10;
    b = 20;
    c = a + b;
    printf("value of a = %d, b = %d and c = %d\n",a,b,c);

    int result;
    result = addtwonum();
    printf("result: %d\n",result);
    return 0;
}
