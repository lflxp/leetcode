#include <stdio.h>
#include <string.h>

// https://www.runoob.com/cprogramming/c-strings.html
void read() {
    char greeting[6] = {'H','e','l','l','o','\0'};

    printf("Greeting message: %s\n",greeting);
}

// 字符串函数
void hanshustring() {
    char str1[12] = "Hello";
    char str2[12] = "World";
    char str3[12];
    int len;

    char *ptr;

    /*复制str1到str3 */
    strcpy(str3,str1);
    printf("strcpy(str3,str1): %s\n", str3);

    /*连接str1和str2 */
    strcat(str1,str2);
    printf("strcat(str1,str2): %s\n",str1);

    /*连接后，str1的总长度 */
    len = strlen(str1);
    printf("strlen(str1):%d\n",len);

    // 读取指针
    ptr = strchr(str1,'lo');
    printf("ptr %c %s %p %p\n",str1,ptr,&str1,&ptr);
}