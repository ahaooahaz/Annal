#include <iostream>
#include <new>
#include <math.h>

int main() {
    char *p = nullptr;
    unsigned long long int size = 2;
    while (size *= 2) {
        try{
            p = new char[size];
        }catch(std::bad_alloc){
            std::cout << "size: " << size << " throw bad_alloc" << std::endl;
            break;
        };

        std::cout << "size: " << size << " ok" << std::endl;
        delete[] p;
    }

    p = new (std::nothrow) char[size];
    if (p == nullptr) {
        std::cout << "with no_throw p == nullptr" << std::endl;
    }

    std::cout << "safe return" << std::endl;
    return 0;
}