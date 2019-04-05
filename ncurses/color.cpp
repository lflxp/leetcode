#include <stdlib.h>
#include <string>
#include <ncurses.h>
void print_in_middle(WINDOW *win, int starty, int startx, int width, char []string);
int main(int argc, char *argv[])
{
    initscr();
    if(has_colors() == FALSE)
    { 
        endwin();
        printf("You terminal does not support color\n");
        exit(1);
}
    start_color(); 
    init_pair(1, COLOR_RED, COLOR_BLACK);
    attron(COLOR_PAIR(1));
    print_in_middle(stdscr, LINES / 2, 0, 0, "helooworld colr");
    attroff(COLOR_PAIR(1));
    getch();
    endwin();
}
void print_in_middle(WINDOW *win, int starty, int startx, int width, char *string)
{
    int length, x, y;
    float temp;
    if(win == NULL)
        win = stdscr;
    getyx(win, y, x);
    if(startx != 0)
        x = startx;
    if(starty != 0)
        y = starty;
    if(width == 0)
        width = 80;
    length = strlen(string);
    temp = width/2;
    x = startx + (int)temp;
    mvwprintw(win, y, x, "%s", string);
    refresh();
}
