/*

You are given a feature to display the median price of a given beverage type. Given an n sized unsorted array of integer prices, write a function findMedian([]) to return the median price. You can assume that the array will contain at least one element.
    Tests:
     - findMedian([1,6,3,5,8,9,4,10,2]) => 5
     - findMedian([1,6,3,5,8,9,4,10,2,7]) => 5.5

*/

function findMedian(arr){
    arrlen = arr.length;
    if (arrlen < 2){
        return arr[0];
    }

    //sort array
    arr.sort(function(a, b) {
      return a - b;
    });

    //calculate the key of the middle element in the array
    middle = Math.floor(arrlen/2);

    //check if odd or even number of elements
    if (arrlen % 2 > 0) {
        //array with an odd number of elements
        //the median is the middle number
        return arr[middle];        
    } else {
        //array with an even number of elements
        //the median is the average of the middle two numbers
        num1 = arr[middle -1];
        num2 = arr[middle];
        median = (num1 + num2) / 2;
        return median;
    }
}

console.log(findMedian([1,6,3,5,8,9,4,10,2]));
console.log(findMedian([1,6,3,5,8,9,4,10,2,7]));

