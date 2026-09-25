#include <vector>
#include <iostream>
#include "fenwick.h"

FenwickTree::FenwickTree(int n)
{
    for (int i = 0; i < n; ++i)
        bit.push_back(0);
}

FenwickTree::FenwickTree(std::vector<int> const &d)
{
    for (int i = 0; i < d.size(); ++i)
        bit.push_back(0);
    for (int i = 0; i < d.size(); ++i)
        add(i, d.at(i));
}

int FenwickTree::sum(int r)
{
    int res = 0;
    for (; r >= 0; r = (r & (r + 1)) - 1) {
        res += bit.at(r);
    }
    return res;
}

int FenwickTree::sum(int l, int r)
{
    return sum(r) - sum(l - 1);
}

void FenwickTree::add(int idx, int delta)
{
    for (; idx < bit.size(); idx = idx | (idx + 1))
        bit[idx] += delta;
}