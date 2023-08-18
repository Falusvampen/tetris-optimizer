#!/bin/bash

prepend_folder() {
    local folder=$1
    shift
    local files=("$@")
    for i in "${!files[@]}"; do
        files[$i]="${folder}/${files[$i]}"
    done
    echo "${files[@]}"
}

# List of test files
bad_examples=($(prepend_folder "examples" "badexample00.txt" "badexample01.txt" "badexample02.txt" "badexample03.txt" "badexample04.txt"))
bad_format=($(prepend_folder "examples" "badformat.txt"))
good_examples=($(prepend_folder "examples" "goodexample00.txt" "goodexample01.txt" "goodexample02.txt" "goodexample03.txt"))
hard_example=($(prepend_folder "examples" "hardexample.txt"))

# Initialize an empty string to accumulate errors
errors=""

# Test with bad examples
for file in "${bad_examples[@]}"; do
    echo "Testing with $file"
    go run . "$file"
    echo "----------------------"
    echo ""
done

# Test with bad format
for file in "${bad_format[@]}"; do
    echo "Testing with $file"
    go run . "$file"
    echo "----------------------"
    echo ""
done

# Test with good examples
for file in "${good_examples[@]}"; do
    echo "Testing with $file"
    output=$(go run . "$file" 2>&1)
    if echo "$output" | grep -qi "error"; then
        errors+="Unexpected error for $file:\n$output\n\n"
    fi
    echo "$output"
    echo "----------------------"
    echo ""
done

# Test with hard example
for file in "${hard_example[@]}"; do
    echo "Testing with $file"
    output=$(go run . "$file" 2>&1)
    if echo "$output" | grep -qi "error"; then
        errors+="Unexpected error for $file:\n$output\n\n"
    fi
    echo "$output"
    echo "----------------------"
    echo ""
done

# Print accumulated errors at the end
if [ ! -z "$errors" ]; then
    echo -e "\n\nERROR SUMMARY:"
    echo -e "$errors"
else
    echo -e "\nAll tests passed without unexpected errors!"
fi
