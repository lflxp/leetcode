#include "utils/func.c"
#include "utils/hello.c"
#include "stringstudy/test1.c"
#include "utils/structs.c"
#include "utils/share.c"
#include "utils/position.c"
#include "utils/typedef.c"
#include "utils/malloc.c"
#include "utils/scanfs.c"
#include "pre/p1.c"

void backupfunc() {
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

    HelloWorld();
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

int main() 
{
    
    // read();
    // hanshustring();

    // printf("title: %s \n author: %s \nsubject: %s\n book_id: %d\n",book2.title,book2.author,book2.subject,book2.book_id);

    // structsFunc();
    // structFunc2();
    // useDomain();

    // unionTest();

    // unionTest1();

    // positionTest();

    // posTest();

    // testBooks();
    // testDefine();

    // memTest();

    // testScanfs();

    p1(1,1);
    return 0;
}

