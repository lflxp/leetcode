#include "stdio.h"
#include "string.h"

/*定义简单的结构 */
struct 
{
    unsigned int widthValidated;
    unsigned int heightValidated;
} status1;

/*定义位域结构 */
struct {
    unsigned int widthValidated:1;
    unsigned int heightValidated:1;
} status2;

// 位域测试
void positionTest() {
    printf("Memory size occupied by status1: %lu\n",sizeof(status1));
    printf("Memory size occupied by status1: %lu\n",sizeof(status2));
}

struct {
    unsigned int age:3;
} Age;

// 上面的结构定义指示 C 编译器，age 变量将只使用 3 位来存储这个值，如果您试图使用超过 3 位，则无法完成。让我们来看下面的实例
void posTest() {
    Age.age = 4;
    printf( "Sizeof( Age ) : %lu\n", sizeof(Age) );
   printf( "Age.age : %d\n", Age.age );
 
   Age.age = 7;
   printf( "Age.age : %d\n", Age.age );
 
   Age.age = 8; // 二进制表示为 1000 有四位，超出
   printf( "Age.age : %d\n", Age.age );
}
