/*
 * Copyright 2015-2023 the original author or authors.
 *
 * All rights reserved. This program and the accompanying materials are
 * made available under the terms of the Eclipse Public License v2.0 which
 * accompanies this distribution and is available at
 *
 * http://www.eclipse.org/legal/epl-v20.html
 */

package me.fritx.strings;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.lang.reflect.InvocationTargetException;

import org.junit.jupiter.params.ParameterizedTest;
import org.junitpioneer.jupiter.json.JsonSource;

import com.fasterxml.jackson.core.JsonProcessingException;

class CapitalizeTheTitle_2129_Tests {

    // @ParameterizedTest(name = "data={0}")
    @ParameterizedTest
    // The Java feature 'Text Blocks' is only available with source level 15 and
    // above
    @JsonSource("""
            [
                ["capiTalIze tHe titLe", "Capitalize The Title"],
                ["First leTTeR of EACH Word", "First Letter of Each Word"],
                ["i lOve leetcode", "i Love Leetcode"]
            ]
            """)
    void test(String[] caseData)
            throws JsonProcessingException, ClassNotFoundException, InstantiationException, IllegalAccessException,
            IllegalArgumentException, InvocationTargetException, NoSuchMethodException, SecurityException {

        String[] classNames = {
                CapitalizeTheTitle_2129.class.getName(),
        };
        for (String className : classNames) {
            Class<?> Clz = Class.forName(className);
            // fix: Type mismatch: cannot convert from capture#2-of ? to CodecBT_297
            CapitalizeTheTitle_2129 solution = (CapitalizeTheTitle_2129) Clz.getDeclaredConstructor().newInstance();
            String input = caseData[0];
            String want = caseData[1];
            String ans = solution.capitalizeTitle(input);
            assertEquals(ans, want);
        }

    }
}
