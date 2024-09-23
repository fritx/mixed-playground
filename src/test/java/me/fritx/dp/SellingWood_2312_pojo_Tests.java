package me.fritx.dp;

import org.junit.jupiter.params.ParameterizedTest;
import org.junitpioneer.jupiter.json.JsonSource;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.lang.reflect.InvocationTargetException;

public class SellingWood_2312_pojo_Tests {

    @ParameterizedTest
    @JsonSource("""
            [
                {"m": 3, "n": 5, "prices": [[1, 4, 2], [2, 2, 7], [2, 1, 3]], "want": 19},
                {"m": 4, "n": 6, "prices": [[3, 2, 10], [1, 4, 2], [4, 1, 3]], "want": 32},
                {"m": 7, "n": 7, "prices": [], "want": 0},
                {"m": 0, "n": 0, "prices": [[1, 2, 3]], "want": 0},
                {"m": 0, "n": 0, "prices": [], "want": 0}
            ]
            """)
    public void test(TestInput input)
            throws ClassNotFoundException, NoSuchMethodException, InstantiationException, IllegalAccessException,
            InvocationTargetException {

        int m = input.getM();
        int n = input.getN();
        int[][] prices = input.getPrices();
        int want = input.getWant();

        String[] classNames = {
                SellingWood_2312.class.getName(),
        };
        for (String className : classNames) {
            Class<?> Clz = Class.forName(className);
            // fix: Type mismatch: cannot convert from capture#2-of ? to CodecBT_297
            SellingWood_2312 solution = (SellingWood_2312) Clz.getDeclaredConstructor().newInstance();
            long ans = solution.sellingWood(m, n, prices);
            assertEquals(ans, want);
        }
    }
}

class TestInput {
    private int m;
    private int n;
    private int[][] prices;
    private int want;

    // Getters and Setters
    public int getM() {
        return m;
    }

    public void setM(int m) {
        this.m = m;
    }

    public int getN() {
        return n;
    }

    public void setN(int n) {
        this.n = n;
    }

    public int[][] getPrices() {
        return prices;
    }

    public void setPrices(int[][] prices) {
        this.prices = prices;
    }

    public int getWant() {
        return want;
    }

    public void setWant(int want) {
        this.want = want;
    }
}
