#include <stdio.h>

int main()
{
    int arr[5] = {1, 2, 3, 4, 5};
    printf("%d\n", arr[0]);
    printf("arr=\t%p\narr[0]=\t%p\narr[1]=\t%p\n", &arr, &arr[0], &arr[1]);
    return 0;
}