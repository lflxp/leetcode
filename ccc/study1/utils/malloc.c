#include "stdio.h"
#include "stdlib.h"
#include "string.h"

void memTest() {
    char name[100];
    char *description;

    strcpy(name,"Zara Ali");

    /*动态分配内存 */
    description = (char *)malloc(40 * sizeof(char));
    if (description == NULL) {
        fprintf(stderr,"Error - unable to allocate required memory\n");
    } else {
        strcpy(description,"Zara ali a GPS student in class 10th");
    }
    printf("Name = %s\n", name);
    printf("Description : %s\n", description);

    // strcat(description,"1234567890987654321");
    /*假设你想要存储更大的描述信息 */
    description = (char *) realloc(description, 100*sizeof(char));
    if (description == NULL) {
        fprintf(stderr,"Error - unable to allocate required memory\n");
    } else {
        strcat(description,"She is in class 10th");
    }
    printf("Name = %s\n", name);
    printf("Description : %s\n", description);

    /*使用free()函数释放内存 */
    free(description);
}