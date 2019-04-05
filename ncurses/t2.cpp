#include <stdlib.h>
#include <ncurses.h>

/* https://blog.csdn.net/Zhanganliu/article/details/79938102 */
int main(int argc, char *argv[])
{
    int ch, prev;
    FILE *fp;
    int goto_prev = FALSE, y, x;
    if(argc != 2)
    {
    printf("Usage: %s <a c file name>\n", argv[0]);
    exit(1);
    }
    fp = fopen(argv[1], "r"); /* 在这里检测文件是否成功打开*/
    if(fp == NULL)
    {
    perror("Cannot open input file");
    exit(1);
    }
    initscr(); /* 初始化并进入curses 模式*/
    prev = EOF;
    while((ch = fgetc(fp)) != EOF)
    {
    if(prev == '/' && ch == '*') /* 当读到字符“/”和“*”的时候开启修饰*/
    {
        attron(A_BOLD); /* 将“/”和“*”及以后输出的文字字体加粗*/
        goto_prev = TRUE;
    }
    if(goto_prev == TRUE) /* 回到“/”和“*”之前开始输出*/
    {
        getyx(stdscr, y, x);
        move(y, x);
        printw("%c%c", '/', ch);/* 打印实际内容的部分*/
        ch = 'a'; /* 避免下次读取变量错误，这里赋一个任意值*/
        goto_prev = FALSE; /* 让这段程序只运行一次*/
    }
    else printw("%c", ch);
        refresh(); /* 将缓冲区的内容刷新到屏幕上*/
 
    if(prev == '*' && ch == '/')
        attroff(A_BOLD); /* 当读到字符“*”和“/”的时候关闭修饰*/
    prev = ch;
}
getch();
endwin(); /* 结束并退出curses 模式*/
return 0;
}