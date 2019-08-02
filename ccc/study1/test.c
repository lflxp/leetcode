#include "func.c"
#include "utils/hello.c"

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

    HelloWorld();

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