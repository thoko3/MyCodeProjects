#!bin/bash
# A shell script to add two numbers

echo "Enter First Number:"
read num1
echo "Enter Second Number:"
read num2
#perform addition
sum=$((num1 + num2))
#Display result
echo "The Sum of Num1 and Num2 is: #sum"