//go:build ignore

// 方法一：链地址法
class MyHashSet {
private:
    static const int base = 769;
    vector<list<int>> data;
    int hash(int key) {
        return key % base;
    }
public:
    MyHashSet(): data(base) {}
    void add(int key) {
        if (contains(key)) return;
        int h = hash(key);
        data[h].push_back(key);
    }
    void remove(int key) {
        int h = hash(key);
        for (auto it = data[h].begin(); it != data[h].end(); ++it) {
            if ((*it) == key) {
                data[h].erase(it);
                return;
            }
        }
    }
    bool contains(int key) {
        int h = hash(key);
        for (auto it = data[h].begin(); it != data[h].end(); ++it) {
            if ((*it) == key) {
                return true;
            }
        }
        return false;
    }
};

/**
 * Your MyHashSet object will be instantiated and called as such:
 * MyHashSet* obj = new MyHashSet();
 * obj->add(key);
 * obj->remove(key);
 * bool param_3 = obj->contains(key);
 */