struct Node {
    int value;
    int rank;
    Node* parent;

    Node(int v): value{v}, rank{0}, parent{nullptr} {};
};

void make_set(Node *v);
Node *find_set(Node *v);
void union_sets(Node *a, Node *b);