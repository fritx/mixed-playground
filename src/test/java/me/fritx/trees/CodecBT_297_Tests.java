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
import org.junit.jupiter.params.ParameterizedTest;
import org.junitpioneer.jupiter.json.JsonSource;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

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
    void test(Integer[] treeData) throws JsonProcessingException {

        // 伪代码：我想 for循环遍历多个Codec类，比如 Codec1, Codec2 这里咋写？
        CodecBT_297 ser = new CodecBT_297();
        CodecBT_297 deser = new CodecBT_297();
        TreeNode root = TreeUtils.binaryTreeFromArray(treeData);
        TreeNode ans = deser.deserialize(ser.serialize(root));

        // // 使用 Arrays.stream() 和 map() 将 int[] 转换为 String[]
        // String[] numberStrings = Arrays.stream(treeData)
        // .mapToObj(String::valueOf)
        // .toArray(String[]::new);
        // // 使用 String.join() 将字符串数组用逗号连接
        // String treeStr = String.join(",", numberStrings);
        ObjectMapper mapper = new ObjectMapper();
        String treeStr = mapper.writeValueAsString(treeData);

        assertTrue(TreeUtils.isSameTree(ans, root),
                () -> treeStr + " should pass");
    }
}
