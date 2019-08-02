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
// int max(int num1, int num2) 
// {
//    /* 局部变量声明 */
//    int result;
 
//    if (num1 > num2)
//       result = num1;
//    else
//       result = num2;
 
//    return result; 
// }

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
// C 从函数返回指针
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

int Basic() {
/*first prom */
    printf("Hello world\n");
    printf("Storage size for int: %lu\n", sizeof(int));
    printf("float 最大值: %d\n", sizeof(float));
    printf("float %E\n",FLT_MAX);
    printf("精确度: %d\n",FLT_DIG);

    int aa,bb;
    int cc;

    aa = 10;
    bb = 20;
    cc = aa + bb;
    printf("value of a = %d, b = %d and c = %d\n",aa,bb,cc);

    int result;
    result = addtwonum();
    printf("result: %d\n",result);
    printf("result %d\n", OK);

    const int ABC = 20;
    const int Son = 99;
    const char BCD = 'FGHJKL';

    printf("result %d  ---- %c\n", ABC * Son, BCD);

    while (count--)
    {
        /* code */
        func1();
    }
    printf("address %p\n",&result);

    /*打印指针（地址）的值*/
    int i = 0;
    int *p = &i;

    printf("指针（地址）的值为：OX%p\n",p);
    printf("变量的值为：%d\n",i);

    unsigned int a = 60;    /* 60 = 0011 1100 */  
    unsigned int b = 13;    /* 13 = 0000 1101 */
    int c = 0;           
    
    c = a & b;       /* 12 = 0000 1100 */ 
    printf("Line 1 - c 的值是 %d\n", c );
    
    c = a | b;       /* 61 = 0011 1101 */
    printf("Line 2 - c 的值是 %d\n", c );
    
    c = a ^ b;       /* 49 = 0011 0001 */
    printf("Line 3 - c 的值是 %d\n", c );
    
    c = ~a;          /*-61 = 1100 0011 */
    printf("Line 4 - c 的值是 %d\n", c );
    
    c = a << 2;     /* 240 = 1111 0000 */
    printf("Line 5 - c 的值是 %d\n", c );
    
    c = a >> 2;     /* 15 = 0000 1111 */
    printf("Line 6 - c 的值是 %d\n", c );

    a = 21;
    c ;
 
    c =  a;
    printf("Line 1 - =  运算符实例，c 的值 = %d\n", c );
    
    c +=  a;
    printf("Line 2 - += 运算符实例，c 的值 = %d\n", c );
    
    c -=  a;
    printf("Line 3 - -= 运算符实例，c 的值 = %d\n", c );
    
    c *=  a;
    printf("Line 4 - *= 运算符实例，c 的值 = %d\n", c );
    
    c /=  a;
    printf("Line 5 - /= 运算符实例，c 的值 = %d\n", c );
    
    c  = 200;
    c %=  a;
    printf("Line 6 - %= 运算符实例，c 的值 = %d\n", c );
    
    c <<=  2;
    printf("Line 7 - <<= 运算符实例，c 的值 = %d\n", c );
    
    c >>=  2;
    printf("Line 8 - >>= 运算符实例，c 的值 = %d\n", c );
    
    c &=  2;
    printf("Line 9 - &= 运算符实例，c 的值 = %d\n", c );
    
    c ^=  2;
    printf("Line 10 - ^= 运算符实例，c 的值 = %d\n", c );
    
    c |=  2;
    printf("Line 11 - |= 运算符实例，c 的值 = %d\n", c );

    int cbd;

    cbd = (1== 2) ? 20:30;
    printf("cbd is %d\n",cbd);

    a = 4;
    short ba;
    double ca;
    int* ptr;
    
    /* sizeof 运算符实例 */
    printf("Line 1 - 变量 a 的大小 = %lu\n", sizeof(a) );
    printf("Line 2 - 变量 b 的大小 = %lu\n", sizeof(ba) );
    printf("Line 3 - 变量 c 的大小 = %lu\n", sizeof(ca) );
    
    /* & 和 * 运算符实例 */
    ptr = &a;    /* 'ptr' 现在包含 'a' 的地址 */
    printf("a 的值是 %d\n", a);
    printf("*ptr 是 %d\n", *ptr);
    
    /* 三元运算符实例 */
    a = 10;
    ba = (a == 1) ? 20: 30;
    printf( "b 的值是 %d\n", ba );
    
    ba = (a == 10) ? 20: 30;
    printf( "b 的值是 %d\n", ba );
}

void isDouble(int a) {
    (a%2 ==0)?printf("偶数\n"):printf("奇数\n");
}

void funcGoto() {
    int a= 10;

    LOOP: do {
        if (a == 15) {
            a += 1;
            goto LOOP;
        }
        printf("a的值： %d\n",a);
        a++;
    }while(a< 20);
}

// 指针地址跨字节存储int * pointer（4）类型数据
// int 类型默认四个字节 所以内存地址加1 按int类型右移4个位置而不是一个位置 只有char字符类型为初始值进行指针运算时才会按一个字节进行移动
void pointers() {
    // unsigned char c = '257';
    unsigned int c = 257;
    unsigned int* a = &c;
    unsigned char* b = (unsigned char*)&c;
    unsigned char* b1 = (unsigned char*)&c;
    unsigned char* b2 = (unsigned char*)(&c+1);
    unsigned char* b3 = (unsigned char*)(&c+2);
    unsigned char* b4 = (unsigned char*)(&c+3);

    printf("%d\n", *a);
    printf("%d %p\n", *b, a);
    printf("%p %p %p %p\n",&c,(int*)(&c+1),(&c+2),(&c+3));
    printf("%d %d %d %d\n",*b1,*b2,*b3,*b4);
    printf("%d\n",c >> 1);

    // ptr++
    printf("%p %p %p\n",a,a++,a++);
    printf("%p %p %p %p\n",b1,b1++,b1++,b1++);
}

//C 指向数组的指针
/*double *p;
double balance[10];

p = balance;
使用数组名作为常量指针是合法的，反之亦然。因此，*(balance + 4) 是一种访问 balance[4] 数据的合法方式。

一旦您把第一个元素的地址存储在 p 中，您就可以使用 *p、*(p+1)、*(p+2) 等来访问数组元素。下面的实例演示了上面讨论到的这些概念： */
void arrayPointers() {
    double balance[5] = {5.88,4.6,3.2,2.2,1.1};
    double *p;
    int i;

    p = balance;
    /*输出数组中每个元素的值 */
    printf("使用指针的数组值\n");
    for (i=0;i<5;i++) {
        printf("*(p+%d) : %f\n",i,*(p + i) );
    }

    printf("使用balance作为地址的数组值\n");
    for (i=0;i<5;i++) {
        printf("*(balance + %d) : %f\n",i,*(balance + i));
    }
}

#define MON  1
#define TUE  2
#define WED  3
#define THU  4
#define FRI  5
#define SAT  6
#define SUN  7

enum DAY {
    Mon = 1,
    Tue,
    Wed,
    Thu,
    Fri,
    Sat,
    Sun
};

enum season {spring, summer=3,autumn,winter};

void enumPrintf() {
    enum DAY day;
    day = Wed;
    printf("%d\n",day);

    enum season spr;
    spr = autumn;
    printf("%d\n", spr);
}

// 一定条件下遍历枚举元素
void enumUnit() {
    enum DAY day;
    // 遍历枚举元素
    for (day = Mon; day <= Sun; day ++) {
        printf("枚举元素: %d\n",day);
    }
}

void enumSwitch() {
    enum color {red=1,green,blue};

    enum color favorite_color;

    /*ask user to choose color */
    printf("请输入你喜欢的颜色:(1.red 2.green 3.blue):");
    scanf("%d",&favorite_color);

    /*输出结果 */
    switch (favorite_color)
    {
    case red:
        printf("你喜欢的颜色是红色\n");
        break;
    case green:
        printf("你喜欢的颜色是绿色\n");
        break;
    case blue:
        printf("你喜欢的颜色是蓝色\n");
        break;
    default:
        printf("你没有选择你喜欢的颜色");
    }
}

// 将整数转换为枚举
void enumToInt() {
    enum day {
        s1,
        s2,
        m1,
        t1,
        w1,
        t2,
        f1
    } workday;

    int a = 1;
    enum day weekend;
    weekend = (enum day) a; // 类型转换
    printf("weekend: %d\n",weekend);
} 

// 打印指针
// https://www.runoob.com/cprogramming/c-pointers.html
void printPtr() {
    int var1;
    char var2[10];

    printf("var1 变量的地址: %p\n",&var1);
    printf("var2 变量的地址: %p\n",&var2);
}

void ptrTest() {
    int var = 20; /*实际变量的声明 */
    int *ip;

    ip = &var; /*在指针变量中存储var的地址 */

    printf("Address of var variable: %p\n",&var);

    /*在指针变量中存储的地址 */
    printf("Address stored in ip variable: %p\n",ip);

    /*使用指针访问值 */
    printf("Value of *ip variable: %d\n",ip);
}

// 空指针
void NullPtr() {
    int *ptr = NULL;

    printf("ptr 的地址是 %p\n", ptr);
}

const int MAX = 3;
// 指针的比较
void comparePtr() {
    int var[] = {10,100,200};
    int i,*ptr;

    /*指针中第一个元素的地址 */
    ptr = var;
    i = 0;

    printf("ptr address %d\n",*ptr);
    // 数组第一个首地址与数组最后一个指针地址的值进行比较
    while (ptr <= &var[MAX-1])
    {
        printf("Address of var[%d] = %x\n",i,ptr);
        printf("Value of var[%d] = %d\n",i,*ptr);

        /*指向上一个位置 */
        ptr++;
        i++;  
    }
}

// C 传递指针给函数
// https://www.runoob.com/cprogramming/c-passing-pointers-to-functions.html
void getSeconds(unsigned long *par) {
    /*获取当前的秒数 */
    *par = time(NULL);
    return;
}

/*函数声明 */
double getAverage(int *arr,int size);

void testAverage() {
    /*带有5个元素的整形数组 */
    int balance[5] = {1000,2,5,123,60};
    double avg;

    /*传递一个指向数组的指针作为参数 */
    avg = getAverage(balance,5);

    /*输出返回值 */
    printf("Average value is: %f\n",avg);
}

int max(int x,int y) {
    return x>y?x:y;
}

/*
函数指针
函数指针是指向函数的指针变量。

通常我们说的指针变量是指向一个整型、字符型或数组等变量，而函数指针是指向函数。

函数指针可以像一般函数一样，用于调用函数、传递参数。

函数指针变量的声明：

typedef int (*fun_ptr)(int,int); // 声明一个指向同样参数、返回值的函数指针类型
 */
void testPtrFunc() {
    /*p是函数指针 */
    int (* p)(int,int) = &max; // &可以省略
    int a,b,c,d;

    printf("请输入三个数字：");
    scanf("%d %d %d",&a,&b,&c);

    /*与直接调用函数等价，d=max(max(a,b),c) */
    d = p(p(a,b),c);

    printf("最大的数字是：%d\n",d);
}

/*
回调函数
函数指针作为某个函数的参数
函数指针变量可以作为某个函数的参数来使用的，回调函数就是一个通过函数指针调用的函数。

简单讲：回调函数是由别人的函数执行时调用你实现的函数。

以下是来自知乎作者常溪玲的解说：

你到一个商店买东西，刚好你要的东西没有货，于是你在店员那里留下了你的电话，过了几天店里有货了，店员就打了你的电话，然后你接到电话后就到店里去取了货。在这个例子里，你的电话号码就叫回调函数，你把电话留给店员就叫登记回调函数，店里后来有货了叫做触发了回调关联的事件，店员给你打电话叫做调用回调函数，你到店里去取货叫做响应回调事件。
 */
// 回调函数
void populate_array(int *array,size_t arraySize,int (*genNextValue)(void)) {
    for (size_t i = 0;i<arraySize;i++)
        array[i] = genNextValue();
}

// 获取随机值
int getNextRandomValue(void) {
    return rand();
}

// 测试回调函数
void testInnerFunc() {
    int myarray[10];
    populate_array(myarray, 10,getNextRandomValue);
    for (int i=0;i<10;i++) {
        printf("%d ",myarray[i]);
    }
    printf("\n");
}

int main() 
{
    // arrayDouble();
    // int one = 100;
    // int two = 200;
    // int ret;

    // ret = max(one,two);

    // printf("Max value is : %d\n", ret);

    // dowhile();
    
    // // funcday();
    // funcdo();
    // isZhiShu();
    // whileInstance();

    /*一个指向整数的指针 */
    // int *p;
    // int i;

    // p = genRandom();
    // for (i=0;i<10;i++) {
    //     printf("*(p + %d) : %d\n", i, *(p + i));
    // }

    // // 判断
    // int num;
    // printf("输入一个数字：");
    // scanf("%d",&num);

    // isDouble(num);

    // funcGoto();

    // pointers();

    // arrayPointers();

    // enumPrintf();

    // enumUnit();

    // enumSwitch();

    // enumToInt();

    // printPtr();

    // ptrTest();

    // NullPtr();

    // comparePtr();

    // unsigned long sec;
    // getSeconds(&sec);
    // /*输出实际值 */
    // printf("Number of seconds: %ld\n",sec);

    // testAverage();

    // testPtrFunc();

    testInnerFunc();

    return 0;
}

// 遍历数组指针首地址对应的常量数组值
double getAverage(int *arr,int size) {
    int i,sum = 0;
    double avg;

    for (i=0; i<size;++i) {
        sum += arr[i];
    }

    avg = (double)sum/size;

    return avg;
}