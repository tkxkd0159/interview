#include <stdio.h>

int doubleNum(int a)
{
    return 2 * a;
}

int square(int a)
{
    return a * a;
}

int calculation(int a, int b)
{
    // square(3) = 9
    // double(5) = 10
    // return 90;
    return square(a) * doubleNum(b);
}

int main()
{
    int v = calculation(3, 5);
    printf("%d\n", v);
    return v;
}
