#include "dsu.h"

void make_set(Node *v)
{
    v->parent = v;
    v->rank = 0;
}

Node *find_set(Node *v)
{
    if (v->parent == v)
        return v;
    return v->parent = find_set(v->parent);
}

void union_sets(Node *a, Node *b)
{
    a = find_set(a);
    b = find_set(b);
    if (a != b)
    {
        if (a->rank < b->rank)
        {
            auto temp = a;
            a = b;
            b = temp;
        }
        b->parent = a;
        if (b->rank == a->rank)
            a->rank++;
    }
}