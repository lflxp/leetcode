#include <stdio.h>
#include <limits.h>
#include <float.h>
#include <time.h>
#include <stdlib.h>

int x=1;
int y=2;
int addtwonum() {
    return x+y;
};

#define OK 6
static int count = 10;

/* 函数声明 */
void func1(void);

void func1(void) {
    static int things = 5;
    things++;
    printf("things %d count %d\n",things, count);
}


void funcday() {
    auto int a;
    printf("input integer number: ");
    scanf("%d",&a);
    switch(a)
    {
        case 1:printf("Monday\n");
        break;
        case 2:printf("Tuesday\n");
        break;
        case 3:printf("Wednesday\n");
        break;
        case 4:printf("Thursday\n");
        break;
        case 5:printf("Friday\n");
        break;
        case 6:printf("Saturday\n");
        break;
        case 7:printf("Sunday\n");
        break;
        default:printf("error\n");
    }
}

void isZhiShu() {
    /* 局部变量定义 */
   int i, j;
   
   for(i=2; i<100; i++) {
      for(j=2; j <= (i/j); j++)
        if(!(i%j)) break; // 如果找到，则不是质数
      if(j > (i/j)) printf("%d 是质数\n", i);
   }
}

void whileInstance() {
    int i =1,j;
    while (i <= 5) {
        j = 1;
        while (j<=i) {
            printf("%d ", j);
            j++;
        }
        printf("\n");
        i++;
    }
}

void dowhile() {
    int i=1,j;
    do {
        j = 1;
        do {
            printf("*");
            j++;
        }while (j<=i);
        i++;
        printf("\n");
    }while (i<=5);    
}

void funcdo() {
    auto int a = 10;
    do {
        printf("a 的值为: %d\n", a);
        a += 1;
    }while ( a< 20);
}

/* 函数返回两个数中较大的那个数 */
int max(int num1, int num2) 
{
   /* 局部变量声明 */
   int result;
 
   if (num1 > num2)
      result = num1;
   else
      result = num2;
 
   return result; 
}

void arrayDouble() {
    int n[10];
    int i,j;

    /* 初始化数组元素 */         
    for ( i = 0; i < 10; i++ )
    {
        n[ i ] = i + 100; /* 设置元素 i 为 i + 100 */
    }
    
    /* 输出数组中每个元素的值 */
    for (j = 0; j < 10; j++ )
    {
        printf("Element[%d] = %d\n", j, n[j] );
    }
}

/*要生成和返回随机数的函数 */
int * genRandom() {
    static int r[10];
    int i;

    /*设置种子 */
    srand((unsigned)time(NULL));
    for (i=0;i<10;++i) {
        r[i] = rand();
        printf("r[%d] = %d\n", i, r[i]);
    }

    return r;
}

int main() 
{
    arrayDouble();
    int one = 100;
    int two = 200;
    int ret;

    ret = max(one,two);

    printf("Max value is : %d\n", ret);

    dowhile();
    // /*first prom */
    // printf("Hello world\n");
    // printf("Storage size for int: %lu\n", sizeof(int));
    // printf("float 最大值: %d\n", sizeof(float));
    // printf("float %E\n",FLT_MAX);
    // printf("精确度: %d\n",FLT_DIG);

    // int aa,bb;
    // int cc;

    // aa = 10;
    // bb = 20;
    // cc = aa + bb;
    // printf("value of a = %d, b = %d and c = %d\n",aa,bb,cc);

    // int result;
    // result = addtwonum();
    // printf("result: %d\n",result);
    // printf("result %d\n", OK);

    // const int ABC = 20;
    // const int Son = 99;
    // const char BCD = 'FGHJKL';

    // printf("result %d  ---- %c\n", ABC * Son, BCD);

    // while (count--)
    // {
    //     /* code */
    //     func1();
    // }
    // printf("address %p\n",&result);

    // /*打印指针（地址）的值*/
    // int i = 0;
    // int *p = &i;

    // printf("指针（地址）的值为：OX%p\n",p);
    // printf("变量的值为：%d\n",i);

    // unsigned int a = 60;    /* 60 = 0011 1100 */  
    // unsigned int b = 13;    /* 13 = 0000 1101 */
    // int c = 0;           
    
    // c = a & b;       /* 12 = 0000 1100 */ 
    // printf("Line 1 - c 的值是 %d\n", c );
    
    // c = a | b;       /* 61 = 0011 1101 */
    // printf("Line 2 - c 的值是 %d\n", c );
    
    // c = a ^ b;       /* 49 = 0011 0001 */
    // printf("Line 3 - c 的值是 %d\n", c );
    
    // c = ~a;          /*-61 = 1100 0011 */
    // printf("Line 4 - c 的值是 %d\n", c );
    
    // c = a << 2;     /* 240 = 1111 0000 */
    // printf("Line 5 - c 的值是 %d\n", c );
    
    // c = a >> 2;     /* 15 = 0000 1111 */
    // printf("Line 6 - c 的值是 %d\n", c );

    // a = 21;
    // c ;
 
    // c =  a;
    // printf("Line 1 - =  运算符实例，c 的值 = %d\n", c );
    
    // c +=  a;
    // printf("Line 2 - += 运算符实例，c 的值 = %d\n", c );
    
    // c -=  a;
    // printf("Line 3 - -= 运算符实例，c 的值 = %d\n", c );
    
    // c *=  a;
    // printf("Line 4 - *= 运算符实例，c 的值 = %d\n", c );
    
    // c /=  a;
    // printf("Line 5 - /= 运算符实例，c 的值 = %d\n", c );
    
    // c  = 200;
    // c %=  a;
    // printf("Line 6 - %= 运算符实例，c 的值 = %d\n", c );
    
    // c <<=  2;
    // printf("Line 7 - <<= 运算符实例，c 的值 = %d\n", c );
    
    // c >>=  2;
    // printf("Line 8 - >>= 运算符实例，c 的值 = %d\n", c );
    
    // c &=  2;
    // printf("Line 9 - &= 运算符实例，c 的值 = %d\n", c );
    
    // c ^=  2;
    // printf("Line 10 - ^= 运算符实例，c 的值 = %d\n", c );
    
    // c |=  2;
    // printf("Line 11 - |= 运算符实例，c 的值 = %d\n", c );

    // int cbd;

    // cbd = (1== 2) ? 20:30;
    // printf("cbd is %d\n",cbd);

    // a = 4;
    // short ba;
    // double ca;
    // int* ptr;
    
    // /* sizeof 运算符实例 */
    // printf("Line 1 - 变量 a 的大小 = %lu\n", sizeof(a) );
    // printf("Line 2 - 变量 b 的大小 = %lu\n", sizeof(ba) );
    // printf("Line 3 - 变量 c 的大小 = %lu\n", sizeof(ca) );
    
    // /* & 和 * 运算符实例 */
    // ptr = &a;    /* 'ptr' 现在包含 'a' 的地址 */
    // printf("a 的值是 %d\n", a);
    // printf("*ptr 是 %d\n", *ptr);
    
    // /* 三元运算符实例 */
    // a = 10;
    // ba = (a == 1) ? 20: 30;
    // printf( "b 的值是 %d\n", ba );
    
    // ba = (a == 10) ? 20: 30;
    // printf( "b 的值是 %d\n", ba );
    // funcday();
    funcdo();
    isZhiShu();
    whileInstance();

    /*一个指向整数的指针 */
    int *p;
    int i;

    p = genRandom();
    for (i=0;i<10;i++) {
        printf("*(p + %d) : %d\n", i, *(p + i));
    }
    return 0;
}
