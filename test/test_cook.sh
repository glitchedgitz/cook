#!/bin/bash

# Function to compare output with expected result
assert_equal() {
    local actual="$1"
    local expected="$2"
    local test_name="$3"
    
    # Trim whitespace from both actual and expected
    actual=$(echo "$actual" | tr -d '[:space:]')
    expected=$(echo "$expected" | tr -d '[:space:]')
    
    if [ "$actual" = "$expected" ]; then
        echo "✅ PASS: $test_name"
    else
        echo "❌ FAIL: $test_name"
        echo "Expected: $expected"
        echo "Actual:   $actual"
        exit 1
    fi
}

# Function to check if output contains expected string
assert_contains() {
    local output="$1"
    local expected="$2"
    local test_name="$3"
    
    if echo "$output" | grep -q "$expected"; then
        echo "✅ PASS: $test_name"
    else
        echo "❌ FAIL: $test_name"
        echo "Expected to find: $expected"
        echo "Actual output: $output"
        exit 1
    fi
}

# Function to check if output contains all expected strings
assert_contains_all() {
    local output="$1"
    shift
    local test_name="$1"
    shift
    
    for expected in "$@"; do
        if ! echo "$output" | grep -q "$expected"; then
            echo "❌ FAIL: $test_name"
            echo "Expected to find: $expected"
            echo "Actual output: $output"
            exit 1
        fi
    done
    echo "✅ PASS: $test_name"
}

# Test 1: Basic pattern matching
echo ""
echo "Test 1: Basic pattern matching"
output=$(cook "test")
assert_equal "$output" "test" "Basic pattern matching"

# Test 2: Method application
echo ""
echo "Test 2: Method application"
output=$(cook "test" -m "upper")
assert_equal "$output" "TEST" "Method application (upper)"

# Test 3: Multiple methods
echo ""
echo "Test 3: Multiple methods"
output=$(cook "test" -m "upper.reverse")
assert_equal "$output" "TSET" "Multiple methods (upper,reverse)"

# Test 4: File input
echo ""
echo "Test 4: File input"
echo "test1" > test_input.txt
echo "test2" >> test_input.txt
output=$(cook -f: test_input.txt f)
assert_contains_all "$output" "File input content" "test1" "test2"
rm test_input.txt

# Test 5: Pattern with parameters
echo ""
echo "Test 5: Pattern with parameters"
output=$(cook test 1,2,3)
assert_contains_all "$output" "Pattern with parameters" "test1" "test2" "test3"

# Test 6: Numeric Ranges
echo ""
echo "Test 6: Numeric Ranges"
output=$(cook test 1-3)
assert_contains_all "$output" "Range generation" "test1" "test2" "test3"

# Test 7: String Ranges
echo ""
echo "Test 7: String Ranges"
output=$(cook test a-c)
assert_contains_all "$output" "String range generation" "testa" "testb" "testc"

# Test 8: Param Approach
echo ""
echo "Test 8: Param Approach"
output=$(cook -start intigriti,bugcrowd -sep _,- -end users.rar,secret.zip / start sep end)
assert_contains_all "$output" "Param approach with separators" "intigriti_users.rar" "intigriti_secret.zip" "bugcrowd_users.rar" "bugcrowd_secret.zip"

# Test 9: Multiple Encoding (Overlapping)
echo ""
echo "Test 9: Multiple Encoding (Overlapping)"
output=$(cook "test" -m "md5.b64e")
# MD5 of "test" is 098f6bcd4621d373cade4e832627b4f6
# Base64 of MD5 is MDk4ZjZiY2Q0NjIxZDM3M2NhZGU0ZTgzMjYyN2I0ZjY=
assert_contains "$output" "MDk4ZjZiY2Q0NjIxZDM3M2NhZGU0ZTgzMjYyN2I0ZjY=" "Overlapping encoding (md5 -> base64)"

# Test 10: Multiple Encoding (Different)
echo ""
echo "Test 10: Multiple Encoding (Different)"
output=$(cook "test" -m "md5,sha1,sha256")
# MD5: 098f6bcd4621d373cade4e832627b4f6
# SHA1: a94a8fe5ccb19ba61c4c0873d391e987982fbbd3
# SHA256: 9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08
assert_contains_all "$output" "Different encoding" \
    "098f6bcd4621d373cade4e832627b4f6" \
    "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3" \
    "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"

# Test 11: Smart Break
echo ""
echo "Test 11: Smart Break"
output=$(cook "adminNew,admin_new" -m "smart")
assert_contains_all "$output" "Smart break" "admin" "New" "new"

# Test 12: Smart Join
echo ""
echo "Test 12: Smart Join"
output=$(cook "adminNew,admin-old" -m "smartjoin[:_]")
assert_contains_all "$output" "Smart join" "admin_New" "admin_old"

# Test 13: String Operations
echo ""
echo "Test 13: String Operations"
output=$(cook "test1:test2:test3" -m "split[:]")
assert_contains_all "$output" "String split operation" "test1" "test2" "test3"

# Test 14: Functions
echo ""
echo "Test 14: Functions"
output=$(cook -dob "date[17,Sep,1994]" elliot _,-, dob)
assert_contains_all "$output" "Function usage with date" "elliot" "17" "Sep" "1994"

# Test 15: Local File with Parameter
echo ""
echo "Test 15: Local File with Parameter"
echo "test1" > test_input.txt
echo "test2" >> test_input.txt
output=$(cook -f: test_input.txt f)
assert_contains_all "$output" "Local file with parameter" "test1" "test2"
rm test_input.txt

# Test 16: Repeat Operator
echo ""
echo "Test 16: Repeat Operator"
output=$(cook "test*3")
assert_contains_all "$output" "Repeat operator" "test" "test" "test"

echo "All tests completed successfully!"
