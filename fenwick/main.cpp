#include <iostream>
#include "fenwick.h"

int main() {
    std::vector<int> d{1, 5, 8, 2, 4, 8, 3, 8, 12};
    FenwickTree fw{d};
    std::cout << fw.sum(2) << "\n";
    std::cout << fw.sum(4) << "\n";
    fw.add(0, 1);
    std::cout << fw.sum(2) << "\n";
    std::cout << fw.sum(4) << "\n";
    return 0;
};