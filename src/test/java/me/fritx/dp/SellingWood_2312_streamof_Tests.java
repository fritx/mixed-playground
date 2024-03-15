/*
 * Copyright 2015-2023 the original author or authors.
 *
 * All rights reserved. This program and the accompanying materials are
 * made available under the terms of the Eclipse Public License v2.0 which
 * accompanies this distribution and is available at
 *
 * http://www.eclipse.org/legal/epl-v20.html
 */

package me.fritx.dp;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.lang.reflect.InvocationTargetException;
import java.util.stream.Stream;

import org.junit.jupiter.api.extension.ExtensionContext;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.ArgumentsProvider;
import org.junit.jupiter.params.provider.ArgumentsSource;

import com.fasterxml.jackson.core.JsonProcessingException;

class SellingWood_2312_streamof_Tests {
    @ParameterizedTest
    @ArgumentsSource(CustomJsonProvider.class)
    void test(int m, int n, int[][] prices, int want)
            throws JsonProcessingException, ClassNotFoundException, InstantiationException, IllegalAccessException,
            IllegalArgumentException, InvocationTargetException, NoSuchMethodException, SecurityException {

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
        @Override
        public Stream<? extends Arguments> provideArguments(ExtensionContext context) {
            // 在这里解析你的JSON数组，并返回一个Arguments的流
            // 你可能需要使用一个JSON库，如Jackson或Gson，来帮助你解析JSON
            return Stream.of(
                    // 这里是一个例子，你需要替换为你的实际代码
                    Arguments.of(3, 5, new int[][] { { 1, 4, 2 }, { 2, 2, 7 }, { 2, 1, 3 } }, 19),
                    Arguments.of(4, 6, new int[][] { { 3, 2, 10 }, { 1, 4, 2 }, { 4, 1, 3 } }, 32),
                    Arguments.of(7, 7, new int[][] {}, 0),
                    Arguments.of(0, 0, new int[][] { { 1, 2, 3 } }, 0),
                    Arguments.of(0, 0, new int[][] {}, 0));
        }
    }
}
