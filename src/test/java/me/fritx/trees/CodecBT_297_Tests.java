/*
 * Copyright 2015-2023 the original author or authors.
 *
 * All rights reserved. This program and the accompanying materials are
 * made available under the terms of the Eclipse Public License v2.0 which
 * accompanies this distribution and is available at
 *
 * http://www.eclipse.org/legal/epl-v20.html
 */

package me.fritx.trees;

import static org.junit.jupiter.api.Assertions.assertTrue;

import java.lang.reflect.InvocationTargetException;

import org.junit.jupiter.params.ParameterizedTest;
import org.junitpioneer.jupiter.json.JsonSource;

import com.fasterxml.jackson.core.JsonProcessingException;
// import com.fasterxml.jackson.databind.ObjectMapper;

class CodecBT_297_Tests {

    // @ParameterizedTest(name = "{0} + {1} = {2}")
    // @ParameterizedTest(name = "treeData={0}")
    @ParameterizedTest
    // The Java feature 'Text Blocks' is only available with source level 15 and
    // above
    @JsonSource("""
            [
                [1,2,3,null,null,4,5],
                [1,null,2],
                [1,2],
                []
            ]
            """)
    // void test(int[] treeData) throws JsonProcessingException {
    void test(Integer[] treeData)
            throws JsonProcessingException, ClassNotFoundException, InstantiationException, IllegalAccessException,
            IllegalArgumentException, InvocationTargetException, NoSuchMethodException, SecurityException {

        String[] classNames = {
                CodecBT_297_1.class.getName(),
                CodecBT_297_2.class.getName(),
        };
        for (String className : classNames) {
            Class<?> codecClass = Class.forName(className);
            // fix: Type mismatch: cannot convert from capture#2-of ? to CodecBT_297
            CodecBT_297 ser = (CodecBT_297) codecClass.getDeclaredConstructor().newInstance();
            CodecBT_297 deser = (CodecBT_297) codecClass.getDeclaredConstructor().newInstance();

            TreeNode root = TreeUtils.binaryTreeFromArray(treeData);
            TreeNode ans = deser.deserialize(ser.serialize(root));

            // // 使用 Arrays.stream() 和 map() 将 int[] 转换为 String[]
            // String[] numberStrings = Arrays.stream(treeData)
            // .mapToObj(String::valueOf)
            // .toArray(String[]::new);
            // // 使用 String.join() 将字符串数组用逗号连接
            // String treeStr = String.join(",", numberStrings);
            // ObjectMapper mapper = new ObjectMapper();
            // String treeStr = mapper.writeValueAsString(treeData);
            // assertTrue(TreeUtils.isSameTree(ans, root),
            // () -> treeStr + " should pass");
            assertTrue(TreeUtils.isSameTree(ans, root));
        }

    }
}
