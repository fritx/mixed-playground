package me.fritx.dp;

import org.junit.jupiter.params.ParameterizedTest;
import org.junitpioneer.jupiter.json.JsonSource;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.lang.reflect.InvocationTargetException;
import java.util.ArrayList;

public class SellingWood_2312_objectarray_Tests {

    @ParameterizedTest
    @JsonSource("""
            [
                [3, 5, [[1, 4, 2], [2, 2, 7], [2, 1, 3]], 19],
                [4, 6, [[3, 2, 10], [1, 4, 2], [4, 1, 3]], 32],
                [7, 7, [], 0],
                [0, 0, [[1, 2, 3]], 0],
                [0, 0, [], 0]
            ]
            """)
    public void test(Object[] input)
            throws ClassNotFoundException, NoSuchMethodException, InstantiationException, IllegalAccessException,
            InvocationTargetException {

        int m = (int) input[0];
        int n = (int) input[1];
        int want = (int) input[3];

        // 报错: ClassCast class java.util.ArrayList cannot be cast to class [[I
        // (java.util.ArrayList and [[I are in module java.base of loader 'bootstrap')
        // int[][] prices = (int[][]) input[2];
        @SuppressWarnings("unchecked")
        ArrayList<ArrayList<Integer>> listOfLists = (ArrayList<ArrayList<Integer>>) input[2];

        // 转换为 int[][]
        int[][] prices = new int[listOfLists.size()][];
        for (int i = 0; i < listOfLists.size(); i++) {
            ArrayList<Integer> innerList = listOfLists.get(i);
            prices[i] = new int[innerList.size()];
            for (int j = 0; j < innerList.size(); j++) {
                prices[i][j] = innerList.get(j);
            }
        }

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
