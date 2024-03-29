//go:build ignore

char* maximumOddBinaryNumber(char* s){
    int cnt = 0, len = strlen(s);
    for (int i = 0; s[i] != '\0'; i++) {
        if (s[i] == '1') {
            cnt++;
        }
    }
    char *res = (char *)malloc(sizeof(char) * (len + 1));
    for (int pos = 0; pos < len - 1; pos++) {
        if (pos < cnt - 1) {
            res[pos] = '1';
        } else {
            res[pos] = '0';
        }
    }
    res[len - 1] = '1';
    res[len] = '\0';
    return res;
}
