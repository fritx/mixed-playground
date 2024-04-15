//go:build ignore

// 方法一：链地址法
class MyHashMap {
private:
    static const int base = 769;
    vector<list<pair<int, int>>> data;
    int hash(int key) {
        return key % base;
    }
public:
    MyHashMap(): data(base) {}
    void put(int key, int value) {
        int h = hash(base);
        for (auto it = data[h].begin(); it != data[h].end(); ++it) {
            if (it->first == key) {
                it->second = value;
                return;
            }
        }
        data[h].push_back(make_pair(key, value));
    }
    int get(int key) {
        int h = hash(base);
        for (auto it = data[h].begin(); it != data[h].end(); ++it) {
            if (it->first == key) {
                return it->second;
            }
        }
        return -1;
    }
    void remove(int key) {
        int h = hash(base);
        for (auto it = data[h].begin(); it != data[h].end(); ++it) {
            if (it->first == key) {
                data[h].erase(it);
                return;
            }
        }
    }
};

/**
 * Your MyHashMap object will be instantiated and called as such:
 * MyHashMap* obj = new MyHashMap();
 * obj->put(key,value);
 * int param_2 = obj->get(key);
 * obj->remove(key);
 */