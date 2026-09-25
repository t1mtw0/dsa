#include <vector>

class FenwickTree
{
public:
    FenwickTree(int n);
    FenwickTree(std::vector<int> const &d);
    int sum(int r);
    int sum(int l, int r);
    void add(int idx, int delta);

private:
    std::vector<int> bit;
};