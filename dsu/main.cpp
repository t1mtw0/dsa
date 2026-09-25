#include "dsu.h"
#include <iostream>

int main() {
    std::vector<Node*> nodes;
    for (int i = 0; i < 10; ++i) {
        Node *n = new Node(i);
        n->value = i;
        nodes.push_back(n);
    }
    for (int i = 0; i < 10; ++i) {
        make_set(nodes.at(i));
    }

    Node *n;
    union_sets(nodes.at(0), nodes.at(1));
    n = find_set(nodes.at(0));
    std::cout << n->value << "\n";
    n = find_set(nodes.at(1));
    std::cout << n->value << "\n";

    union_sets(nodes.at(0), nodes.at(2));
    n = find_set(nodes.at(2));
    std::cout << n->value << "\n";

    std::cout << nodes.at(0)->rank << "\n";
    std::cout << nodes.at(1)->rank << "\n";
    std::cout << nodes.at(2)->rank << "\n";
};