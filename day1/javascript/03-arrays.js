let arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
console.log("all array's elements", arr);

console.log("element by element");
for (let i = 0; i < arr.length; i++) {
  console.log(arr[i]);
}

arr[0] = 99;
console.log(arr);

console.log("using array's forEach method");
arr.forEach((element) => console.log(element));

let doubleVals = arr.map((element) => {
  return element * 2;
});
console.log(doubleVals);

console.log(
  arr.findIndex((element) => {
    return element === 99;
  })
);
