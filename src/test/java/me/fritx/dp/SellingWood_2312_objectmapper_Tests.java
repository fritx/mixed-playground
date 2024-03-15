package me.fritx.dp;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.jupiter.api.extension.ExtensionContext;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.ArgumentsProvider;
import org.junit.jupiter.params.provider.ArgumentsSource;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.io.IOException;
import java.lang.reflect.InvocationTargetException;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class SellingWood_2312_objectmapper_Tests {

    @ParameterizedTest
    @ArgumentsSource(CustomJsonProvider.class)
    public void test(int m, int n, int[][] prices, int want)
            throws ClassNotFoundException, InstantiationException, IllegalAccessException, IllegalArgumentException,
            InvocationTargetException, NoSuchMethodException, SecurityException {

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

    static class CustomJsonProvider implements ArgumentsProvider {
        private static final String JSON = """
                [
                    [3, 5, [[1, 4, 2], [2, 2, 7], [2, 1, 3]], 19],
                    [4, 6, [[3, 2, 10], [1, 4, 2], [4, 1, 3]], 32],
                    [7, 7, [], 0],
                    [0, 0, [[1, 2, 3]], 0],
                    [0, 0, [], 0]
                    ]
                """;

        @Override
        public Stream<? extends Arguments> provideArguments(ExtensionContext context) {
            ObjectMapper mapper = new ObjectMapper();
            try {
                @SuppressWarnings("unchecked")
                List<List<Object>> testCases = mapper.readValue(JSON, List.class);
                return testCases.stream().map(testCase -> {
                    int m = (int) testCase.get(0);
                    int n = (int) testCase.get(1);

                    // int[][] prices = Arrays.stream((List<List<Integer>>) testCase.get(2))
                    // .map(list -> list.toArray(new Integer[0]))
                    // .toArray(int[][]::new);
                    @SuppressWarnings("unchecked")
                    List<List<Integer>> nestedLists = (List<List<Integer>>) testCase.get(2);
                    // 将 List<List<Integer>> 转换为 List<int[]>
                    List<int[]> intermediateList = nestedLists.stream()
                            .map(list -> list.stream().mapToInt(Integer::intValue).toArray())
                            .collect(Collectors.toList());
                    // 将 List<int[]> 转换为 int[][]
                    int[][] prices = intermediateList.toArray(int[][]::new);

                    int want = (int) testCase.get(3);
                    return Arguments.of(m, n, prices, want);
                });
            } catch (IOException e) {
                throw new RuntimeException(e);
            }
        }
    }
}
