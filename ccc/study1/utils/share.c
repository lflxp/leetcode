#include <stdio.h>
#include <string.h>

// C 共用体
// 共用体是一种特殊的数据类型，允许您在相同的内存位置存储不同的数据类型。您可以定义一个带有多成员的共用体，但是任何时候只能有一个成员带有值。共用体提供了一种使用相同的内存位置的有效方式。
union Data {
    int i;
    float f;
    char str[20];
};

// 访问共用体成员
// 为了访问共用体的成员，我们使用成员访问运算符（.）。成员访问运算符是共用体变量名称和我们要访问的共用体成员之间的一个句号。您可以使用 union 关键字来定义共用体类型的变量。下面的实例演示了共用体的用法：
// 在这里，我们可以看到共用体的 i 和 f 成员的值有损坏，因为最后赋给变量的值占用了内存位置，这也是 str 成员能够完好输出的原因。
void unionTest() {
    union Data data;

    printf("Memory size occupied by data: %lu \n", sizeof(data));

    data.i = 10;
    data.f = 220.5;
    strcpy(data.str,"C Programming");

    printf("data.i : %d\n",data.i);
    printf("data.f : %f\n",data.f);
    printf("data.str : %s\n",data.str);
}

// 现在让我们再来看一个相同的实例，这次我们在同一时间只使用一个变量，这也演示了使用共用体的主要目的：
// 在这里，所有的成员都能完好输出，因为同一时间只用到一个成员
void unionTest1() {
    union Data data;
    data.i = 10;
    printf("data.i : %d\n",data.i);

    data.f = 220.5;
    printf("data.f : %f\n", data.f);

    strcpy(data.str,"C Hello World");
    printf("data.str : %s\n",data.str);
}